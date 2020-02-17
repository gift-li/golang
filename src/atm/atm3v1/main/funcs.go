package main

import "fmt"

// Log : 登入系統
func Log(a *Acc) {
	var (
		input int
		route int = 1
	)
	fmt.Println("歡迎使用ATM3.0")
	for route > 0 {
		fmt.Println("請輸入欲執行動作\n1. 新增帳號\n2. 登入系統\n3. 離開系統")
		fmt.Scanln(&input)
		switch input {
		case 1:
			Build(a)
		case 2:
			Login(a)
		case 3:
			Leave(&route)
		default:
			fmt.Println("輸入無效，請重新輸入")
		}
	}
}

// List : 清單系統
func List(a *Acc) {
	var (
		input int
		route int = 1
	)
	for route > 0 {
		fmt.Println("請選擇清單項目\n1. 個人資料顯示\n2. 修改資料\n3. 金流\n4. 登出")
		fmt.Scanln(&input)
		switch input {
		case 1:
			Show(a)
		case 2:
			Modify(a)
		case 3:
			Cashflow(a)
		case 4:
			Logout(a, &route)
		}
	}
	return
}
