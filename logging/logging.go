package logging

import (
	"go.uber.org/zap"
)

// Setup 配置日志
func Setup(debug bool, appName string) {
	var conf zap.Config
	if debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}
	logger, _ := conf.Build()
	logger.WithOptions(zap.AddCaller())
	logger = logger.With(zap.String("app", appName))
	zap.ReplaceGlobals(logger)
}
