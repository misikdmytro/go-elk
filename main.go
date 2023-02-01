package main

import (
	"os"
	"time"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func main() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller()).With(zap.String("app", "go-elk")).With(zap.String("environment", "local"))

	var i int
	for {
		logger.Info("application log",
			zap.Int("times", i),
		)

		i++
		time.Sleep(5 * time.Second)
	}
}
