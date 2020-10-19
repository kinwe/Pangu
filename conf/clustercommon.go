package conf

type ClusterCommons struct {
	Mastercommons []string `mapstructure:"mastercommons" json:"mastercommons" yaml:"mastercommons"`
	Slavecommons []string `mapstructure:"slavecommons" json:"slavecommons" yaml:"slavecommons"`
}
type SqlCommons struct {
	GtidCommons ClusterCommons `mapstructure:"gtidcommons" json:"gtidcommons" yaml:"gtidcommons"`
	FilePosCommons ClusterCommons `mapstructure:"fileposcommons" json:"fileposcommons" yaml:"fileposcommons"`
}
