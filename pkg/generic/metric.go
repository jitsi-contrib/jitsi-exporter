package generic

import (
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

var Namespace = "jitsi"

type MetricGetter func(*prometheus.Desc, string) prometheus.Metric

type Metric struct {
	desc *prometheus.Desc
	get  MetricGetter
}

func NewMetric(desc *prometheus.Desc, getter MetricGetter) Metric {
	return Metric{
		desc: desc,
		get:  getter,
	}
}

// func NewSimpleMetric(path string) Metric {
// 	return NewMetric(
// 		prometheus.NewDesc(prometheus.BuildFQName(Namespce, "", PathToMetricName(path)), "", nil, nil),
// 	)
// }

// func NewPathDesc(path string,) *prometheus.Desc {
// 	return prometheus.NewDesc(prometheus.BuildFQName(Namespce, "", PathToMetricName(path)), "", nil, nil),
// }

type LabelGetter func(string) []string

func FloatGetter(path string, valueType prometheus.ValueType, getLabel LabelGetter) MetricGetter {
	return newMetricGetter(path, valueType, getLabel, func(result gjson.Result) float64 {
		return result.Float()
	})
}

func BoolGetter(path string, valueType prometheus.ValueType, getLabel LabelGetter) MetricGetter {
	return newMetricGetter(path, valueType, getLabel, func(result gjson.Result) float64 {
		var value float64 = 0
		if result.Bool() {
			value = 1
		}
		return value
	})
}

type valueGetter func(gjson.Result) float64

func newMetricGetter(path string, valueType prometheus.ValueType, getLabel LabelGetter, getValue valueGetter) MetricGetter {
	return func(desc *prometheus.Desc, data string) prometheus.Metric {
		result := gjson.Get(data, path)
		if !result.Exists() {
			return nil
		}
		var labels []string
		if getLabel != nil {
			labels = getLabel(data)
		}
		return prometheus.MustNewConstMetric(desc, valueType, getValue(result), labels...)
	}
}

func PathToMetricName(path string) string {
	regex := regexp.MustCompile(`\.|-`)
	return regex.ReplaceAllString(path, "_")
}
