package main

import (
    "fmt"
    "net"
    "time"
)

const (
    CONN_HOST = "129.241.187.161"
    CONN_PORT = "33546"
    CONN_TYPE = "tcp"
)

func main() {
	conn,err:= net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil{
		fmt.Println("Error connecting")	
	
	}
	defer conn.Close()


	go readFromConn(conn)

	conn.Write([]byte("Connect to: 129.241.187.103:7770 \x00"))

	conn_s,err :=net.Listen("tcp","129.241.187.103:7770")
	if err!=nil{
		fmt.Println("Error listening")
	}

	fmt.Println("Accepting")
	c, err_a := conn_s.Accept()
	if err_a!=nil{
		fmt.Println("Error accepting")
	}

	fmt.Println("Accepted")
	go server(c)
	time.Sleep(time.Second * 100)

}

func server(conn_s net.Conn){
	
	fmt.Println("You are now running as a server")
	conn_s.Write([]byte("You are now connected to workspace 22\x00"))
	defer conn_s.Close()
	go readFromConn(conn_s)
	var p string
	for{
		
		
		fmt.Scanln(&p)
		p=p+("\x00")
		arr:=[]byte (p)
		conn_s.Write(arr)

		//buf:=make([]byte,1024)
		//reqLen, _ := conn_s.Read(buf)
		//s:=string(buf[:reqLen])
		//fmt.Println("Client wrote "+s+"\n")
		//s="you wrote: " +s+"\x00"
		//arr:=[]byte(s)
		//conn_s.Write(arr)
	
	}
}


func readFromConn(conn net.Conn){
	for {
		time.Sleep(time.Millisecond * 1)
		buf:=make([]byte,1024)
		reqLen, _ := conn.Read(buf)
		s:=string(buf[:reqLen])
		fmt.Println("Server said: "+s)
fmt.Println("Write a new message: ")
	}
}
