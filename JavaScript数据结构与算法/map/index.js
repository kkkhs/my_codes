const m = new Map()

// 字典内部不重复

// 增
m.set('a', 'aa')
m.set('b', 'bb')

//删
m.delete('b')
// m.clear()  //全部删除

//改
m.set('a', 'aaa')

//查
m.has('a')

//取
m.get('a')

//长度
m.size

// 补充知识string
const s = '123'
let l = 1, r = 2
s.substring(l, r + 1)