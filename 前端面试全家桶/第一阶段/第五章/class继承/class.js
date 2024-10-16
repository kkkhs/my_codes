// 父类
class People {
  constructor(name){
    this.name = name
  }

  eat(){
    console.log(`${this.name} is eating`)
  }
}


//子类
class Student extends People{
  constructor(name, number){
    super(name)
    this.number = number
    // this.gender = male
  }
  sayHi(){
    console.log(
      `姓名 ${ this.name } , 学号 ${this.number}.`
    )
  }
}

class Teacher extends People{
  constructor(name, major){
    super(name)
    this.major = major
  }
  teach(){
    console.log(`${this.name} is teaching ${this.major}`)
  }
}

//通过类声明对象/实例
const xialuo = new Student('xialuo', 1)
xialuo.sayHi()
xialuo.eat()

const wang = new Teacher('wang', 'Chinese')
wang.eat()
wang.teach()
