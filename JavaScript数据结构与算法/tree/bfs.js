const tree1 = {
  val: 'a',
  children: [
    {
      val: 'b',
      children: [
        {
          val: 'd',
          children: [],
        },
        {
          val: 'e',
          children: [],
        },
      ],
    },
    {
      val: 'c',
      children: [
        {
          val: 'f',
          children: [],
        },
        {
          val: 'g',
          children: [],
        },
      ],
    },
  ],
}

const bfs = (root) => {
  const queue = [root]

  while (queue.length > 0) {
    const t = queue.shift()
    console.log(t.val)
    t.children.forEach((child) => queue.push(child))
  }
}

bfs(tree1)

