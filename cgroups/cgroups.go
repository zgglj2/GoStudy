//go:build linux
// +build linux

package main

import (
	"fmt"

	"github.com/containerd/cgroups"
	cgroupsv2 "github.com/containerd/cgroups/v2"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/shirou/gopsutil/process"
)

func GetProcessesPidByPath(path string) (pids []int32, err error) {
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("get Processes failed, err:%v\n", err)
		return nil, err
	}
	for _, process := range processes {
		processPath, err := process.Exe()
		if err != nil {
			continue
		}

		if processPath == path {
			pids = append(pids, process.Pid)
		}

	}

	return pids, nil
}
func main() {
	// shares := uint64(100)
	control1, err := cgroups.Load(cgroups.V1, cgroups.StaticPath("/test"))
	if err == nil {
		fmt.Printf("%v\n", control1)
		control1.Delete()
	}
	quota := int64(10000)
	memLimit := int64(300 * 1024 * 1024)
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath("/test"), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota: &quota,
		},
		Memory: &specs.LinuxMemory{
			Limit: &memLimit,
		},
	})

	// got error when creating new cgroup with systemd
	// control2, err := cgroups.Load(cgroups.Systemd, cgroups.Slice("system.slice", "sdcgroup.slice"))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n", control2)
	// defer control2.Delete()

	// control, err := cgroups.New(cgroups.Systemd, cgroups.Slice("system.slice", "sdcgroup1.slice"), &specs.LinuxResources{
	// 	CPU: &specs.LinuxCPU{
	// 		Shares: &shares,
	// 	},
	// })

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", control)

	// cgroups dir will not delete until Delete() called
	defer control.Delete()

	// get process pid by path
	pids, err := GetProcessesPidByPath("/usr/bin/stress")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("pids: %v\n", pids)
	for _, pid := range pids {
		if err := control.Add(cgroups.Process{Pid: int(pid)}); err != nil {
			fmt.Println(err)
			continue
		}
	}
	// shares = uint64(200)
	// if err := control.Update(&specs.LinuxResources{
	// 	CPU: &specs.LinuxCPU{
	// 		Shares: &shares,
	// 	},
	// }); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	quota = int64(20000)
	memLimit = int64(100 * 1024 * 1024)
	err = control.Update(&specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota: &quota,
		},
		Memory: &specs.LinuxMemory{
			Limit: &memLimit,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("freeze")
	if err := control.Freeze(); err != nil {
		fmt.Println(err)
		return
	}

	// time.Sleep(time.Second * 20)

	fmt.Println("Thaw")
	if err := control.Thaw(); err != nil {
		fmt.Println(err)
		return
	}
	processes, err := control.Processes(cgroups.Devices, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", processes)

	stats, err := control.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", stats)

	stats, err = control.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", stats)

	// err = control.MoveTo(destination)

	// event := cgroups.MemoryThresholdEvent(50*1024*1024, false)
	// efd, err := control.RegisterMemoryEvent(event)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v\n", efd)

	// event = cgroups.MemoryPressureEvent(cgroups.MediumPressure, cgroups.DefaultMode)
	// efd, err = control.RegisterMemoryEvent(event)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v\n", efd)

	// efd, err = control.OOMEventFD()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v\n", efd)

	// // or by using RegisterMemoryEvent
	// event = cgroups.OOMEvent()
	// efd, err = control.RegisterMemoryEvent(event)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v\n", efd)

	var cgroupV2 bool
	if cgroups.Mode() == cgroups.Unified {
		cgroupV2 = true
	}
	fmt.Printf("cgroupV2: %v\n", cgroupV2)

	if cgroupV2 {
		res := cgroupsv2.Resources{}
		// dummy PID of -1 is used for creating a "general slice" to be used as a parent cgroup.
		// see https://github.com/containerd/cgroups/blob/1df78138f1e1e6ee593db155c6b369466f577651/v2/manager.go#L732-L735
		m, err := cgroupsv2.NewSystemd("/", "my-cgroup-abc.slice", -1, &res)
		if err != nil {
			return
		}
		err = m.DeleteSystemd()
		if err != nil {
			return
		}
		m1, err := cgroupsv2.LoadSystemd("/", "my-cgroup-abc.slice")
		if err != nil {
			return
		}
		err = m1.DeleteSystemd()
		if err != nil {
			return
		}
	}
}
