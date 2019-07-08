package setting

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JWTSecret string
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = viper.GetString(`general.RUN_MODE`)
}

func loadServer() {
	HTTPPort = viper.GetInt(`server.HTTP_PORT`)
	ReadTimeout = time.Duration(viper.GetInt(`server.READ_TIMEOUT`))
	WriteTimeout = time.Duration(viper.GetInt(`server.WRITE_TIMEOUT`))
}

func loadApp() {
	JWTSecret = viper.GetString(`app.JWT_SECRET`)
	PageSize = viper.GetInt(`app.PAGE_SIZE`)
}
