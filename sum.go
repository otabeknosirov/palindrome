package main

import "fmt"

func main(){
	
	var num int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
  
	
    var evenSum, oddSum  = sum(num)
	fmt.Println("even sum is ", evenSum)
	fmt.Println("odd sum is ", oddSum)
}

func sum(num int) (int, int){

	var evenS int = 0
	var oddS int = 0

    for i:=1; i <= num; i++ { 
		if(i % 2 == 0){
           evenS += i
		} else if(i % 2 != 0){
			oddS = oddS + i
		}
	}

	return evenS, oddS
}