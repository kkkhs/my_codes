class jQuery{
  constructor(selector){
    const result = document.querySelectorAll(selector)  // 选择器选择的所有dom元素
    const length = result.length
    for(let i = 0; i < length; i ++){
      this[i] = result[i]
    }
    this.length = length
    this.selector = selector
    //类数组
  }
  get(index){
    return this[index]
  }
  each(fn){
    for(let i = 0; i < this.length; i ++){
      const elem = this[i]
      fn(elem)
    }
  }
  on(type, fn){
    return this.each(elem => {
      elem.addEventListener(type, fn, false)
    })
  } 

  // 拓展很多方法...
}

// 1、插件机制
jQuery.prototype.dialog = function (info){
  alert(info)
}

// 2、'造轮子'
class myJQuery extends jQuery{
  constructor(selector){
    super(selector)
  }

  //拓展自己的方法...
  addClass(className){}
  style(data){}
}


// const $p = new jQuery('p')
// $p.get(1)
// $p.each((elem) => console.log(elem.nodeName))
// $p.on('click', () => alert('cliked'))
