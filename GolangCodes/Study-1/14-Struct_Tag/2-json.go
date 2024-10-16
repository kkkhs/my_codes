package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string   `json:"title"`
	Price int      `json:"price"`
	Year  int      `json:"year"`
	Actor []string `json:"actor"`
}

func main() {
	movie := Movie{"喜剧之王", 10, 2000, []string{"星爷", "张柏芝"}}

	//编码的过程： 结构体 ---> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程： jsonStr ---> 结构体
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}

	fmt.Println(my_movie)
}
