package conf

type Service struct {
	Dbhost dbshost `mapstructure:"dbshost" json:"dbshost" yaml:"dbshost"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	SqlCommons SqlCommons  `mapstructure:"sqlcommons" json:"sqlcommons" yaml:"sqlcommons"`
}
