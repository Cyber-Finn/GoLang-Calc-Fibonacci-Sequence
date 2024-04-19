package main

import "fmt"

type FibonacciNumber struct {
	previous_number int
	current_number  int
}

var (
	continue_operation bool
	fib_num            FibonacciNumber
)

func main() {
	init_struct()
	calc_next_num()
}

func init_struct() {
	continue_operation = true
	fib_num.current_number = 0
	print_data()
}

func calc_next_num() {
	for continue_operation {
		if increment_data() {
			print_data()
		}
	}
}

func print_data() {
	fmt.Printf("%d,", fib_num.current_number)
}

func increment_data() bool {
	if fib_num.current_number < 130 { //number here is arb/adjustable, currently prints: "0,1,1,2,3,5,8,13,21,34,55,89,144,"
		temp_num := 0
		if fib_num.current_number == 0 {
			temp_num = fib_num.current_number + 1
		} else {
			temp_num = fib_num.current_number + fib_num.previous_number
		}

		fib_num.previous_number = fib_num.current_number
		fib_num.current_number = temp_num
		return true
	} else {
		handle_exception()
		return false
	}
}

func handle_exception() {
	continue_operation = false
}
