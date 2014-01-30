// Go 1.2
// go run helloworld_go.go

package main

import (
     "fmt" // Using '.' to avoid prefixing functions with their package names
     "net" // This is probably not a good idea for large projects...
     "time"
)

const (
    CONN_HOST = "129.241.187.161"
    CONN_PORT = "20022"
    CONN_TYPE = "udp"
)
//129.241.187.255
func main() {

		addr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20022")
		sock, _ := net.DialUDP("udp",nil, addr)

		defer sock.Close()
        if err != nil{
                fmt.Println("Error connecting")        
        
        }

		

		
		l, err := net.ListenUDP("udp", addr)

		if err!=nil{
			fmt.Println("Error listening")		
		}

		go write_to_udp(sock)
		go read_from_udp(l,addr)
		
		c := make(chan int)
		<-c
}

func write_to_udp(socket *net.UDPConn){
		fmt.Println("Sender started")
		time.Sleep(time.Millisecond * 1)
		var p string
		p="Hei, server"+("\x00")
        arr:=[]byte (p)
        socket.Write(arr)

}

func read_from_udp(sock *net.UDPConn,addr *net.UDPAddr){
	fmt.Println("Reader started") 
	buf:=make([]byte,1024)
	
	for {
		time.Sleep(time.Millisecond * 1)
	
		rlen,UDPAddr, _ := (&sock).ReadFromUDP(buf)
		s:=string(buf[:rlen])
		

		p:=UDPAddr.String()[:15]

		if ("129.241.187.118"!=p){	

			fmt.Println(s)
		}else{	
		}
	}
}
