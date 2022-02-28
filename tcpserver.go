package main

import "net"
import "fmt"

func handler(c net.Conn) {
	c.Write([]byte("Trym Falkum"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	fmt.Println(l.Addr())

	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}