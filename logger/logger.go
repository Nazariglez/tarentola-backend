// Created by nazarigonzalez on 2/10/17.

package logger

import (
	_logger "github.com/nazariglez/logger"
	"github.com/nazariglez/tarentola-backend/config"
	"strings"
)

var Log = func() *_logger.Logger {
	l := _logger.New()
	l.SetLevel(_logger.LogLevel(config.Data.Logger.Level))

	if config.Data.Logger.File {
		err := l.EnableFileOutput(strings.ToLower(config.Data.Name), config.Data.Logger.Path, _logger.LogLevel(config.Data.Logger.FileLevel))
		if err != nil {
			l.DisableFileOutput()
			l.Error(err.Error())
		}
	}

	return l
}()
