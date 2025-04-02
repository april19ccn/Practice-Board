// 规则 1：自动递增的常数项 -----------------------------
enum AutoIncrement {
    A,          // 常数项 (默认 0)
    B,          // 常数项 (B = A + 1 = 1)
    C = 5,      // 常数项 (显式赋值 5)
    D           // 常数项 (D = C + 1 = 6)
}
console.log(AutoIncrement.D); // 输出 6

// 规则 2：常数枚举表达式类型 -----------------------------
// - 1. 数字字面量
enum NumericLiteral {
    A = 10       // 常数项（直接数字字面量）
}
console.log(NumericLiteral.A); // 输出 10

// - 2. 引用其他常数项
enum ReferenceOther {
    A = 1,
    B = A,        // 常数项（引用同枚举的常数项）
    C = ReferenceOther.A, // 常数项（引用其他枚举的常数项需限定名）
    D = NumericLiteral.A  // 常数项（引用其他枚举的常数项）
}
console.log(ReferenceOther.B); // 输出 1

// - 3. 带括号的表达式
enum Parenthesized {
    A = (1 + 2) * 3   // 常数项（括号内为常数表达式）
}
console.log(Parenthesized.A); // 输出 9

// - 4. 一元运算符 (+, -, ~)
enum UnaryOperators {
    A = +10,      // 常数项 (+10 → 10)
    B = -20,      // 常数项 (-20)
    C = ~0        // 常数项 (~0 → -1)
}
console.log(UnaryOperators.C); // 输出 -1

// - 5. 二元运算符 (+, -, *, /, %, << 等)
enum BinaryOperators {
    A = 2 + 3,        // 常数项 (5)
    B = 8 / 4,        // 常数项 (2)
    C = 7 % 5,        // 常数项 (2)
    D = 1 << 3,       // 常数项 (8)
    E = 1 / 0         // 编译错误：结果为 Infinity
}

// 规则 3：计算所得项的陷阱 -----------------------------
// enum ComputedMember {
//     A = "Hello".length,  // 计算所得项 (值为 5)
//     B                    // 编译错误：B 无法自动递增
// }
// 原因：A 是计算所得项，后面的 B 无法自动初始化。

// 规则 4：混合类型枚举 -----------------------------
enum MixedEnum {
    A,                  // 常数项 (0)
    B = "TS".length,    // 计算所得项 (2)
    C = B + 1,          // 计算所得项（B 是计算所得项）
    D = 10,             // 常数项（显式赋值）
    E                   // 常数项 (E = D + 1 = 11)
}
console.log(MixedEnum.E); // 输出 11

// 规则 5：NaN 和 Infinity 的编译检查
enum InvalidValues {
    A = 1 / 0,      // 编译错误：Infinity
    B = 0 / 0       // 编译错误：NaN
}
// 实际并不会编译错误，有编译输出