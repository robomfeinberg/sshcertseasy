package main

import (
  "fmt"
  "github.com/okta/okta-sdk-golang/okta"
 // "github.com/okta/okta-sdk-golang/okta/query"
)

// Okta golang sdk is not so easy. Documentation is terrible. Maybe I just need more golang experience.

func main() {
  user := "matthew.feinberg"
  //client := okta.NewClient(config, nil, nil)
  //client := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))
  //config := okta.NewConfig().WithOrgUrl().WithToken()
  client := okta.NewClient(context, okta.WithOrgUrl("https://OKTADOMAIN.okta.com/"), okta.WithToken("OKTAAPIKEY"))
  groups, resp, err := client.User.ListUserGroups(user.Id, nil)
  fmt.Println(user.Id)
 // filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))
 // users, _, _ := client.User.ListUsers(filter)
 // for _, user := range users {
 //   fmt.Println(user.Id)
 // }
}