package balancer
 
import (
  "net/http"
  "net/http/httputil"
  "net/url"
  "strings"
  "sync/atomic"
)

const (
  MAX_COUNTER_VALUE = int32(1000)
  COUNTER_STEP      = int32(1)
)

type BalancerNode struct {
  counter *int32
  prefix  string
  nodes   []*url.URL
}

func (self *BalancerNode) NewMultipleHostReverseProxy() *httputil.ReverseProxy {
  director := func(req *http.Request) {
    atomic.AddInt32(self.counter, COUNTER_STEP)
    target := self.nodes[int(*self.counter) % len(self.nodes)]
    req.URL.Scheme = target.Scheme
    req.URL.Host = target.Host
    req.URL.Path = strings.TrimPrefix(req.URL.Path, self.prefix)
  }

  return &httputil.ReverseProxy{Director: director}
}

func NewBalancerNode(prefix string, hosts []string) *BalancerNode {
  var i *int32 = new(int32)
  nodes := make([]*url.URL, len(hosts))
  *i = 0

  for i, e := range hosts {
    u, err := url.Parse(e)
    if err != nil {
      panic(err)
    }

    nodes[i] = u
  }
  return &BalancerNode{i, prefix, nodes}
}

func NewBalancerWithoutPrefix(hosts []string) *BalancerNode {
  return NewBalancerNode("", hosts)
}
