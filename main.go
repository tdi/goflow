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

const version string = "0.2"
const address string = "127.0.0.1"
const port string = "2055"

func print_help() {
	fmt.Printf("goflow version: %s\n", version)
	fmt.Println("usage: goflow [-h] [-H HOST_NAME] [-p PORT]")
	os.Exit(0)
}

func main() {

	addressFlag := flag.String("H", address, "address to listen on default: localhost")
	portFlag := flag.String("p", port, "port to listen on, default: 2055")
	helpFlag := flag.Bool("h", false, "help message")
	flag.Parse()

	if *helpFlag != false {
		print_help()
	}

	connString := *addressFlag + ":" + *portFlag

	udpAddress, err := net.ResolveUDPAddr("udp4", connString)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp4", udpAddress)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening on ", connString)

	defer ln.Close()

	fmt.Printf("%11s%25s%25s", "Duration", "SrcAddr:SrcPort", "DstAddr:DstPort")
	fmt.Printf("%10s%10s%10s\n", "Proto", "Packets", "Octets")

	for {
		var header NetFlow5
		var buf []byte = make([]byte, 1500)
		_, _, err := ln.ReadFromUDP(buf)
		if err != nil {
			log.Print("Cannot read from UDP socket")
			continue
		}
		p := bytes.NewBuffer(buf)
		err = binary.Read(p, binary.BigEndian, &header)
		if err != nil {
			log.Print("Cannot read header from datagram")
			continue
		}
		if header.Version != uint16(5) {
			log.Fatal("goflow only support netflow 5")
		}
		//log.Printf("dgram size %d from %s [%d] records", n, adr, header.Count)
		// iterate over records, we are set after header after 24byte
		// record is 48 bytes
		for i := 0; i < int(header.Count); i++ {
			var record NetFlow5Record
			err = binary.Read(p, binary.BigEndian, &record)
			record.Print()
		}
	}
}
