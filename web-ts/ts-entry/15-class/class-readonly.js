var AnimalRedonly = /** @class */ (function () {
    function AnimalRedonly(name) {
        this.name = name;
    }
    return AnimalRedonly;
}());
var aAnimalRedonly = new AnimalRedonly('Jack');
console.log(aAnimalRedonly.name); // Jack
// aAnimalRedonly.name = 'Tom'; // 无法为“name”赋值，因为它是只读属性。
// index.ts(10,3): TS2540: Cannot assign to 'name' because it is a read-only property.
var AnimalRedonly2 = /** @class */ (function () {
    // public readonly name;
    function AnimalRedonly2(name) {
        this.name = name;
        // this.name = name;
    }
    return AnimalRedonly2;
}());
