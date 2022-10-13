package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string `json:"name"` //标记json名字为name
	Age  int    `json:"age"`
	Time int64  `json:"-"` // 标记忽略该字段

}

func main() {
	person := Person{"ss", 18, time.Now().Unix()}
	//person := Person{Name: "xiaoming", Age: 18, Time: time.Now().Unix()}
	if result, err := json.Marshal(&person); err == nil {
		fmt.Println(string(result))
	}
}
