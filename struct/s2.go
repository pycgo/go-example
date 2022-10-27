package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "标题",
		Year:   2000,
		Price:  20,
		Actors: []string{"zhangsan", "lisi"},
	}
	jsonstr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(string(jsonstr))
		//可以使用Printf 格式化输出，而不需要上面的转类型
		fmt.Printf("%s", jsonstr)
	}
}
