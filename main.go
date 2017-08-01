package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

type Schedule map[string]TimeRange

func loadSchedule(path string) (Schedule, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	schedule := Schedule{}
	err = json.Unmarshal(data, &schedule)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configPath := path.Join(usr.HomeDir, ".config/barista/schedule.json")
	schedule, err := loadSchedule(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sup := NewSupervisor(schedule)
	if err = sup.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
