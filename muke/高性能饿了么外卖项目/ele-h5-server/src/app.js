const path = require('path')
const jsonServer = require('json-server')
const router = require('./router')
const db = require('./db')()

const server = jsonServer.create()

const middlewares = jsonServer.defaults({
  static: path.join(__dirname, '../public')
})
server.use(middlewares)
//req.body
server.use(jsonServer.bodyParser)

router(server)
// 使用路由配置
const jsonRouter = jsonServer.router(db)
server.use('/api', jsonRouter)

server.listen(8000, () => {
  console.log("json Server is running at port:8000")
})