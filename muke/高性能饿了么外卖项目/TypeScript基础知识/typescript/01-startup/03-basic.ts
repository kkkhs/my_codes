/**
 * 基础概念
 * 
 * 基础类型：字符串、数字、布尔
 * 复杂类型：数组、对象
 * 
 * 对象类型：匿名、接口interface、类型别名type
 *    函数function
 *    类class
 *    枚举enum
 * 
 * 类型推断 type interface
 * 类型断言
 * 泛型 generic: 对象类型、函数、类
 * 
 */


/**
 * 概览
 */

// 字符串
const str: string = 'hello world'
// 数字
const num: number = 123
// 布尔
const bool: boolean = true
// 数组
const arr: number[] = [123, 123]
arr[0] = 'a'
// 对象
const obj: { x: number } = { x: 1, y: 1 }
// 可选字段
const optionalObj: { x: number, y?: number } = { x: 1 }
optionalObj.y = 'x'
optionalObj.y = 123
optionalObj.z = 123

// 对象类型
// 接口
interface iType {
  str: string
  name: string
  fc: () => string
}
const add: iType = {
  str: 'khh',
  name: 'khs',
  fc: (): string => { return 'abc' }
}
// 类型别名
type TType = {
  str: string
}
const str2: TType = { str: '123' }
str2.str

// 类型判断
let n: unknown = 1

// 类型断言: 格式
// <类型>值
// 或:
// 值 as 类型
console.log((<string>n).toLowerCase())
console.log((n as string).toLowerCase())

// 泛型, 包括泛型对象类型、泛型函数、泛型类
interface IGeneric<T> {
  test: T
}

// 特殊类型: void、any、unknown、never

// 1、any： 任意值是 TypeScript 针对编程时类型不明确的变量使用的一种数据类型
// 2、null
//  在 JavaScript 中 null 表示 "什么都没有"。
//  null是一个只有一个值的特殊类型。表示一个空对象引用。
//  用 typeof 检测 null 返回是 object。
// 3、undefined
//  在 JavaScript 中, undefined 是一个没有设置值的变量。
//  typeof 一个没有值的变量会返回 undefined。
//  Null 和 Undefined 是其他任何类型（包括 void）的子类型，可以赋值给其它类型，如数字类型，此时，赋值后的类型会变成 null 或 undefined。而在TypeScript中启用严格的空校验（--strictNullChecks）特性，就可以使得null 和 undefined 只能被赋值给 void 或本身对应的类型，

// never 是其它类型（包括 null 和 undefined）的子类型，代表从不会出现的值。这意味着声明为 never 类型的变量只能被 never 类型所赋值，在函数中它通常表现为抛出异常或无法执行到终止点（例如无限循环）


