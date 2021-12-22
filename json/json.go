package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MyData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyArrayData struct {
	MyArrays []MyData `json:"results"`
}

func main() {

	//First
	//getJsonDatas()
	//Second
	//getJsonDatasOfArray()
	//Third
	fetchJsonDatasOfArray()

}

func getJsonDatas() {
	myJson := string(`{"id": 12, "name": "sefa ün", "email": "sefa.un@hotmail.com", "password": "ıruhşfdldfdfkl"}`)
	myDataStruct := MyData{}
	err := json.Unmarshal([]byte(myJson), &myDataStruct)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(myDataStruct.ID, myDataStruct.Name, myDataStruct.Email, myDataStruct.Password)
}

func getJsonDatasOfArray() {
	arrayJson := `[{"id": 12, "name": "sefa ün", "email": "hatice.un@hotmail.com", "password": "ıruhşfdldfdfkl"},
	{"id": 12, "name": "sefa ün", "email": "sefa.un@hotmail.com", "password": "ıruhşfdldfdfkl"},
	{"id": 12, "name": "sefa ün", "email": "erdem.un@hotmail.com", "password": "ıruhşfdldfdfkl"},
	{"id": 12, "name": "sefa ün", "email": "bekir.un@hotmail.com", "password": "ıruhşfdldfdfkl"}]`
	var myArrayDataStruct MyArrayData

	if myArrayErr := json.Unmarshal([]byte(arrayJson), &myArrayDataStruct.MyArrays); myArrayErr != nil {
		fmt.Println(myArrayErr.Error())
	}

	fmt.Println(len(myArrayDataStruct.MyArrays))
	fmt.Println(myArrayDataStruct.MyArrays[0].ID)
}

func fetchJsonDatasOfArray() {

	arrayJson := `[{"id": 12, "name": "sefa ün", "email": "hatice.un@hotmail.com", "password": "ıruhşfdldfdfkl"},
	{"id": 12, "name": "sefa ün", "email": "sefa.un@hotmail.com", "password": "ıruhşfdldfdfkl"},
	{"id": 12, "name": "sefa ün", "email": "erdem.un@hotmail.com", "password": "ıruhşfdldfdfkl"},
	{"id": 12, "name": "sefa ün", "email": "bekir.un@hotmail.com", "password": "ıruhşfdldfdfkl"}]`
	var myArrayDataStruct MyArrayData

	if myArrayErr := json.Unmarshal([]byte(arrayJson), &myArrayDataStruct.MyArrays); myArrayErr != nil {
		fmt.Println(myArrayErr.Error())
	}

	for i, v := range myArrayDataStruct.MyArrays {
		//fmt.Println(i, v.Email)
		fmt.Println(i, strings.Index(v.Email, "efa"))
	}

	for i := 0; i < len(myArrayDataStruct.MyArrays); i++ {

		//String Contains
		fmt.Println(strings.Contains(myArrayDataStruct.MyArrays[i].Name, "sefa"))

	}

}
