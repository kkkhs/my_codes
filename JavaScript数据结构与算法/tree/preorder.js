// @ts-ignore
const bt = require('./bt')

// 递归版
const preorder = (root) => {
  if(!root){
    return
  }
  console.log(root.val)
  preorder(root.left)
  preorder(root.right)
}

// 非递归版
const preorder1 = (root) => {
  if(!root){
    return
  }
  const stack = [root]
  while(stack.length){
    const n = stack.pop()
    console.log(n.val)
    if(n.right) stack.push(n.right) // 先进后出
    if(n.left) stack.push(n.left)
  }
}

preorder(bt)
preorder1(bt)