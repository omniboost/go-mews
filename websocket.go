package mews

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/omniboost/go-mews/commands"
	"github.com/omniboost/go-mews/reservations"
	"github.com/omniboost/go-mews/resources"
)

var (
	WebsocketURL = &url.URL{
		Scheme: "wss",
		Host:   "ws.mews.com",
		Path:   "/ws/connector",
	}
	WebsocketURLDemo = &url.URL{
		Scheme: "wss",
		Host:   "ws.mews-demo.com",
		Path:   "/ws/connector",
	}
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

type Websocket struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// Base URL for API requests
	baseURL *url.URL

	// Debugging flag
	debug bool

	// Disallow unknown json fields
	disallowUnknownFields bool

	accessToken string
	clientToken string

	connection *websocket.Conn
	cancelFunc context.CancelFunc

	// msgChan  chan []byte
	errChan chan error

	cmdChan         chan CommandEvent
	resChan         chan ReservationEvent
	resourceChan    chan ResourceEvent
	priceUpdateChan chan PriceUpdateEvent
}

func NewWebsocket(httpClient *http.Client, accessToken string, clientToken string) *Websocket {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	ws := &Websocket{}
	ws.SetAccessToken(accessToken)
	ws.SetClientToken(clientToken)
	ws.SetDebug(false)
	ws.SetBaseURL(WebsocketURL)

	return ws
}

func (ws Websocket) AccessToken() string {
	return ws.accessToken
}

func (ws *Websocket) SetAccessToken(accessToken string) {
	ws.accessToken = accessToken
}

func (ws Websocket) ClientToken() string {
	return ws.clientToken
}

func (ws *Websocket) SetClientToken(clientToken string) {
	ws.clientToken = clientToken
}

func (ws *Websocket) BaseURL() *url.URL {
	return ws.baseURL
}

func (ws *Websocket) SetBaseURL(baseURL *url.URL) {
	ws.baseURL = baseURL
	ws.baseURL.Scheme = "wss"
}

func (ws *Websocket) Debug() bool {
	return ws.debug
}

func (ws *Websocket) SetDebug(debug bool) {
	ws.debug = debug
}

func (ws *Websocket) CommandEvents() chan (CommandEvent) {
	ws.cmdChan = make(chan CommandEvent)
	return ws.cmdChan
}

func (ws *Websocket) ReservationEvents() chan (ReservationEvent) {
	ws.resChan = make(chan ReservationEvent)
	return ws.resChan
}

func (ws *Websocket) ResourceEvents() chan (ResourceEvent) {
	ws.resourceChan = make(chan ResourceEvent)
	return ws.resourceChan
}

func (ws *Websocket) PriceUpdateEvents() chan (PriceUpdateEvent) {
	ws.priceUpdateChan = make(chan PriceUpdateEvent)
	return ws.priceUpdateChan
}

func (ws *Websocket) Errors() chan (error) {
	ws.errChan = make(chan error)
	return ws.errChan
}

func (ws *Websocket) Connect(ctx context.Context) error {
	var err error
	var resp *http.Response

	d := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	d.Jar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}

	cookies := []*http.Cookie{
		&http.Cookie{Name: "ClientToken", Value: ws.ClientToken()},
		&http.Cookie{Name: "AccessToken", Value: ws.AccessToken()},
	}

	u := *ws.BaseURL()
	u.Scheme = "https"
	d.Jar.SetCookies(&u, cookies)

	ws.connection, resp, err = d.Dial(ws.BaseURL().String(), nil)
	if err != nil {
		if ws.debug && resp != nil {
			b, _ := httputil.DumpResponse(resp, true)
			log.Println(string(b))
		}
		return err
	}

	// Time allowed to read the next pong message from the peer.
	ws.connection.SetReadDeadline(time.Now().Add(pongWait))
	// After receiving a pong: reset the read deadline
	ws.connection.SetPongHandler(func(data string) error {
		if ws.Debug() {
			log.Printf("Received pong message; %s", data)
		}
		ws.connection.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	ws.connection.SetPingHandler(func(data string) error {
		if ws.Debug() {
			log.Printf("Received ping message; %s", data)
		}
		return nil
	})

	// Send ping messages. Stop doing that when context is canceled
	go func() {
		err := ws.KeepAlive(ctx)
		if err != nil {
			ws.errChan <- err
		}
	}()

	// Receive close messages from the peer
	// ws.connection.SetCloseHandler(func(code int, text string) error {
	// 	log.Printf("Received close message from peer: %s (%d)", text, code)
	// 	return ws.Close()
	// })

	// read messages
	go func() {
		for {
			select {
			case <-ctx.Done():
				if ws.debug {
					log.Println("stopping reading messages: context is canceled")
				}
				break
			default:
				if ws.debug {
					log.Println("waiting to receive message")
				}
				msgType, msg, err := ws.connection.ReadMessage()
				if ws.debug {
					log.Printf("received msg on websocket: %s (%d)", msg, msgType)
				}
				if err != nil {
					if ws.errChan != nil {
						ws.errChan <- err
					}
					return
				}

				message := Message{}
				err = json.Unmarshal(msg, &message)
				if err != nil {
					if ws.errChan != nil {
						ws.errChan <- err
					}
					return
				}

				for _, b := range message.Events {
					event := Event{}
					err = json.Unmarshal(b, &event)
					if err != nil {
						if ws.errChan != nil {
							ws.errChan <- err
						}
						return
					}

					if event.Type == EventTypeDeviceCommand && ws.cmdChan != nil {
						cmdEvent := CommandEvent{}
						err := json.Unmarshal(b, &cmdEvent)
						if err != nil {
							if ws.errChan != nil {
								ws.errChan <- err
							}
							return
						}
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushing command %s to cmd channel", cmdEvent.ID))
						}
						ws.cmdChan <- cmdEvent
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushed command %s to cmd channel", cmdEvent.ID))
						}
					}

					if event.Type == EventTypeReservation && ws.resChan != nil {
						resEvent := ReservationEvent{}
						err := json.Unmarshal(b, &resEvent)
						if err != nil {
							if ws.errChan != nil {
								ws.errChan <- err
							}
							return
						}
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushing command %s to res channel", resEvent.ID))
						}
						ws.resChan <- resEvent
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushed command %s to res channel", resEvent.ID))
						}
					} else if event.Type == EventTypeResource && ws.resourceChan != nil {
						resourceEvent := ResourceEvent{}
						err := json.Unmarshal(b, &resourceEvent)
						if err != nil {
							if ws.errChan != nil {
								ws.errChan <- err
							}
							return
						}
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushing command %s to resource channel", resourceEvent.ID))
						}
						ws.resourceChan <- resourceEvent
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushed command %s to resource channel", resourceEvent.ID))
						}
					} else if event.Type == EventTypePriceUpdate && ws.priceUpdateChan != nil {
						priceUpdateEvent := PriceUpdateEvent{}
						err := json.Unmarshal(b, &priceUpdateEvent)
						if err != nil {
							if ws.errChan != nil {
								ws.errChan <- err
							}
							return
						}
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushing command %s to price update channel", priceUpdateEvent.ID))
						}
						ws.priceUpdateChan <- priceUpdateEvent
						if ws.debug {
							log.Println(fmt.Sprintf("websocket: pushed command %s to prive update channel", priceUpdateEvent.ID))
						}
					}
				}
			}
		}
	}()

	return nil
}

func (ws *Websocket) KeepAlive(ctx context.Context) error {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if ws.Debug() {
				log.Println("sending keep alive ping message")
			}
			err := ws.connection.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait))
			if err != nil {
				return err
			}
		case <-ctx.Done():
			if ws.Debug() {
				log.Println("keep alive stopped")
			}
			return nil
		}
	}
}

func (ws *Websocket) Close() error {
	// Send close message to the peer
	if ws.Debug() {
		log.Println("send close message to peer")
	}
	message := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	err := ws.connection.WriteControl(websocket.CloseMessage, message, time.Now().Add(writeWait))
	if err != nil {
		return err
	}

	// wait for a specified time before force-closing the connection
	time.Sleep(closeGracePeriod)

	// Close closes the underlying network connection without sending or waiting for a close message.
	return ws.connection.Close()
}

func (ws *Websocket) ReadMessages() {
}

func (ws *Websocket) Stop() {
	ws.connection.Close()
}

type Message struct {
	Events []json.RawMessage `json:"Events"`
}

type Events []Event

type Event struct {
	Type  EventType             `json:"Type"`  // Type of the event.
	ID    string                `json:"Id"`    // Unique identifier of the Command.
	State commands.CommandState `json:"State"` // State of the command.
}

type EventType string

const (
	EventTypeDeviceCommand EventType = "DeviceCommand"
	EventTypeReservation   EventType = "Reservation"
	EventTypeResource      EventType = "Resource"
	EventTypePriceUpdate   EventType = "PriceUpdate"
)

type CommandEvent struct {
	Event
}

type ReservationEvent struct {
	Event

	ID                 string                        `json:"Id"`                 // Unique identifier of the Reservation.
	State              reservations.ReservationState `json:"State"`              // State of the reservation.
	StartUTC           time.Time                     `json:"StartUtc"`           // Start of the reservation (arrival) in UTC timezone in ISO 8601 format.
	EndUTC             time.Time                     `json:"EndUtc"`             // End of the reservation (departure) in UTC timezone in ISO 8601 format.
	AssignedResourceID string                        `json:"AssignedResourceId"` // Unique identifier of the operations/enterprises#resource assigned to the reservation.
}

type ResourceEvent struct {
	Event

	State resources.ResourceState `json:"State"` // State of the resource.

}

type PriceUpdateEvent struct {
	Event

	StartUTC           time.Time `json:"StartUtc"`           // Start of the price update interval in UTC timezone in ISO 8601 format.
	EndUtc             time.Time `json:"EndUtc"`             // End of the price update interval in UTC timezone in ISO 8601 format.
	RateID             string    `json:"RateId"`             // Unique identifier of the Rate assigned to the update price event.
	ResourceCategoryID string    `json:"ResourceCategoryId"` // Unique identifier of the Resource category assigned to the update price event.
}
