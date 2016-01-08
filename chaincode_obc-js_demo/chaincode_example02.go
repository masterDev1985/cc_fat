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
	"encoding/json"

	"github.com/openblockchain/obc-peer/openchain/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

/* should we make something like this?
@-',---
{
	"consumptionSpec":{
		"vars":{
			"A": "string",
			"B": "string"
		}
		func:{
			"invoke":{
				"from_var": string,
				"to_var": string,
				"value": integer
			},
			"init":{
				
			}
		}
	}
}
--,'-@
*/

/*
type Person struct{
	UserId string
	FullName string
	Address string
	PublicKey string
}
*/

type CarData struct{
	Vin string
	Year string
	Make string
	Model string
	License string
}

type CarUsers struct{
	UserId string
	Permissions []string
}

type Car struct{
	Data CarData
	Users []CarUsers
}

func (t *SimpleChaincode) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var A, B string    																// Entities
	var Aval, Bval int 																// Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ============================================================================================================================
// Run
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "init" {													// Initialize the entities and their asset holdings
		return t.init(stub, args)
	} else if function == "delete" {										// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "write" {											// Writes a value to the chaincode state
		return t.Write(stub, args)
	} else if function == "init_person" {									//init_person
		return t.init_person(stub, args)
	} else if function == "init_car" {										//init car
		return t.init_car(stub, args)
	} else if function == "attach_license" {								//attach license
		return t.attach_license(stub, args)
	} else if function == "get_permissions" {								//get permissions
		return t.get_permissions(stub, args)
	} else if function == "set_user_perms" {								//set user permissions
		return t.set_user_perms(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

//----------------------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------------------

// ============================================================================================================================
// Write var into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) Write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var name string // Entities
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the variable and value to set")
	}

	name = args[0]

	// Write the state back to the ledger
	err = stub.PutState(name, []byte(args[1]))
	if err != nil {
		return nil, err
	}
	//t.remember_me(name, args[0])

	return nil, nil
}

// ============================================================================================================================
// Init Person 
// ============================================================================================================================
func (t *SimpleChaincode) init_person(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	str := `{"userid": "` + args[0] + `", "fullname": "` + args[1] + `", "address": "` + args[2] + `", "publickey": "` + args[3] + `"}`

	// Write the state back to the ledger
	err = stub.PutState(args[0], []byte(str))							//store person with userid as key
	if err != nil {
		return nil, err
	}
	//t.remember_me(stub, args[0])
	
	return nil, nil
}


// ============================================================================================================================
// Init Car 
// ============================================================================================================================
func (t *SimpleChaincode) init_car(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error

	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
	}

	str := `{
				"data": {
					"vin": "`   + args[0] + `",
					"year": "`  + args[1] + `"
					"make": "`  + args[2] + `",
					"model": "` + args[3] + `",
					"license": "-"
				},
				"users": [{
					"userid": "` + args[4] + `",
					"permissions":["owner"]
				}]
			}`
			

	// Write the state back to the ledger
	err = stub.PutState(args[0], []byte(str))			//store car with vin# as key
	if err != nil {
		return nil, err
	}
	t.remember_me(stub, args[0]);
	
	return nil, nil
}

// ============================================================================================================================
// Remeber Me - remember the name of variables we stored in ledger 
// ============================================================================================================================
func (t *SimpleChaincode) remember_me(stub *shim.ChaincodeStub, name string) ([]byte, error) {		//dsh - to do, should probably not exist here, move to stub
	var err error
	var ben = "_ben_knows"
	var storedNames string
	
	storeNamesAsBytes, err := stub.GetState(ben)
	if err != nil {
		return nil, errors.New("Failed to get ben")
	}
	storedNames = string(storeNamesAsBytes)
	
	// Write the state back to the ledger
	err = stub.PutState(ben, []byte(storedNames + "," + name))										//dsh - to do, should probably be json
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Attach License - edit license field in car profile
// ============================================================================================================================
func (t *SimpleChaincode) attach_license(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	carAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New("Failed to get car profile")
	}
	res := Car{}
	json.Unmarshal(carAsBytes, &res)
	fmt.Println(res)
	
	res.Data.License = args[1]
	jsonAsBytes, _ := json.Marshal(res)
	// Write the state back to the ledger
	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Get Permissions - get permissions form the car profile  
// ============================================================================================================================
func (t *SimpleChaincode) get_permissions(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	carAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New("Failed to get car profile")
	}
	res := Car{}
	json.Unmarshal(carAsBytes, &res)
	fmt.Println(res)
	
	jsonAsBytes, _ := json.Marshal(res.Users)
	return jsonAsBytes, nil
}

// ============================================================================================================================
// Set User Permissions 
// ============================================================================================================================
func (t *SimpleChaincode) set_user_perms(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	carAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New("Failed to get car profile")
	}
	res := Car{}
	json.Unmarshal(carAsBytes, &res)
	fmt.Println(res)
	
	for i,perm := range res.Users{
		if perm.UserId == args[1]{							//find the correct user
			res.Users[i].Permissions[0] = args[2]			//set new perm, dsh - to do make this input as array of all perms
			fmt.Println(res.Users[i].Permissions)
			break
		}
	}
	
	// Write the state back to the ledger
	jsonAsBytes, _ := json.Marshal(res)
	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}