package main

import "fmt"

// Option : 清單
func Option(a *Account) {
	fmt.Println("進入系統清單，請輸入數字選擇功能")
	var (
		count int = 1
		input int
	)
	for count > 0 {
		fmt.Println("1.更改帳號名稱\n2.更改密碼\n3.提領/存取帳戶金額\n4.離開本系統")
		fmt.Scanln(&input)
		if input == 4 {
			count--
		} else {
			switch input {
			case 1:
				acc.Name()
			case 2:
				acc.Pass()
			case 3:
				acc.Money()
			default:
				fmt.Println("輸入無效")
			}
		}
	}
	return
}

// Build : 初次創帳
func Build(a *Account) {
	var (
		inp1 string
		inp2 int
	)
	fmt.Println("歡迎使用本系統，將進行初次設定")
	fmt.Println("請輸入帳號:")
	fmt.Scanln(&inp1)
	a.name = inp1
	fmt.Println("請輸入密碼(只接受數字):")
	fmt.Scanln(&inp2)
	a.pass = inp2
	fmt.Println("請輸入首次存入金錢:")
	fmt.Scanln(&inp2)
	a.money = inp2
	fmt.Printf("建檔完成\n您的帳號為: %s\n密碼為: %d\n帳戶金額為: %d\n", a.name, a.pass, a.money)
	return
}

// Name : 更改帳號名稱
func (a *Account) Name() {
	var n string
	fmt.Println("進入改帳號名稱系統\n舊帳號為:", a.name)
	fmt.Println("請輸入要更改的帳號名稱")
	fmt.Scanln(&n)
	a.name = n
	fmt.Println("帳號名稱已成功更改，新名稱為:", a.name)
	fmt.Println("離開改帳號名稱系統")
}

// Pass : 更改密碼
func (a *Account) Pass() {
	var (
		input  string
		p1, p2 int
	)
	fmt.Println("進入改密碼系統\n舊密碼為", a.pass)
	fmt.Println("你要更改密碼嗎? y/n")
	fmt.Scanln(&input)
	if input == "y" {
		fmt.Println("請輸入要更改的密碼")
		fmt.Scanln(&p1)
		fmt.Println("請再輸入一次新密碼")
		fmt.Scanln(&p2)
		if p1 == p2 {
			a.pass = p1
			fmt.Println("密碼輸入正確，新密碼為:", a.pass)
		} else {
			fmt.Println("前後密碼不符合")
		}
	}
	fmt.Println("離開改密碼系統")
	return
}

// Money : 更改帳戶金額
func (a *Account) Money() {
	var (
		opt        string
		inputmoney int
	)
	fmt.Println("進入提領款系統，帳戶內有$", a.money)
	fmt.Println("請確認提款或領款 (+ / -):")
	fmt.Scanln(&opt)
	fmt.Println("請輸入金額:")
	fmt.Scanln(&inputmoney)
	if opt != "" && inputmoney > 0 {
		switch opt {
		case "+":
			a.money += inputmoney
			fmt.Println("存款成功")
		case "-":
			if a.money >= inputmoney {
				a.money -= inputmoney
				fmt.Println("提款成功")
			} else {
				fmt.Println("帳戶金額不足")
			}
		}
	} else {
		fmt.Println("此次輸入無效")
	}
	fmt.Println("餘額為$", a.money)
	return
}

// func (a *Account) writeName() {
// 	fmt.Printf("原帳號名稱為: %s, 請輸入新帳號名稱: ", a.name)
// 	n, err := fmt.Scanln()
// 	if err != nil {
// 		a.name = string(n)
// 		fmt.Println("成功更改，新帳號名稱為:", a.name)
// 		return
// 	}
// }

// func (a *Account) writePass() {
// 	fmt.Printf("原密碼為: %d, 請輸入新密碼: ", a.pass)
// 	p, err := fmt.Scanln()
// 	if err != nil {
// 		a.pass = p
// 		fmt.Println("成功更改，新帳號名稱為:", a.pass)
// 		return
// 	}
// }

// func (a *Account) writeMoney() {
// 	fmt.Printf("原帳戶金額為: %d, 請先輸入存取或提領的金額: ", a.money)
// 	p, err1 := fmt.Scanln()
// 	if err1 != nil {
// 		fmt.Println("請輸入+/-決定存款或提領:")
// 		action, err2 := fmt.Scanln()
// 		if err2 != nil {
// 			switch string(action) {
// 			case "+":
// 				a.money += p
// 				fmt.Printf("已存款%d元\n帳戶金額為%d\n回到主頁面", p, a.money)
// 			case "-":
// 				a.money -= p
// 				fmt.Printf("已提領%d元\n帳戶金額為%d\n回到主頁面", p, a.money)
// 			}
// 		}
// 	}
// }
