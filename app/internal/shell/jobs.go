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

func (s *Shell) Job() {

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

	}
	

}

func (s *Shell) ReapJobs() {

	var doneIndexes []int
	jobCount := len(s.Jobs)
	// fmt.Println("before reap:", len(s.Jobs))

	for i, job := range s.Jobs {
		if job.Status == "Done" {
			doneIndexes = append(doneIndexes, i)
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
	// fmt.Println("removing:", doneIndexes)
	for i := len(doneIndexes) - 1; i >= 0; i-- {
		idx := doneIndexes[i]
		s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
	}
	// fmt.Println("after reap:", len(s.Jobs))

}

// func (s *Shell) ReapJobs() {
//     var doneIndexes []int

//     for i, job := range s.Jobs {
//         if job.Status == "Done" {
// 			fmt.Printf(
// 				"[%d]  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)
//             doneIndexes = append(doneIndexes, i)
//         }
//     }

//     for i := len(doneIndexes)-1; i >= 0; i-- {
//         idx := doneIndexes[i]
//         s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
//     }
// }
