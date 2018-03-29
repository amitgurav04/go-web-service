package main

import (
	"fmt"
	"net/http"
	"math"
	"math/rand"
	"io/ioutil"
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

func Multiple(a, b string) (string, string) {
	return a, b
}

func TypeCast(n int) float64 {
	var num float64 = float64(n)
	fmt.Println("Float64 value is:",num)
	return num
}

func PointerExample(n int) {
	x := n
	a := &x
	fmt.Println("Address: ", a)
	fmt.Println("a: ", *a)

	*a = 5
	fmt.Println("a:", *a)
	fmt.Println("x:", x)

	*a = *a**a
        fmt.Println("a:", *a)
        fmt.Println("x:", x)

}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Version 1.0.1: Welcome you are at thi path %s! Amit Gurav<h1>
<p>Good to go</p>`, r.URL.Path[1:])
}

func about_handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Version 1.0.1: This is About page", r.URL.Path[1:])
}

type car struct {
	gas_pedal int
	brake_pedal int
	steering_wheel int
	top_speed_km int
}

const usixteenbitmax int = 1
const kmh_multiple int = 1

func (c car) kmh() float64 {
	return float64(c.gas_pedal*c.top_speed_km/usixteenbitmax)
}

func (c car) mph() float64 {
        return float64(c.gas_pedal*c.top_speed_km/usixteenbitmax/kmh_multiple)
}

func (c *car) new_speed(speed int) {
        c.top_speed_km = speed
}

func accessInternet() {
        resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
        bytes, _ := ioutil.ReadAll(resp.Body)
        string_body := string(bytes)
        fmt.Println(string_body)
}

func main() {
        Welcome()
	PrintSqrt(25)
	PrintRandomNumber(1000)

	var num1 float64 = 5.6
	var num2 float64 = 9.5
	fmt.Println("Addition is:",add(num1,num2))

        var numb1, numb2 float64 = 5.6, 9.5
        fmt.Println("Addition is:",add(numb1,numb2))

        str1, str2 := Multiple("Hello","amit")
        fmt.Println("String combination is: ",str1,str2)

	var num int  = 45
	fmt.Println("TypeCast of num:", TypeCast(num))

	PointerExample(5)

	car_a := car{gas_pedal: 65, brake_pedal: 1, steering_wheel: 1, top_speed_km: 2}
	fmt.Println("Gas_pedal is:", car_a.gas_pedal)
        fmt.Println("KMH is:", car_a.kmh)
        fmt.Println("MPH is:", car_a.mph)
	fmt.Println("Top speed:", car_a.top_speed_km)
	car_a.new_speed(500)
	fmt.Println("Top speed:", car_a.top_speed_km)

	accessInternet()
 
	http.HandleFunc("/", index_handler)
        http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8080", nil)
}
