package runner

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yohamta/dagu/internal/controller"
	"github.com/yohamta/dagu/internal/scheduler"
)

func TestJobStart(t *testing.T) {
	file := path.Join(testdataDir, "start.yaml")
	dr := controller.NewDAGReader()
	dag, _ := dr.ReadDAG(file, false)
	c := controller.New(dag.DAG)

	j := &job{
		DAG:    dag.DAG,
		Config: testConfig,
		Next:   time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC),
	}

	go func() {
		_ = j.Start()
	}()

	time.Sleep(time.Millisecond * 100)

	err := j.Start()
	require.Equal(t, ErrJobRunning, err)

	c.Stop()
	time.Sleep(time.Millisecond * 200)

	s, _ := c.GetLastStatus()
	require.Equal(t, scheduler.SchedulerStatus_Cancel, s.Status)

	err = j.Start()
	require.Equal(t, ErrJobFinished, err)
}

func TestJobSop(t *testing.T) {
	file := path.Join(testdataDir, "stop.yaml")
	dr := controller.NewDAGReader()
	dag, _ := dr.ReadDAG(file, false)

	j := &job{
		DAG:    dag.DAG,
		Config: testConfig,
		Next:   time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC),
	}

	go func() {
		_ = j.Start()
	}()

	time.Sleep(time.Millisecond * 100)

	err := j.Stop()
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	c := controller.New(dag.DAG)
	s, _ := c.GetLastStatus()
	require.Equal(t, scheduler.SchedulerStatus_Cancel, s.Status)
}
