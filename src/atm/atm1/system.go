package atm1

import (
	"fmt"
)

// System : main prog
func System(A *Acc) {
	for {
		var inp int
		fmt.Println("輸入指令以進行動作\n1:改密碼\n2:改名稱\n3:提領錢")
		fmt.Scanln(&inp)
		switch inp {
		case 1:
			Pschange(A)
		case 2:
			Nachange(A)
		case 3:
			Money(A)
		default:
			goto finish
		}
	}
finish:
}
