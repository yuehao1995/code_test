package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init ...
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)

}

// Invoke ...
// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### example_cc Invoke ###########")
	args := stub.GetStringArgs()
	if len(args) == 0 {
		return shim.Error("Function not provided")
	}

	function := args[0]

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting at least 2")
	}

	if function == "put" {
		return t.put(stub, args[1:])
	}

	if function == "get" {
		return t.get(stub, args[1:])
	}

	if function == "delete" {
		return t.delete(stub, args[1:])
	}

	return shim.Error(fmt.Sprintf("Unknown function call: %s", function))
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	key := args[1]

	// Delete the key from the state in ledger
	err := stub.DelState(key)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) put(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Invalid args. Expecting key and value")
	}

	key := args[0]
	value := args[1]

	existingValue, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error getting data for key [%s]: %s", key, err))
	}
	if existingValue != nil {
		value = string(existingValue) + "-" + value
	}

	if err := stub.PutState(key, []byte(value)); err != nil {
		return shim.Error(fmt.Sprintf("Error putting data for key [%s]: %s", key, err))
	}

	return shim.Success([]byte(value))
}

func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid args. Expecting key")
	}

	key := args[0]

	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error getting data for key [%s]: %s", key, err))
	}

	return shim.Success([]byte(value))
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
