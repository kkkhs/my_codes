// 类可以实现接口，使用关键字 implements，并将 interest 字段作为类的属性使用。

interface ILoan {
  interest: number
}

class AgriLoan implements ILoan {
  interest: number
  rebate: number

  constructor(interest: number, rebate: number) {
    this.interest = interest
    this.rebate = rebate
  }
}

var obj = new AgriLoan(10, 1)
console.log("利润为 : " + obj.interest + "，抽成为 : " + obj.rebate)