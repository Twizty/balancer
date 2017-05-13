package balancer
 
import (
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "strings"
)

type BalancerNode struct {
  counter *int
  prefix  string
  nodes   []*url.URL
}

func (self *BalancerNode) dropCounter() {
  if *self.counter > 1000 {
    *self.counter = 0
  }
}

func (self *BalancerNode) NewMultipleHostReverseProxy() *httputil.ReverseProxy {
  director := func(req *http.Request) {
    defer self.dropCounter()

    *self.counter++
    target := self.nodes[*self.counter % len(self.nodes)]
    req.URL.Scheme = target.Scheme
    req.URL.Host = target.Host
    req.URL.Path = strings.TrimPrefix(req.URL.Path, self.prefix)
  }

  return &httputil.ReverseProxy{Director: director}
}

func NewBalancerNode(prefix string, hosts []string) *BalancerNode {
  var i *int = new(int)
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
