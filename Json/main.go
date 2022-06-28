package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func encode() {
	//use json.Marshal编码
	person1 := Person{"嘟嘟", 3}
	byte1, err := json.Marshal(person1)
	if err != nil {
		log.Println("encode failed,err:", err)
	}
	fmt.Println(string(byte1))
}

func decode() {
	//使用json.Unmarshal
	str := `{"name":"嘟嘟","age":3}`
	var person2 Person
	err := json.Unmarshal([]byte(str), &person2)
	if err != nil {
		log.Println("decode failed,err:", err)
	}
	fmt.Println(person2.Name, person2.Age)

}

func main() {
	encode()
	decode()
}
