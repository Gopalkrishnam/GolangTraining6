package timer

import (
	"errors"
	"fmt"
	"time"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	"github.com/google/uuid"
)

type TimerCB func(timerID string, interval int)

type timer struct {
	timerID  string
	interval int
	expiry   int64
	cb       TimerCB
}
type timerEvent uint8

const (
	start timerEvent = iota
	cancel
	exit
)

var (
	timerStoreTree = rbt.NewWith(utils.StringComparator)
	timerIndexTree = rbt.NewWith(utils.Int64Comparator)
	yield          = make(chan timerEvent)
	isIntialised   = false
)

func Init() {
	isIntialised = true
	execute()
}

func StartTimer(interval int, cb TimerCB) (string, error) {
	if isIntialised {
		err := fmt.Errorf("timer is not initalized, please call Init before StartTimer")
		return "", err
	}
	timerID := uuid.Must(uuid.NewRandom()).String()

	tnow := time.Now()
	texp := tnow.Add(time.Duration(interval) * time.Second)
	expiry := texp.Unix()
	tm := timer{timerID: timerID, interval: interval, expiry: expiry, cb: cb}
	timerStoreTree.Put(tm.timerID, tm)
	timerIndexTree.Put(tm.expiry, tm.timerID)
	yield <- start
	return timerID, nil
}

func CancelTimer(timerId string) error {
	if !isIntialised {
		err := errors.New("timer is not initalize, please call Init before CancelTimer ")
		return err
	}
	tm := timerStoreTree.GetNode(timerId)
	tmr := tm.Value.(timer)
	timerIndexTree.Remove(tmr.expiry)
	timerStoreTree.Remove(tmr.timerID)
	yield <- cancel
	return nil
}

func handleExpiry(tId string) {
	nd := timerStoreTree.GetNode(tId)
	tm := nd.Value.(timer)
	timerStoreTree.Remove(tId)
	timerIndexTree.Remove(tm.expiry)
	tm.cb(tm.timerID, tm.interval)
}

func execute() {
	go func() {
		for {

			nd := timerIndexTree.Left()
			for nd == nil {
				<-yield
				nd = timerIndexTree.Left()
			}
			expiry := nd.Key.(int64)
			tm := time.Unix(expiry, 0)
			duration := time.Until(tm)
			tmtimer := time.NewTimer(duration)
			select {
			case <-tmtimer.C:
				handleExpiry(nd.Value.(string))
			case <-yield:
				tmtimer.Stop()
			}
		}

	}()
}
