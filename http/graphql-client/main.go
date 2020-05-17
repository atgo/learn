package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func main() {
	url := "https://qa/query"
	client := graphql.NewClient(url)

	req := graphql.NewRequest(`
		query($email: String) {
		  findUser(
			filters: {
			  email: { eq: $email }
			}
		  ) {
			legacyId
			email
		  }
		}`,
	)

	req.Var("email", "andy@go1.com")
	res := struct {
		Data struct {
			Id    int    `json:"legacyId"`
			Email string `json:"email"`
		} `json:"findUser"`
	}{}

	err := client.Run(context.Background(), req, &res)
	if nil != err {
		panic(err)
	} else {
		fmt.Println("result: ", res.Data)
	}
}
