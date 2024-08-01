package logger

import (
	"io"
	"log"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/utilities/fmtrs"
)

func fileLogger() *os.File {
	logFile, err := os.OpenFile("./logs/system.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return logFile
}

func LogDebug(str string) {
	file := fileLogger()
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	formatted := fmtrs.FormatLogString(str, "", 0)

	logger.Println(formatted)
}

func LogInfo(str string) {
	file := fileLogger()
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	formatted := fmtrs.FormatLogString(str, "", 0)

	logger.Println(formatted)
}

func LogWarn(str string, warnCode int) {
	file := fileLogger()
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	formatted := fmtrs.FormatLogString(str, "", 0)

	logger.Println(formatted)
}

func LogErr(str string, errCode int) {
	file := fileLogger()
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	formatted := fmtrs.FormatLogString(str, "err", errCode)

	logger.Println(formatted)
}

func LogFatal(str string, fatalErrCode int) {
	file := fileLogger()
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := log.New(mw, "FATAL ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	formatted := fmtrs.FormatLogString(str, "ftl", fatalErrCode)

	logger.Println(formatted)
	os.Exit(1)
}
