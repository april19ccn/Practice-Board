enum Days {Sun, Mon, Tue, Wed, Thu, Fri, Sat};

console.log(Days["Sun"] === 0); // true
console.log(Days["Mon"] === 1); // true
console.log(Days["Tue"] === 2); // true
console.log(Days["Sat"] === 6); // true

// 手动赋值的枚举项可以不是数字，此时需要使用类型断言来让 tsc 无视类型检查 (编译出的 js 仍然是可用的)：
enum Days2 {Sun = 7, Mon =10, Tue, Wed, Thu, Fri, Sat = <any>"S"};

// 但如果这样会报错
// enum Days3 {Sun = 7, Mon = 'b', Tue, Wed, Thu, Fri, Sat = <any>"S"};
// Tue, Wed, Thu, Fri 会报错提示 枚举成员必须具有初始化表达式。 所以需要之后的每个都赋值。


// 常数枚举
const enum Directions {
    Up,
    Down,
    Left,
    Right
}

let directions = [Directions.Up, Directions.Down, Directions.Left, Directions.Right];


// const enum Color {Red, Green, Blue = "blue".length}; // const 枚举成员初始值设定项必须是常量表达式。


declare enum Directions1 {
    Up,
    Down,
    Left,
    Right
}

let directions1 = [Directions.Up, Directions.Down, Directions.Left, Directions.Right];

