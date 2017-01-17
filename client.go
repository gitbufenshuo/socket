package main

import (
	"encoding/binary"
	"log"
	"net"
	"os"
)

var blockSize = 1024 * 1024 // 1M
var headerSize = 1024       // 1K

func client(host, port string, file string) {
	filesize, filename := fileSizeAndName(file)
	if filesize <= 0 {
		log.Printf("[%s] not exists!!\n", file)
		os.Exit(0)
	}
	log.Printf("[file size byte] -> [%d byte] OR [%d k-byte] OR [%d M-byte]\n", filesize, filesize/1024, filesize/(1024*1024))
	f, err := os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
	}
	conn, err := net.Dial("tcp4", host+":"+port)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	defer conn.Close()

	// begin send
	// [][][][]{[]}[][][][]{[]}
	// 文件名长度 文件名 文件长度 文件
	filenamelength := len(filename)
	filenamebyte := []byte(filename)
	b := make([]byte, blockSize, blockSize)
	bheader := make([]byte, headerSize, headerSize)
	binary.BigEndian.PutUint32(bheader, uint32(filenamelength))
	copy(bheader[4:], filenamebyte)
	binary.BigEndian.PutUint32(bheader[4+filenamelength:], uint32(filesize))
	headerLen := 8 + filenamelength
	nw, _ := conn.Write(bheader[:8+filenamelength])
	if nw != headerLen {
		log.Println("[Fuck Fatal Header]")
		os.Exit(0)
	}
	var fileBytesCount uint32

	for {
		if fileBytesCount == uint32(filesize) {
			break
		}
		nr, _ := f.Read(b)
		fileBytesCount += uint32(nr)
		var sendBytesCount uint32
		for {
			if nr == int(sendBytesCount) {
				break
			}
			nw, _ := conn.Write(b[sendBytesCount:nr])
			sendBytesCount += uint32(nw)
		}
	}
	log.Printf("[fileBytesCount] -> [%d]\n", fileBytesCount)
}
func fileSizeAndName(path string) (int64, string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return -1, ""
		}
	}
	return fileInfo.Size(), fileInfo.Name()
}
