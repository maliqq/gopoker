package engine

type StageProcess struct {
	Name string
	If     func() bool
	Before func()
	Run    func()
	After  func()
	Notify bool
}

type StageProcesses []StageProcess

func (procs StageProcesses) Process() {
	for _, proc := range procs {
		proc.Process()
	}
}

func (proc *StageProcess) Process() {
	// stage condition
	if proc.If != nil && !proc.If() {
		return
	}

	if proc.Before != nil {
		proc.Before()
	}

	if proc.Run != nil {
		proc.Run()
	}

	if proc.After != nil {
		proc.After()
	}
}
