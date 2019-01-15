package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type T struct {
	Mysql struct {
		Host     string
		DbName   string
		User     string
		Password string
		Port     string
	}

	Memcache struct {
		Host string
		Port string
	}
}

var Config T

func init() {
	//获得配置
	getConf()

}

func getConf() {
	Config = T{}
	data, err := ioutil.ReadFile("conf/config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.UnmarshalStrict([]byte(data), &Config)
	if err != nil {
		log.Fatalln(err)
	}

}
