package shell



type Shell struct{
	Jobs []*Job
	NextJobID int
}

func New() *Shell {
	return &Shell{
		Jobs: []*Job{},
		NextJobID: 1,

	}
}
