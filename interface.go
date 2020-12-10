package main

import "fmt"
import "math"

type Shape interface {    //메소드를 가지고 있지 않은 인터페이스는 어떠한 타입도 담을 수 있는 용기 
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}
 
func (r Rect) area() float64 {    //사각형 구조체 변수를 받고 해당 변수의 width, height 사용 
	return r.width * r.height
}	
func (r Rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s Shape) {  //shape 인터페이스타입의 개체들 사용 area, perimeter shpae 인터페이스를 이용하는 Rect , Circle의 객체들을 다 가져올 수 있음 
	fmt.Println(s)   //s가 circle이라면 circle의 객체들의 값을 가져올 수 있음, s가 rect이라면 rect의 객체들의 값을 가져올 수 있음 
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func showArea(shapes ...Shape) {  //도형이 몇개가 오든 들어오는 개수만큼 처리할 수 있도록 매개변수를 가변형태로 받음  ... 해서 처리
	for _, shape := range shapes {  //range를 이용하여 shpapes에 들어있는 값을 하나씩 처리 
		fmt.Println(shape.area())
	}
}

func printIf(itf interface{}){
	fmt.Println(itf)
}

func main() {
	var x interface{}
	r := Rect{width:5,height:2} //rect 구조체 변수를 선언 하고 해당 구조체의 값을 넣어줌 
	c := Circle{radius:10} //c라는 Circle 구조체 변수 생성 및 해당 구조체의 멤버 값을 넣어줌 
	
	x = 55
	x = "test"
	
	printIf(x)

	showArea(c,r)

	measure(r)
	measure(c)
}




