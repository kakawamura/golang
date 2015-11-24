// 引数で指定されたURLへHTTPアクセスして、その結果を表示する。
package main

import (
  "os"
  "fmt"
  "bufio"
)

func main() {
  grepString := os.Args[1]
  filePath := os.Args[2]

  file, err := os.Open(filePath)

  if err != nil {
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords)

  count := 0
  for scanner.Scan() {
    text := scanner.Text()
    if grepString == text {
      fmt.Println(text)
      count++
    }
  }
  fmt.Println(count)

}
