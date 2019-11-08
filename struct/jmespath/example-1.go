package main

import (
    "encoding/json"
    "fmt"

    "github.com/jmespath/go-jmespath"
)

// output:
//   row map[ID:123]
//   row map[ID:222]

func main() {
    hits := []byte(`
    [
            { "id": 123, "timestamp": 456 },
            { "id": 222, "timestamp": 222 }
        ]
`)

    in := []interface{}{}
    _ = json.Unmarshal(hits, &in)

    node, err := jmespath.Compile(`[*].{ID: id}`)
    if nil != err {
        panic(err)
    }

    results, err := node.Search(in)
    if nil != err {
        panic(err)
    }

    if rows, ok := results.([]interface{}); ok {
        for _, row := range rows {
            fmt.Println("row", row)
        }
    }
}
