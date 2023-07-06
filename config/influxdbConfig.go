package config

import "github.com/spf13/viper"

type InfluxdbConfig struct {
	Token  string
	Url    string
	Org    string
	Bucket string
}

func LoadInfluxdbConfig() InfluxdbConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read configuration file")
	}

	return InfluxdbConfig{
		Token:  viper.GetString("influxdb.token"),
		Url:    viper.GetString("influxdb.url"),
		Org:    viper.GetString("influxdb.org"),
		Bucket: viper.GetString("influxdb.bucket"),
	}
}
