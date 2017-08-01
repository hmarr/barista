package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Schedule map[string]TimeRange

//    Monday    string `json:"monday"`
//    Tuesday   string `json:"tuesday"`
//    Wednesday string `json:"wednesday"`
//    Thursday  string `json:"thursday"`
//    Friday    string `json:"friday"`
//    Saturday  string `json:"saturday"`
//    Sunday    string `json:"sunday"`
//}

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
	schedule, err := loadSchedule("config.json")
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
