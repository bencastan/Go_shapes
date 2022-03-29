package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangleShape struct{
	height  float64
	base float64
}
type squareShape struct{
	sideLength  float64
}

func main() {
	ts := triangleShape{base: 10, height:10}
	ss := squareShape{sideLength: 10}

	printArea(ts)
	printArea(ss)
}


// func printGreeting(eb englishBot){
// 	fmt.Println((eb.getGreeting()))
// }

// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (ts triangleShape) getArea() float64 {
	// Very custom logic for generating an english greeting
	return (0.5 * ts.base * ts.height)
}

func (ss squareShape) getArea() float64{
	// Very custom logic for generating a spanish greeting
	return (ss.sideLength * ss.sideLength)

}
