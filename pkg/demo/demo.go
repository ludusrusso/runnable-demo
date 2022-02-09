package demo

import (
	"context"
	"sync"
	"time"

	"github.com/ludusrusso/runnable-demo/pkg/runnable"
)

func NewDemoRunnable(data []Data) runnable.Runnable {
	return &demo{
		data:     data,
		dataChan: make(chan *Data, 1),
	}
}

type demo struct {
	data     []Data
	dataChan chan *Data
}

func (d *demo) loopPickData(ctx context.Context) {
	done := false
	go func() {
		<-ctx.Done()
		done = true
	}()

	for !done {
		for idx, datum := range d.data {
			if datum.Status == DataStatusPending {
				d.data[idx].Status = DataStatusLoading
				d.dataChan <- &d.data[idx]
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (d *demo) loopElaborate(ctx context.Context) {
	for {
		select {
		case datum := <-d.dataChan:
			datum.ComputedValue = datum.value * 2
			datum.Status = DataStatusDone
		case <-ctx.Done():
			return
		}
	}
}

func (d *demo) Run(ctx context.Context) error {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		d.loopPickData(ctx)
		wg.Done()
	}()

	go func() {
		d.loopElaborate(ctx)
		wg.Done()
	}()

	wg.Wait()

	return nil
}
