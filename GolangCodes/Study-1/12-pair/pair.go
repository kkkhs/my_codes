package main

import "fmt"

func main1() {
	var a string
	//pair<statictype:string, value:"aceld">
	a = "aceld"

	//pair<type:string, value:"aceld">
	var alltype interface{}
	alltype = a

	str, _ := alltype.(string)
	fmt.Println(str)
}
