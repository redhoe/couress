package confer

type Redis struct {
	Host  string `mapstructure:"host" json:"host"  yaml:"host"`
	Port  int    `mapstructure:"port" json:"port" yaml:"port"`
	Auth  string `mapstructure:"auth" json:"auth" yaml:"auth"`
	Index int    `mapstructure:"index" json:"index" yaml:"index"`
}
