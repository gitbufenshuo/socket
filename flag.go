package main

import (
	"flag"
	"log"
	"os"
)

var clientOrServer string
var port string
var host string
var filepath string
var storepath string

func initFlag() {
	flag.StringVar(&clientOrServer, "type", "c", "c for client and s for server")
	flag.StringVar(&port, "port", "8090", "the port to connet or listen")
	flag.StringVar(&host, "host", "", "the remote server host to connect")
	flag.StringVar(&filepath, "file", "", "which file to send")
	flag.StringVar(&storepath, "storepath", "", "absolute path : where to store the recieved files")
	flag.Parse()
	if clientOrServer == "c" {
		if host == "" {
			log.Println("[--host] need be specified!!!")
			os.Exit(0)
		}
		if filepath == "" {
			log.Println("[--file] need be specified!!!")
			os.Exit(0)
		}
	} else {
		if storepath == "" {
			log.Println("[--storepath] need be specified!!!")
			os.Exit(0)
		}
	}
	log.Printf("[type]->[%s] [port]->[%s]\n", clientOrServer, port)
}

func getType() string {
	return clientOrServer
}
func getHost() string {
	return host
}
func getPort() string {
	return port
}
func getFile() string {
	return filepath
}
func getStorepath() string {
	return storepath
}
