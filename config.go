package prom

type PrometheusConfig struct {
	Path    string   `yaml:"path"`
	Listen  string   `yaml:"listen"`
	Metrics []Metric `yaml:"metrics"`
}

func (c *PrometheusConfig) GetPath() string {
	return c.Path
}

func (c *PrometheusConfig) GetListen() string {
	return c.Listen
}

type Metric struct {
	Type string   `yaml:"type"`
	Key  string   `yaml:"key"`
	Spec YamlNode `yaml:"spec"`
}

func (m *Metric) GetType() string {
	return m.Type
}

func (m *Metric) GetKey() string {
	return m.Key
}

func (m *Metric) GetSpec() *YamlNode {
	return &m.Spec
}

type OptsConfig struct {
	Namespace string `yaml:"namespace"`
	Subsystem string `yaml:"subsystem"`
	Name      string `yaml:"name"`

	Help string `yaml:"help"`
}

func (o *OptsConfig) GetNamespace() string {
	return o.Namespace
}

func (o *OptsConfig) GetSubsystem() string {
	return o.Subsystem
}

func (o *OptsConfig) GetName() string {
	return o.Name
}

func (o *OptsConfig) GetHelp() string {
	return o.Help
}
