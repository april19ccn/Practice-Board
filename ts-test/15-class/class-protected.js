var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
// 工厂模式 (Factory)
var AnimalFactory = /** @class */ (function () {
    function AnimalFactory(name) {
        this.name = name;
    }
    // 工厂方法：根据类型创建不同实例
    AnimalFactory.create = function (name, type) {
        switch (type) {
            case 'cat': return new CatFactory(name);
            default: throw new Error('Invalid type');
        }
    };
    return AnimalFactory;
}());
var CatFactory = /** @class */ (function (_super) {
    __extends(CatFactory, _super);
    function CatFactory(name) {
        return _super.call(this, name) || this; // ✅ 允许访问父类 protected 构造函数
    }
    return CatFactory;
}(AnimalFactory));
// 使用
var catFactory = AnimalFactory.create('Whiskers', 'cat');
