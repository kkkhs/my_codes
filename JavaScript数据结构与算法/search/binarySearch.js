/**
 *   二分搜索
 */

Array.prototype.binarySearch = function (item) {
  let low = 0
  let high = this.length - 1
  while (low <= high) {
    const mid = Math.floor((low + high) / 2)
    const element = this[mid]
    if (element < item) {
      low = mid + 1
    } else if (element > item) {
      high = mid - 1
    } else {
      return mid
    }
  }
  return -1
}

const obj = {
  sayThis: () => {
    console.log(this)
  },
}

obj.sayThis() // window 因为 JavaScript 没有块作用域，所以在定义 sayThis 的时候，里面的 this 就绑到 window 上去了
const globalSay = obj.sayThis
globalSay() // window 浏览器中的 global 对象
