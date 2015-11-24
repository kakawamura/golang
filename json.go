// 引数で指定されたURLへHTTPアクセスして、その結果を表示する。
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var data []struct {
		Todo string
		At   string
	}

	url := "https://gist.githubusercontent.com/Iwark/ba6f1f73a059f7f7f675/raw/39bde97ad72096bdf23ccb027b915c5e8c71966b/golang20151124json.json"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		if response.StatusCode == 200 {
			fmt.Println(response.Header["Content-Type"])
			body, _ := ioutil.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			for _, item := range data {
				fmt.Println(item.Todo + "を" + item.At + "にやる")
			}
		}
	}
}
