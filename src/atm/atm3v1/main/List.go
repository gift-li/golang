package main

import "fmt"

// Show : 顯示帳戶資訊
func Show(a *Acc) {
	fmt.Printf("顯示帳戶資訊:\n帳號: %s\n密碼: %d\n餘額: %d\n", a.name[Peracc], a.pass[Peracc], a.money[Peracc])
	return
}

// Modify : 修改資料
func Modify(a *Acc) {
	var (
		input int
	)
	fmt.Println("請選擇欲修改項目\n1. 帳號\n2. 密碼")
	fmt.Scanln(&input)
	switch input {
	case 1:
		a.Modifyname()
	case 2:
		a.Modifypass()
	default:
		fmt.Println("輸入錯誤，回到上一頁")
	}
}

// Cashflow : 金流系統
func Cashflow(a *Acc) {
	var input int
	fmt.Println("目前餘額: ", a.money[Peracc])
	fmt.Println("請選擇動作:\n1. 存錢\n2. 領錢")
	fmt.Scanln(&input)
	switch input {
	case 1:
		a.Deposit()
	case 2:
		a.Withdraw()
	}
}

// Logout : 登出系統
func Logout(a *Acc, route *int) {
	var input string
	fmt.Println("即將登出，你確定要登出嗎? y/n")
	fmt.Scanln(&input)
	switch input {
	case "y":
		fmt.Scanf("%s成功登出\n", a.name[Peracc])
		Peracc = 0
		*route--
	case "n":
		fmt.Scanln("將回到清單")
	}
	return
}
