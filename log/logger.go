package logger

import (
	"log"
	"os"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("./log/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

	//before deployment uncomment the following six lines
	// logf, err := os.OpenFile("./log/prodLogs.txt", os.O_WRONLY|os.O_CREATE,
    //     0640)
    //  if err != nil {
	// 	    log.Fatalln(err)
	//    }
	// log.SetOutput(logf)

	
    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}
