// 引数で指定されたURLへHTTPアクセスして、その結果を表示する。
package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  url := "https://gist.githubusercontent.com/Iwark/8a294789add68b9a61fa/raw/f720b43e12539d03e2550e680b54baf8825f2db1/golang20151124.md"
  response, err := http.Get(url);
  if err != nil {
    fmt.Println(err)
  } else {
    if response.StatusCode == 200 {
      fmt.Println(response.Header["Content-Type"])
      body, _ := ioutil.ReadAll(response.Body)
      fmt.Println(string(body))
    }
  }

}
