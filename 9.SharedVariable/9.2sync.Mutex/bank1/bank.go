package main

//模拟一个bank
func main() {

}

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func deposits(amount int) {
	sema <- struct{}{} //获取令牌
	balance += amount
	<-sema //释放令牌
}
func balances() int {
	sema <- struct{}{}
	money := balance
	<-sema
	return money
}
