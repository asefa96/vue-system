package main

import (
	"fmt"
	"net"
	"os"

	. "../models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {

	var info SystemInfo
	hostInfo, err := host.Info()

	if err != nil {
		fmt.Println(err)
	}

	cpuInfo, err := cpu.Info()

	v, _ := mem.VirtualMemory()

	info.HostName, _ = os.Hostname()
	info.SystemOS = hostInfo.Platform + " " + hostInfo.PlatformVersion
	info.HostIP = getMyIP().String()
	info.Processors = cpuInfo[0].ModelName

	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Println(info)

}

func getMyIP() net.IP {
	var addr net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				address := ipnet.IP.To4()
				addr = address
				// os.Stdout.WriteString(ipnet.IP.String() + "\n")
				//

			}
		}
	}
	return addr
}
