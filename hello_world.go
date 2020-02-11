package main

import "fmt"

func main(){

	// Golang Data Types
	// fmt.Println("Hello World Kehinde");

	// const pi = 3.1424459990
	// fmt.Println(pi)

	// var Name string = "HelloTesting"
	// fmt.Println(len(Name))


	// var a int = 67
	// var b int = 43 
	// var c float32 = 32.4
	// var pis float64 = 13868959

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(pis)

	// fmt.Printf("%f \n", pi)

	// Golang Decision

	age:= 32

	if age > 30{
		fmt.Println("You are eligible to vote")
	}else{
		fmt.Println("You cannot vote")
	}

	fmt.Println("Golang Arrays")
	// Golang Arrays
	 var EventsNum[5] int

	 EventsNum[0] = 12
	 EventsNum[1] = 98
	 EventsNum[2] = 7
	 EventsNum[3] = 12
	 EventsNum[4] = 6

	 fmt.Println(EventsNum[1])


	 Events := [] int {44, 55, 32, 21, 10}
	 fmt.Println(Events[3])

	 fmt.Println("")
	 for i, value := range Events{
	 	fmt.Println(value, i)
	fmt.Printf("\n", value)
	// }

	// sliced := Events[1:]
	// fmt.Println(sliced)

	// fmt.Println("")


	// newSlice := append(Events, 100, 5)
	// fmt.Println(newSlice)

	// var studentCount[10] int;

	// for i:=0; i<10; i++{
	// 	studentCount[i] = i + 10
	// 	fmt.Println(studentCount[i])
	// }

	// fmt.Println("Golang Maps")
	// StudentAge := make(map[string] int)

	// StudentAge["kehinde"] = 45
	// StudentAge["tolu"] = 21
	// StudentAge["taye"] = 34
	// StudentAge["jones"] = 54
	// StudentAge["mercy"] = 11
	// StudentAge["emma"] = 87

	// fmt.Println(StudentAge["tolu"])

	// fmt.Println("Golang Map inside Map")

	// superHero := map[string]map[string]string{
	// 	"Superman" : map[string] string{
	// 		"RealName": "Clark Kent",
	// 		"City": "Metropolis",
	// 	},

	// 	"Batman" : map[string] string{
	// 		"RealName": "Bruce Wayne",
	// 		"City": "Gotham",
	// 	},
	// }

	// if temp, hero:= superHero["Superman"]; hero{
	// 	fmt.Println(temp["RealName"], temp["City"])
	// }

	// fmt.Println("Golang Functions")

	// x,y:= 6,7

	// fmt.Println(add(x,y))

	 num:= 8
	 fmt.Println(factorial(num))

	// fmt.Println("Defer, Recover and Panice")

	 addList := [5] int {20,34,55,24,65}
	
	fmt.Println(addList)

	fmt.Println(addUp(20,30,40,50,60))

	fmt.Println("Golang Struct")
	rect1 := Rectangle{10,5}
	fmt.Println("Area of a Rectangle", rect1.area())
}

type Rectangle struct {
	height float64
	width float64
}

func addUp(args ...int) int {
	sum := 0
	for _, value := range args{
		sum += value
	}
	return sum
}

func add(numa, numb int) int {
	return numa + numb
}

func factorial(num int) int {
	if num == 0 { return 1 }
	return num * factorial(num - 1)
}

func (rect *Rectangle) area() float64 {
 return	rect.height * rect.width
}