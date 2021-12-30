package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	ContextUserClaimKey = "userClaim"
	ContextDBKey        = "DB"
	ContextLoggingDBKey = "LOGGING_DB"
)

var Config = struct {
	HttpPort    string
	Environment string
	Database    struct {
		Driver           string
		User             string
		Connection       string
		ConnectionString string
	}
	Service struct {
		Name string
	}
	Log struct {
		ShowSql    bool
		Path       string
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
	}
	Encrypt struct {
		EncryptKey string
	}
}{}

func InitConfig(cfg string) {
	configor.Load(&Config, cfg)
}

// 서비스 무관하게 공통으로 사용하는 부분
func ConfigureEnvironment(path string, env ...string) {
	configor.Load(&Config, path+"config/config.json")
	properties := make(map[string]string)

	for _, key := range env {
		arg := os.Getenv(key)
		if len(arg) == 0 {
			panic(fmt.Errorf("No %s system env variable\n", key))
		}
		properties[key] = arg
	}

	afterPropertiesSet(properties)
}

// 서비스별 처리 로직이 달라지는 부분.
func afterPropertiesSet(properties map[string]string) {
	Config.Encrypt.EncryptKey = properties["STUDY_GENIE_ENCRYPT_KEY"]
	if properties["STUDY_GENIE_DB_PASSWORD"] != "" {
		Config.Database.ConnectionString = fmt.Sprintf("%s:%s%s", Config.Database.User, properties["STUDY_GENIE_DB_PASSWORD"], Config.Database.Connection)
	} else {
		Config.Database.ConnectionString = Config.Database.Connection
	}
}

func ConfigureLogger() {
	lum := &lumberjack.Logger{
		Filename:   Config.Log.Path,
		MaxSize:    Config.Log.MaxSize,
		MaxBackups: Config.Log.MaxBackups,
		MaxAge:     Config.Log.MaxAge,
		Compress:   Config.Log.Compress,
	}

	logrus.SetOutput(lum)

	environment := Config.Environment
	if environment == "production" || environment == "dev" || environment == "qa" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceQuote:    true,
		})
	}

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.Infoln("Log Service Started")
}
