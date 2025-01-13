package config

// 其 App 属性对应 config.yaml 中的 app 部分
type App struct {
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	Mode    string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Host    string `mapstructure:"host" json:"host" yaml:"host"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Debug   bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
}
