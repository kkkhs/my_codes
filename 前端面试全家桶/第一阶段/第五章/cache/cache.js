// 闭包隐藏数据，只提供API
function createCache() {
  const data = {} //闭包中的数据，被影藏，不被外界访问
  return{
    set: function (key, val){
      data[key] = val
    },
    get: function(key){
      return data[key]
    }
   }
}

const c = createCache()
c.set('a', 100)
console.log(c.get('a'))

// 手写用Promise加载图片
function loadImg(src){
  return new Promise(
    (resolve, reject) => {
      const img = document.createElement('img')
      img.onload = () => {
        resolve(img)  // then中接收resolve的引用对象
      }
      img.onerror = () => {
        reject(new Error(`图片加载失败 ${src}`))  // catch中接收reject的引用对象
      }
      img.src = src
    }
  )
}

const url1= 'https://www.baidu.com/favicon.ico'
loadImg(url1).then(img => {
  console.log(img.width)
  return img  //传到第二个then
}).then(img => {
  console.log(img.height)
}).catch(ex => {console.error(ex)})

const url2= 'https://www.baidu.com/favicon.ico'
loadImg(url1).then(img => {
  console.log(img.width)
  return img  // 1、返回普通对象
}).then(img => {
  console.log(img.height)
  return loadImg(url2) // 2、返回promise实例
}).then(img2 => {
  console.log(img2.width)
})