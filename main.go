package main

import (
	"fmt"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
const name = "Kostya"

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
	hello := "hello"

	fmt.Println(hello)
	fmt.Println(name)
	Greet()
	PersonalGreet("angelo4ek")
	FioGreet("Vladimir", "Putin")
	first, second := 2, 4
	sum := first + second
	fmt.Println(sum)
	summa, multiplay := SumAndMultiplay(first, second)
	fmt.Println(summa, multiplay)
	_, multiplay64 := namedSumAndMultiplay(first, second)
	fmt.Println(multiplay64)
}
func Greet() {
	fmt.Println("Hello, angelochek!")
}

func PersonalGreet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func FioGreet(name, surname string) {
	fmt.Printf("President, %s %s\n", name, surname)
}

func Sum(first, second int) int {
	sum := first + second
	return sum
}

func SumAndMultiplay(first, second int) (int, int) {
	return first + second, first * second
}

func namedSumAndMultiplay(first, second int) (sum int64, multiplay int64) {
	sum = int64(first + second)
	multiplay = int64(first) * int64(second)
	return sum, multiplay: 10
}

func FioGreet(name, surname string) {
	fmt.Printf("Foot Ball, %s %s\n", name, surname)