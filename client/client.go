// Team members
// Gerson Fialho | jgfn1
// Arthur Frade | afa4
// César Silva | accs2

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/afa4/golang-tcp/util"
)

var newSource = rand.NewSource(time.Now().UnixNano())
var newRand = rand.New(newSource)

func main() {
	numberOfRequests, err := strconv.Atoi(os.Args[1])
	util.FailOnError(err)
	tcpClient(numberOfRequests)
}

func tcpClient(numberOfRequests int) {
	serverAddr, err := net.ResolveTCPAddr("tcp", "localhost:1313")

	if err != nil {
		fmt.Println("Fatal Error: Could not resolve server address")
		return
	}

	connection, err := net.DialTCP("tcp", nil, serverAddr)

	for i := 0; i < numberOfRequests; i++ {
		number := strconv.Itoa(newRand.Intn(9))
		response := tcpRequest(connection, number+"\n")
		fmt.Printf("Client: Is %s divisible by 2? \nServer: %s\n", number, response)
	}
	connection.Close()
}

func tcpRequest(connection *net.TCPConn, message string) string {
	_, err := fmt.Fprint(connection, message)
	util.FailOnError(err)

	response, err := bufio.NewReader(connection).ReadString('\n')
	util.FailOnError(err)

	return response
}
