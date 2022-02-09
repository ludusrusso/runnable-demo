package demo_test

import (
	"context"
	"testing"
	"time"

	"github.com/ludusrusso/runnable-demo/pkg/demo"
	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	data := []demo.Data{demo.NewData(2), demo.NewData(4)}
	s := demo.NewDemoRunnable(data)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	s.Run(ctx)

	assert.Equal(t, demo.DataStatusDone, data[0].Status)
	assert.Equal(t, 4, data[0].ComputedValue)
}
