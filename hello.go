package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var userInput int
	var secertNumber int
	rand.Seed(time.Now().Unix())
	secertNumber = rand.Intn(10)
	fmt.Println("Secert Number is :", secertNumber)
	fmt.Printf("Please enter a number:")
	fmt.Scan(&userInput)
	fmt.Println("You Entered:", userInput)	
	if userInput == secertNumber{
		fmt.Println("You won!!")
	}else if userInput < secertNumber{
		fmt.Println("Less than")
	}else if userInput >secertNumber{
		fmt.Println("Great than")
	}
}