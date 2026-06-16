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
		if job.ID == len(s.Jobs)-1 {
			fmt.Printf(
			"[%d]+  %-24s%s\n",
			job.ID,
			job.Status,
			job.Command,
			)
		}else if job.ID == len(s.Jobs)-2 {
			fmt.Printf(
			"[%d]-  %-24s%s\n",
			job.ID,
			job.Status,
			job.Command,
			)

		}else {
			fmt.Printf(
			"[%d]  %-24s%s\n",
			job.ID,
			job.Status,
			job.Command,
			)

		}
		
		
	}
}

