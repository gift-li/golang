package main

import "fmt"

// Steps : 規範基本動作
type Steps interface {
	Modifyname()
	Modifypass()
	Deposit()
	Withdraw()
	Search(string) int
}

// Acc : 定義帳號格式
type Acc struct {
	name  []string
	pass  []int
	money []int
}

// Modifyname : 修改帳號
func (a *Acc) Modifyname() {
	var input string
	fmt.Println("目前帳號:", a.name[Peracc])
	fmt.Println("請輸入新名稱:")
	fmt.Scanln(&input)
	a.name[Peracc] = input
	fmt.Println("成功更改帳號，新帳號為:", a.name[Peracc])
	return
}

// Modifypass : 修改密碼
func (a *Acc) Modifypass() {
	var (
		check string
	)
	fmt.Println("目前密碼:", a.pass[Peracc])
	fmt.Println("你確定要修改密碼嗎? y/n")
	fmt.Scanln(&check)
	switch check {
	case "y":
		fmt.Printf("請輸入新密碼:")
		fmt.Scanln(a.pass[Peracc])
		fmt.Println("成功更改密碼，新密碼為:", a.pass[Peracc])
	case "n":
		fmt.Println("回到上一頁")
	}
}

// Deposit : 存款
func (a *Acc) Deposit() {
	var input int
	fmt.Println("請輸入存款金額:")
	fmt.Scanln(&input)
	a.money[Peracc] += input
	fmt.Println("成功存款，餘額:", a.money[Peracc])
	return
}

// Withdraw : 提款
func (a *Acc) Withdraw() {
	var input int
	fmt.Println("請輸入提款金額:")
	fmt.Scanln(&input)
	if a.money[Peracc] >= input {
		a.money[Peracc] -= input
		fmt.Println("成功提款，餘額:", a.money[Peracc])
	} else {
		fmt.Println("餘額不足")
	}
}

// Search : 尋找帳戶索引值
func (a *Acc) Search(s string) int {
	var i int = 0
	for i = range a.name {
		if a.name[i] == s {
			break
		}
	}
	return i
}
