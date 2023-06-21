/*
READY->EXECUTING->WAITING
EXECUTING->PREEMPTED->Ready

Ready -0
Executing -1
Waiting 2
*/

package main

import "fmt"

type ThreadState int

const (
	NONE ThreadState = iota
	READY
	EXECUTING
	WAITING
)

const (
	none  = 1
	Ready = iota
	Executing
	Waiting
)

func (t *ThreadState) Update(newState ThreadState) bool {
	switch *t {
	case NONE:
		if newState == READY {
			*t = newState
			return true
		}
	case READY:
		if newState == EXECUTING {
			*t = newState
			return true

		}
	case EXECUTING:
		if newState == WAITING || newState == READY {
			*t = newState
			return true
		}
	case WAITING:
		if newState == EXECUTING {
			*t = newState
			return true
		}
	}
	return false
}

func (t ThreadState) String() string {
	fmt.Println("Calling String")
	switch t {
	case NONE:
		return "None"
	case READY:
		return "Ready"
	case EXECUTING:
		return "Executing"
	case WAITING:
		return "Waiting"
	default:
		return "Invalid"
	}
}

func main() {
	fmt.Println("Defining enumeration")
	var tstate ThreadState = NONE

	res := tstate.Update(WAITING)
	fmt.Println(res)

	res = tstate.Update(READY)
	fmt.Println(res)

	fmt.Println(tstate)
}
