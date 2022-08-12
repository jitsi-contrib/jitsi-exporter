package generic

import "github.com/prometheus/client_golang/prometheus"

type BasicMetricOptions struct {
	component      string
	variableLabels []string
	constLabels    prometheus.Labels
	getLabels      LabelGetter
}

type BasicMetricOption func(*BasicMetricOptions)

func NewBasicMetric(path string, valueType prometheus.ValueType, help string, opts ...BasicMetricOption) Metric {
	options := &BasicMetricOptions{}
	for _, apply := range opts {
		apply(options)
	}
	return NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, options.component, PathToMetricName(path)), help, options.variableLabels, options.constLabels),
		FloatGetter(path, valueType, options.getLabels),
	)
}

func WithComponent(component string) BasicMetricOption {
	return func(opts *BasicMetricOptions) {
		opts.component = component
	}
}

func WithVariableLabels(variableLabels []string) BasicMetricOption {
	return func(opts *BasicMetricOptions) {
		opts.variableLabels = variableLabels
	}
}

func WithConstLabels(constLabels prometheus.Labels) BasicMetricOption {
	return func(opts *BasicMetricOptions) {
		opts.constLabels = constLabels
	}
}

func WithLabelGetter(getLabels LabelGetter) BasicMetricOption {
	return func(opts *BasicMetricOptions) {
		opts.getLabels = getLabels
	}
}
