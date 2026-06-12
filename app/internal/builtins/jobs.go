package builtins





func (s *Shell) jobs(args []string) {
	for _, job := range s.Jobs {
		fmt.Printf(
			"[%d]+  %-24s%s\n",
			job.ID,
			job.Status,
			job.Command,
		)
	}
}