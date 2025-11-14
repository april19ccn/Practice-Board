本小节，是我把 Newsqueak 写的素数筛改成了 Go 版本

```Newsqueak
// 向通道输出从2开始的自然数序列
counter := prog(c:chan of int) {
    i := 2;
    for(;;) {
        c <-= i++;
    }
};
  
// 针对listen通道获取的数列，过滤掉是prime倍数的数
// 新序列输出到send通道
filter := prog(prime:int, listen, send:chan of int) {
    i:int;
    for(;;) {
        if((i = <-listen)%prime) {
            send <-= i;
        }
    }
};
  
// 主函数
// 每个通道第一个流出的数必然是素数
// 然后基于这个新素数构建新的素数过滤器
sieve := prog() of chan of int {
    c := mk(chan of int);
    begin counter(c);
    prime := mk(chan of int);
    begin prog(){
        p:int;
        newc:chan of int;
        for(;;){
            prime <-= p =<- c;
            newc = mk();
            begin filter(p, c, newc);
            c = newc;
        }
    }();
    become prime;
};
  
// 启动素数筛
prime := sieve();
```

## 并发
每个新过滤器都使用上一个过滤器的输出通道作为自己的输入通道。

### 具体连接过程
初始状态：

counter → 通道 c0

发现素数2：

创建 filter(2)，连接：c0 → filter(2) → c1

现在主程序使用 c1 作为当前通道

发现素数3：

创建 filter(3)，连接：c1 → filter(3) → c2

现在主程序使用 c2 作为当前通道

发现素数5：

创建 filter(5)，连接：c2 → filter(5) → c3

现在主程序使用 c3 作为当前通道

### 为什么这样设计？
这种"接力"方式的好处：

- 逐步过滤：每个素数只负责过滤掉自己的倍数，剩下的交给后续素数

- 职责分离：每个过滤器只关心一个任务

- 自动扩展：发现新素数就自动添加新的过滤器

- 并发高效：所有过滤器并行工作

### 实际的数据流
数字经历的完整旅程：

2,3,4,5,6,7,8,9,10,11,... → filter(2) → 3,5,7,9,11,... → filter(3) → 5,7,11,... → filter(5) → 7,11,... → ...