package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    original, reversed := x, 0
    
    for x > 0 {
        digit := x % 10
        reversed = reversed*10 + digit
        x /= 10
    }
    
    return original == reversed
}

func main() {
    fmt.Println(isPalindrome(121))  
	fmt.Println(isPalindrome(10)) 
}












