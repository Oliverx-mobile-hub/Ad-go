package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false, // 禁用、启动时间戳
		FullTimestamp:    true,  // 完整时间戳
		TimestampFormat:  "2006-01-02 15:04:05",
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}

func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)
}

func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Warning(msg)
}

func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Error(msg)
}
