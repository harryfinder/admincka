package configs

import (
	"github.com/activ-capital/partner-service/internal/models"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func InitConfig() (cfg models.Configuration, err error) {
	//viper.AddConfigPath("")
	//viper.SetConfigName("adminConfig")
	//viper.SetConfigType("yml") // Explicitly set the config file type
	viper.SetConfigFile("./adminConfig.yml")
	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %s\n", err)
		return cfg, err
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err = viper.UnmarshalKey("configuration", &cfg); err != nil {
		log.Printf("Error unmarshalling configuration: %s\n", err)
		return cfg, err
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    2,  // megabytes
		MaxAge:     40, // days
		MaxBackups: 30,
		Compress:   true,
	})

	log.Println("---*--- Starting Logging ---*---")
	return cfg, nil
}
