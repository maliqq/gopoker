package engine

import "testing"

func TestStages(t *testing.T) {

	var run1, run2, run3 bool
	stages := Stages{
		StageDo{
			Do: func() {
				run1 = true
			},
		},
		StageSkip{
			Do: func(skip chan bool) {
				run2 = true
				close(skip)
			},
		},
		StageDo{
			Do: func() {
				run3 = true
			},
		},
	}
	stages.Run()
	if !(run1 && run2 && run3) {
		t.Fatalf("run1=%t run2=%t run3=%t", run1, run2, run3)
	}
	run1, run2, run3 = false, false, false

	stagesSkip := Stages{
		StageDo{
			Do: func() {
				run1 = true
			},
		},
		StageSkip{
			Do: func(skip chan bool) {
				run2 = true
				skip <- true
			},
		},
		StageDo{
			Do: func() {
				run3 = true
			},
		},
	}
	stagesSkip.Run()
	if run3 {
		t.Fatalf("run3 should be skipped!")
	}

}
