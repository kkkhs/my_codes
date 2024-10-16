function fn1(a, b){
  console.log('this:', this)
  console.log(a, b)
  return 'this is fn1'
}

const fn2 = fn1.bind({x:100}, 10, 20)
const res = fn2()
console.log(res)

/**
 * 手写bind
 */
Function.prototype.bind1 = function(){
  // 将参数列表拆解为数组
  const args = Array.prototype.slice.call(arguments)

  // 获取 bind 的 this (数组第一项)
  const t = args.shift()

  // 获取 fn1.bind(...) 中的 fn1
  const self = this

  //返回一个函数
  return function () {
    return self.apply(t, args)
  }
}

const fn3 = fn1.bind1({x:100}, 11, 22)
const res2 = fn3()
console.log(res2)