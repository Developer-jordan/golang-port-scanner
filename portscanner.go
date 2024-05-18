package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ch := make(chan int, 100)
	var IP string
	fmt.Printf("enter ip address here:")
	fmt.Scan(&IP)
	start := time.Now()
	for port := 0; port < 65535; port++ {
		ch <- port
		go portScanner(IP, "No discription", port)
		<-ch
	}
	close(ch)
	fmt.Println(time.Now().Sub(start))
}
func portScanner(ip string, portDiscription string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err == nil {
		fmt.Printf("\n[+] port: %v\n-->  for : %v \n", port, portDiscription)
	} else {
		return
	}
	defer conn.Close()
	return
}
