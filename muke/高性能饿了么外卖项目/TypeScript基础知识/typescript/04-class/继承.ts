// TypeScript 一次只能继承一个类，不支持继承多个类,但TypeScript 支持多重继承（A 继承 B，B 继承 C）
// class child_class_name extends parent_class_name

class Shape {
  Area: number

  constructor(a: number) {
    this.Area = a
  }
}

class Circle extends Shape {
  disp(): void {
    console.log("圆的面积:  " + this.Area)
  }
}

var obj3 = new Circle(223);
obj3.disp()


//继承类的方法重写
// super 关键字是对父类的直接引用，该关键字可以引用父类的属性和方法
class PrinterClass {
  doPrint(): void {
    console.log("父类的 doPrint() 方法。")
  }
}

class StringPrinter extends PrinterClass {
  doPrint(): void {
    super.doPrint() // 调用父类的函数
    console.log("子类的 doPrint()方法。")
  }
}