package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/serialt/crab"
)

func service() {
	slog.Debug("debug msg")
	slog.Info("info msg")
	slog.Error("error msg")

	for _, hw := range config.Huaweicloud {
		for _, region := range hw.Region {
			fmt.Printf("\n\n[ %s    %s]\n", hw.Name, region)
			ListServer(hw, region)
			ListEIP(hw, region)
		}

	}

}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}

func (c *Config) DecryptConfig() {
	if c.Encrypt {
		crab.AESDecryptCBCBase64(c.Token, AesKey)
		slog.Debug(c.Token)
	}
}

func ListServer(hw Huaweicloud, region string) {
	_ecsClient := NewECSClient(hw, region)
	servers := HWListServers(_ecsClient)
	fmt.Printf("%-25s\t %-18s\t %-18s  %-22s\t\t\n", "ECSName", "IP", "Pub IP", "OS Type")
	for _, server := range *servers {
		pubIP := ""
		if len(server.Addresses[server.Metadata["vpc_id"]]) > 1 {
			pubIP = server.Addresses[server.Metadata["vpc_id"]][1].Addr
		}
		fmt.Printf("%-25s\t %-18s\t %-18s  %-20s\t\n", server.Name, server.Addresses[server.Metadata["vpc_id"]][0].Addr, pubIP, server.Metadata["image_name"])
	}
}

func ListEIP(hw Huaweicloud, region string) {
	_eipClient := NewEIPClient(hw, region)
	pubips := HWEIPListIP(_eipClient)
	fmt.Printf("\n%-25s\t %-18s\t %-18s\t \n", "EIPName", "Pub IP", "IP")
	for _, ip := range *pubips {
		fmt.Printf("%-25s\t %-18s\t %-18s\t\n", *ip.BandwidthName, *ip.PublicIpAddress, *ip.PrivateIpAddress)
	}

}
