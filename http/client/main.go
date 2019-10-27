package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	{
		ctx, can := context.WithTimeout(ctx, 5*time.Second)
		defer can()

		req, err := http.NewRequest("GET", "https://github.com", nil)
		if nil != err {
			panic(err)
		}

		req.WithContext(ctx)

		client := http.Client{}
		res, err := client.Do(req)
		if nil != err {
			panic(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(res.StatusCode, body, err)
	}
}
