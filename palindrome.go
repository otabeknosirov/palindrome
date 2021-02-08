package main

import "fmt"

func main()  {
	var palindrome string
	fmt.Println("Enter the expression")
	fmt.Scan(&palindrome)
	
	var isPal bool
	isPal = isPalindrome(palindrome)

	if (isPal == true) {
		fmt.Println(palindrome, "is Palindrome")
	} else if (isPal == false){
		fmt.Println(palindrome,"is not Palindrome")
	}
}

func isPalindrome(exp string)bool{

	var signal bool
	var n int = len(exp)
	fmt.Printf("len %d\n",n)
	for i:=0; i < n; i++ { 
       if(exp[i] == exp[n - i - 1]){
		   signal = true
	   }
	   if(exp[i] != exp[n - i - 1]){
		   signal = false
		   break
	   }
	}
	return signal
}