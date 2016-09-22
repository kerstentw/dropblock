package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type DocumentStorage struct {

}


func main(){

}

func (t *DocumentStorage) Init(stub *shim.ChaincodeShim, function string, args []string) ([]byte, error){

        if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting 1")
        }

        err := stub.PutState("test_throw", []byte(args[0]))
        if err != nil {
                return nil, err
        }

        return nil, nil

}


func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

        // Handle different functions
        } if function == "write" {
                return t.write_to_document(stub, args)
        }
        fmt.Println("invoke did not find func: " + function)

        return nil, errors.New("Received unknown function invocation")
}


func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

        // Handle different functions
        if function == "read" { //read a variable
                return t.read_document(stub, args)
        }
        fmt.Println("query did not find func: " + function)

        return nil, errors.New("Received unknown function query")
}

func (t *SimpleChaincode) write_to_document(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
        var key, value string
        var err error
        fmt.Println("running write()")

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
        }

        key = args[0] //rename for funsies
        value = args[1]
        err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
        if err != nil {
                return nil, err
        }
        return nil, nil
}


func (t *SimpleChaincode) read_document(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
        var key, jsonResp string
        var err error

        if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
        }

        key = args[0]
        valAsbytes, err := stub.GetState(key)
        if err != nil {
                jsonResp = "{\"Error\":\"No State recorded for:" + key + "\"}"
                return nil, errors.New(jsonResp)
        }

        return valAsbytes, nil
}

