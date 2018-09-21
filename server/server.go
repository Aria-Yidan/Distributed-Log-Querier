package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"strings"
)

type GrepStr string

type GreReq struct {
	RegPat   string
	Filename string
}

//run grep command on server and get the results: https://blog.csdn.net/qq_36874881/article/details/78234005
func (s *GrepStr) GrepResult(req GreReq, reply *string) error {

	commandName := "grep"
	params := []string{"-n"}
	path := "/home/yidanli2/MP1/log/" + req.Filename
	fmt.Println("Path = ", path)
	params = append(params, req.RegPat, path)
	cmd := exec.Command(commandName, params...)

	//output the grep results
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Grep Error")
		fmt.Println(err)
	}
	*reply = string(output)
	*reply = strings.TrimSpace(*reply) //delete the blank line
	return nil
}

// rpc based on TCP protocol, wait for the client to call: https://blog.csdn.net/qq_34777600/article/details/81159443
func main() {
	strMessage := new(GrepStr)
	rpc.Register(strMessage)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9000")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("Start Listening!")
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err.Error())
		os.Exit(1)
	}
}
