package prom

import (
	"errors"
	"sync"

	"github.com/jealone/sli4go"
	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/yaml.v3"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type (
	CounterOpts   = prometheus.CounterOpts
	GaugeOpts     = prometheus.GaugeOpts
	HistogramOpts = prometheus.HistogramOpts
	SummaryOpts   = prometheus.SummaryOpts
	Opts          = prometheus.Opts
	Labels        = prometheus.Labels
	Timer         = prometheus.Timer

	Gauge        = prometheus.Gauge
	GaugeVec     = prometheus.GaugeVec
	Counter      = prometheus.Counter
	CounterVec   = prometheus.CounterVec
	Histogram    = prometheus.Histogram
	HistogramVec = prometheus.HistogramVec
	Summary      = prometheus.Summary
	SummaryVec   = prometheus.SummaryVec

	YamlNode = yaml.Node
)

const (
	GaugeType        = "gauge"
	GaugeVecType     = "gauge_vector"
	CounterType      = "counter"
	CounterVecType   = "counter_vector"
	HistogramType    = "histogram"
	HistogramVecType = "histogram_vector"
	SummaryType      = "summary"
	SummaryVecType   = "summary_vector"
)

var (
	RemindOpts = Opts{
		Namespace: "remind",
	}
	ErrorNotInit        = errors.New("please init prometheus registry first")
	ErrorParseConfig    = errors.New("prometheus parse config error ")
	ErrorInvalidConfig  = errors.New("incorrect prometheus config for key ")
	ErrorMetricNotFound = errors.New("prometheus metric not found for key ")
)

type ConfigDecoder interface {
	Decode(v interface{}) (err error)
}

type Registry struct {
	ConfigMap map[string]Metric
	mu        sync.RWMutex
}

var (
	once            sync.Once
	defaultRegistry *Registry
)

func InitRegistry(conf *PrometheusConfig) {

	once.Do(func() {
		defaultRegistry = NewRegistry(conf)
	})
}

func GetRegistry() *Registry {
	once.Do(func() {
		sli4go.Fatal(ErrorNotInit)
	})
	return defaultRegistry
}

func NewRegistry(conf *PrometheusConfig) *Registry {
	re := &Registry{}
	re.ConfigMap = make(map[string]Metric)
	re.mu.Lock()
	for _, c := range conf.Metrics {
		re.ConfigMap[c.GetKey()] = c
	}
	re.mu.Unlock()
	return re

}

func (r *Registry) GetGauge(name string) Gauge {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if GaugeType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &GaugeConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewGauge(GaugeOpts{
		Namespace: conf.GetOpts().GetNamespace(),
		Subsystem: conf.GetOpts().GetSubsystem(),
		Name:      conf.GetOpts().GetName(),
		Help:      conf.GetOpts().GetHelp(),
	})

	prometheus.MustRegister(collector)

	return collector
}

func (r *Registry) GetGaugeVec(name string) *GaugeVec {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if GaugeVecType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &GaugeVecConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewGaugeVec(GaugeOpts{
		Namespace: conf.GetOpts().GetNamespace(),
		Subsystem: conf.GetOpts().GetSubsystem(),
		Name:      conf.GetOpts().GetName(),
		Help:      conf.GetOpts().GetHelp(),
	}, conf.GetLabels())

	prometheus.MustRegister(collector)

	return collector
}

func (r *Registry) GetCounter(name string) Counter {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if CounterType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &CounterConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewCounter(CounterOpts{
		Namespace: conf.GetOpts().GetNamespace(),
		Subsystem: conf.GetOpts().GetSubsystem(),
		Name:      conf.GetOpts().GetName(),
		Help:      conf.GetOpts().GetHelp(),
	})

	prometheus.MustRegister(collector)

	return collector
}

func (r *Registry) GetCounterVec(name string) *CounterVec {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if CounterVecType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &CounterVecConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewCounterVec(CounterOpts{
		Namespace: conf.GetOpts().GetNamespace(),
		Subsystem: conf.GetOpts().GetSubsystem(),
		Name:      conf.GetOpts().GetName(),
		Help:      conf.GetOpts().GetHelp(),
	}, conf.GetLabels())

	prometheus.MustRegister(collector)

	return collector
}

func (r *Registry) GetHistogram(name string) Histogram {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if HistogramType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &HistogramConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewHistogram(HistogramOpts{
		Namespace: conf.GetOpts().GetNamespace(),
		Subsystem: conf.GetOpts().GetSubsystem(),
		Name:      conf.GetOpts().GetName(),
		Help:      conf.GetOpts().GetHelp(),
		Buckets:   conf.GetOpts().GetBuckets(),
	})
	prometheus.MustRegister(collector)
	return collector
}

func (r *Registry) GetHistogramVec(name string) *HistogramVec {
	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if HistogramVecType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &HistogramVecConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewHistogramVec(HistogramOpts{
		Namespace: conf.GetOpts().GetNamespace(),
		Subsystem: conf.GetOpts().GetSubsystem(),
		Name:      conf.GetOpts().GetName(),
		Help:      conf.GetOpts().GetHelp(),
		Buckets:   conf.GetOpts().GetBuckets(),
	}, conf.GetLabels())

	prometheus.MustRegister(collector)
	return collector
}

func (r *Registry) GetSummary(name string) Summary {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if SummaryType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &SummaryConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewSummary(SummaryOpts{
		Namespace:  conf.GetOpts().GetNamespace(),
		Subsystem:  conf.GetOpts().GetSubsystem(),
		Name:       conf.GetOpts().GetName(),
		Help:       conf.GetOpts().GetHelp(),
		Objectives: conf.GetOpts().GetObjectives(),
		MaxAge:     conf.GetOpts().GetMaxAge(),
		AgeBuckets: conf.GetOpts().GetAgeBuckets(),
		BufCap:     conf.GetOpts().GetBufCap(),
	})

	prometheus.MustRegister(collector)
	return collector
}

func (r *Registry) GetSummaryVec(name string) *SummaryVec {

	r.mu.RLock()
	metric, ok := r.ConfigMap[name]
	r.mu.RUnlock()

	if !ok {
		sli4go.Error(ErrorMetricNotFound, name)
		return nil
	}

	if SummaryVecType != metric.GetType() {
		sli4go.Error(ErrorInvalidConfig, name)
		return nil
	}

	conf := &SummaryVecConfig{}

	err := metric.GetSpec().Decode(conf)

	if nil != err {
		sli4go.Error(ErrorParseConfig)
		return nil
	}

	collector := prometheus.NewSummaryVec(SummaryOpts{
		Namespace:  conf.GetOpts().GetNamespace(),
		Subsystem:  conf.GetOpts().GetSubsystem(),
		Name:       conf.GetOpts().GetName(),
		Help:       conf.GetOpts().GetHelp(),
		Objectives: conf.GetOpts().GetObjectives(),
		MaxAge:     conf.GetOpts().GetMaxAge(),
		AgeBuckets: conf.GetOpts().GetAgeBuckets(),
		BufCap:     conf.GetOpts().GetBufCap(),
	}, conf.GetLabels())

	prometheus.MustRegister(collector)
	return collector
}
