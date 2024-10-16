/**
 *   顺序搜索
 */

Array.prototype.sequentialSearch = function (item) {
  for(let i = 0; i < this.length; i ++){
    if(this[i] === item){
      return i
    }
  }
  return -1
} 