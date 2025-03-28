// 「类型 + 方括号」表示法
let fibonacci: number[] = [1, 1, 2, 3, 5];

// 数组泛型
let fibonacci1: Array<number> = [1, 1, 2, 3, 5];

// 用接口表示数组
interface NumberArray {
    [index: number]: number;
}

let fibonacci2: NumberArray = [1, 1, 2, 3, 5];

// 类数组
// function sum() {
//     let args: number[] = arguments;
// }

// 为什么  [propName: string]: string | number; 会限制确定和可选属性，而 [index: number]: number; 不会呢？
// 在 TypeScript 中，字符串索引签名和数字索引签名对属性的约束行为不同，这与 JavaScript 的底层对象模型和类型系统的设计有关
// 关键区别总结
// 索引签名类型	         约束范围	                      与显式属性的关系
// 字符串索引签名	所有字符串键属性（包括显式属性）	显式属性必须兼容索引签名类型
// 数字索引签名	        仅数字键属性	                显式字符串键属性不受影响

interface NumberArray1 {
    [index: number]: number;
    length: number;
    callee: Function;
}


function sum() {
    let args: {
        [index: number]: number;
        length: number;
        callee: Function;
    } = arguments;
}