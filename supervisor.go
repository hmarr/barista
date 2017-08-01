package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const tickInterval = 5 * time.Second

type Supervisor struct {
	Schedule Schedule
	Process  *os.Process
	sync.Mutex
}

func NewSupervisor(schedule Schedule) *Supervisor {
	return &Supervisor{Schedule: schedule}
}

func (s *Supervisor) Start() error {
	for {
		if err := s.tick(); err != nil {
			return err
		}
		time.Sleep(tickInterval)
	}
}

func (s *Supervisor) tick() error {
	if s.withinSchedule() {
		if err := s.ensureRunning(); err != nil {
			return err
		}
	} else {
		if err := s.ensureNotRunning(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Supervisor) withinSchedule() bool {
	now := time.Now()
	weekday := strings.ToLower(now.Weekday().String())

	timeRange, ok := s.Schedule[weekday]
	if !ok {
		return false
	}

	year, month, day := now.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	start := startOfDay.Add(timeRange.Start)
	if now.Before(start) {
		return false
	}

	end := startOfDay.Add(timeRange.End)
	if now.After(end) {
		return false
	}

	return true
}

func (s *Supervisor) ensureRunning() error {
	if s.Process != nil {
		return nil
	}

	log.Println("Caffeinating")

	s.Lock()
	defer s.Unlock()

	cmd := exec.Command("caffeinate", "-i")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	go s.monitorCommand(cmd)

	s.Process = cmd.Process
	return nil
}

func (s *Supervisor) monitorCommand(cmd *exec.Cmd) {
	cmd.Wait()

	s.Lock()
	defer s.Unlock()

	s.Process = nil
}

func (s *Supervisor) ensureNotRunning() error {
	if s.Process == nil {
		return nil
	}

	log.Println("Decaffeinating")
	return s.Process.Kill()
}
