package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)
import "os"

import "time"

func server(port string, storepath string) {
	log.Printf("[storepath] -> [%s]\n", storepath)
	makeDirOK(storepath)
	hostPort := "0.0.0.0:" + port
	l, e := net.Listen("tcp4", hostPort)
	if e != nil {
		log.Println(e.Error())
		os.Exit(0)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err.Error())
			os.Exit(0)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	bheader := make([]byte, headerSize, headerSize)
	conn.Read(bheader)
	var filename string
	var filenamelen uint32
	var filesize uint32
	filenamelen = binary.BigEndian.Uint32(bheader)
	log.Printf("TEMP:[filenamelen] -> [%d]\n", filenamelen)
	filename = string(bheader[4 : 4+filenamelen])
	filesize = binary.BigEndian.Uint32(bheader[4+filenamelen:])
	log.Printf("[%s]->[%d] = [AC]\n", filename, filesize)

	b := make([]byte, blockSize, blockSize)
	f, _ := os.Create(getlastfilename(storepath, filename))

	var recvBytesCount uint32
	for {
		if recvBytesCount == filesize {
			break
		}
		nr, _ := conn.Read(b)
		recvBytesCount += uint32(nr)
		nw, _ := f.Write(b[:nr])
		if nw != nr {
			log.Println("[Fuck Fatal Write File So Slow]")
			os.Exit(0)
		}
	}
	f.Close()
}

func makeDirOK(storepath string) {
	_, err := os.Stat(storepath)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(storepath, 0777)
		}
	}
}

func getlastfilename(storepath, filename string) string {
	return storepath + "/" + strconv.Itoa(time.Now().Nanosecond()) + "_" + filename
}
