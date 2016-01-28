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
	strval = []byte(fmt.Sprintf("\"%s\"", strvalBytes))
    fmt.Printf("String %s = %s\n", strkey, string(strval))

	// Write the state to the ledger
	err = stub.PutState(strkey, strval)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) initTest(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    
    if len(args) != 0 {
        return nil, errors.New("Test initialization doesn't take in arguments.")
    }
    
    // Create an example owner, asset, and licensee
    
}

// Run callback representing the invocation of a chaincode
func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "newOwner" {
		// Initialize
		return t.init(stub, args)
	} else if function == "newLicensee" {
		// Create a new photographer
		return t.json(stub, args)
	} else if function == "testInit" {
        return t.
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
    // Add quotes to the string or it will cause an invalid json to be created.
    strvalBytes = []byte(fmt.Sprintf("\"%s\"", string(strvalBytes)))
	
	jsonResp := "{\"StringKey\":\"" + strkey + "\",\"Value\":\"" + string(strvalBytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return strvalBytes, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "string" {
		return t.getstring(stub, args)
	} else if function == "number" {
		return nil, errors.new("number querying isn't implemented yet")
	} else if function != "number" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}