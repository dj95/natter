package listen

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// OnInterface Start listening on the specified interface with the given filter
func OnInterface(intrfc, filter string) {
	handle, err := pcap.OpenLive(intrfc, 1600, false, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	err = handle.SetBPFFilter(filter)

	if err != nil {
		panic(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	go startListening(intrfc, packetSource)
}

func startListening(intrfc string, packetSource *gopacket.PacketSource) {
	for packet := range packetSource.Packets() {
		handlePacket(intrfc, packet)
	}
}

func handlePacket(intrfc string, pkt gopacket.Packet) {
	nlayer := pkt.NetworkLayer()
	srcIP := nlayer.NetworkFlow().Src().String()
	dstIP := nlayer.NetworkFlow().Dst().String()

	transportType := pkt.Layers()[2].LayerType()

	fmt.Printf(
		"%s :: %s :: %s -> %s\n",
		intrfc,
		transportType,
		srcIP,
		dstIP,
	)
}
