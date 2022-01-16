package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"github.com/kardianos/service"
)

var (
	installBool   *bool
	uninstallBool *bool

	InstallPath string
	M           *sync.Mutex
	Cmd         *exec.Cmd
	Service     service.Service
	AgentStatus bool
)

func init() {
	InstallPath = "../agent/"
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	var agentFilePath string
	if runtime.GOOS == "windows" {
		agentFilePath = InstallPath + "agent.exe"
	} else {
		agentFilePath = InstallPath + "agent"
	}
	for {
		M.Lock()
		log.Println("Start Agent")
		Cmd = exec.Command(agentFilePath, "--nousearg")
		err := Cmd.Start()
		M.Unlock()
		if err == nil {
			AgentStatus = true
			log.Println("Start Agent successful")
			err = Cmd.Wait()
			if err != nil {
				AgentStatus = false
				log.Println("Agent to exit：", err.Error())
			}
		} else {
			log.Println("Startup Agent failed", err.Error())
		}
		time.Sleep(time.Second * 10)
	}
}

func (p *program) Stop(s service.Service) error {
	return KillAgent()
}

func KillAgent() error {
	if AgentStatus {
		return Cmd.Process.Kill()
	}
	return nil
}

func main() {
	installBool = flag.Bool("install", false, "Install service")
	uninstallBool = flag.Bool("uninstall", false, "Remove service")
	flag.Parse()
	svcConfig := &service.Config{
		Name:        "agent-svc",
		DisplayName: "agent-svc",
		Description: "agent service",
		Arguments:   []string{"-nousearg"},
	}
	prg := &program{}
	var err error
	Service, err = service.New(prg, svcConfig)
	if err != nil {
		log.Println("New a service error:", err.Error())
		return
	}
	if *uninstallBool {
		fmt.Println("start to uninstall service")
		KillAgent()
		if err := Service.Uninstall(); err != nil {
			log.Println("Uninstall error:", err.Error())
		}
		fmt.Println("end to uninstall service")
		return
	}
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		return
	}

	if *installBool {
		fmt.Println("start to install agent")
		err = Service.Install()
		if err != nil {
			log.Println("Install daemon as service error:", err.Error())
		} else {
			if err = Service.Start(); err != nil {
				log.Println("Service start error:", err.Error())
			} else {
				log.Println("Install as a service", "ok")
			}
		}
		fmt.Println("end to install agent")
		return
	}
	err = Service.Run()
	if err != nil {
		log.Println("Service run error:", err.Error())
	}
}
