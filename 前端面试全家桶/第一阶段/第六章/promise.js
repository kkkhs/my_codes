/**
 * 手写Promise
 * 
 * 功能:
 * 1、初始化 & 异步调用
 * 2、then catch链式调用
 * 3、 API .resolve .reject .all .race
 */
 
class MyPromise{
  state = 'pending'  // 3种状态: 'pending'  'fulfilled' 'rejected'
  value = undefined  //成功后的值
  reason = undefined  //失败后的原因

  resolveCallbacks = [] // pending状态下， 存储成功的回调
  rejectCallbacks = []  // pending状态下， 存储失败的回调
  
  constructor(fn){
    const resolveHandler = (value) => {
      if(this.state === 'pending'){
        this.state = 'fulfilled'
        this.value = value
        this.resolveCallbacks.forEach(fn => fn(this.value ))
      }
    }
    
    const rejectHandler = (reason) => {
      if(this.state === 'pending'){
        this.state = 'rejected'
        this.reason = reason
        this.rejectCallbacks.forEach(fn => fn(this.reason))
      }
    }
  
    try{
      fn(resolveHandler, rejectHandler)
    }catch (err){
      rejectHandler(err)
    }
  }

  then(fn1, fn2){
     fn1 = typeof fn1 === 'function' ? fn1 : (v) => v
     fn2 = typeof fn2 === 'function' ? fn2 : (e) => e

     if(this.state === 'pending'){
      return new MyPromise((resolve, reject) => {
        this.resolveCallbacks.push(() => {
          try{
            const newValue = fn1(this.value)
            resolve(newValue)
          }catch(err){
            reject(err)
          }                         
        })      

        this.rejectCallbacks.push(() => {
          try{
            const newError = fn2(this.reason)
            reject(newError)
          }catch(err){
            reject(err)
          }
        })
      })
     }

     if(this.state === 'fulfilled'){
      return new MyPromise((resolve, reject) => {
        try{
          const v = fn1(this.value)
          resolve(v)
        }catch (err){
          reject(err)
        }
      })
     }

     if(this.state === 'rejected'){
      return new MyPromise((resolve, reject) => {
        try{
          const NewReason = fn2(this.reason)
          reject(NewReason)
        }catch (err){
          reject(err)
        }
      })
     }

  }

  // 就是 then 的一个语法糖，简单模式
  catch(fn){
    return this.then(null, fn)
  }
}


// API实现
MyPromise.resolve = function(value){
  return new MyPromise((resolve, reject) => resolve(value))
}

MyPromise.reject = function(reason){
  return new MyPromise((resolve, reject) => reject(reason))
}

// all 等待全部promise完成后一起返回所有结果
MyPromise.all = function (promiseList = []) {
  return new MyPromise((resolve, reject) => {
    const result = []
    const length = promiseList.length
    let resolvedCount = 0

    promiseList.forEach(p => { 
      p.then(data => {
        result.push(data)

        // resolvedCount 必须在 then 里面做 ++
        // 不能用index
        resolvedCount ++
        if(resolvedCount === length){
          // 已经遍历到了最后一个 promise
          resolve(result)
        }
      }).catch(err => {
        reject(err)
      }) 
    })
  })
}

// race只返回第一个fulfilled的Promise结果
MyPromise.race = function (promiseList = []) {
  return new MyPromise((resolve, reject) => {
    let resolved = false //标记

    promiseList.forEach(p => { 
      p.then(data => {
        if(resolved === false){
          resolve(data)
          resolved = true
        }
      }).catch(err => {
        reject(err)
      }) 
    })
  })
}
