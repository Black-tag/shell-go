package shell

import "fmt"




type Job struct {
	ID int
	PID int
	Command string
	Status string

}


func (s *Shell)jobs(args []string) {
	for _, job := range s.Jobs {
		fmt.Printf(
			"[%d]+  %-24s%s%s\n",
			job.ID,
			job.Status,
			job.Command,
			args
		)
	}
}

