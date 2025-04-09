interface Alarm16 {
    address: string;
    alert(): void;
}

class Door16 {
}

class SecurityDoor16 extends Door16 implements Alarm16 {
    address = "2"
    alert() {
        console.log('SecurityDoor alert' + this.address);
    }
}

class Car16 implements Alarm16 {
    // address: "23";  // ❌ 错误：这是类型注解，而非赋值！
    address = "23"
    alert() {
        console.log('Car alert' + this.address);
    }
}

const car16 = new Car16();
car16.alert();