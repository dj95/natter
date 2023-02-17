package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	initializeCommandFlags()
}

func initializeCommandFlags() {
	// create a new flag for docker health checks
	pflag.StringP("interfaces", "i", "", "comma separated list of interfaces to listen on")
	pflag.StringP("filter", "f", "icmp", "bpf filter")
	pflag.StringP("target", "t", "", "target ip, to monitor")

	// parse the pflags
	pflag.Parse()

	// bind the pflags
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
}

func main() {
	for _, intrfc := range getInterfaces() {
		handle, err := pcap.OpenLive(intrfc, 1600, false, pcap.BlockForever)

		if err != nil {
			panic(err)
		}

		err = handle.SetBPFFilter(getFilter())

		if err != nil {
			panic(err)
		}

		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		go startListening(intrfc, packetSource)
	}

	for {
		time.Sleep(60 * time.Second)
	}
}

func getInterfaces() []string {
	intrfcsString := viper.GetString("interfaces")

	if intrfcsString == "" {
		panic("no interface given")
	}

	return strings.Split(intrfcsString, ",")
}

func getFilter() string {
	filter := viper.GetString("filter")
	target := viper.GetString("target")

	if target != "" {
		filter = fmt.Sprintf(
			"(%s) and (dst host %s)",
			filter,
			target,
		)
	}

	return filter
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
