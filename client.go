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
//sendBuf:= make([]byte,1024)
//sendBuf="Hei, Server\x00"
//begin:="Hello, server!"
//arr:=[]byte (sendBuf)
//conn.Write(arr)


var p string
for i := 0; i < 5; i++ {
	
	fmt.Println("Write a new message: ")
	fmt.Scanln(&p)
	p=p+("\x00")
	time.Sleep(time.Second * 1)

	arr:=[]byte (p)
	conn.Write(arr)
	time.Sleep(time.Second * 1)
	

	buf := make([]byte, 1024)

	reqLen, _ := conn.Read(buf)
	s:=string(buf[:reqLen])
	fmt.Println(s)
	
	


}
conn.Close	()
}
