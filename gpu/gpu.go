package gpu

import (
	"time"

	"github.com/coroot/coroot-node-agent/flags"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
)

var (
	gpuInfo = prometheus.NewDesc(
		"node_gpu_info",
		"Meta information about the GPU",
		[]string{"gpu_uuid", "name"}, nil,
	)
	gpuMemoryTotal = prometheus.NewDesc(
		"node_resources_gpu_memory_total_bytes",
		"Total memory available on the GPU in bytes",
		[]string{"gpu_uuid"}, nil,
	)
	gpuMemoryUsed = prometheus.NewDesc(
		"node_resources_gpu_memory_used_bytes",
		"GPU memory currently in use in bytes",
		[]string{"gpu_uuid"}, nil,
	)
	gpuMemoryUsageAvg = prometheus.NewDesc(
		"node_resources_gpu_memory_utilization_percent_avg",
		"Average GPU memory utilization (percentage) over the collection interval",
		[]string{"gpu_uuid"}, nil,
	)
	gpuMemoryUsagePeak = prometheus.NewDesc(
		"node_resources_gpu_memory_utilization_percent_peak",
		"Peak GPU memory utilization (percentage) over the collection interval",
		[]string{"gpu_uuid"}, nil,
	)
	gpuTemperature = prometheus.NewDesc(
		"node_resources_gpu_temperature_celsius",
		"Current temperature of the GPU in Celsius",
		[]string{"gpu_uuid"}, nil,
	)
	gpuPowerWatts = prometheus.NewDesc(
		"node_resources_gpu_power_usage_watts",
		"Current power usage of the GPU in watts",
		[]string{"gpu_uuid"}, nil,
	)
	gpuUsageAvg = prometheus.NewDesc(
		"node_resources_gpu_utilization_percent_avg",
		"Average GPU core utilization (percentage) over the collection interval",
		[]string{"gpu_uuid"}, nil,
	)
	gpuUsagePeak = prometheus.NewDesc(
		"node_resources_gpu_utilization_percent_peak",
		"Peak GPU core utilization (percentage) over the collection interval",
		[]string{"gpu_uuid"}, nil,
	)
)

type Collector struct {
	ProcessUsageSampleCh chan ProcessUsageSample
}

type ProcessUsageSample struct {
	UUID          string
	Pid           uint32
	Timestamp     time.Time
	GPUPercent    uint32
	MemoryPercent uint32
}

func NewCollector() (*Collector, error) {
	c := &Collector{
		ProcessUsageSampleCh: make(chan ProcessUsageSample, 100),
	}
	if *flags.DisableGPUMonitoring {
		return c, nil
	}

	klog.Infoln("GPU monitoring is enabled but NVML library not available on this host - skipping GPU monitoring")
	return c, nil
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	// no-op: when disabled, no GPU metrics are exposed
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	// no-op: when disabled, no GPU metrics are collected
}

func (c *Collector) Close() {
	// no-op
}
