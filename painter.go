package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Find(slice []string, sliceItem string) bool {
	for _, item := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the paint calculation tool.\n")
	fmt.Println("How many walls would you like to paint? ")
	scanner.Scan()
	numWalls := scanner.Text()
	n, err1 := strconv.Atoi(numWalls)
	var tWalls int
	if err1 != nil {
		panic("Invalid number input, run code again. ")
	} else if n == 0 {
		fmt.Println("Fine, keep your secrets. ")
		os.Exit(0)
	} else {
		tWalls = n
	}
	wallsAndDimensions := make(map[string]float64)

	for totalWalls := 0; totalWalls < tWalls; totalWalls++ {
		fmt.Println("Please input the name of your wall, if you are inputting multiple walls, they must have unique names: ")
		scanner.Scan()
		nWalls := scanner.Text()
		if _, found := wallsAndDimensions[nWalls]; found {
			fmt.Println("Oops! That name is not unique. Please start again.")
			panic("duplicate names")
		}
		fmt.Println("Please input the dimensions (m2) of your wall as a positive number: ")
		scanner.Scan()
		nDimensions := scanner.Text()
		n, err2 := strconv.ParseFloat(nDimensions, 64)
		if err2 != nil {
			fmt.Println("Your input is not a number. Please only input a positive number. ")
			panic("err")
		} else {
			wallsAndDimensions[nWalls] = n
		}
	}

	fmt.Println("How many types of paint will you use? ")
	scanner.Scan()
	typesOfPaint := scanner.Text()
	n2, err1 := strconv.Atoi(typesOfPaint)
	if err1 != nil {
		panic("Invalid number input, run code again. ")
	} else if n == 0 {
		fmt.Println("Fine, keep your secrets. ")
		os.Exit(0)
	}
	if n2 == 1 {
		fmt.Println("How much does the paint cost? ($ per liter)")
		scanner.Scan()
		costOfPaint := scanner.Text()
		n3, err3 := strconv.ParseFloat(costOfPaint, 64)
		if err3 != nil {
			panic("Your input is not a number. Please only input a positive number. ")
		} else {
			fmt.Println("How many coats of paint would you like? ")
			scanner.Scan()
			coatsOfPaint := scanner.Text()
			n4, err4 := strconv.ParseFloat(coatsOfPaint, 64)
			if err4 != nil {
				panic("Your input is not a number. Please only input a positive number. ")
			} else {
				for wallName, wallArea := range wallsAndDimensions {
					paintAmt := (wallArea / 11 * n4)
					paintCost := (paintAmt * n3)
					fmt.Printf("Wall %s will need %.2f liters of paint. It will cost %2.f to paint.\n", wallName, paintAmt, paintCost)
				}
			}
		}
	} else if n2 > 1 {

		paintList := make(map[string]float64)

		for paintTypes := 0; paintTypes < n2; paintTypes++ {
			fmt.Println("Please input the name of your paint, they must be unique:")
			scanner.Scan()
			nPaint := scanner.Text()
			if _, found := paintList[nPaint]; found {
				fmt.Println("Oops! That name already exists. I know this is annoying but please start again.")
				panic("duplicate paint names")
			}
			fmt.Println("How much does the paint cost? ($ per liter):")
			scanner.Scan()
			cPaint := scanner.Text()
			n4, err4 := strconv.ParseFloat(cPaint, 64)
			if err4 != nil {
				fmt.Println("Your input is not a number. Please only input a positive number. ")
				panic("err")
			} else {
				paintCost := n4
				paintList[nPaint] = paintCost
			}
		}

		for wallName, wallArea := range wallsAndDimensions {
			fmt.Printf("Which type of paint would you like to use for wall %s? Please use the exact name of the paint.\n", wallName)
			scanner.Scan()
			xPaint := scanner.Text()
			if _, found := paintList[xPaint]; found {
				fmt.Printf("How many coats of paint would you like for wall %s?", wallName)
				scanner.Scan()
				coatsOfPaint := scanner.Text()
				n5, err5 := strconv.ParseFloat(coatsOfPaint, 64)
				cPaint := paintList[xPaint]
				if err5 != nil {
					panic("Your input is not a number. Please only input a positive number. ")
				} else {
					paintAmt := (wallArea / 11 * n5)
					paintCost := (paintAmt * cPaint)
					fmt.Printf("Wall %s will need %.2f liters of paint and will cost %.2f to paint.\n", wallName, paintAmt, paintCost)
				}
			} else {
				panic("Could not find paint in paint list. Please try again and ensure that you are using the exact name of paint specified.\n")
			}
		}

	}
	// comment for pushin
}
