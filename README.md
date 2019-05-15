# go-slack-scimapi

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

You can customize client's behavior with methods `Client.WithXXX`.

```go
client = client.WithParseResp(func(resp *http.Response, output interface{}) error {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return json.Unmarshal(b, output)
})
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
