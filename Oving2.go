package main

import (
    . "fmt" // Using '.' to avoid prefixing functions with their package names
    . "runtime" // This is probably not a good idea for large projects...
    . "time"
)

//var i = 0
//var p="0"
//var message=make(chan string,1)
//var m=""

func adder(i chan int) {
        
	for x := 0; x < 1000010; x++ {
                c := <- i
        	c++
		
                i <- c
	}
}
func adder2(i chan int){
	
	for x := 0; x < 1000000; x++ {
                d := <- i
                d--
		
                i <- d
	}

}

func main() {
	var i chan int = make(chan int, 1)
	i <- 0
	GOMAXPROCS(NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	go adder(i) // This spawns adder() as a goroutine
	go adder2(i)
	// No way to wait for the completion of a goroutine (without additional syncronization)
	// We'll come back to using channels in Exercise 2. For now: Sleep
	Sleep(Millisecond * 500)
	a := <-i
	Println("Done:", a);
}
