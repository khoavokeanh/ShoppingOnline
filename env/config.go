package env

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func (c *Config) GetConf() string {

	yamlFile, err := ioutil.ReadFile("env/docker-compose.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)
	fmt.Println(psqlInfo)
	return psqlInfo
}
