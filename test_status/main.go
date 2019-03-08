/**
 * @author zhangyuehao
 * @date 2019-03-05 14:30
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:36136/ping")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
