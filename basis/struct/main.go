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

type dog struct {
}

type fakedog struct {
}

type truedog struct {
}

func (f fakedog) say() {
	fmt.Println("fake dog")
}

func (f truedog) say() {
	fmt.Println("true dog")
}

func (t Teacher) say() {
	fmt.Println("I am a teacher")
}

func (s Student) say() {
	fmt.Println("I am a student")
}

// 结构体接收器
func (person Student) changeValueByStruct(newName string) {
	person.Name = newName
}

// 指针接收器
func (person *Teacher) changeValueByPointer(newname string) {
	person.Name = newname
}

func main() {
	var person1 Student
	person1.Name = "Feng"
	person1.Num = 1
	fmt.Println(person1.Name, person1.Num)

	// Marshal＆Unmarshal
	person2 := Teacher{"teacher1", "math"}
	m, err := json.Marshal(person2)
	if err == nil {
		fmt.Println(string(m))
	}

	//改名字
	person1.changeValueByStruct("hello")
	p2 := &person2
	p2.changeValueByPointer("world")
	//只有指针才能改名字

	fmt.Println(person1.Name, person2.Name)

	empJsonData := `{"Name":"Xu","Subjects":"Math"}`
	empBytes := []byte(empJsonData)
	var person3 Teacher
	json.Unmarshal(empBytes, &person3)
	fmt.Println(person3.Name)
	fmt.Println(person3.Subjects)

	//	Dogs
	f1 := fakedog{}
	f1.say()

	f2 := truedog(f1)
	f2.say()
}
