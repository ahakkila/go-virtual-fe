package main

import (
  "net/http"
  "net/http/httputil"
  "fmt"
  "time"
  "net/url"
)

var Proxies map[string]*httputil.ReverseProxy

func proxyHandler(w http.ResponseWriter, r *http.Request) {
  p := Proxies[r.Host]
  if p != nil {
    p.ServeHTTP(w, r)
  } else {
    http.NotFound(w, r)
  }
  return
}

func main() {
  // frontend Host string : backend port
  vhosts := map[string]string{
    "www1.example.com" : "1234",
    "www2.example.com:8080" : "1235",
  }
  Proxies = make(map[string]*httputil.ReverseProxy)

  for siteurl, backendport := range vhosts {
    fmt.Println("Adding backend " + siteurl + " on port " + backendport)

    backendurl, err := url.Parse("http://localhost:"+ backendport)
    if err != nil {
      panic(err)
    }
    proxy := httputil.NewSingleHostReverseProxy(backendurl)
    Proxies[siteurl] = proxy
  }

  mux := http.NewServeMux()
  mux.HandleFunc("/", proxyHandler)

  srv := &http.Server{
    Handler: mux,
    Addr: ":80",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }
  fmt.Println(srv.ListenAndServe())
}
 
