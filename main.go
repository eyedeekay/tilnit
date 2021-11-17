package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/reiver/go-telnet"
)

var (
	ahost = flag.String("host", "localhost", "hostname")
	aport = flag.Int("port", 23, "port")
)

func main() {
	flag.Parse()
	host := *ahost
	port := fmt.Sprintf("%d", *aport)
	fmt.Printf("host: %s, port: %d\n", *ahost, *aport)
	switch len(flag.Args()) {
	case 0:
		flag.Usage()
	case 1:
		host = flag.Args()[0]
	case 2:
		host = flag.Args()[0]
		port = flag.Args()[1]
	}
	if strings.HasSuffix(host, "i2p") {
		var caller telnet.Caller = telnet.StandardCaller
		hostport := host
		fmt.Printf("I2P telnet - host: %s\n", hostport)
		if err := telnet.DialToI2PAndCall(hostport, caller); err != nil {
			panic(err)
		}
	} else {
		var caller telnet.Caller = telnet.StandardCaller
		hostport := fmt.Sprintf("%s:%s", host, port)
		fmt.Printf("Regular telnet -  host:port=%s\n", hostport)
		if err := telnet.DialToAndCall(hostport, caller); err != nil {
			panic(err)
		}
	}

}
