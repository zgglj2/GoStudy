package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"

	// "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/host"
)

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for i := 0; i < 1; i++ {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}
func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

// mem info
func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

// host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
	KernelVersion, _ := host.KernelVersion()
	Platform, PlatformFamily, PlatformVersion, _ := host.PlatformInformation()
	Arch, _ := host.KernelArch()

	fmt.Printf("KernelVersion:%v Platform:%v PlatformFamily:%v PlatformVersion:%v Arch:%v\n", KernelVersion, Platform, PlatformFamily, PlatformVersion, Arch)
}

// disk info
func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

func getNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

// func GetLocalIP() (ip string, err error) {
// 	addrs, err := net.InterfaceAddrs()
// 	net.Get
// 	if err != nil {
// 		return
// 	}
// 	for _, addr := range addrs {
// 		ipAddr, ok := addr.(*net.IPNet)
// 		if !ok {
// 			continue
// 		}
// 		if ipAddr.IP.IsLoopback() {
// 			continue
// 		}
// 		if !ipAddr.IP.IsGlobalUnicast() {
// 			continue
// 		}
// 		return ipAddr.IP.String(), nil
// 	}
// 	return
// }

// // Get preferred outbound ip of this machine
// func GetOutboundIP() string {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	localAddr := conn.LocalAddr().(*net.UDPAddr)
// 	fmt.Println(localAddr.String())
// 	return localAddr.IP.String()
// }

func getProcessesInfo() {
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("get Processes failed, err:%v\n", err)
		return
	}
	for _, process := range processes {
		fmt.Println(process)
		fmt.Println(process.Name())
		fmt.Println(process.Exe())
	}
}

func main() {
	getCpuInfo()
	getCpuLoad()
	getMemInfo()
	getHostInfo()
	getDiskInfo()
	getNetInfo()
	getProcessesInfo()
}
