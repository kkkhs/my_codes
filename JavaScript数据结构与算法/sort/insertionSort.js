/**
 *  插入排序
 *  O(n2)
 * 1 < i < length
 * 0 < j < i
 */

Array.prototype.insertionSort = function () {
  for (let i = 1; i < this.length; i++) {
    const temp = this[i]
    let j = i
    while (j > 0) {
      if (this[j - 1] > temp) {
        this[j] = this[j - 1]
      } else {
        break
      }
      j--
    }
    this[j] = temp
  }
}
