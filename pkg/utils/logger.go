package utils

import (
	time "time"

	logrus "github.com/sirupsen/logrus"
	//gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)

var GlobalLog = logrus.New()

func Log_initContainerLogger(){

	GlobalLog.SetLevel( logrus.DebugLevel);
	GlobalLog.SetFormatter(&logrus.JSONFormatter{})
}

func Log_log(message string, transactionID string, level int, msgData string, details string ){

	GlobalLog.WithFields( logrus.Fields{
		"MethodName": message,
		"TransactionId": transactionID,
		"ErrorLevel": level,
		"Data": msgData,
		"Info": details,
		"Timestamp":  time.Now().UTC().UnixNano(),

	}).Info("Info");

}
