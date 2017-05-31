### Balancer

This is a simple and lightweight implementation of reverse-proxy balancer based on round-tripper algorithm. Work is in progress, but it's ready to use (if you'd want to use if for some reason).

If you have something to add, please fork and make a pull-request.

### Example

```go
func main() {
  b := balancer.NewBalancerNode("/foo", []string{"http://localhost:9091", "http://localhost:9092", "http://localhost:9093"})
  proxy := b.NewMultipleHostReverseProxy()

  log.Fatal(http.ListenAndServe(":9090", proxy))
}
```

Then all your request to `localhost:9090/foo/**` will balance between `localhost:9091/**`, `localhost:9092/**`, `localhost:9093/**`.

If you do not want to use prefix, use `NewBalancerWithoutPrefix`.
