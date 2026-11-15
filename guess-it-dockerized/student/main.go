package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Error(err error) string {
	return err.Error()
}

func MaxMin(Data []float64) (float64, float64) {
	slices.Sort(Data)
	return Data[0], Data[len(Data)-1]
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run program.go || ./program_name")
		return
	}

	// read data file
	fileData, err := os.ReadFile("data.txt")
	if err != nil {
		Error(err)
	}

	// fields automatically handles spaces, tabs and newlines.
	// split will fail if there is empty element in the end.
	fileStr := strings.Fields(string(fileData))

	if len(fileStr) == 0 {
		fmt.Println("Error: empty file")
		return
	}

	// transform to numbers.
	Data := make([]float64, 0)
	for _, val := range fileStr {
		val, err := strconv.ParseFloat(val, 64)
		if err != nil {
			Error(err)
		}
		Data = append(Data, val)
	}

	numbers := make([]float64, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			Error(err)
		}
		numbers = append(numbers, number)
		if len(numbers) >= 2 {
			av := Average(numbers)

			va := Variance(Data, av)
			if va < 0{
				va = -va
			}

			sd := StandardDeviation(va)

			Lower := av - (sd+40)
			Upper := av + (sd+40)

			fmt.Printf("%0.f %0.f\n", Lower, Upper)
		}else{
			fmt.Printf("%.0f %0.0f\n",number-40, number+40)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

/*
//calculate the slope a
		var seg1, seg2 float64
		for i := 0; i < len(xAxis); i++ {
			seg1 += (xAxis[i] - avX) * (numbers[i] - avY)
			seg2 += math.Pow(xAxis[i]-avX, 2)
		}
		a := seg1 / seg2
		fmt.Printf("a=%.2f ", a)
		//calculate the intercept b
		b := avY - (a * avX)
		fmt.Printf("b=%.2f \n", b)

		//calcul Linear Regression
		y := (a * float64(len(numbers))) + b
		fmt.Println("y=", y)
*/
