// 布尔值 -------------

let isDone: boolean = false;

// new Boolean() 返回的是一个 Boolean 对象
// let createdByNewBoolean: boolean = new Boolean(1);
// 不能将类型“Boolean”分配给类型“boolean”。 “boolean”是基元，但“Boolean”是包装器对象。如可能首选使用“boolean”。
let createdByNewBoolean: Boolean = new Boolean(1);

// 直接调用 Boolean 也可以返回一个 boolean 类型：
let createdByBoolean: boolean = Boolean(1);


// 数值 -------------
let decLiteral: number = 6;

let hexLiteral: number = 0xf00d;

// ES6 中的二进制表示法
let binaryLiteral: number = 0b1010;

// ES6 中的八进制表示法
let octalLiteral: number = 0o744;

let notANumber: number = NaN;

let infinityNumber: number = Infinity;


// 字符串 -------------
let myName: string = 'Tom';
let myAge: number = 25;

// 模板字符串
let sentence: string = `Hello, my name is ${myName}.
I'll be ${myAge + 1} years old next month.`;


// 空值 -------------
function alertName(): void {
    alert('My name is Tom');
}

let unusable: void = undefined;


// Null 和 Undefined -------------
let u: undefined = undefined;
let n: null = null;

// undefined 和 null 是所有类型的子类型
let num: number = undefined;
let num1: number = null

let u1: undefined;
let num2: number = u

// void 类型的变量不能赋值给 number 类型的变量

let u2: void;
let num3: number = u;
// Type 'void' is not assignable to type 'number'.