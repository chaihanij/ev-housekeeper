package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ReadConf write file from config.yaml file
func ReadConf() *viper.Viper {
	v := viper.New()

	flag.String("path", ".", "configuration path")
	flag.String("config", "conf.yaml", "configuration file")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	v.BindPFlags(pflag.CommandLine)

	confFile := v.GetString("config")

	log.Printf("config name: %s", confFile)

	//set default configuration
	// filename to write
	v.SetDefault("filename", "untitled")

	// credential to mongodb
	v.SetDefault("mongodb.host", "localhost")
	v.SetDefault("mongodb.port", "27017")
	v.SetDefault("mongodb.user", "user")
	v.SetDefault("mongodb.password", "pass")
	v.SetDefault("mongodb.collection", "evhousekeeper")

	v.SetConfigName(confFile)
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AddConfigPath(v.GetString("path"))

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return v
}
