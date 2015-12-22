/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/openblockchain/obc-peer/openchain/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

/**
 * Initializes the blockchain state with a string that we are going to attempt to run DRM on.
 * Args:
 * 0 strkey: The key that we are associating with our protected string in the ledger
 * 1 strval: The value that we want to assign to the string (make it one-of-a-kind!)
 */
func (t *SimpleChaincode) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var strkey string  // Key associated with the string
	var strval []byte  // Actual value of the string
	var err error	   // The error that gets returned, if there is one

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	// Parse the protected string from the arguments
	strkey = args[0]
	strval = []byte(args[1])

	// Write the state to the ledger
	err = stub.PutState(strkey, strval)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

/**
 * Creates a new Party that can own Images
 * Args:
 * 0 key: The key that will be associated with the owner
 * 1 name: The owner's name
 */
func (t *SimpleChaincode) initImgSrc(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	return nil, nil
}

// Run callback representing the invocation of a chaincode
func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "init" {
		// Initialize
		return t.init(stub, args)
	} else if function == "initImgSrc" {
		// Create a new photographer
		return t.initImgSrc(stub, args)
	}

	return nil, nil
}

/**
 * Should allow us to query the value of the protected strings in the ledger
 * Args:
 * stub: The blockchain that holds the string we're want.
 * strkey: The key under which the string is stored.
 */
func (t *SimpleChaincode) getstring(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var strkey string
	var strvalBytes []byte  // The actual value of the string
	
	if(len(args) != 1) {
		return nil, errors.New("Incorrect number of arguments. Expecting 1");
	}
	
	// Which string do we want?
	strkey = args[0];
	
	// Get the value of the string from the ledger
	strvalBytes, err := stub.GetState(strkey)
	if err != nil {
		jsonResp := "{\"Error\": \"Failed to get state for " + strkey + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	if strvalBytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + strkey + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	jsonResp := "{\"StringKey\":\"" + strkey + "\", \"Value\":\"" + string(strvalBytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return strvalBytes, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "string" {
		return t.getstring(stub, args);
	} else if function == "number" {
		
	} else if function != "number" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entity
	var Aval int // Asset holding
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d\n", Aval)

	// Write the state to the ledger - this put is illegal within Run
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		jsonResp := "{\"Error\":\"Cannot put state within chaincode query\"}"
		return nil, errors.New(jsonResp)
	}

	fmt.Printf("Something is wrong. This query should not have succeeded")
	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}