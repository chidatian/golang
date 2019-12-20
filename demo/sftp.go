package main

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"log"
	"bytes"
	"net"
	// "github.com/pkg/sftp"
)

// func Dial(network, addr string, config *ClientConfig) (*Client, error)
func main() {
	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	// type HostKeyCallback func(hostname string, remote net.Addr, key PublicKey) error
	config := &ssh.ClientConfig{
	    User: "root",
	    Auth: []ssh.AuthMethod{
	        ssh.Password("iwehave2305!"),
	    },
	    HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
	    	return nil
	    },
	}
	client, err := ssh.Dial("tcp", "192.168.1.86:22", config)
	if err != nil {
	    log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
	    log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
	    log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}