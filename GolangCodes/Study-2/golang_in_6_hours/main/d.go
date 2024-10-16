package main

//init函数在import该包时就会调用
import (
	"fmt"
	"g6/util"
	maths "g6/util/math" //引用时引用文件名(前面可加包的别名)
	_ "os"               //前加_ 调包但不使用

	"github.com/bytedance/sonic" //使用go get url或go mod tidy命令初始化第三方包
)

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(util.Name)
	fmt.Println(util.Add(a, b))
	fmt.Println(maths.Add(a, b, c)) //使用时使用包名
	//fmt.Println(maths.sub(a, b))
	bytes, _ := sonic.Marshal("hello")
	fmt.Println(string(bytes))

}
