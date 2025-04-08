// 工厂模式 (Factory)
class AnimalFactory {
    protected constructor(public name: string) { }

    // 工厂方法：根据类型创建不同实例
    static create(name: string, type: 'cat' | 'dog'): Animal {
        switch (type) {
            case 'cat': return new CatFactory(name);
            default: throw new Error('Invalid type');
        }
    }
}

class CatFactory extends AnimalFactory {
    constructor(name: string) {
        super(name); // ✅ 允许访问父类 protected 构造函数
    }
}

// 使用
const catFactory = AnimalFactory.create('Whiskers', 'cat');


