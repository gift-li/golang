package main

// Acc funtion :account struct set
type Acc struct {
	name  string
	pass  int
	money int
}

func main() {
	A := Acc{"amy", 99, 100}
	// account-verifying process
	if Pscheck(A) {
		System(&A)
	}
}
