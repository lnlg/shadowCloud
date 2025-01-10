package config

// 其 Database 属性对应 config.yaml 中的 app 部分
type Database struct {
	Driver   string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" json:"charset" yaml:"charset"`
}
