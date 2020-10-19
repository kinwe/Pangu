package model

import (
	"database/sql"
	"fmt"
	"test/conf"
	"test/global"
)

type Mysql struct {
	Host         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Relo         Relo `mapstructure:"relo" json:"relo" yaml:"relo"`
	Connect *sql.DB
}

func (this *Mysql)BuildConnect()  {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", this.Username, this.Password, this.Host ,this.Dbname)

	db ,err := sql.Open("mysql",url)

	if err != nil {
		return
	}
	this.Connect = db
}
type Cluster struct {

	masterServer *Mysql

	slaveServer []*Mysql

	changemodel string

	Commonds conf.ClusterCommons
}

func (this * Cluster)AddMaster(masterserver *Mysql)  {
	if this.masterServer != nil {
		fmt.Println("Master 已存在")
		return
	}
	this.masterServer = masterserver
}

func (this * Cluster)AddSlave(slaveserver *Mysql)  {
	this.slaveServer = append(this.slaveServer,slaveserver)
}

func (this * Cluster)SetChangeModel(model string)  {
	if model == "GTID"  || model == "gtid" {
		this.changemodel = "GTID"
	}
	this.changemodel = "FILEPOS"
}

func (this * Cluster)BuildCluster(model string)  {

	if model == "GTID" {
		this.Commonds =global.GVA_Serve.SqlCommons.GtidCommons
	}
	if this.masterServer != nil {
		this.masterServer.BuildConnect()
	}
	if len(this.slaveServer) != 0 {
		for _, value :=  range this.slaveServer{
			if value != nil{
				value.BuildConnect()
			}
		}
	}
}


type Relo int

const (
	master Relo = iota
	slave
)
