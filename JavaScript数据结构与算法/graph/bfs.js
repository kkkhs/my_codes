const graph = require('./graph')

const visited = new Set()
const bfs = (n) => {
  visited.add(n)  // 手动添加队头visited
  const q = [n]
  while(q.length){
    const t = q.shift()
    console.log(t)
    // visited.add(t)   // 可能会导致队列元素重复(因为是子元素入队)
    graph[t].forEach(c => {
      if(!visited.has(c)){
        q.push(c) 
        visited.add(c)
      }
    })
  }
}

bfs(2)