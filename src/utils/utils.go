package utils

import (
	"net"
	"os"
	"strings"
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

		flags := netInterface.Flags.String()

		// get the interface adapter

		address, _ := netInterface.Addrs()

		// check if the address is ipv4
		for _, addr := range address {
			if ipnet, ok := addr.(*net.IPNet); ok &&
				!ipnet.IP.IsLoopback() &&
				// check if up flag is set
				strings.Contains(flags, "up") &&
				!strings.Contains(ipnet.IP.String(), "169.254") &&
				!strings.Contains(interfaceName, "VirtualBox") &&
				!strings.Contains(interfaceName, "Virtual") {
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
