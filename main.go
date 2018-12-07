package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8172")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			fmt.Fprintln(c, "Your name please?")
			scanner := bufio.NewScanner(c)
			scanner.Scan()
			fmt.Fprintf(c, "Hello %s!\n", scanner.Text())
			c.Close()
		}(conn)
	}

}
