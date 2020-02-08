// interface 和 struct 的連接練習
package main

import "fmt"

// adjust : 閱讀和撰寫
type adjust interface {
	read() (string int)
	write(string, int)
}

// struct : 格式建立
type account struct {
	name string
	pass int
}

func (a *account) read() (string, int) {
	return a.name, a.pass
}
func (a *account) write(b string, c int) {
	a.name = b
	a.pass = c
}
func change(a *account) {
	var (
		b string
		c int
	)
	fmt.Scanln(&b, &c)
	a.write(b, c)
	fmt.Println(a.name, a.pass)
}
func main() {
	a := account{"Apple", 20}
	change(&a)
}
