/**
 * O(n2)
 *  i < length - 1
 *  j < length - 1 - i
 */

Array.prototype.bubleSort = function () {
  for (let i = 0; i < this.length - 1; i++) {
    for (let j = 0; j < this.length - 1 - i; j++) {
      if (this[j] > this[j + 1]) {
        const temp = this[j]
        this[j] = this[j + 1]
        this[j + 1] = temp
      }
    }
  }
}

var reverseKGroup = function (head, k) {
  const myReverse = (l, r) => {
    let prev = r.next
    let p = l
    while (prev !== r) {
      const next = p.next
      p.next = prev
      prev = p
      p = next
    }
    return [r, l]
  }

  const dummpy = new ListNode(0, head)
  let pre = dummpy

  while (head) {
    let tail = pre
    // 查看剩余部分长度是否大于等于 k
    for (let i = 0; i < k; i++) {
      tail = tail.next
      if (tail === null) {
        return dummpy.next
      }
    }

    ;[head, tail] = myReverse(head, tail)
    // 把子链表重新接回原链表
    pre.next = head

    // 迭代
    pre = tail
    head = pre.next
  }
  return dummpy.next
}

function ListNode(val, next) {
  this.val = val
  this.next = next
}

const node5 = new ListNode(5, null)
const node4 = new ListNode(4, node5)
const node3 = new ListNode(3, node4)
const node2 = new ListNode(2, node3)
const node1 = new ListNode(1, node2)

console.log(reverseKGroup(node1, 2))
