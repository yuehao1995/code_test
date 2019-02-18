/**
 * Created by g7tianyi on 23/9/2018
 */

package main

import (
	"fmt"

	"bytes"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type AbstractChainCode struct{}

func (c *AbstractChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("########### Invoking ChainCode ###########")

	function, args := stub.GetFunctionAndParameters()
	if function != "invoke" {
		return shim.Error("Unknown function call")
	}

	if len(args) < 1 {
		return shim.Error("Insufficient arguments")
	}

	if args[0] == "create" {
		if len(args) != 3 {
			return shim.Error("Insufficient arguments for create")
		}
		return c.create(stub, args[1], args[2])
	}

	if args[0] == "update" {
		if len(args) != 3 {
			return shim.Error("Insufficient arguments for update")
		}
		return c.update(stub, args[1], args[2])
	}

	if args[0] == "delete" {
		if len(args) != 2 {
			return shim.Error("Insufficient arguments for delete")
		}
		return c.delete(stub, args[1])
	}

	if args[0] == "queryByKey" {
		if len(args) != 2 {
			return shim.Error("Insufficient arguments for queryByKey")
		}
		return c.queryByKey(stub, args[1])
	}

	if args[0] == "queryBySelector" {
		if len(args) != 2 {
			return shim.Error("Insufficient arguments for queryBySelector")
		}
		return c.queryBySelector(stub, args[1])
	}

	if args[0] == "queryHistoryByKey" {
		if len(args) != 2 {
			return shim.Error("Insufficient arguments for queryHistoryByKey")
		}
		return c.queryHistoryByKey(stub, args[1])
	}

	return shim.Error("Unknown action, check the first argument")
}

func (c *AbstractChainCode) create(stub shim.ChaincodeStubInterface, key, value string) peer.Response {
	b, err := stub.GetState(key)
	if err != nil {
		return shim.Error(formatErrorMessage("Failed to get key", key))
	}
	if b != nil {
		return shim.Error(formatErrorMessage("Key already exists ", key))
	}

	err = stub.PutState(key, []byte(value))
	if err != nil {
		return formatErrorResponse(formatErrorMessage("Failed to create for key ", key), err)
	}

	return shim.Success(nil)
}

// 在调用此函数之前，请确保value包含了整个链上数据结构所要求的全部字段，可以理解为REST-ful的PUT而不是PATCH动作
func (c *AbstractChainCode) update(stub shim.ChaincodeStubInterface, key, value string) peer.Response {
	b, err := stub.GetState(key)
	if err != nil {
		return shim.Error(formatErrorMessage("Failed to get key", key))
	}
	if b == nil {
		return shim.Error(formatErrorMessage("Key not found exists ", key))
	}

	err = stub.PutState(key, []byte(value))
	if err != nil {
		return formatErrorResponse(formatErrorMessage("Failed to update key ", key), err)
	}

	return shim.Success(nil)
}

func (c *AbstractChainCode) delete(stub shim.ChaincodeStubInterface, key string) peer.Response {
	b, err := stub.GetState(key)
	if err != nil {
		return shim.Error(formatErrorMessage("Failed to get key", key))
	}
	if b == nil {
		return shim.Error(formatErrorMessage("Key not found exists ", key))
	}

	err = stub.DelState(key)
	if err != nil {
		return formatErrorResponse(formatErrorMessage("Failed to delete key ", key), err)
	}

	return shim.Success(nil)
}

func (c *AbstractChainCode) queryByKey(stub shim.ChaincodeStubInterface, key string) peer.Response {

	fmt.Println("########### queryByKey ###########")

	state, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Failed to queryByKey " + key)
	}
	return shim.Success(state)
}

func (c *AbstractChainCode) queryBySelector(stub shim.ChaincodeStubInterface, couchQuery string) peer.Response {

	fmt.Println("########### queryBySelector ###########")

	resultsIterator, err := stub.GetQueryResult(couchQuery)
	if err != nil {
		return formatErrorResponse(formatErrorMessage("Failed to queryHistoryByKey with couchQuery", couchQuery), err)
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	arrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return formatErrorResponse("Failed to read resultsIterator", err)
		}

		if arrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		arrayMemberAlreadyWritten = true

		buffer.WriteString("{\"key\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"value\":")
		buffer.WriteString(string(response.Value))

		buffer.WriteString("}")
	}

	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (c *AbstractChainCode) queryHistoryByKey(stub shim.ChaincodeStubInterface, key string) peer.Response {

	fmt.Println("########### queryHistoryByKey ###########")

	resultsIterator, err := stub.GetHistoryForKey(key)
	if err != nil {
		return formatErrorResponse(formatErrorMessage("Failed to queryHistoryByKey with key", key), err)
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	arrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return formatErrorResponse("Failed to read resultsIterator", err)
		}

		if arrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		arrayMemberAlreadyWritten = true

		buffer.WriteString("{\"txid\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"value\":")
		// 如果是删除操作，则将对应的历史值记为null
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"isDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
	}

	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func formatErrorMessage(message, highlight string) string {
	return fmt.Sprintf("%s [%s]", message, highlight)
}

func formatErrorResponse(message string, err error) peer.Response {
	return shim.Error(fmt.Sprintf("%s: %v", message, err))
}
