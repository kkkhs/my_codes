package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secrectNumber := rand.Intn(maxNum)
	//fmt.Println("The secrect number is", secrectNumber)

	fmt.Println("Please input your guess:")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n') //读取一行输入
		if err != nil {
			fmt.Println("An err occured while redding input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n") //去掉换行符

		guess, err := strconv.Atoi(input) //转换为数字
		if err != nil {
			fmt.Println("Please try again")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secrectNumber {
			fmt.Println("bigger !")
		} else if guess < secrectNumber {
			fmt.Println("smaller !")
		} else {
			fmt.Println("correct !")
			break
		}
	}
}
