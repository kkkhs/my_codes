/*
1、同一目录下go文件的包名称保持一致，
2、包名称与目录名无关，
3、同一个package内部可以自由随意调用
4、
*/
package maths

//大写字母开头的函数在其他包内也可见
func Add(a, b, c int) int {
	return a + sub(b, c)
}
