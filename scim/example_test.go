package scim_test

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/h2non/gock.v1"

	"github.com/suzuki-shunsuke/go-slack-scimapi/scim"
)

const (
	token = "Slack Token"
)

func Example() {
	// use mock for example
	defer gock.Off()
	gock.New("https://api.slack.com").
		Get("/scim/v1/Users").
		MatchType("json").Reply(200).
		BodyString(`{"totalResults": 1, "Resources": [{"userName": "other_username"}]}`)

	client := scim.NewClient(token)
	ctx := context.Background()
	users, resp, err := client.GetUsers(ctx, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d %d\n", resp.StatusCode, len(users.Resources))
	// Output:
	// 200 1
}
