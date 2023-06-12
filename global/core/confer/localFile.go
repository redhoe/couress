package confer

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                   // 本地文件访问路径 Url
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // 本地文件存储路径 path
	Types     string `mapstructure:"types" json:"types" yaml:"types"`                // 支持的文件类型
	MaxSize   int64  `mapstructure:"max-size" json:"max-size" yaml:"max-size"`       // 文件允许上传大小
}
