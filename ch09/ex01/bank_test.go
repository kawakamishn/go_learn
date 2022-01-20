package main

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	Deposit(200)
	balance := Balance()
	if balance != 200 {
		t.Errorf("Error")
	}
}

func TestWithdraw(t *testing.T) {
	var success bool
	Deposit(200)
	success = Withdraw(150) // 200-150 = 50 が成功しているか
	balance := Balance()
	if balance != 50 || success == false {
		t.Errorf("Error")
	}
	success = Withdraw(51) // 50-51 = -1 によって失敗状態になっているか
	if success == true {
		t.Errorf("Error")
	}
}
