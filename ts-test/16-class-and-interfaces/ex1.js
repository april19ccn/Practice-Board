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
var Door16 = /** @class */ (function () {
    function Door16() {
    }
    return Door16;
}());
var SecurityDoor16 = /** @class */ (function (_super) {
    __extends(SecurityDoor16, _super);
    function SecurityDoor16() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.address = "2";
        return _this;
    }
    SecurityDoor16.prototype.alert = function () {
        console.log('SecurityDoor alert' + this.address);
    };
    return SecurityDoor16;
}(Door16));
var Car16 = /** @class */ (function () {
    function Car16() {
        // address: "23";  // ❌ 错误：这是类型注解，而非赋值！
        this.address = "23";
    }
    Car16.prototype.alert = function () {
        console.log('Car alert' + this.address);
    };
    return Car16;
}());
var car16 = new Car16();
car16.alert();
