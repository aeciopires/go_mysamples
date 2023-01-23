package main

// Reference: https://github.com/zcalusic/sysinfo#readme

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

func main() {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("requires superuser privilege")
	}

	var informationSystem sysinfo.SysInfo

	informationSystem.GetSysInfo()

	data, err := json.MarshalIndent(&informationSystem, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
