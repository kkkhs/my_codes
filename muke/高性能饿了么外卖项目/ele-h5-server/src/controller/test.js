const test = require('../../data/test')

module.exports = (req, res) => {
  const testData = test()
  testData.desc = '我是自定义测试数据'
  res.json(testData)
}