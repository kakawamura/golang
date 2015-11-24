
package main

import (
  "fmt"
  "net/http"
  "golang.org/x/net/html"
  "io/ioutil"
  "strings"
)

func main() {

  url := "https://candypot.jp/summaries/1050"
  response, err := http.Get(url);
  if err != nil {
    fmt.Println(err)
  } else {
    if response.StatusCode == 200 {
      fmt.Println(response.Header["Content-Type"])
      body, _ := ioutil.ReadAll(response.Body)

      doc, err := html.Parse(strings.NewReader(string(body)))
      if err != nil {

      } else {
        var f func(*html.Node)
        f = func(n *html.Node) {
          if n.Type == html.ElementNode && n.Data == "h2" {
            fmt.Printf("%q\n", n.FirstChild.Data) 
          }
          for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
          }
        }
        f(doc)
      }
    }

  }

}
