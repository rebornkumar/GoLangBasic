package main

import (
	"fmt"
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

func main() {
	fmt.Println("Hello, playground")
	type Message struct {
    		Name string `json:"name" validate:"required"`
    		Body string `json:"body" validate:"required"`
    		Time int64  `json:"time"`
	}
	myMessage := Message {
	Name : "Anuj",
	Body : "Do Something",
	}
	fmt.Println(myMessage)
		jsonString, err := json.Marshal(myMessage)
	if err != nil {
		return
	}
	
	fmt.Println(string(jsonString))
	var validate *validator.Validate = validator.New()
	err = validate.Struct(myMessage)
	
	if err != nil {
		fmt.Println("struct validation failed")
		return
	}
	//Arbitrary JSON objects
	arbitraryJSON := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez",1234]}`)
	
	var arbitraryStruct interface{}
	err = json.Unmarshal(arbitraryJSON,&arbitraryStruct)
	if err != nil {
		fmt.Println("Error while Unmarshaling arbitraryJSON")
		return
	}
	arbitraryStructMap := arbitraryStruct.(map[string]interface{})
	fmt.Println(arbitraryStructMap["Parents"].([]interface{})[1].(float64))
	
	
}
