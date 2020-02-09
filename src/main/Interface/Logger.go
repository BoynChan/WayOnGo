package main

import logger "main/Interface/Logger"

func createLogger() *logger.Logger {
	//
	cw := logger.NewConsoleWriter()
	fl := logger.NewFileWriter()
	fl.SetFile("log.log")

	log := logger.NewLogger()
	log.RegisterWriter(fl)
	log.RegisterWriter(cw)

	return log
}

func main() {
	l := createLogger()
	l.Log("hello")
}
