package config

import (
	"log"
	"os"
	"github.com/spf13/viper"
)

type Config struct {
    App struct
    {
        Port string `mapstructure:"PORT"`
		Name string `mapstructure:"NAME"`
    }
	Database struct
	{
		Dsn string `mapstructure:"DSN"`
		MaxIdleConns int `mapstructure:"MAX_IDLE_CONNS"`
		MaxOpenConns int `mapstructure:"MAX_OPEN_CONNS"`
		}
	FilePath struct
	{
		BasePath string `mapstructure:"BASE_PATH"`
	}
	}
 var AppConfig *Config

 func InitConfig() {
     viper.SetConfigName("config")
	 viper.SetConfigType("yaml")
	 viper.AddConfigPath("./config")
	cwd,_:= os.Getwd()
	viper.SetDefault("FILE_PATH.BASE_PATH", cwd)

	 if err:= viper.ReadInConfig(); err != nil {
		 log.Fatal(err)
	 }
	AppConfig= &Config{}
	 if err := viper.Unmarshal(AppConfig); err != nil {
		 log.Fatal(err)
	 }
	 InitDB()
 }