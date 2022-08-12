package generic

// func NewMetricGetter(valueType prometheus.ValueType, path, help string) MetricGetter {
// 	return newMetricGetter(valueType, PathToMetricName(path), help, func(data string) float64 {
// 		return gjson.Get(data, path).Float()
// 	})
// }

// func NewBoolMetricGetter(valueType prometheus.ValueType, path, help string) MetricGetter {
// 	return newMetricGetter(valueType, PathToMetricName(path), help, func(data string) float64 {
// 		var value float64
// 		if gjson.Get(data, path).Bool() {
// 			value = 1
// 		}
// 		return value
// 	})
// }

// func PathToMetricName(path string) string {
// 	regex := regexp.MustCompile(`\.|-`)
// 	return regex.ReplaceAllString(path, "_")
// }

// func newMetricGetter(valueType prometheus.ValueType, name, help string, valueGetter func(data string) float64) MetricGetter {
// 	return func(data string, p Prober) prometheus.Metric {
// 		labels := getLabels(data, p)
// 		return prometheus.MustNewConstMetric(
// 			prometheus.NewDesc(prometheus.BuildFQName("jitsi", p.Name(), name), help, p.Labels(), nil),
// 			prometheus.GaugeValue,
// 			valueGetter(data), labels...)
// 	}
// }

// func getLabels(data string, p Prober) []string {
// 	labels := make([]string, len(p.Labels()))
// 	for k, v := range p.Labels() {
// 		labels[k] = gjson.Get(data, v).String()
// 	}
// 	return labels
// }
