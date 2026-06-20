package shell

import (
	"fmt"
	"os/exec"
	"strings"
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
	// fmt.Println("FROM JOB")
	var doneIndexes []int

	// s.ReapJobs()
	jobCount := len(s.Jobs)

	for i, job := range s.Jobs {

		switch i {
		case jobCount - 1:
			fmt.Printf(
				"[%d]+  %-24s%s\n",
				job.ID,
				job.Status,
				job.Command,
			)

		case jobCount - 2:
			fmt.Printf(
				"[%d]-  %-24s%s\n",
				job.ID,
				job.Status,
				job.Command,
			)

		default:
			fmt.Printf(
				"[%d]  %-24s%s\n",
				job.ID,
				job.Status,
				job.Command,
			)

		}
		if job.Status == "Done" {
			doneIndexes = append(doneIndexes, i)
		}

	}

}

func (s *Shell) ReapJobs() {

	s.mu.Lock()
	defer s.mu.Unlock()
	// fmt.Println("FROM REAPJOB")

	var doneIndexes []int
	jobCount := len(s.Jobs)
	// fmt.Println("before reap:", len(s.Jobs))

	for i, job := range s.Jobs {
		// fmt.Println("REAPING", job.ID)
		if job.Status == "Running" &&
			job.Cmd.ProcessState != nil &&
			job.Cmd.ProcessState.Exited() {

			job.Status = "Done"
			job.Command = strings.TrimSuffix(job.Command, " &")
			fmt.Print("printing done inside 1")
		}
		if job.Status == "Done" {
			doneIndexes = append(doneIndexes, i)
			switch i {
			case jobCount - 1:
				fmt.Print("printintg inside latest ")
				fmt.Printf(
					"[%d]+  %-24s%s\n",
					job.ID,
					job.Status,
					job.Command,
				)

			case jobCount - 2:
				// fmt.Print("printintg inside second last ")
				fmt.Printf(
					"[%d]-  %-24s%s\n",
					job.ID,
					job.Status,
					job.Command,
				)

			default:
				// fmt.Print("printintg inside defualt")
				fmt.Printf(
					"[%d]  %-24s%s\n",
					job.ID,
					job.Status,
					job.Command,
				)

			}

		}

	}

	// fmt.Println("removing:", doneIndexes)
	for i := len(doneIndexes) - 1; i >= 0; i-- {
		idx := doneIndexes[i]
		s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
	}
	// fmt.Println("after reap:", len(s.Jobs))
	// s.mu.Unlock()
}
