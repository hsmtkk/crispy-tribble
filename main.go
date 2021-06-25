package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/client4"
)

const needsBroadcast = true

func main() {
	printInterfaces()
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s interface", os.Args[0])
	}
	ifName := os.Args[1]
	pkt, err := dhcpv4.NewInformForInterface(ifName, needsBroadcast)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", pkt)

	clt := client4.NewClient()

	sendFd, err := client4.MakeBroadcastSocket(ifName)
	if err != nil {
		log.Fatal(err)
	}
	recvFd, err := client4.MakeListeningSocket(ifName)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := clt.SendReceive(sendFd, recvFd, pkt, dhcpv4.MessageTypeInform)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp)

}

func printInterfaces() error {
	ifs, err := net.Interfaces()
	if err != nil {
		return fmt.Errorf("failed to get interfaces; %w", err)
	}
	for _, iface := range ifs {
		fmt.Println(iface.Name)
	}
	return nil
}
