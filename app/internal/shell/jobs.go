package shell

import (
	"fmt"
)

type Job struct {
	ID      int
	PID     int
	Command string
	Status  string
}

func (s *Shell) jobs(args []string) {
	var doneIndexes []int

	// for i:=0; i < len(s.Jobs)-1; i++ {
	// 	if job.Status == "Done" {
	// 		fmt.Printf(
	// 			"[%d]+  %-24s%s\n",
	// 			job.ID,
	// 			job.Status,
	// 			job.Command,
	// 		)
	// 	}

	// }
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
			fmt.Printf(
				"[%d]+  %-24s%s\n",
				job.ID,
				job.Status,
				job.Command,
			)
        }

	}
	for i := len(doneIndexes) - 1; i >= 0; i-- {
        idx := doneIndexes[i]
        s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
    }
}
