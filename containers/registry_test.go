package containers

import (
	"testing"
	"time"

	"github.com/coroot/coroot-node-agent/ebpftracer"
	"github.com/stretchr/testify/require"
)

func TestRegistrySkipsEbpfStatsWhenTracerUnavailable(t *testing.T) {
	r := &Registry{
		tracer:               ebpftracer.NewTracer(0, 0, false),
		trafficStatsUpdateCh: make(chan *TrafficStatsUpdate, 1),
		nodejsStatsUpdateCh:  make(chan *NodejsStatsUpdate, 1),
		pythonStatsUpdateCh:  make(chan *PythonStatsUpdate, 1),
	}

	require.NotPanics(t, r.updateStatsFromEbpfMapsIfNecessary)
	require.WithinDuration(t, time.Now(), r.ebpfStatsLastUpdated, time.Second)
}
