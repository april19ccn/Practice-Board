// 如果换成type会报错（同名 type 会报错，不会合并）
interface Alarm11111 {
    weight: number;
    alert(s: string, n: number): string;
}
interface Alarm11111 {
    weight: number;
    alert(s: string, n: number): string;
}
