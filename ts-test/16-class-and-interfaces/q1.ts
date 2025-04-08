// interface 能用type替代吗
type Alarm = {
    alert(): void;
}

type LightableAlarm extends Alarm ={
    lightOn(): void;
    lightOff(): void;
}

// interface 和 abstract class 的区别