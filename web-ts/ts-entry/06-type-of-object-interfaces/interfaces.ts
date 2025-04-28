// interface Person {
//     name: string;
//     age: number;
// }

// let tom: Person = {
//     name: 'Tom',
//     age: 25
// };

// interface Person {
//     name: string;
//     age?: number;
//     [propName: string]: string;
// }

// let tom: Person = {
//     name: 'Tom',
//     gender: 'male'
// };

// [propName: string]: any;
// 允许有任意的属性

// [propName: string]: string;
// 一旦定义了任意属性，那么确定属性和可选属性的类型都必须是它的类型的子集
// 故：类型“number”的属性“age”不能赋给“string”索引类型“string”

// 一个接口中只能定义一个任意属性。如果接口中有多个类型的属性，则可以在任意属性中使用联合类型
interface Person {
    name: string;
    age?: number;
    [propName: string]: string | number;
}

interface Person2 {
    name: string;
    age?: number;
    x: () => void;
    // 1: () => void; // 类型“() => void”的属性“1”不能赋给“number”索引类型“string | number”。
    [propName: number]: string | number;
}

let tom: Person = {
    name: 'Tom',
    age: 25,
    gender: 'male'
};
