class AnimalRedonly {
    readonly name;
    public constructor(name) {
        this.name = name;
    }
}

let aAnimalRedonly = new AnimalRedonly('Jack');
console.log(aAnimalRedonly.name); // Jack
// aAnimalRedonly.name = 'Tom'; // 无法为“name”赋值，因为它是只读属性。

// index.ts(10,3): TS2540: Cannot assign to 'name' because it is a read-only property.


class AnimalRedonly2 { // 等价第一种声明
    // public readonly name;
    public constructor(public readonly name) { // 修饰符和readonly还可以使用在构造函数参数中，等同于类中定义该属性同时给该属性赋值，使代码更简洁。
        // this.name = name;
    }
}