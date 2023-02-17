package config

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitializeCommandFlags() {
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

func Interfaces() []string {
	intrfcsString := viper.GetString("interfaces")

	if intrfcsString == "" {
		panic("no interface given")
	}

	return strings.Split(intrfcsString, ",")
}

func Filter() string {
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
