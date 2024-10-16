// @ts-ignore
const bt = require('./bt')

// 递归版
const postorder = (root) => {
  if(!root){
    return
  }
  postorder(root.left)
  postorder(root.right)
  console.log(root.val)
}

// 非递归版
const postorder1 = (root) => {
  if(!root){
    return
  }
  const outputStack = []
  const stack = [root]
  while(stack.length){
    const n = stack.pop()
    outputStack.push(n)     // 根 右 左(类似先序遍历，然后倒序输出结果)
    if(n.left) stack.push(n.left)
    if(n.right) stack.push(n.right)
  }

  while(outputStack.length){
    console.log(outputStack.pop().val)
  }
}

postorder(bt)
postorder1(bt)