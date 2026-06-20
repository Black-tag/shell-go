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

// func (s *Shell) Job() {

// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	// fmt.Println("FROM JOB")
// 	// var doneIndexes []int

// 	// // s.ReapJobs()
// 	var doneIndexes []int
//     for i := range s.Jobs {
//         if s.Jobs[i].Status == "Done" {
//             doneIndexes = append(doneIndexes, i)
//         }
//     }
//     for i := len(doneIndexes) - 1; i >= 0; i-- {
//         idx := doneIndexes[i]
//         s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
//     }
// 	jobCount := len(s.Jobs)
// 	// for i := len(doneIndexes) - 1; i >= 0; i-- {
// 	// 	idx := doneIndexes[i]
// 	// 	s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
// 	// }

// 	for i, job := range s.Jobs {
// 		// if job.Status == "Done" {
// 		// 	doneIndexes = append(doneIndexes, i)
// 		// }

// 		switch i {
// 		case jobCount - 1:
// 			fmt.Printf(
// 				"[%d]+  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)

// 		case jobCount - 2:
// 			fmt.Printf(
// 				"[%d]-  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)

// 		default:
// 			fmt.Printf(
// 				"[%d]  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)

// 		}

// 	}

// }

// func (s *Shell) ReapJobs() {

// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	// fmt.Println("FROM REAPJOB")

// 	var doneIndexes []int
// 	jobCount := len(s.Jobs)
// 	// fmt.Println("before reap:", len(s.Jobs))

// 	for i:= range s.Jobs {
// 		// fmt.Println("REAPING", job.ID)
// 		if s.Jobs[i].Status == "Running" &&
// 			s.Jobs[i].Cmd.ProcessState != nil &&
// 			s.Jobs[i].Cmd.ProcessState.Exited() {

// 			s.Jobs[i].Status = "Done"
// 			s.Jobs[i].Command = strings.TrimSuffix(s.Jobs[i].Command, " &")
// 			// fmt.Print("printing done inside 1")
// 		}
// 		if s.Jobs[i].Status == "Done" {
// 			doneIndexes = append(doneIndexes, i)
// 			switch i {
// 			case jobCount - 1:
// 				// fmt.Print("printintg inside latest ")
// 				fmt.Printf(
// 					"[%d]+  %-24s%s\n",
// 					s.Jobs[i].ID,
// 					s.Jobs[i].Status,
// 					s.Jobs[i].Command,
// 				)

// 			case jobCount - 2:
// 				// fmt.Print("printintg inside second last ")
// 				fmt.Printf(
// 					"[%d]-  %-24s%s\n",
// 					s.Jobs[i].ID,
// 					s.Jobs[i].Status,
// 					s.Jobs[i].Command,
// 				)

// 			default:
// 				// fmt.Print("printintg inside defualt")
// 				fmt.Printf(
// 					"[%d]  %-24s%s\n",
// 					s.Jobs[i].ID,
// 					s.Jobs[i].Status,
// 					s.Jobs[i].Command,
// 				)

// 			}

// 		}

// 	}

// 	// fmt.Println("removing:", doneIndexes)

// 	// fmt.Println("after reap:", len(s.Jobs))
// 	// s.mu.Unlock()
// }

// func (s *Shell) ReapJobs() {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	var doneIndexes []int

// 	for i := range s.Jobs {
// 		job := s.Jobs[i]

// 		if job.Status == "Running" &&
// 			job.Cmd.ProcessState != nil &&
// 			job.Cmd.ProcessState.Exited() {

// 			job.Status = "Done"
// 			job.Command = strings.TrimSuffix(job.Command, " &")

// 			doneIndexes = append(doneIndexes, i)
// 		}
// 	}

// 	jobCount := len(s.Jobs)

// 	for _, idx := range doneIndexes {
// 		job := s.Jobs[idx]

// 		switch idx {
// 		case jobCount - 1:
// 			fmt.Printf(
// 				"[%d]+  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)

// 		case jobCount - 2:
// 			fmt.Printf(
// 				"[%d]-  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)

// 		default:
// 			fmt.Printf(
// 				"[%d]  %-24s%s\n",
// 				job.ID,
// 				job.Status,
// 				job.Command,
// 			)
// 		}
// 	}

// 	for i := len(doneIndexes) - 1; i >= 0; i-- {
// 		idx := doneIndexes[i]
// 		s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
// 	}
// }

// func (s *Shell) ReapJobs() {
//     s.mu.Lock()
//     defer s.mu.Unlock()

//     var doneIndexes []int
//     jobCount := len(s.Jobs)

//     for i := range s.Jobs {
//         if s.Jobs[i].Status == "Done" {
//             doneIndexes = append(doneIndexes, i)

//             switch i {
//             case jobCount - 1:
//                 fmt.Printf("[%d]+  %-24s%s\n", s.Jobs[i].ID, s.Jobs[i].Status, s.Jobs[i].Command)
//             case jobCount - 2:
//                 fmt.Printf("[%d]-  %-24s%s\n", s.Jobs[i].ID, s.Jobs[i].Status, s.Jobs[i].Command)
//             default:
//                 fmt.Printf("[%d]   %-24s%s\n", s.Jobs[i].ID, s.Jobs[i].Status, s.Jobs[i].Command)
//             }
//         }
//     }

//     for i := len(doneIndexes) - 1; i >= 0; i-- {
//         idx := doneIndexes[i]
//         s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
//     }
// }

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
