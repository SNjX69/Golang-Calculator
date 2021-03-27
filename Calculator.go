package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Print("\n\n 0 For Subtract\n\n")

	fmt.Print(" 1 For Addition\n\n")

	fmt.Print(" 2 For Multtiply\n\n")

	fmt.Print(" 3 For Dividing\n\n")

	fmt.Print(" Type Your Choice  : ")

	var input float64
	fmt.Scanln(&input)

	print("\n")

	if input >= 4 {
		print(" Wrong Choice\n\n")

		time.Sleep(500000000)
		print("Programmed ")

		time.Sleep(500000000)
		print("By ")

		time.Sleep(500000000)
		print(" SNjX, iG : ")

		time.Sleep(500000000)
		print("www.instagram.com/6o9s , ")

		time.Sleep(500000000)
		print("Github : www.Github.com/SNjX69 \n\n")

		time.Sleep(500000000)
		print("Press Enter To Exit ")

		fmt.Scanln()

		os.Exit(0)
	}

	time.Sleep(500000000 * 2)
	fmt.Print(" Type The First Number : ")
	var input2 float64
	fmt.Scanln(&input2)

	print("\n\n")

	time.Sleep(500000000 * 2)
	fmt.Print(" Type The Second Number : ")
	var input3 float64
	fmt.Scanln(&input3)

	print("\n\n")

	time.Sleep(500000000 * 2)
	print(" Result Is : ")

	if input == 0 {
		output := input2 - input3
		fmt.Print(output)

		print("\n\n")

	} else if input == 1 {
		output := input2 + input3

		fmt.Print(output)

		print("\n\n")

	} else if input == 2 {
		output := input2 * input3
		fmt.Print(output)

		print("\n\n")

	} else if input == 3 {
		output := input2 / input3
		fmt.Print(output)

		print("\n\n")

	}

	time.Sleep(500000000)
	print(" Programmed ")

	time.Sleep(500000000)
	print("By ")

	time.Sleep(500000000)
	print(" SNjX, iG : ")

	time.Sleep(500000000)
	print("www.instagram.com/6o9s , ")

	time.Sleep(500000000)
	print("Github : www.Github.com/SNjX69 \n\n")

	time.Sleep(500000000)
	print(" Press Enter To Exit ")

	fmt.Scanln()
}
