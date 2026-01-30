package Duration

import "fmt"

type Duration int

func FromMinutes(min int) Duration {
	return Duration(min * 60)
}

func FromHours(h int) Duration {
	return Duration(h * 3600)
}

func FromDays(d int) Duration {
	return Duration(d * 24 * 3600)
}

func Example() {
	//var a Duration
	var a Duration = FromMinutes(15)
	fmt.Println(a)

	//Output:

}
