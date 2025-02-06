# Go Mews API client

go-mews is an API client library for accessing the Mews API v1.

API documentation can be found here:
https://mewssystems.github.io/public/content/developers/api.html

## Usage

``` go
import "github.com/omniboost/go-mews"
```

### Request companies

``` go
// get token
token := os.Getenv("MEWS_TOKEN")

// build client
client := mews.NewClient(nil, token)
client.SetDebug(true)

// request all accounting categories
requestBody := &accountingcategories.AllRequest{}
resp, err := client.AccountingCategories.All(requestBody)
if err != nil {
	panic(err)
}

categories := resp.AccountingCategories
```

### Request all employees for a company

``` go
import "github.com/tim-online/go-nmbrs/employees"

// get id of company
companyID := companies[0].ID

// request all employees for this company ID
resp2, err := client.Employees.ListByCompany(companyID, employees.All)
if err != nil {
	panic(err)
}
```

