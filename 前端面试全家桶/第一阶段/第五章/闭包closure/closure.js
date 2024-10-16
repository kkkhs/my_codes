// // 1、函数作为返回值
// function create(){
//   let a = 100
//   return function(){
//     console.log(a)
//   }
// }

// const fn1 = create()
// const a = 200
// fn1()  // 100

// // 2、函数作为参数传递
// function print(fn2){
//   const b = 200
//   fn2()
// }

// const b = 100
// function fn2(){
//   console.log(b)
// }
// print(fn2)

// // 闭包: 自由变量的查找，是在函数**定义**的地方，向上级作用域查找
// // 不是在执行的地方！！！