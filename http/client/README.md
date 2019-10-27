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
