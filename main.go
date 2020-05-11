package main

import (
	"fmt"
	"github.com/elastic/gosigar"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func main() {
	global := gosigar.Cpu{}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}

	times, _ := cpu.Times(false)
	global.Get()
	fmt.Println(times[0].User, times[0].System, times[0].Nice, times[0].Iowait, times[0].Idle)
	fmt.Println(global.User, global.Sys, global.Nice, global.Wait)
}
