package utils

import "os"

func GetComputerName() string {
	computername := ""
	// Get computer name from os package
	computername = os.Getenv("COMPUTERNAME")
	return computername
}
