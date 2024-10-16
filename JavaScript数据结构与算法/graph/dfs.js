const graph = require('./graph')

const visited = new Set() // 访问标志
const dfs = (n) => {
  console.log(n)
  visited.add(n)
  graph[n].forEach(c => {
    if(!visited.has(c)){
      dfs(c)
    }
  })
}

dfs(2)