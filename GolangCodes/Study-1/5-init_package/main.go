package main

//_ "khs/Study/5-init_package/lib1" // 1、_ 匿名导包(可以不使用该包)
// mylib2 "khs/Study/5-init_package/lib2" //2、别名导包(可以通过别名使用该包)
//. "khs/Study/5-init_package/lib2" // 3、. 可以直接使用该包

func main() {
	//lib1.Lib1test()   //1、
	//mylib2.Lib2test()  //2、
	//Lib2test() //3、
}
