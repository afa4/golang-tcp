// Team members
// Gerson Fialho | jgfn1
// Arthur Frade | afa4
// César Silva | accs2

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/afa4/golang-tcp/util"
)

func main() {
	tcpServer()
}

func tcpServer() {
	addr := "localhost:1313"
	tcpAddress, err := net.ResolveTCPAddr("tcp", addr)
	util.FailOnError(err)
	listener, err := net.ListenTCP("tcp", tcpAddress)
	util.FailOnError(err)
	fmt.Printf("TCP Server up at %s\n", addr)
	tcpHandleConnections(listener)
}

func tcpHandleConnections(listener *net.TCPListener) {
	for {
		connection, err := (*listener).Accept()
		util.FailOnError(err)
		go tcpHandleMessages(&connection)
	}
}

func tcpHandleMessages(connection *net.Conn) {
	for {
		request, err := bufio.NewReader(*connection).ReadString('\n')
		if err != nil {
			break
		}
		result := tcpExecuteApplication(string(request[0]))
		(*connection).Write([]byte(result + "\n"))
	}
	(*connection).Close()
}

func tcpExecuteApplication(request string) string {
	num, err := strconv.Atoi(request)
	util.FailOnError(err)
	if num%2 == 0 {
		return "Yes"
	} else {
		return "No"
	}
}
