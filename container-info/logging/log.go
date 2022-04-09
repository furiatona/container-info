package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
		},
	})
	Log = logrus.WithFields(logrus.Fields{
		"service": "container-info",
		"country": "ID",
	})
}
