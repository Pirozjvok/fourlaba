package tasksone

import (
	"fmt"
	"githab/greenhell1337/LabaForDB/variables"
)

func Task11(a int){
	if a > 0 {
		a += 1
		fmt.Printf(variables.ResultOneInt, a)
	} else if a < 0 {
		a -= 2
		fmt.Printf(variables.ResultOneInt, a)
	} else {
		a = 10
		fmt.Printf(variables.ResultOneInt, a)
	}
}

func Task12(a, b, c int){
	arr := [...]int{
		a, b, c,
	}

	for _, v := range arr {
		if v > 0 {
			variables.Z += 1
			continue
		}
	}
	fmt.Printf(variables.ResultOneInt, variables.Z)
}

func Task13(a, b, c int) {
	arr := [...]int{
		a, b, c,
	}
	
	for _, v := range arr {
		if v > 0 {
			variables.Z += 1
			continue
		} else{
			variables.X += 1
			continue
		}
	}
	fmt.Printf("Положительные: " + variables.ResultOneInt, variables.Z)
	fmt.Printf("\nОтрицательные: " + variables.ResultOneInt, variables.X)
}

func Task14(a, b, c int){
	arr := [...]int{
		a, b, c,
	}
	variables.Max = arr[0]
	for i := 0; i < (len(arr)-1); i++ {
		if arr[i] > variables.Max {
			variables.Max = arr[i]
		}
	}
	fmt.Printf(variables.ResultOneInt, variables.Max)
}

func Task15(a, b float64){
	variables.TempFloat = a
	a = b
	b = variables.TempFloat
	fmt.Printf(variables.ResultTwoFloat, a, b)
}

func Task16(a, b int){
	variables.TempInt = a
	if a != b{
		a = a + b
		b = variables.TempInt + b
		fmt.Printf(variables.ResultTwoInt, a, b)
	} else if a == b {
		a = 0
		b = 0
		fmt.Printf(variables.ResultTwoInt, a, b)
	}
}

func Task17(a, b int){
	arr := [...]int{
		a, b,
	}
	if a != b {
		max := arr[0]
		for _, element := range arr {
		   if element > max {
			  max = element
		   }
		}
		a = max
		b = max
		fmt.Printf(variables.ResultTwoInt, a, b)
	} else if a == b {
		a = 0
		b = 0
		fmt.Printf(variables.ResultTwoInt, a, b)
	}
}

func Task18(a, b, c int) {
	arr := [...]int{
		a, b, c,
	}
	variables.Min = arr[0]
	for i := 0; i < (len(arr)-1); i++ {
		if arr[i] < variables.Min {
			variables.Min = arr[i]
		}
	}
	fmt.Printf(variables.ResultOneInt, variables.Min)
}

func Task19(a, b, c int){
	arr := [...]int{
		a, b, c,
	}

	for _, v := range arr {
		if v > 0 {
			variables.J += v
			continue
		}
	}
	fmt.Printf(variables.ResultOneInt, variables.J)
}

func Task110(a, b, c int){
	arr := [...]int{
		a, b, c,
	}
	variables.J = 1
	for _, v := range arr {
		if v < 0 {
			variables.J *= v
			continue
		}
	}
	fmt.Printf(variables.ResultOneInt, variables.J)
}