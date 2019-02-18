/**
 * @author zhangyuehao
 * @date 2019-01-17 19:10
 */

package main

import (
	"encoding/json"
	"fmt"
)

var JsonMap map[string]interface{}

func main() {
	JsonMap := make(map[string]interface{})
	content := "{\"code\":\"UUID4\",\"name\":\"合同4\",\"signYear\":2004,\"signMonth\":4,\"signerName\":\"zhangsan4\",\"signerSex\":\"male\"}"
	err := json.Unmarshal([]byte(content), &JsonMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("JsonMap为%v", JsonMap)
}
