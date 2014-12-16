package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type NetFlow5 struct {
	Version          uint16
	Count            uint16
	SysUptime        uint32
	UnixSecs         uint32
	UnixNsecs        uint32
	FlowSequence     uint32
	EngineType       uint8
	EngineID         uint8
	SamplingInterval uint16
}

type NetFlow5Record struct {
	SrcAddr  uint32
	DstAddr  uint32
	NextHop  uint32
	Input    uint16
	Output   uint16
	DPkts    uint32
	DOctets  uint32
	First    uint32
	Last     uint32
	SrcPort  uint16
	DstPort  uint16
	Pad1     uint8
	TcpFlags uint8
	Prot     uint8
	Tos      uint8
	SrcAs    uint16
	DstAs    uint16
	SrcMask  uint8
	DstMask  uint8
	Pad2     uint16
}

// Protocol translation as seen in netinet/in.h
// not add protocols are covered, for the rest UNK(NUMBER) is returned
func (n *NetFlow5Record) ProtoToString(Prot uint8) string {
	switch int(Prot) {
	case 17:
		return "UDP"
	case 6:
		return "TCP"
	case 1:
		return "ICMP"
	case 2:
		return "IGMP"
	case 8:
		return "EGP"
	}
	return fmt.Sprintf("UNK%d", int(Prot))
}

func (n *NetFlow5Record) Print() {
	dur := fmt.Sprintf("%d", int(n.Last-n.First)/1000)
	srcTup := fmt.Sprintf("%s:%d", IPtoString(n.SrcAddr), int(n.SrcPort))
	dstTup := fmt.Sprintf("%s:%d", IPtoString(n.DstAddr), int(n.DstPort))
	fmt.Printf("%10ss%25s%25s", dur, srcTup, dstTup)
	fmt.Printf("%10s%10d%10d\n", n.ProtoToString(n.Prot), int(n.DPkts), int(n.DOctets))
}

func IPtoString(IP uint32) string {
	s := strconv.FormatUint(uint64(IP), 16)
	a, _ := hex.DecodeString(s)
	if len(a) == 0 {
		return fmt.Sprintf("EMPTY\n")
	}
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}
