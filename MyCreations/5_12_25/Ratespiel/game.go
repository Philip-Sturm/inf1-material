package ratespiel

import (
	"fmt"
	"math/rand/v2"
)

func Ratespiel() {
	gesuchteZahl := RandomNum()
	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		if RichtigeZahl(guess, gesuchteZahl) {
			fmt.Println("Richtig geraten")
			return
		} else {
			fmt.Println("Falsch versuch es nochmal!")
		}
	}
	fmt.Println("Zu viele Versuche die richtige Zahl war", gesuchteZahl, "game over")
}
func RichtigeZahl(a int, r int) bool {
	return a == r
}
func ReadNumber() int {
	var n int
	fmt.Print("Rate eine Zahl zwischen 0 und 9: ")
	fmt.Scan(&n)
	return n
}
func RandomNum() int {
	return rand.IntN(10)
}
