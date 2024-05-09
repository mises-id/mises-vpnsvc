package env

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Envs *Env

type Env struct {
	Port             int    `env:"PORT" envDefault:"8080"`
	AppEnv           string `env:"APP_ENV" envDefault:"development"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"INFO"`
	MongoURI         string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	DBUser           string `env:"DB_USER"`
	DBPass           string `env:"DB_PASS"`
	DBName           string `env:"DB_NAME" envDefault:"mises_vpn_test"`
	SyncVpnOrderMode string `env:"SYNC_VPN_ORDER_MODE" envDefault:"close"`
	RootPath         string
}

func init() {
	fmt.Println("vpnsvc env initializing...")
	//_, b, _, _ := runtime.Caller(0)
	appEnv := os.Getenv("APP_ENV")
	projectRootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	envPath := projectRootPath + "/.env"
	appEnvPath := envPath + "." + appEnv
	localEnvPath := appEnvPath + ".local"
	_ = godotenv.Load(filtePath(localEnvPath, appEnvPath, envPath)...)
	Envs = &Env{}
	err = env.Parse(Envs)
	if err != nil {
		panic(err)
	}
	Envs.RootPath = projectRootPath
	fmt.Println("vpnsvc env root " + projectRootPath)
	fmt.Println("vpnsvc env loaded...")
}

func filtePath(paths ...string) []string {
	result := make([]string, 0)
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			result = append(result, path)
		}
	}
	return result
}
