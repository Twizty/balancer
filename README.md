### Balancer

This is a simple and lightweight implementation of reverse-proxy balancer based on round-tripper algorithm. Work is in progress, but it's ready to use (if you'd want to use if for some reason).

If you have something to add, please fork and make a pull-request.

### Example

```go
func main() {
  balancer := NewBalancerNode("/foo", []string{"http://localhost:9091", "http://localhost:9092", "http://localhost:9093"})
  proxy := balancer.NewMultipleHostReverseProxy()

  log.Fatal(http.ListenAndServe(":9090", proxy))
}
```

