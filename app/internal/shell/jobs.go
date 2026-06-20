package shell

import (
	"fmt"
	"os/exec"
	// "strings"
)

type Job struct {
	ID      int
	PID     int
	Command string
	Status  string
	Cmd     exec.Cmd
}


func (s *Shell) Job() {
	s.mu.Lock()
	defer s.mu.Unlock()

	jobCount := len(s.Jobs)
	var doneIndexes []int

	for i, job := range s.Jobs {
		fmt.Printf("[%d]%s  %-24s%s\n", job.ID, jobMarker(i, jobCount), job.Status, job.Command)
		if job.Status == "Done" {
			doneIndexes = append(doneIndexes, i)
		}
	}

	for i := len(doneIndexes) - 1; i >= 0; i-- {
		idx := doneIndexes[i]
		s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
	}
}

func (s *Shell) ReapJobs() {
	s.mu.Lock()
	defer s.mu.Unlock()

	var doneIndexes []int
	jobCount := len(s.Jobs)

	for i := range s.Jobs {
		if s.Jobs[i].Status == "Done" {
			doneIndexes = append(doneIndexes, i)
			fmt.Printf("[%d]%s  %-24s%s\n", s.Jobs[i].ID, jobMarker(i, jobCount), s.Jobs[i].Status, s.Jobs[i].Command)
		}
	}

	for i := len(doneIndexes) - 1; i >= 0; i-- {
		idx := doneIndexes[i]
		s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
	}
}

func jobMarker(index, jobCount int) string {
	if index == jobCount-1 {
		return "+"
	}
	if index == jobCount-2 {
		return "-"
	}
	return " "
}


func (s *Shell) getNextJobID() int {
    used := make(map[int]bool)

    for _, job := range s.Jobs {
        used[job.ID] = true
    }

    id := 1
    for used[id] {
        id++
    }

    return id
}