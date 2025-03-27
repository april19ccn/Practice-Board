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