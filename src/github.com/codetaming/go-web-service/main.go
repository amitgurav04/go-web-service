package main

import (
	"fmt"
	"net/http"
	"math"
	"math/rand"
	"io/ioutil"
	"encoding/xml"
	"html/template"
	"time"
	"sync"
	//"go-web-service/src/github.com/codetaming/go-web-service/stringutil"
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

type SitemapIndex struct {
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	Loc string `xml:"loc"`
}

func (s Sitemap) String() string  {
	return fmt.Sprintf(s.Loc)
}

func (s SitemapIndex) RangeExample() {
        fmt.Println("Example of Range:")
        for _, sitemap := range s.Sitemaps {
                fmt.Printf("\n%s",sitemap)
        }
}

func parseXML() {
        resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
        bytes, _ := ioutil.ReadAll(resp.Body)

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	s.RangeExample()
	fmt.Println(s.Sitemaps)
}

func forLoopExample(n int) {
	fmt.Println("Example of For Loop:")
	for i:=0; i<n; i++ {
		fmt.Println(i)
	}
}

func whileLoopExample(n int) {
        fmt.Println("Example of while Loop using for loop:")

	i := 0
        for i<n {
                fmt.Println(i)
		i++
		i+=2
        }
}

func breakInfiniteLoopExample(n int) {
        fmt.Println("Example of break Infinite Loop using for loop:")

        i := 0
        for {
                fmt.Println(i)
		i+=3

		if i>n {
			break
		}
        }
}

func mapExample() {
	grades := make(map[string]float64)
	grades["Amit"] = 45.3
	grades["Manoj"] = 55.2
	grades["Amol"] = 63.3

	fmt.Println("Grades from Map are:",grades)

	AmsGrade := grades["Amit"]
	fmt.Println("Grade of Amit is:",AmsGrade)

        delete(grades,"Amit")
        fmt.Println("Grades after deleteing Amit:", grades)

	for k, v := range grades {
		fmt.Println(k,":",v)
	}
}

type NewsAggPage struct {
	Title string
	News string
}

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
	defer wg.Done()
}

func say(s string) {
	defer cleanup()
	defer fmt.Println("I am Done")
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2 {
			panic("Ohh I am panic, it's 2")
		}
	}
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amit Gurav", News: "Won the match"}
	t, err := template.ParseFiles("application.html")
	fmt.Println(err)
	t.Execute(w, p)
}
var wg sync.WaitGroup

func goRoutineExample() {
        wg.Add(1)
        go say("Hey")
        wg.Add(1)
        go say("There")
        wg.Wait()
}

func putOnChannel(c chan int, somevalue int) {
	defer wg.Done()
	c <- somevalue * 5
}

func channelExample() {
	fmt.Println("This is example of channel:")
	mychannel := make(chan int, 12)

	wg.Add(1)
	go putOnChannel(mychannel, 3)
	wg.Add(1)
	go putOnChannel(mychannel, 5)

	v1 := <-mychannel
	v2 := <-mychannel

	fmt.Println("v1 from channel:", v1)
	fmt.Println("v2 from channel:", v2)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go putOnChannel(mychannel, i)
	}
	wg.Wait()
	close(mychannel)
	for item := range mychannel {
		fmt.Println("Val from channel is:",item)
	}
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
	parseXML()
	forLoopExample(10)
	whileLoopExample(20)
	breakInfiniteLoopExample(30)
	mapExample()
	goRoutineExample()
	channelExample()
	//fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	http.HandleFunc("/", index_handler)
        http.HandleFunc("/about", about_handler)
	http.HandleFunc("/news",newsAggHandler)
	http.ListenAndServe(":8080", nil)
}
