package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		defer conn.Close()

		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		fmt.Println("Code got here.")
		_, err = io.WriteString(conn, "I see you connected.")
		if err != nil {
			log.Print(err)
		}
		fmt.Println("after write")
		conn.Close()
	}
}
