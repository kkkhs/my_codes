// @ts-ignore
const bt = require('./bt')

// 递归版
const inorder = (root) => {
  if(!root){
    return
  }
  inorder(root.left)
  console.log(root.val)
  inorder(root.right)
} 

// 非递归版
const inorder1 = (root) => {
  if(!root) return

  const stack = []
  let p = root
  while(stack.length || p){
    while(p){
      stack.push(p)
      p = p.left
    }
    const n = stack.pop()
    console.log(n.val)
    p = n.right
  }
}

inorder(bt)
inorder1(bt)