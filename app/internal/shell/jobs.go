package shell

import "fmt"

type Job struct {
	ID      int
	PID     int
	Command string
	Status  string
}

func (s *Shell) jobs(args []string) {
	for i, job := range s.Jobs {
		jobCount := len(s.Jobs)

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

	}
}
