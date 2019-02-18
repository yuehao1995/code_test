package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	InitQueryMap(m)
	Handler(m)
	fmt.Println(m)
	Handler2(m)
	fmt.Println(m)
}

//初始化
func InitQueryMap(m map[string]interface{}) {
	mse := make(map[string]interface{})
	m["seletor"] = mse
	mso := make([]map[string]string, 0)
	m["sort"] = mso
}

func Handler(qMap map[string]interface{}) error {
	seletors := qMap["seletor"].(map[string]interface{})
	fmt.Println(seletors["income"])
	expression, ok := seletors["income"].(map[string]interface{})
	if !ok {
		expression = make(map[string]interface{})
	}
	expression["$eq"] = 89
	seletors["income"] = expression
	return nil
}

func Handler2(qMap map[string]interface{}) error {
	seletors := qMap["seletor"].(map[string]interface{})
	expression, ok := seletors["income"].(map[string]interface{})
	if !ok {
		expression = make(map[string]interface{})
	}
	expression["$gt"] = 60
	seletors["income"] = expression
	return nil
}
