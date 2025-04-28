// 布尔值 -------------
var isDone = false;
// new Boolean() 返回的是一个 Boolean 对象
// let createdByNewBoolean: boolean = new Boolean(1);
// 不能将类型“Boolean”分配给类型“boolean”。 “boolean”是基元，但“Boolean”是包装器对象。如可能首选使用“boolean”。
var createdByNewBoolean = new Boolean(1);
// 直接调用 Boolean 也可以返回一个 boolean 类型：
var createdByBoolean = Boolean(1);
// 数值 -------------
var decLiteral = 6;
var hexLiteral = 0xf00d;
// ES6 中的二进制表示法
var binaryLiteral = 10;
// ES6 中的八进制表示法
var octalLiteral = 484;
var notANumber = NaN;
var infinityNumber = Infinity;
// 字符串 -------------
var myName = 'Tom';
var myAge = 25;
// 模板字符串
var sentence = "Hello, my name is ".concat(myName, ".\nI'll be ").concat(myAge + 1, " years old next month.");
// 空值 -------------
function alertName() {
    alert('My name is Tom');
}
var unusable = undefined;
// Null 和 Undefined -------------
var u = undefined;
var n = null;
// undefined 和 null 是所有类型的子类型
var num = undefined;
var num1 = null;
var u1;
var num2 = u;
// void 类型的变量不能赋值给 number 类型的变量
var u2;
var num3 = u;
// Type 'void' is not assignable to type 'number'.
