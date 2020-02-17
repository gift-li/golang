package main

// Adjust : 定義更改模式接口
type Adjust interface {
	Name()
	Pass()
	Money(int)
}

// Account : 定義帳號格式
type Account struct {
	name  string
	pass  int
	money int
}
