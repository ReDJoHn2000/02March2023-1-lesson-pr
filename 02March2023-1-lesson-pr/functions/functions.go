package functions

import "fmt"

func HalfPyramid(){
	fmt.Println("The half pyramid")
	for i := 1; i <= 6; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func InvertedHalfPyramid(){
	fmt.Println("The reversed half pyramid")
	for i := 6; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func HollowInvertedHalfPyramid(){
	fmt.Println("The reversed hollow half pyramid")
	for i := 6; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if i == 1 || i == 6 || i == 2 || i == j || j == 1{
				fmt.Printf("* ")
				} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func FullPyramid(){
	fmt.Println("The full pyramid")
	for i := 0; i <= 6; i++ {
		for j := 1; j <= 6 - i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func InvertedFullPyramid(){
	fmt.Println("The inverted full pyramid")
	for i := 6; i >= 0; i-- {
		for j := 5; j >= i - 5; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++{
			fmt.Print("* ")
		}
		fmt.Println()
	}

}

func HollowFullPyramid(){
	fmt.Println("The hollow full pyramid")
	for i := 0; i <= 6; i++ {
		for j := 1; j <= 6 - i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++{
			if k == 1 || i == 6 || k == i {
				fmt.Print(" *")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func FibonachiNum(){
	fmt.Println("The fibonachi numbers")
	var max_num int
	fmt.Print("Input the maximum number: ")
	fmt.Scanln(&max_num)
	var second_num int = 1
	var next_num int

	for i := 0; second_num < max_num;{
		next_num = second_num
		second_num = i + second_num;
		i = next_num
		fmt.Printf(" %d", second_num)
	}
	fmt.Println()
}

func FactorialNumber(num int) int {
	if num == 1 || num == 0 {
		return 1
	} else {
		return num * FactorialNumber(num - 1)
	}
}

func FizzBuzz(){

	fmt.Print("Enter integer: ")
	var input int
	fmt.Scanf("%d", &input)

	var fizz string = "Fizz";
	var buzz string = "Buzz";

	for i := 1; i <= input; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(i, fizz + buzz)
		} else if i % 3 == 0 {
			fmt.Println(i, fizz)
		} else if i % 5 == 0 {
			fmt.Println(i, buzz)
		} else {
			fmt.Println(i)
		}
	}
}

func ChessOfficer(){
	var (
		// x1, y1 int
		)
}