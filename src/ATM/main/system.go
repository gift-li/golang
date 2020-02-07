package main

import (
	"fmt"
)

// Acc : 資料規格設定
type Acc struct {
	name  string
	pass  int
	money int
}

// System : 主操作程式
func System() int {
	var a int
	for {
		var input int
		fmt.Println("輸入指令以進行動作\n1:創建帳號\n2:改密碼\n3:改名稱\n4:提領錢\n5:離開")
		fmt.Scanln(&input)
		switch input {
		case 1:
			Build(A)
		case 2:
			Pschange(A)
		case 3:
			Nachange(A)
		case 4:
			Money(A)
		case 5:
			break
		default:
			fmt.Println("無效指令")
		}
	}
}
