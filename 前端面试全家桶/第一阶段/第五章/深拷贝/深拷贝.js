// 手写深拷贝

const obj1 = {
  age: 20,
  name: 'xxx',
  address: {
    city: 'beijing'
  },
  arr: ['a', 'b', 'c']
}

const obj2 = deepClone(obj1)

obj2.address.city = 'shanghai'
console.log(obj1.address.city)


/**
 * 深拷贝
 * @param {object} obj 要拷贝的对象
 */
function deepClone(obj = {}){
  if(typeof obj !== 'object' || obj == null){
    return obj
  }

  //初始化返回结果
  let result
  if(obj instanceof Array){
    result = []
  }else{
    result = {}
  }

  for(let key in obj){
    // 保证 key 不是原型的属性
    if(obj.hasOwnProperty(key)){
      //递归调用!!!
      result[key] = deepClone(obj[key])
    }
  }

  // 返回结果
  return result
}