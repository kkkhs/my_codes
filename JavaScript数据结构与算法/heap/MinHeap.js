class MinHeap{
  constructor(){
    this.heap = []
  }
  // 交换节点
  swap(i, j){
    const temp = this.heap[i]
    this.heap[i] = this.heap[j]
    this.heap[j] = temp
  }
  // 获取父节点
  getParentIndex(i) {
    // return Math.floor((i - 1) / 2)   // 拿到商
    return (i - 1) >> 1
  }
  // 获取左子节点
  getLeftIndex(i) {
    return i * 2 + 1
  }
  // 获取右子节点
  getRightIndex(i) {
    return i * 2 + 2
  }
   // 上移操作
  shiftUp(index){
    if(index === 0) return
    const parentIndex = this.getParentIndex(index)
    if(this.heap[parentIndex] > this.heap[index]){
      this.swap(parentIndex, index)
      this.shiftUp(parentIndex)
    }
  }
  // 下移操作
  shiftDown(index){
    const leftIndex = this.getLeftIndex(index)
    const rightIndex = this.getRightIndex(index)
    if(this.heap[leftIndex] < this.heap[index]){
      this.swap(leftIndex, index)
      this.shiftDown(leftIndex)
    }
    if(this.heap[rightIndex] < this.heap[index]){
      this.swap(rightIndex, index)
      this.shiftDown(rightIndex)
    }
  }
  // 1、插入元素
  insert(value){
    this.heap.push(value)
    this.shiftUp(this.heap.length - 1)
  }
  // 2、删除堆顶
  pop(){
     this.heap[0] = this.heap.pop()
     this.shiftDown(0)
  }
  // 3、获取堆顶
  peek(){
    return this.heap[0]
  }
  // 4、获取堆的大小
  size(){
    return this.heap.length
  }
}

const h = new MinHeap()
h.insert(3)
h.insert(2)
h.insert(1)
h.pop()
h.peek()
h.size()

var mergeKLists = function(lists) {
  
};