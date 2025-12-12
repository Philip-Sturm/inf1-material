package main

func Factorial(n int) int {
	Result := 1
	for i := n; i > 1; i-- {
		Result = Result * i
	}
	return Result
}

//func main() {
//	var Input int
//	fmt.Print("Wähle die Zahl des Factorials: ")
//	fmt.Scan(&Input)
//	fmt.Println(Factorial(Input))
//}
