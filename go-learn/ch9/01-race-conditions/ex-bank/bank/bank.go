// 练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
// 其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
// 这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，
// 这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan struct {
	amount int
	res    chan bool
})

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraw <- struct {
		amount int
		res    chan bool
	}{amount, result}
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case data := <-withdraw:
			if balance >= data.amount {
				balance -= data.amount
				data.res <- true
			} else {
				data.res <- false
			}

		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
