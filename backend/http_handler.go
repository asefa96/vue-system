package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func getInfos(w http.ResponseWriter, r *http.Request) {
	var info SystemInfo
	hostInfo, err := host.Info()

	if err != nil {
		fmt.Println(err)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}

	v, _ := mem.VirtualMemory()

	info.HostName, _ = os.Hostname()
	info.SystemOS = hostInfo.Platform + " " + hostInfo.PlatformVersion
	info.HostIP = getMyIP().String()
	info.Processors = cpuInfo[0].ModelName
	info.DiskUsed = v.UsedPercent

	data, err := json.Marshal(info)
	if err != nil {
		return
	}
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

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
			}
		}
	}
	return addr
}
