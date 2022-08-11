package generic

import (
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

type MetricDecription struct {
	ValueType prometheus.ValueType
	Path      string
	Help      string
}

type MetricGetter func(string, Prober) prometheus.Metric

type Prober interface {
	Name() string
	Probe() (string, bool)
	MetricGetters() []MetricGetter
	Labels() []string
}

type Collector struct {
	prober Prober
}

func NewMetricDecription(valueType prometheus.ValueType, path, help string) MetricDecription {
	return MetricDecription{
		ValueType: valueType,
		Path:      path,
		Help:      help,
	}
}

func NewCollector(prober Prober) *Collector {
	return &Collector{prober: prober}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	data, ok := c.prober.Probe()

	var up float64
	if ok {
		up = 1
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(prometheus.BuildFQName("jitsi", c.prober.Name(), "up"), "", c.prober.Labels(), nil), prometheus.GaugeValue, up, getLabels(data, c.prober)...)

	for _, getter := range c.prober.MetricGetters() {
		ch <- getter(data, c.prober)
	}
}

func NewMetricGetter(valueType prometheus.ValueType, path, help string) MetricGetter {
	return newMetricGetter(valueType, PathToMetricName(path), help, func(data string) float64 {
		return gjson.Get(data, path).Float()
	})
}

func NewBoolMetricGetter(valueType prometheus.ValueType, path, help string) MetricGetter {
	return newMetricGetter(valueType, PathToMetricName(path), help, func(data string) float64 {
		var value float64
		if gjson.Get(data, path).Bool() {
			value = 1
		}
		return value
	})
}

func PathToMetricName(path string) string {
	regex := regexp.MustCompile(`\.|-`)
	return regex.ReplaceAllString(path, "_")
}

func newMetricGetter(valueType prometheus.ValueType, name, help string, valueGetter func(data string) float64) MetricGetter {
	return func(data string, p Prober) prometheus.Metric {
		labels := getLabels(data, p)
		return prometheus.MustNewConstMetric(
			prometheus.NewDesc(prometheus.BuildFQName("jitsi", p.Name(), name), help, p.Labels(), nil),
			prometheus.GaugeValue,
			valueGetter(data), labels...)
	}
}

func getLabels(data string, p Prober) []string {
	labels := make([]string, len(p.Labels()))
	for k, v := range p.Labels() {
		labels[k] = gjson.Get(data, v).String()
	}
	return labels
}
