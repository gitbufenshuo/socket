package main

func main() {
	initFlag()
	var tp = getType()
	if tp == "c" { // client
		client(getHost(), getPort(), getFile())
	} else { // server
		server(getPort(), getStorepath())
	}
}
