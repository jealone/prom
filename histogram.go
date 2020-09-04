package prom

import (
	"sync"

	"github.com/jealone/sli4go"
)

type LazyHistogram struct {
	name  string
	Value Histogram
	once  sync.Once
}

func NewLazyHistogram(name string) *LazyHistogram {
	return &LazyHistogram{
		name: name,
	}
}

func (c *LazyHistogram) Get() (Histogram, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetHistogram(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type HistogramConfig struct {
	Opts HistogramOptsConfig `yaml:"opts"`
}

func (c *HistogramConfig) GetOpts() *HistogramOptsConfig {
	return &c.Opts
}

type HistogramOptsConfig struct {
	OptsConfig `yaml:",inline"`
	Buckets    string `yaml:"buckets"`
}

func (o *HistogramOptsConfig) GetBuckets() []float64 {
	var buckets []float64
	//for _, s := range strings.Split(o.Buckets, ",") {
	//	f, err := strconv.ParseFloat(s, 64)
	//
	//	if nil != err {
	//		panic(err)
	//	}
	//
	//	buckets = append(buckets, f)
	//}
	err := json.UnmarshalFromString(o.Buckets, &buckets)
	if nil != err {
		sli4go.Errorf("unmarshal json(%s) error %s", o.Buckets, err)
		return nil
	}
	return buckets
}
