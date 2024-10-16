interface IListNode {
  value: any
  key: string
  prev?: IListNode
  next?: IListNode
}

class LRUCache {
  private length: number
  private data: { [key: string]: IListNode } = {}
  private dataLength: number = 0
  private listHead: IListNode | null = null
  private listTail: IListNode | null = null

  constructor(length: number) {
    if (length < 1) throw new Error('invalid length')
    this.length = length
  }

  // 移动到末尾(最新)
  private moveToTail(curNode: IListNode) {
    const tail = this.listTail
    if (tail === curNode) return

    // 1、让 preNode 和 nextNode 建立关系
    const preNode = curNode.prev
    const nextNode = curNode.next
    if (preNode) {
      if (nextNode) {
        preNode.next = nextNode
      } else {
        delete preNode.next
      }
    }
    if (nextNode) {
      if (preNode) {
        nextNode.prev = preNode
      } else {
        delete nextNode.prev
      }

      // 头指针更新
      if (this.listHead === curNode) this.listHead = nextNode
    }
    // 2、让 curNode 断绝与 preNode 和 nextNode 的关系
    delete curNode.prev
    delete curNode.next
    // 3、在 list 末尾建立与 curNode 的新关系
    if (tail) {
      // 先后再前
      tail.next = curNode
      curNode.prev = tail
    }
    this.listTail = curNode
  }

  private tryClean() {
    while (this.dataLength > this.length) {
      const head = this.listHead
      if (head == null) throw new Error('head is null')
      const headNext = head.next
      if (headNext == null) throw new Error('headNext is null')

      // 1、断绝 head 和 next 的关系
      delete headNext?.prev
      delete head.next

      // 2、重新赋值 listHead
      this.listHead = headNext

      // 3、清理data, 减小长度
      delete this.data[head.key]
      this.dataLength--
    }
  }

  get(key: string): any {
    const data = this.data
    const curNode = data[key]

    if (curNode == null) return null

    if (this.listTail === curNode) {
      // 在末尾直接返回
      return curNode.value
    }

    // 不在末尾则需要移动到末尾
    this.moveToTail(curNode)

    return curNode.value
  }

  set(key: string, value: any) {
    const data = this.data
    const curNode = data[key]

    if (curNode == null) {
      // 新增数据
      const newNode: IListNode = { key, value }
      // 移动到末尾
      this.moveToTail(newNode)

      data[key] = newNode
      this.dataLength++

      // 首个元素的处理
      if (this.dataLength === 1) this.listHead = newNode
    } else {
      // 修改现有数据
      curNode.value = value
      this.moveToTail(curNode)
    }

    // 尝试清理长度
    this.tryClean()
  }
}
