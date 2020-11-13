package main

import (
	"bufio"
	"fmt"
	"os"
)

type multipe struct {
	multipeNumber int
	multipeWord   string
}

func changeMultipeValue(multipe *multipe, nameOfMultiple string) {
	var answer string = ""

	// If you want change multipe
	fmt.Printf("Would you like change multipe of  %s ? (Y/N)", nameOfMultiple)
	readLine(&answer)

	if answer == "Y" {
		fmt.Printf("Enter the new value of multiple (%d) ? ", multipe.multipeNumber)
		fmt.Scanf("%d", &multipe.multipeNumber)

		fmt.Printf("Would you like change the word of multiple (%s) (Y/N) ? ", multipe.multipeWord)
		readLine(&answer)

		if answer == "Y" {
			fmt.Printf("Enter the new word of multiple (%s) ? ", multipe.multipeWord)
			fmt.Scanf("%d", &multipe.multipeWord)
		}
	}
}

func readLine(answer *string) {
	*answer = ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	*answer = scanner.Text()
}

func main() {
	var limit int = 200
	var answer string = ""
	multipeSeven := multipe{7, "Fizz"}
	multipleNine := multipe{9, "Buzz"}

	fmt.Printf("Would you like change limit default (%d) ? (Y/N)", limit)
	readLine(&answer)

	if answer == "Y" {
		fmt.Printf("Enter the new limit (%d) ? ", limit)
		fmt.Scanf("%d", &limit)
	}

	answer = ""
	changeMultipeValue(&multipeSeven, "seven")
	changeMultipeValue(&multipleNine, "Nine")

	for i := 1; i <= limit; i++ {
		if (i % multipeSeven.multipeNumber) == 0 {
			fmt.Print(multipeSeven.multipeWord)
		} else if (i % multipleNine.multipeNumber) == 0 {
			fmt.Print(multipleNine.multipeWord)
		}
	}

}
