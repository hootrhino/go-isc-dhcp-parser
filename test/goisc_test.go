package main

import (
	"github.com/hootrhino/go-isc-dhcp/iscdhcp"
	"testing"
)

// go test -timeout 30s -run ^Test_parse_isc_conf github.com/hootrhino/go-isc-dhcp/test -v -count=1
func Test_parse_isc_conf(t *testing.T) {
	dhcp := iscdhcp.NewDHCP("./iscdhcp.conf")
	err := dhcp.ReloadConfig()
	if err != nil {
		panic(err)
	}
	err = dhcp.Start()
	if err != nil {
		panic(err)
	}
}

/*
import (
	"encoding/json"
	"fmt"
	"github.com/xaionaro-go/iscDhcp/cfg"
	"os"
)

func main() {
	cfg := cfg.NewConfig()
	err := cfg.LoadFrom("/etc/dhcp/dhcpd.conf")
	if err != nil {
		panic(err)
	}

	jsonEncoder := json.NewEncoder(os.Stderr)
	jsonEncoder.SetIndent("", "  ")
	jsonEncoder.Encode(cfg)

	fmt.Printf("\n\n--- Regenerating the configuration:\n\n")

	cfg.ConfigWrite(os.Stdout)
}
*/
