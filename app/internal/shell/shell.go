package shell

import "sync"

type Shell struct {
	Jobs      []*Job
	NextJobID int
	mu        sync.Mutex
}

func New() *Shell {
	return &Shell{
		Jobs:      []*Job{},
		NextJobID: 1,
	}
}
