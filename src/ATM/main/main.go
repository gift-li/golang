package main

// Acc funtion :account struct set
type Acc struct {
	name  string
	pass  int
	money int
}

func main() {
	A := [3]Acc{}
	// account-verifying process
	if Pscheck(A) {
		System(&A)
	}
}
