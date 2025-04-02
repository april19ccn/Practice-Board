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
function isFish(animal) {
    if (typeof animal.swim === 'function') {
        return true;
    }
    return false;
}
// 将一个父类断言为更加具体的子类
var ApiError = /** @class */ (function (_super) {
    __extends(ApiError, _super);
    function ApiError() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.code = 0;
        return _this;
    }
    return ApiError;
}(Error));
var HttpError = /** @class */ (function (_super) {
    __extends(HttpError, _super);
    function HttpError() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.statusCode = 200;
        return _this;
    }
    return HttpError;
}(Error));
function isApiError(error) {
    if (typeof error.code === 'number') {
        return true;
    }
    return false;
}
function isApiError1(error) {
    if (error instanceof ApiError) {
        return true;
    }
    return false;
}
// 将任何一个类型断言为 any
window.foo = 1;
// 补充：如何扩展window
function getLocalApiUrl(url) {
    var _a;
    var _window = window;
    var baseUrl = ((_a = _window === null || _window === void 0 ? void 0 : _window.globalConfig) === null || _a === void 0 ? void 0 : _a.api) || '';
    //  let baseUrl = window?.globalConfig?.api || ''
    return "".concat(baseUrl).concat(url);
}
function testAnimal(animal) {
    return animal;
}
function testCat(cat) {
    return cat;
}
