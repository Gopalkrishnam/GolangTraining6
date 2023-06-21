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
	NONE      ThreadState = -1
	READY     ThreadState = 0
	EXECUTING ThreadState = 1
	WAITING   ThreadState = 2
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

func main() {
	fmt.Println("Defining enumeration")
	var tstate ThreadState = NONE

	res := tstate.Update(WAITING)
	fmt.Println(res)

	res = tstate.Update(READY)
	fmt.Println(res)
}
