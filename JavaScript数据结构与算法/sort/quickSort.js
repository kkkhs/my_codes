/**
 *  快 排
 *  O(nlogn) 到	O(n2)
 */

Array.prototype.quickSort = function () {
  const rec = (arr) => {
    if (arr.length === 1) {
      return arr
    }

    const left = []
    const right = []
    const mid = arr[0]
    for (let i = 1; i < arr.length; i++) {
      if (arr[i] < mid) {
        left.push(arr[i])
      } else {
        right.push(arr[i])
      }
    }
    return [...rec(left), mid, ...rec(right)] // 返回合并后的数组
  }

  const res = rec(this)
  res.forEach((n, i) => {
    // 拷贝结果到this上(改变原数组)
    this[i] = n
  })
}
