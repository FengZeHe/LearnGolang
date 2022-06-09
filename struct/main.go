package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Num  int
}

type Teacher struct {
	Name     string `json:Name`
	Subjects string `json:Subjects`
}

func main() {
	var person1 Student
	person1.Name = "Feng"
	person1.Num = 1
	fmt.Println(person1.Name, person1.Num)

	person2 := Teacher{"teacher1", "math"}
	m, err := json.Marshal(person2)
	if err == nil {
		fmt.Println(string(m))
	}

	empJsonData := `{"Name":"Xu","Subjects":"Math"}`
	empBytes := []byte(empJsonData)
	var person3 Teacher
	json.Unmarshal(empBytes, &person3)
	fmt.Println(person3.Name)
	fmt.Println(person3.Subjects)

}
