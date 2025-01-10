package config

type Logger struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	FilePath   string `mapstructure:"file_path" json:"file_path" yaml:"file_path"`
	FileName   string `mapstructure:"file_name" json:"file_name" yaml:"file_name"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
