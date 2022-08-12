package generic

import (
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type Collector struct {
	client  *http.Client
	target  string
	metrics []Metric
}

func NewCollector(client *http.Client, target string, metrics []Metric) *Collector {
	return &Collector{
		client:  client,
		target:  target,
		metrics: metrics,
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m.desc
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	data, err := c.probe()

	if err != nil {
		log.Error(err)
		return
	}

	for _, m := range c.metrics {
		pm := m.get(m.desc, data)
		if pm != nil {
			ch <- pm
		}
	}
}

func (c *Collector) probe() (string, error) {
	resp, err := c.client.Get(c.target)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
