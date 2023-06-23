package timer

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	utl "github.com/emirpasic/gods/utils"
)

type TimerCallBk func(timerID string, interval int)

type timer struct {
	timerId   string
	ineterval int
	callback  TimerCallBk
	expiry    int64
}

var (
	timerTree      = rbt.NewWith(utl.StringComparator)
	timerIndexTree = rbt.NewWith(utl.Int64Comparator)
	yield          = make(chan int)
	exit           = make(chan int)
)

func Init() {
	execute()
}

func DeInit() {

}

func StartTimer(interval int, cb TimerCallBk) (timerID string) {
	timerID = uuid.Must(uuid.NewRandom()).String()
	tExpiry := time.Now().Add(time.Duration(interval * int(time.Second))).Unix()
	tm := timer{timerId: timerID, ineterval: interval, callback: cb, expiry: tExpiry}
	timerTree.Put(timerID, tm)
	timerIndexTree.Put(tExpiry, timerID)
	yield <- 0
	return timerID
}

func CancelTimer(timerID string) {
	//Search for timerID idx
	nd := timerTree.GetNode(timerID)
	tm := nd.Value.(timer)
	fmt.Println("", tm)
	timerTree.Remove(timerID)
	timerIndexTree.Remove(tm.expiry)
	yield <- 1
}

func handleExpiry(tmId string) {
	nd := timerTree.GetNode(tmId)
	tm := nd.Value.(timer)
	tm.callback(tm.timerId, tm.ineterval)
	timerIndexTree.Remove(tm.expiry)
	timerTree.Remove(tm.timerId)
}

func execute() {
	go func() {
		<-yield
		for {
			nd := timerIndexTree.Left()
			if nd == nil {
				<-yield
				nd = timerIndexTree.Left()
			}
			timer := time.NewTimer(time.Until(time.Unix(nd.Key.(int64), 0)))
			select {
			case <-timer.C:
				handleExpiry(nd.Value.(string))
			case <-yield:
				fmt.Println("Yielding...")
			}
		}
	}()
}
