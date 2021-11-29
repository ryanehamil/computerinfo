package utils

import (
	"net"
	"os"
)

func GetComputerName() string {
	computername := ""
	// Get computer name from os package
	computername = os.Getenv("COMPUTERNAME")
	return computername
}
func GetIPAddress() []string {
	// ipaddress := ""
	// Get all interfaces from net package
	netInterfaces, _ := net.Interfaces()

	// Get all interface names and addresses from each interface
	var addrs []string
	for _, netInterface := range netInterfaces {
		// get the interface name
		interfaceName := netInterface.Name

		// get the interface adapter
		

		address, _ := netInterface.Addrs()

		// check if the address is ipv4
		for _, addr := range address {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.String()[0:7] != "169.254" {
				if ipnet.IP.To4() != nil {
					addrs = append(addrs, interfaceName+": "+ipnet.IP.String())
				}
			}
		}

	}

	return addrs
}

func ConvertIPAddress(addrs []net.Addr) []string {
	var ipaddress []string
	for _, addr := range addrs {
		ipaddress = append(ipaddress, addr.String())
	}
	return ipaddress
}
