# socket is a seed for socket programming with Golang .Uhmm, more that.

# go get github.com/gitbufenshuo/socket
# Then find the binary executable file in you $gobin
# Then $gobin/socket

---
# This socket seed code implements a file-uploader with c/s which server listen on one specified port that client should connect to. 
# flags
- type:['s' | 'c'] -> "s for server; c for client"
- host:[Any valid hostname or Ip address] -> "127.0.0.1"
- port:[Any available port number] -> "8080"
- storepath:<with --type s>[In the server, specify one path to hold all files] -> "/tmp/socketfiles"
- file:<with --type c>[In the client, specify one file to upload] -> "./goodthing.avi"
