package logging

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jepbura/go-server/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger used for the whole application.
func LoggerInit(cnf config.Env) (*zap.Logger, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("LoggerInit\n")
	fmt.Print("*********************************************\n")
	config := zap.NewDevelopmentConfig()

	if cnf.Environment == "live" {
		os.Mkdir("log", os.ModePerm)
		year, month, day := time.Now().Date()
		path := "log/" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + ".json"
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{path, "stderr"}
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return config.Build()
}
