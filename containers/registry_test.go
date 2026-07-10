package containers

import (
	"testing"
	"time"

	"github.com/coroot/coroot-node-agent/ebpftracer"
	"github.com/coroot/coroot-node-agent/proc"
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

func TestNewProcessPids(t *testing.T) {
	seen := map[uint32]struct{}{}

	pids, seen := newProcessPids(seen, []uint32{1, 2})
	require.Equal(t, []uint32{1, 2}, pids)

	pids, seen = newProcessPids(seen, []uint32{2, 3})
	require.Equal(t, []uint32{3}, pids)

	pids, _ = newProcessPids(seen, []uint32{1, 2, 3})
	require.Equal(t, []uint32{1}, pids)
}

func TestProcessSnapshotEvents(t *testing.T) {
	events := processSnapshotEvents(42, []proc.Fd{
		{Fd: 3, Dest: "/var/log/app.log"},
		{Fd: 4, Dest: "/data/file"},
		{Fd: 5, Dest: "socket:[123]"},
		{Fd: 6, Dest: "pipe:[456]"},
	})

	require.Equal(t, []ebpftracer.Event{
		{Type: ebpftracer.EventTypeProcessStart, Pid: 42},
		{Type: ebpftracer.EventTypeFileOpen, Pid: 42, Fd: 3, Log: true},
		{Type: ebpftracer.EventTypeFileOpen, Pid: 42, Fd: 4},
	}, events)
}
