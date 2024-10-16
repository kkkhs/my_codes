// 字段 − 字段是类里面声明的变量。字段表示对象的有关数据。
// 构造函数 − 类实例化时调用，可以为类的对象分配内存。
// 方法 − 方法为对象要执行的操作。
class Car {
  // 字段 
  engine: string;

  // 构造函数 : 类实例化时会调用构造函数
  constructor(engine: string) {
    this.engine = engine
  }

  // 方法 
  disp(): void {
    console.log("发动机为 :   " + this.engine)
  }
}

// 创建实例化对象
// 我们使用 new 关键字来实例化类的对象，语法格式如下：
// var object_name = new class_name([ arguments ])

// 创建一个对象
var obj1 = new Car("XXSY1")

// 访问字段
console.log("读取发动机型号 :  " + obj1.engine)

// 访问方法
obj1.disp()