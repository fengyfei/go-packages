package runner

import (
	"runtime"
	"testing"
)

func TestChanCloserRunner_Go(t *testing.T) {
	var (
		chanCloserRunner1  = NewChanCloserRunner()
		chanCloserRunner2  = NewChanCloserRunner()
		chanCloseGoRoutine chanCloseGoRoutine
		goRoutineNum       = runtime.NumGoroutine()
	)

	chanCloseGoRoutine = func(<-chan bool) {
		t.Log("start goroutine")
	}

	err := chanCloserRunner1.Go(chanCloseGoRoutine)
	if err != nil {
		t.Error("ChanCloserRunner_Go starts a goroutine error")
	}
	if runtime.NumGoroutine()-goRoutineNum != 1 {
		t.Error("ChanCloserRunner_Go starts a goroutine failed")
	}

	err = chanCloserRunner1.Go(chanCloseGoRoutine)
	if err != nil {
		t.Error("ChanCloserRunner_Go starts a goroutine error")
	}
	if runtime.NumGoroutine()-goRoutineNum != 2 {
		t.Error("ChanCloserRunner_Go starts a goroutine failed")
	}
	chanCloserRunner1.Wait()

	/* ----- */

	chanCloserRunner2.closed = true

	err = chanCloserRunner2.Go(chanCloseGoRoutine)
	if err != errRunnerClosed {
		t.Error("runner closed catched unright error: ", err)
	}
}

func TestChanCloserRunner_Interrupt(t *testing.T) {
	var (
		chanCloserRunner   = NewChanCloserRunner()
		chanCloseGoRoutine chanCloseGoRoutine
		goRoutineNum       = runtime.NumGoroutine()
	)

	chanCloseGoRoutine = func(<-chan bool) {
		t.Log("start goroutine")
	}

	err := chanCloserRunner.Go(chanCloseGoRoutine)
	if err == nil && runtime.NumGoroutine()-goRoutineNum == 1 {
		chanCloserRunner.Interrupt()
		if chanCloserRunner.closed != true {
			t.Error("ChanCloserRunner after interrupt closed is not true")
		}
		if runtime.NumGoroutine() != goRoutineNum {
			t.Error("ChanCloserRunner after interrupt goroutines is not close")
		}
	} else {
		t.Error("ChanCloserRunner_Go starts a goroutine error")
		if runtime.NumGoroutine()-goRoutineNum != 1 {
			t.Error("ChanCloserRunner_Go starts a goroutine failed")
		}
	}
}
