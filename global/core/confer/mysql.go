package confer

import "fmt"

type Mysql struct {
	Host   string `mapstructure:"host" json:"host" yaml:"host"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	User   string `mapstructure:"user" json:"user" yaml:"user"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Name   string `mapstructure:"name" json:"name" yaml:"name"`
}

func (d Mysql) MysqlDns() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s", d.User, d.Secret, d.Host, d.Port, d.Name, "utf8mb4", "true", "Asia%2FShanghai")
}
