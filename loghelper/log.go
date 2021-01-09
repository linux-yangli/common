package loghelper

import (
	"github.com/phachon/go-logger"
)
var InfoLogPath = "./"
var ErrorLogPath = "./"
var DebugLogPath = "./"
func Logs()  {
	logger := go_logger.NewLogger()

	logger.Detach("console")

	// console adapter config
	consoleConfig := &go_logger.ConsoleConfig{
		Color: true, // Does the text display the color
		JsonFormat: true, // Whether or not formatted into a JSON string
		Format: "", // JsonFormat is false, logger message output to console format string
	}
	// add output to the console
	logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

	// file adapter config
	fileConfig := &go_logger.FileConfig {
		Filename : "./test.log", // The file name of the logger output, does not exist automatically
		// If you want to separate separate logs into files, configure LevelFileName parameters.
		LevelFileName : map[int]string {
			logger.LoggerLevel("error"): ErrorLogPath,    // The error level log is written to the error.log file.
			logger.LoggerLevel("info"): InfoLogPath,      // The info level log is written to the info.log file.
			logger.LoggerLevel("debug"): InfoLogPath,    // The debug level log is written to the debug.log file.
		},
		MaxSize : 1024 * 1024,  // File maximum (KB), default 0 is not limited
		MaxLine : 100000, // The maximum number of lines in the file, the default 0 is not limited
		DateSlice : "d",  // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
		JsonFormat: true, // Whether the file data is written to JSON formatting
		Format: "", // JsonFormat is false, logger message written to file format string
	}
	// add output to the file
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)


	logger.Info("this is a info log!")
	logger.Errorf("this is a error %s log!", "format")
}