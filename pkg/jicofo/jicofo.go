package jicofo

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jitsi-contrib/jitsi-exporter/pkg/generic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

var Component = "jicofo"

var Metrics = []generic.Metric{
	generic.NewBasicMetric("bridge_failures.participants_moved", prometheus.GaugeValue, "TODO"),

	generic.NewBasicMetric("bridge_selector.bridge_count", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.lost_bridges", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.in_shutdown_bridge_count", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.operational_bridge_count", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_least_loaded", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_least_loaded_in_conference", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_least_loaded_in_region", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_least_loaded_in_region_group", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_least_loaded_in_region_group_in_conference", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_least_loaded_in_region_in_conference", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_not_loaded_in_region", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_not_loaded_in_region_group", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_not_loaded_in_region_group_in_conference", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_not_loaded_in_region_in_conference", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_split_due_to_load", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("bridge_selector.total_split_due_to_region", prometheus.CounterValue, "TODO"),

	generic.NewBasicMetric("conference_sizes.average", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("conference_sizes.max", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("conference_sizes.min", prometheus.GaugeValue, "TODO"),
	newBucketGetter("conference_sizes"),
	newQuantilesGetter("conference_sizes"),

	generic.NewBasicMetric("conferences", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("endpoint_pairs", prometheus.GaugeValue, "TODO"),
	generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, Component, "healthy"), "TODO", nil, nil),
		generic.BoolGetter("healthy", prometheus.GaugeValue, nil),
	),

	generic.NewBasicMetric("jibri.live_streaming_active", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri.live_streaming_pending", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri.recording_active", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri.recording_pending", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri.sip_call_active", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri.sip_call_pending", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri.total_live_streaming_failures", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("jibri.total_recording_failures", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("jibri.total_sip_call_failures", prometheus.CounterValue, "TODO"),

	generic.NewBasicMetric("jibri_detector.available", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jibri_detector.count", prometheus.GaugeValue, "TODO"),

	// TODO jigasi

	generic.NewBasicMetric("jingle.received.session-accept", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jingle.received.transport-info", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jingle.sent.session-initiate", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("jingle.sent.source-add", prometheus.GaugeValue, "TODO"),

	generic.NewBasicMetric("largest_conference", prometheus.GaugeValue, "TODO"),

	generic.NewBasicMetric("participant_notifications.ice_failed", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("participant_notifications.request_restart", prometheus.GaugeValue, "TODO"),

	generic.NewBasicMetric("participants", prometheus.GaugeValue, "TODO"),

	generic.NewBasicMetric("queues.jibri-iq-queue.added_packets", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("queues.jibri-iq-queue.dropped_packets", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("queues.jibri-iq-queue.removed_packets", prometheus.GaugeValue, "TODO"),

	generic.NewBasicMetric("queues.jibri-iq-queue.queue_size_at_remove.average_queue_size_at_remove", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("queues.jibri-iq-queue.queue_size_at_remove.discarded", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("queues.jibri-iq-queue.queue_size_at_remove.max_queue_size_at_remove", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("queues.jibri-iq-queue.queue_size_at_remove.min_queue_size_at_remove", prometheus.GaugeValue, "TODO"),
	newBucketGetter("queues.jibri-iq-queue.queue_size_at_remove"),
	newQuantilesGetter("queues.jibri-iq-queue.queue_size_at_remove"),

	generic.NewBasicMetric("slow_health_check", prometheus.GaugeValue, "TODO"),
	generic.NewBasicMetric("threads", prometheus.GaugeValue, "TODO", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_conferences_created", prometheus.CounterValue, "TODO"),
	generic.NewBasicMetric("total_participants", prometheus.CounterValue, "TODO"),
}

func newBucketGetter(path string) generic.Metric {
	return generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, "", generic.PathToMetricName(path)+"_quantile"), "TODO", nil, nil),
		func(desc *prometheus.Desc, data string) prometheus.Metric {
			sum := gjson.Get(data, fmt.Sprintf("%s.total_value", path)).Float()
			count := gjson.Get(data, fmt.Sprintf("%s.total_count", path)).Int()
			buckets := map[float64]uint64{}
			regex := regexp.MustCompile(`_to_(\d+)$`)
			for k, v := range gjson.Get(data, fmt.Sprintf("%s.buckets", path)).Map() {
				matches := regex.FindStringSubmatch(k)
				if matches != nil {
					b, _ := strconv.ParseFloat(matches[1], 64)
					buckets[b] = v.Uint()
				}
			}
			return prometheus.MustNewConstHistogram(
				desc,
				uint64(count), sum, buckets)
		})
}

func newQuantilesGetter(path string) generic.Metric {
	return generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, "", generic.PathToMetricName(path)), "TODO", nil, nil),
		func(desc *prometheus.Desc, data string) prometheus.Metric {
			sum := gjson.Get(data, fmt.Sprintf("%s.total_value", path)).Float()
			count := gjson.Get(data, fmt.Sprintf("%s.total_count", path)).Int()
			quantiles := map[float64]float64{
				0.99:  gjson.Get(data, fmt.Sprintf("%s.buckets.p99_upper_bound", path)).Float(),
				0.999: gjson.Get(data, fmt.Sprintf("%s.buckets.p999_upper_bound", path)).Float(),
			}
			return prometheus.MustNewConstSummary(
				desc,
				uint64(count), sum, quantiles)
		})
}
