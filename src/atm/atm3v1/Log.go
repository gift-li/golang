package atm3v1

import (
	"fmt"
)

// Build : 創建帳號
func Build(a *Acc) {
	var (
		str string
		in  int
		lng int = len(a.name)
	)
	fmt.Println("進入創建帳戶功能\n請輸入帳號名稱")
	fmt.Scanln(&str)
	a.name = append(a.name, str)
	fmt.Println("請輸入密碼")
	fmt.Scanln(&in)
	a.pass = append(a.pass, in)
	fmt.Println("請輸入開戶金")
	fmt.Scanln(&in)
	a.money = append(a.money, in)
	fmt.Printf("創建帳戶成功!~\n帳號: %s\n密碼: %d\n餘額: %d\n", a.name[lng], a.pass[lng], a.money[lng])
}

// Login : 登入系統
func Login(a *Acc) {
	var (
		tempacc   int
		inputname string
		inputpass int
	)
	fmt.Println("進入登入介面\n請輸入帳號:")
	fmt.Scanln(&inputname)
	tempacc = a.Search(inputname)
	fmt.Println("請輸入密碼:")
	fmt.Scanln(&inputpass)
	if a.pass[tempacc] == inputpass {
		fmt.Println("登入成功，歡迎回來", a.name[tempacc])
		Peracc = tempacc
		List(a)
	} else {
		fmt.Println("登入失敗，將回到首頁")
	}
	return
}

// Leave : 離開/結束程式
func Leave(n *int) {
	fmt.Println("感謝使用本系統~~")
	*n--
	return
}
