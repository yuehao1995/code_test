/**
 * @author zhangyuehao
 * @date 2019-01-09 15:08
 */

package main

import (
	"encoding/json"
	"fmt"
)

type contractEntity struct {
	Content string `json:"content"`
	Digest  string `json:"digest"`
}

type Contract struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	SignYear   int    `json:"signYear"`
	SignMonth  int    `json:"signMonth"`
	SignerName string `json:"signerName"`
	SignerSex  string `json:"signerSex"`
}

func main() {
	test := []*Contract{
		&Contract{Code: "UUID1", Name: "合同1", SignYear: 2008, SignMonth: 2, SignerName: "zhangsan1", SignerSex: "male"},
		&Contract{Code: "UUID2", Name: "合同2", SignYear: 2010, SignMonth: 5, SignerName: "zhangsan2", SignerSex: "female"},
		&Contract{Code: "UUID3", Name: "合同3", SignYear: 2012, SignMonth: 8, SignerName: "zhangsan3", SignerSex: "male"},
		&Contract{Code: "UUID4", Name: "合同4", SignYear: 2014, SignMonth: 11, SignerName: "zhangsan4", SignerSex: "female"},
		&Contract{Code: "UUID5", Name: "合同5", SignYear: 2016, SignMonth: 3, SignerName: "zhangsan5", SignerSex: "female"},
	}
	entity := contractEntity{}
	for i := 0; i < len(test); i++ {
		byte, err := json.Marshal(test[i])
		if err != nil {
		}
		fmt.Println(string(byte))
		entity.Content = string(byte)
		entity.Digest = "shakjkdsadgjhd1311"
		byte2, err := json.Marshal(entity)
		if err != nil {
		}
		fmt.Println(string(byte2))
	}

}
