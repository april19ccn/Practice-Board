// 规则 1：自动递增的常数项 -----------------------------
var AutoIncrement;
(function (AutoIncrement) {
    AutoIncrement[AutoIncrement["A"] = 0] = "A";
    AutoIncrement[AutoIncrement["B"] = 1] = "B";
    AutoIncrement[AutoIncrement["C"] = 5] = "C";
    AutoIncrement[AutoIncrement["D"] = 6] = "D"; // 常数项 (D = C + 1 = 6)
})(AutoIncrement || (AutoIncrement = {}));
console.log(AutoIncrement.D); // 输出 6
// 规则 2：常数枚举表达式类型 -----------------------------
// - 1. 数字字面量
var NumericLiteral;
(function (NumericLiteral) {
    NumericLiteral[NumericLiteral["A"] = 10] = "A"; // 常数项（直接数字字面量）
})(NumericLiteral || (NumericLiteral = {}));
console.log(NumericLiteral.A); // 输出 10
// - 2. 引用其他常数项
var ReferenceOther;
(function (ReferenceOther) {
    ReferenceOther[ReferenceOther["A"] = 1] = "A";
    ReferenceOther[ReferenceOther["B"] = 1] = "B";
    ReferenceOther[ReferenceOther["C"] = 1] = "C";
    ReferenceOther[ReferenceOther["D"] = 10] = "D"; // 常数项（引用其他枚举的常数项）
})(ReferenceOther || (ReferenceOther = {}));
console.log(ReferenceOther.B); // 输出 1
// - 3. 带括号的表达式
var Parenthesized;
(function (Parenthesized) {
    Parenthesized[Parenthesized["A"] = 9] = "A"; // 常数项（括号内为常数表达式）
})(Parenthesized || (Parenthesized = {}));
console.log(Parenthesized.A); // 输出 9
// - 4. 一元运算符 (+, -, ~)
var UnaryOperators;
(function (UnaryOperators) {
    UnaryOperators[UnaryOperators["A"] = 10] = "A";
    UnaryOperators[UnaryOperators["B"] = -20] = "B";
    UnaryOperators[UnaryOperators["C"] = -1] = "C"; // 常数项 (~0 → -1)
})(UnaryOperators || (UnaryOperators = {}));
console.log(UnaryOperators.C); // 输出 -1
// - 5. 二元运算符 (+, -, *, /, %, << 等)
var BinaryOperators;
(function (BinaryOperators) {
    BinaryOperators[BinaryOperators["A"] = 5] = "A";
    BinaryOperators[BinaryOperators["B"] = 2] = "B";
    BinaryOperators[BinaryOperators["C"] = 2] = "C";
    BinaryOperators[BinaryOperators["D"] = 8] = "D";
    BinaryOperators[BinaryOperators["E"] = Infinity] = "E"; // 编译错误：结果为 Infinity
})(BinaryOperators || (BinaryOperators = {}));
// 规则 3：计算所得项的陷阱 -----------------------------
// enum ComputedMember {
//     A = "Hello".length,  // 计算所得项 (值为 5)
//     B                    // 编译错误：B 无法自动递增
// }
// 原因：A 是计算所得项，后面的 B 无法自动初始化。
// 规则 4：混合类型枚举 -----------------------------
var MixedEnum;
(function (MixedEnum) {
    MixedEnum[MixedEnum["A"] = 0] = "A";
    MixedEnum[MixedEnum["B"] = "TS".length] = "B";
    MixedEnum[MixedEnum["C"] = MixedEnum.B + 1] = "C";
    MixedEnum[MixedEnum["D"] = 10] = "D";
    MixedEnum[MixedEnum["E"] = 11] = "E"; // 常数项 (E = D + 1 = 11)
})(MixedEnum || (MixedEnum = {}));
console.log(MixedEnum.E); // 输出 11
// 规则 5：NaN 和 Infinity 的编译检查
var InvalidValues;
(function (InvalidValues) {
    InvalidValues[InvalidValues["A"] = Infinity] = "A";
    InvalidValues[InvalidValues["B"] = NaN] = "B"; // 编译错误：NaN
})(InvalidValues || (InvalidValues = {}));
// 实际并不会编译错误，有编译输出
