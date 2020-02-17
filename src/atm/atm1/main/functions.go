package main

import "fmt"

//Pscheck : password-checking procedure
func Pscheck(A Acc) bool {
	for {
		var typeps int
		fmt.Println("請輸入密碼: ")
		fmt.Scanln(&typeps)
		if A.pass == typeps {
			fmt.Println("密碼正確")
			break
		} else {
			fmt.Println("密碼錯誤，請重試")
		}
	}
	return true
}

// Pschange : change password
func Pschange(A *Acc) {
	fmt.Println("你選擇模式:更改密碼")
	var psnew1, psnew2 int
	for {
		fmt.Println("請輸入新密碼")
		fmt.Scanln(&psnew1)
		fmt.Println("請再輸入一次新密碼")
		fmt.Scanln(&psnew2)
		if psnew1 == psnew2 {
			A.pass = psnew1
			fmt.Println("密碼更改成功，新密碼為", A.pass)
			break
		} else {
			fmt.Println("前後密碼不符，請重新輸入")
		}
	}
}

// Nachange : change name
func Nachange(A *Acc) {
	fmt.Println("你選擇模式:更改帳號名稱")
	var nanew1, nanew2 string
	for {
		fmt.Println("請輸入新帳號名稱")
		fmt.Scanln(&nanew1)
		fmt.Println("請再輸入一次新帳號名稱")
		fmt.Scanln(&nanew2)
		if nanew1 == nanew2 {
			A.name = nanew1
			fmt.Println("名稱更改成功，新名稱為", A.name)
			break
		} else {
			fmt.Println("前後名稱不符，請重新輸入")
		}
	}
}

// Money : manage account's money
func Money(A *Acc) {
	var (
		adjmoney int
		status   rune
	)
	fmt.Println("請輸入更動金額")
	fmt.Scanln(&adjmoney)
	fmt.Println("請輸入+/-調整模式")
	fmt.Scanln(&status)
	switch status {
	case '+':
		goto save
	case '-':
		goto withdraw
	}
save:
	A.money += adjmoney
	fmt.Println("成功存款", adjmoney, "餘額為", A.money)
withdraw:
	A.money -= adjmoney
	fmt.Println("成功提領", adjmoney, "餘額為", A.money)
}
