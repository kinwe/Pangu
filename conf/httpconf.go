package conf

type dbshost struct {
	Dbsurl string `mapstructure:"dbs-url" json:"dbsurl" yaml:"dbsurl"`
}
