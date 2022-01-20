package main

type WithdrawInfo struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)            // send amount to deposit
var balances = make(chan int)            // receive balance
var withdrawCh = make(chan WithdrawInfo) // 引き出し情報を乗せるチャンネル

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	successCh := make(chan bool)
	wi := WithdrawInfo{amount: amount, success: successCh}
	withdrawCh <- wi
	return <-successCh
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case WithdrawInfo := <-withdrawCh:
			if balance <= WithdrawInfo.amount {
				WithdrawInfo.success <- false
			} else {
				balance -= WithdrawInfo.amount
				WithdrawInfo.success <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
