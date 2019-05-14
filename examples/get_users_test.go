package examples

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/h2non/gock.v1"

	"github.com/suzuki-shunsuke/go-slack-scimapi/scim"
)

func Example() {
	defer gock.Off()
	gock.New("https://api.slack.com").
		Get("/scim/v1/Users").
		MatchType("json").Reply(200).
		BodyString(`{
  "totalResults": 1,
  "Resources": [
    {
      "userName": "other_username"
    }
  ]
}`)

	client := scim.NewClient("XXX")
	ctx := context.Background()
	users, resp, err := client.GetUsers(ctx, nil, "")
	if resp != nil {
		fmt.Printf("status code: %d\n", resp.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total results: %d\n", users.TotalResults)
	for _, user := range users.Resources {
		fmt.Printf("user name: %s\n", user.UserName)
	}
	gock.Off()

	gock.New("https://api.slack.com").
		Get("/scim/v1/Users").
		MatchType("json").Reply(401).
		BodyString(`{
  "Errors": {
    "description": "invalid_authentication",
    "code": 401
  }
}`)
	users, resp, err = client.GetUsers(ctx, nil, "")
	if resp != nil {
		fmt.Printf("status code: %d\n", resp.StatusCode)
	}
	if err != nil {
		if e, ok := err.(*scim.Error); ok {
			fmt.Printf("description: %s, code: %d\n", e.Description, e.Code)
		}
		fmt.Println(err)
	}

	// Output:
	// status code: 200
	// total results: 1
	// user name: other_username
	// status code: 401
	// description: invalid_authentication, code: 401
	// invalid_authentication
}
