# go-slack-scimapi

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-slack-scimapi/scim)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-slack-scimapi/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-slack-scimapi)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-slack-scimapi/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-slack-scimapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-slack-scimapi)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-slack-scimapi)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-slack-scimapi.svg)](https://github.com/suzuki-shunsuke/go-slack-scimapi)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-slack-scimapi/master/LICENSE)

Slack SCIM API client for Golang

https://api.slack.com/scim

## How to use

### Create a client

```go
client := scim.NewClient("Slack API Token")
```

### Get users

Note that returned *http.Response.Body is closed.

```go
ctx := context.Background()
users, resp, err := client.GetUsers(ctx, nil, "")
```

#### Pagination

```go
users, resp, err := client.GetUsers(ctx, &scim.Pagination{
	Count: 5000,
	StartInex: 1,
}, "")
```

#### Filter

```go
users, resp, err := client.GetUsers(ctx, nil, `email eq "foo@example.com"`)
```

### Customize response handling

You can customize client's behavior with methods `Client.WithXXX` and `Client.SetXXX` .
`Client.SetXXX` changes the receiver itself.
On the other hand, `Client.WithXXX` changes a shallow copy of client and returns it.
`Client.WithXXX` is useful for the method chain.

```go
client.SetParseResp(func(resp *http.Response, output interface{}) error {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return json.Unmarshal(b, output)
})
```

```go
users, resp, err := client.WithParseErrorResp(func(resp *http.Response) error {
	fmt.Println("customize")
	return ParseErrorRespDefault(resp)
}).GetUsers(ctx, nil, "")
```

### client.XXXResp

`Client.GetUsers` parses response body and returns users.
On the other hand, `Client.GetUsersResp` doesn't parse response body.
Returned response body is open.

```go
resp, err := client.GetUsersResp(ctx, nil, "")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()
b, err := ioutil.ReadAll(resp.Body)
```

```go
resp, err := client.CreateUser(ctx, &scim.User{
	UserName: "foo",
	Emails: []scim.Email{
		Value: "foo@example.com",
	},
})
```

## License

[MIT](LICENSE)
