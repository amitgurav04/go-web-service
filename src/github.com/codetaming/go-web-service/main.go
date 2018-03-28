package main

import (
	"fmt"
	"net/http"
	"math"
	"math/rand"
)

func Welcome() {
  fmt.Println("Welcome")
}

func PrintRandomNumber(n int) {
	fmt.Println("Random number: ", rand.Intn(n))
}
func PrintSqrt(n float64) {
	fmt.Println("The square root of ", n, "is:", math.Sqrt(n))
}

func add(x float64, y float64) float64 {
	return x+y
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version 1.0.1: Welcome you are at thi path %s! Amit Gurav", r.URL.Path[1:])
}


func main() {
        Welcome()
	PrintSqrt(25)
	PrintRandomNumber(1000)

	var num1 float64 = 5.6
	var num2 float64 = 9.5
	fmt.Println("Addition is:",add(num1,num2))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
