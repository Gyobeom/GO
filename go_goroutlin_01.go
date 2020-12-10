package main
import ("fmt"
       "time" // java Thread.sleep(5000);
)

func x(){
	for i:= 0 ; i<10; i++{
		fmt.Print("x")
	}
}
func y() {
	for i:= 0; i<10; i++{
		fmt.Print("y")
	}
}

func z() { 
	fmt.Println("z")
}
func k(){
	fmt.Println("K")
}

func main(){
	defer z() //defer main 함수가 끝날때 까지 지연
	defer k() //last in firstout으로 K가 먼저 출력 
	go x()  //x()의 경우 순차 처리 go x() go 루틴으로 처리됨 
	go y()  //y()의 경우 순차 처리 go y() go 루틴으로 처리됨 
	time.Sleep(time.Second * 2)
	fmt.Println("end main()");
}
