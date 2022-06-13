package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	person1 := Studnet{Name:"feng",Id:2,Grade:1}
	fmt.Println(person1)
	person1ToJson ,_:= json.Marshal(person1)
	fmt.Println(string(person1ToJson))

	JsonTemp := `{"name":"xu","id":3,"grade":4}`
	person := Studnet{}
	err := json.Unmarshal( []byte(JsonTemp),&person)
	if err == nil{
		fmt.Println(person,person.Name)
	}

}

type Studnet struct{
	Name string `json:"name"`
	Id int `json:"id"`
	Grade int `json:"grade"`
}