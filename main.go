package main

import (
	"fmt"
	"github.com/habakke/router/server"

	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
)

const (
	CONFIG_FILE_NAME     = "./conf/app.toml"
	LOG_CONFIG_FILE_NAME = "./conf/log.xml"
)

var (
	version   string // build version number
	commit    string // sha1 revision used to build the program
	buildTime string // when the executable was built
	buildBy   string
)

func GetVersionString(name string) string {
	return fmt.Sprintf("%s %s (%s at %s by %s)", name, version, commit, buildTime, buildBy)
}

func init() {
	viper.Reset()
	viper.SetConfigType("toml")
	viper.SetConfigFile(CONFIG_FILE_NAME)
	if e := viper.ReadInConfig(); e != nil {
		panic(e)
	}

	logger, err := log.LoggerFromConfigAsFile(LOG_CONFIG_FILE_NAME)
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)

}
func main() {
	fmt.Println(GetVersionString("router"))

	port := viper.GetString("service.port")
	topic := "broker000100101info"
	srv := server.NewServer(port, topic)
	srv.Start()
}
