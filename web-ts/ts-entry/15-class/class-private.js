// ts把构造函数私有化 有啥用，以及有什么应用场景?
var Animal15 = /** @class */ (function () {
    function Animal15(name) {
        this.name = name;
    }
    return Animal15;
}());
// class Cat15 extends Animal15 { // 无法扩展类“Animal15”。类构造函数标记为私有。
//     constructor(name) {
//         super(name);
//     }
// }
// let a = new Animal('Jack');
// 一、核心用途
// 禁止直接实例化
// 当类构造函数为 private 时，外部无法通过 new 直接创建实例，但允许通过类内部的静态方法创建实例（如单例模式）。
// 强制通过工厂方法创建
// 可以要求用户必须通过特定静态方法（如工厂方法）创建实例，隐藏实现细节，增强控制。
// 二、
// 1. 单例模式 (Singleton)
var Singleton = /** @class */ (function () {
    function Singleton() {
    } // 私有构造函数
    Singleton.getInstance = function () {
        if (!Singleton.instance) {
            Singleton.instance = new Singleton();
        }
        return Singleton.instance;
    };
    return Singleton;
}());
// 使用
var s1 = Singleton.getInstance();
var s2 = Singleton.getInstance();
console.log(s1 === s2); // true，确保唯一性
// 2. 工具类
var MathUtils = /** @class */ (function () {
    function MathUtils() {
    } // 禁止实例化
    MathUtils.add = function (a, b) {
        return a + b;
    };
    return MathUtils;
}());
// 使用
MathUtils.add(1, 2); // ✅ 正确
//   new MathUtils();      // ❌ 编译错误：类“MathUtils”的构造函数是私有的，仅可在类声明中访问。
