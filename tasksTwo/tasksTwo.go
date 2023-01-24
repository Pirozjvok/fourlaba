package taskstwo

import (
	"fmt"
	"githab/greenhell1337/LabaForDB/variables"
	"log"
)

func Task21(k, n int){
	if (n < 0) || (n == 0){
		log.Println("Ошибка!")
	} else{
		for i := 1; i < n; i++ {
			fmt.Printf(variables.ResultOneInt, k)
		}
	}

}

func Task22(a, b int){
	if (a > b) || (a == b){
		log.Println("Ошибка!")
	} else {
		for i := a; i <= b; i++ {
			fmt.Printf(variables.ResultOneInt, i)
		}
		fmt.Printf("Всего: %d", b)
	}
}

func Task23(a, b int){
	if (a > b) || (a == b){
		log.Println("Ошибка!")
	} else {
		for i := (b-1); i > a; i-- {
			fmt.Printf(variables.ResultOneInt, i)
		}
		fmt.Printf("Всего: %d", b-2)
	}
}

func Task24(a int){
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d КГ конфет стоит %d\n", i, (a*i))
	}
}

func Task25(a float64){
	for i := 0.1; i <= 1; i=i+0.1 {
		fmt.Printf("%g КГ конфет стоит %g\n", i, (a*i))
	}
}

func Task26(a float64){
	for i := 1.2; i <= 2; i=i+0.2 {
		fmt.Printf("%g КГ конфет стоит %g\n", i, (a*i))
	}
}

func Task27(a, b int){
	if (a > b) || (a == b){
		log.Println("Ошибка!")
	} else {
		for i := a; i <= b; i++ {
			variables.TempInt += i
			variables.Z = variables.TempInt
			continue
		}
		fmt.Printf(variables.ResultOneInt, variables.Z)
	}
}

func Task28(a, b int){
	if (a > b) || (a == b){
		log.Println("Ошибка!")
	} else {
		for i := a; i <= b; i++ {
			variables.TempInt += (i*i)
			variables.Z = variables.TempInt
			continue
		}
		fmt.Printf(variables.ResultOneInt, variables.Z)
	}
}

func Task29(a, b int){
	if (a > b) || (a == b){
		log.Println("Ошибка!")
	} else {
		for i := a; i <= b; i++ {
			fmt.Printf(variables.ResultOneInt, i)
		}
	}
}

func Task210(a, b int){
	if (a > b) || (a == b){
		log.Println("Ошибка!")
	} else {
		for i := a; i <= b; i++ {
			if i%2==0 {
				fmt.Printf(variables.ResultOneInt, i)
			}
		}
	}
}