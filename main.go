package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	version string = "0.4"
	address string = "127.0.0.1"
	port    string = "2055"
)

func printHelp() {
	fmt.Printf("goflow version: %s\n", version)
	fmt.Println("usage: goflow [-h] [-H HOST_NAME] [-p PORT]")
	os.Exit(0)
}

func setupUDPServer(connString string, c chan string) {
	listenAddress, err := net.ResolveUDPAddr("udp4", connString)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", listenAddress)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		var header NetFlow5Header
		buf := make([]byte, 1464)
		_, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Cannot read from UDP socket")
			continue
		}
		p := bytes.NewBuffer(buf)
		err = binary.Read(p, binary.BigEndian, &header)
		// log.Printf("%v\n", header)
		if err != nil {
			log.Fatal("Cannot read header from datagram")
			continue
		}
		if header.Version != uint16(5) {
			log.Println("Invalid packet, goflow only supports netflow v5")
			continue
		}
		// log.Printf("dgram size %d from %s [%d] records", n, adr, header.Count)
		// iterate over records, we are set after header after 24byte
		// record is 48 bytes long
		var record NetFlow5Record
		for i := 0; i < int(header.Count); i++ {
			err = binary.Read(p, binary.BigEndian, &record)
			if err != nil {
				log.Printf("Could not read a packet")
			}
			c <- record.String()
		}
	}
}

func main() {

	addressFlag := flag.String("H", address, "address to listen on default: localhost")
	portFlag := flag.String("p", port, "port to listen on, default: 2055")
	helpFlag := flag.Bool("h", false, "help message")
	flag.Parse()

	if *helpFlag != false {
		printHelp()
	}
	connString := *addressFlag + ":" + *portFlag
	c := make(chan string)
	go setupUDPServer(connString, c)
	log.Printf("Listening on %s", connString)
	fmt.Printf("%11s%25s%25s", "Duration", "SrcAddr:SrcPort", "DstAddr:DstPort")
	fmt.Printf("%10s%10s%10s\n", "Proto", "Packets", "Octets")
	for {
		select {
		case a := <-c:
			fmt.Print(a)
		}
	}
}
