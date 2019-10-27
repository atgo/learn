HTTP client
====

## With timeout

```go

{
  ctx, can := context.WithTimeout(ctx, 5*time.Second)
  defer can()
  
  req, _ := http.NewRequest("GET", "https://github.com", nil)
  req.WithContext(ctx)
}
```

## Read response body

```go
client := http.Client{}
res, err := client.Do(req)
defer res.Body.Close()

body, err := ioutil.ReadAll(res.Body)
fmt.Println(res.StatusCode, body, err)
```
