package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	client()
}
func client() {
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	go mustCopy(os.Stdout, con)
	mustCopy(con, os.Stdin)
}
func mustCopy(dest io.Writer, src io.Reader) { //消息传输
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
}
