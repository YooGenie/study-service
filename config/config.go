package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

const (
	ContextUserClaimKey = "userClaim"
	ContextDBKey        = "DB"
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
	Kakao struct {
		RestApiKey        string
		RedirectURL       string
		LogoutRedirectURL string
	}
	Encrypt struct {
		EncryptKey string
	}
	JwtSecret string
	Mail      struct {
		Host     string
		Port     int
		User     string
		Password string
		Sender   string
		Content  struct {
			Subject string
			Path    struct {
				MailBody   string
				Attachment string
			}
		}
	}
}{}

func InitConfig(cfg string) {
	configor.Load(&Config, cfg)
}

// 서비스 무관하게 공통으로 사용하는 부분
func ConfigureEnvironment(path string, env ...string) {
	// 	"github.com/jinzhu/configor" import 해서 설정파일을 읽을 때 사용한다.
	configor.Load(&Config, path+"config/config.json") //배포 환경에 따른 설정 파일(json)을 로딩한다.
	properties := make(map[string]string)

	//env가 1) STUDY_GENIE_DB_PASSWORD 2) STUDY_GENIE_DB_PASSWORD => 키값이 된다.
	for _, key := range env { //환경변수의 키-값 쌍을 설정
		arg := os.Getenv(key) //환경변수의 키-값 쌍을 설정하고 키에 따른 값을 가져온다. => 환경변수 읽기  // os.Setenv() 환경변수 쓰기
		if len(arg) == 0 {
			panic(fmt.Errorf("No %s system env variable\n", key))
		}
		properties[key] = arg //키에 값을 저장한다. key값은 STUDY_GENIE_DB_PASSWORD이고 value는 내가 컴퓨터에 저장해 놓는 환경변수 값을 가져온다.
	}

	afterPropertiesSet(properties)
}

// 서비스별 처리 로직이 달라지는 부분.
func afterPropertiesSet(properties map[string]string) { //환경변수를 가지고 와서 Config 구조체 안에 값을 넣어준다.
	Config.Encrypt.EncryptKey = properties["STUDY_GENIE_ENCRYPT_KEY"] //Config 구조체 안에서 Encrypt안에 EncryptKey에 값을 넣어준다.
	Config.Kakao.RestApiKey = properties["KAKAO_API_KEY"]

	if properties["STUDY_GENIE_DB_PASSWORD"] != "" {
		Config.Database.ConnectionString = fmt.Sprintf("%s:%s%s", Config.Database.User, properties["STUDY_GENIE_DB_PASSWORD"], Config.Database.Connection)
	} else {
		Config.Database.ConnectionString = Config.Database.Connection
	}
}

//func ConfigureLogger() {
//	lum := &lumberjack.Logger{
//		Filename:   Config.Log.Path,
//		MaxSize:    Config.Log.MaxSize,
//		MaxBackups: Config.Log.MaxBackups,
//		MaxAge:     Config.Log.MaxAge,
//		Compress:   Config.Log.Compress,
//	}
//
//	logrus.SetOutput(lum)
//
//	environment := Config.Environment
//	if environment == "production" || environment == "dev" || environment == "qa" {
//		logrus.SetFormatter(&logrus.JSONFormatter{})
//	} else {
//		logrus.SetFormatter(&logrus.TextFormatter{
//			FullTimestamp: true,
//			ForceQuote:    true,
//		})
//	}
//
//	// Only log the warning severity or above.
//	logrus.SetLevel(logrus.TraceLevel)
//
//	logrus.SetReportCaller(true)
//
//	logrus.Infoln("Log Service Started")
//}
