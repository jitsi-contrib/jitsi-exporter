package jicofo

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jitsi-contrib/jitsi-exporter/pkg/generic"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type prober struct {
	client *http.Client
	target string
}

func NewProber(client *http.Client, target string) generic.Prober {
	return &prober{
		client: client,
		target: target,
	}
}

func (p *prober) Name() string {
	return "jicofo"
}

func (p *prober) Labels() []string {
	return []string{}
}

func (c *prober) Probe() (string, bool) {
	resp, err := c.client.Get(c.target)
	if err != nil {
		log.Error(err)
		return "", false
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "", false
	}

	return string(data), true
}

func (p *prober) MetricGetters() []generic.MetricGetter {
	return []generic.MetricGetter{
		generic.NewMetricGetter(prometheus.GaugeValue, "bridge_failures.participants_moved", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "bridge_selector.bridge_count", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "bridge_selector.lost_bridges", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "bridge_selector.in_shutdown_bridge_count", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "bridge_selector.operational_bridge_count", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_least_loaded", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_least_loaded_in_conference", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_least_loaded_in_region", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_least_loaded_in_region_group", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_least_loaded_in_region_group_in_conference", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_least_loaded_in_region_in_conference", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_not_loaded_in_region", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_not_loaded_in_region_group", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_not_loaded_in_region_group_in_conference", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_not_loaded_in_region_in_conference", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_split_due_to_load", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "bridge_selector.total_split_due_to_region", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "conference_sizes.average", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "conference_sizes.max", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "conference_sizes.min", "TODO"),
		newBucketGetter("conference_sizes"),
		newQuantilesGetter("conference_sizes"),

		generic.NewMetricGetter(prometheus.GaugeValue, "conferences", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoint_pairs", "TODO"),
		generic.NewBoolMetricGetter(prometheus.GaugeValue, "healthy", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "jibri.live_streaming_active", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jibri.live_streaming_pending", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jibri.recording_active", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jibri.recording_pending", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jibri.sip_call_active", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jibri.sip_call_pending", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "jibri.total_live_streaming_failures", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "jibri.total_recording_failures", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "jibri.total_sip_call_failures", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "jibri_detector.available", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jibri_detector.count", "TODO"),

		// TODO jigasi

		generic.NewMetricGetter(prometheus.GaugeValue, "jingle.received.session-accept", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jingle.received.transport-info", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jingle.sent.session-initiate", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jingle.sent.source-add", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "largest_conference", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "participant_notifications.ice_failed", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "participant_notifications.request_restart", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "participants", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.added_packets", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.dropped_packets", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.removed_packets", "TODO"),

		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.queue_size_at_remove.average_queue_size_at_remove", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.queue_size_at_remove.discarded", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.queue_size_at_remove.max_queue_size_at_remove", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "queues.jibri-iq-queue.queue_size_at_remove.min_queue_size_at_remove", "TODO"),
		newBucketGetter("queues.jibri-iq-queue.queue_size_at_remove"),
		newQuantilesGetter("queues.jibri-iq-queue.queue_size_at_remove"),

		generic.NewMetricGetter(prometheus.GaugeValue, "slow_health_check", "TODO"),
		generic.NewMetricGetter(prometheus.GaugeValue, "threads", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_conferences_created", "TODO"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_participants", "TODO"),
	}
}

func newBucketGetter(path string) generic.MetricGetter {
	return func(data string, p generic.Prober) prometheus.Metric {
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
			prometheus.NewDesc(prometheus.BuildFQName("jitsi", p.Name(), generic.PathToMetricName(path)), "TODO", nil, nil),
			uint64(count), sum, buckets)
	}
}

func newQuantilesGetter(path string) generic.MetricGetter {
	return func(data string, p generic.Prober) prometheus.Metric {
		sum := gjson.Get(data, fmt.Sprintf("%s.total_value", path)).Float()
		count := gjson.Get(data, fmt.Sprintf("%s.total_count", path)).Int()
		quantiles := map[float64]float64{
			0.99:  gjson.Get(data, fmt.Sprintf("%s.buckets.p99_upper_bound", path)).Float(),
			0.999: gjson.Get(data, fmt.Sprintf("%s.buckets.p999_upper_bound", path)).Float(),
		}
		return prometheus.MustNewConstSummary(
			prometheus.NewDesc(prometheus.BuildFQName("jitsi", p.Name(), fmt.Sprintf("%s_quantile", generic.PathToMetricName(path))), "TODO", nil, nil),
			uint64(count), sum, quantiles)
	}
}
