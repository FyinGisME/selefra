package ui

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/hashicorp/go-hclog"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra/global"
	"github.com/selefra/selefra/pkg/logger"
	"github.com/selefra/selefra/pkg/ws"
	"runtime"
	"time"
)

var defaultLogger *logger.Logger

const (
	prefixManaged   = "managed"
	prefixUnmanaged = "unmanaged"
	defaultAlias    = "default"
)

func InitLogger() {
	defaultLogger, _ = logger.NewLogger(logger.Config{
		FileLogEnabled:    true,
		ConsoleLogEnabled: false,
		EncodeLogsAsJson:  true,
		ConsoleNoColor:    true,
		Source:            "client",
		Directory:         "logs",
		Level:             "info",
	})
}

var (
	ErrorColor   = color.New(color.FgRed, color.Bold)
	WarningColor = color.New(color.FgYellow, color.Bold)
	InfoColor    = color.New(color.FgWhite, color.Bold)
	SuccessColor = color.New(color.FgGreen, color.Bold)
	BaseColor    = color.New(color.FgBlack, color.Bold)
)

type LogJOSN struct {
	Cmd   string    `json:"cmd"`
	Stag  string    `json:"stag"`
	Msg   string    `json:"msg"`
	Time  time.Time `json:"time"`
	Level string    `json:"level"`
}

func createLog(msg string, c *color.Color) string {
	var level string
	switch c {
	case ErrorColor:
		level = "error"
	case WarningColor:
		level = "warn"
	case InfoColor:
		level = "info"
	case SuccessColor:
		level = "success"
	case BaseColor:
		level = "base"
	default:
		level = "info"
	}
	l := LogJOSN{
		Cmd:   global.CMD,
		Stag:  global.STAG,
		Msg:   msg,
		Time:  time.Now(),
		Level: level,
	}
	b, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(b)
}

func PrintErrorF(format string, a ...interface{}) {
	InitLogger()
	_, f, l, ok := runtime.Caller(1)
	if ok {
		if defaultLogger != nil {
			defaultLogger.Log(hclog.Error, "%s %s:%d", fmt.Sprintf(format, a...), f, l)
		}
	}
	PrintCustomizeF(ErrorColor, format, a...)
}

func PrintWarningF(format string, a ...interface{}) {
	InitLogger()
	if defaultLogger != nil {
		defaultLogger.Log(hclog.Warn, format, a...)
	}
	PrintCustomizeF(WarningColor, format, a...)

}

func PrintSuccessF(format string, a ...interface{}) {
	InitLogger()
	if defaultLogger != nil {
		defaultLogger.Log(hclog.Info, format, a...)
	}
	PrintCustomizeF(SuccessColor, format, a...)

}

func PrintInfoF(format string, a ...interface{}) {
	InitLogger()
	if defaultLogger != nil {
		defaultLogger.Log(hclog.Info, format, a...)
	}
	PrintCustomizeF(InfoColor, format, a...)
}

func PrintErrorLn(a ...interface{}) {
	InitLogger()
	_, f, l, ok := runtime.Caller(1)
	if ok {
		if defaultLogger != nil {
			defaultLogger.Log(hclog.Error, "%s %s:%d", fmt.Sprintln(a...), f, l)
		}
	}
	PrintCustomizeLn(ErrorColor, a...)
}

func PrintWarningLn(a ...interface{}) {
	InitLogger()
	if defaultLogger != nil {
		defaultLogger.Log(hclog.Warn, fmt.Sprintln(a...))
	}
	PrintCustomizeLn(WarningColor, a...)

}

func PrintSuccessLn(a ...interface{}) {
	InitLogger()
	if defaultLogger != nil {
		defaultLogger.Log(hclog.Info, fmt.Sprintln(a...))
	}
	PrintCustomizeLn(SuccessColor, a...)

}

func PrintInfoLn(a ...interface{}) {
	InitLogger()
	if defaultLogger != nil {
		defaultLogger.Log(hclog.Info, fmt.Sprintln(a...))
	}
	PrintCustomizeLn(InfoColor, a...)
}

func PrintCustomizeF(c *color.Color, format string, a ...interface{}) {
	ws.SendLog(createLog(fmt.Sprintf(format, a...), c))
	_, _ = c.Printf(format+"\n", a...)
}

func PrintCustomizeFNotN(c *color.Color, format string, a ...interface{}) {
	ws.SendLog(createLog(fmt.Sprintf(format, a...), c))
	_, _ = c.Printf(format, a...)
}

func PrintCustomizeLn(c *color.Color, a ...interface{}) {
	ws.SendLog(createLog(fmt.Sprintln(a...), c))
	_, _ = c.Println(a...)
}

func PrintCustomizeLnNotShow(a ...interface{}) {
	ws.SendLog(createLog(fmt.Sprintln(a...), InfoColor))
}
func SaveLogToDiagnostic(diagnostics []*schema.Diagnostic) error {
	InitLogger()
	var err error
	for i := range diagnostics {
		switch diagnostics[i].Level() {
		case schema.DiagnosisLevelError:
			_, f, l, ok := runtime.Caller(1)
			if ok {
				if defaultLogger != nil {
					defaultLogger.Log(hclog.Error, "%s %s:%d", diagnostics[i].Content(), f, l)
				}
			}
			err = errors.New(diagnostics[i].Content())
		case schema.DiagnosisLevelWarn:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Warn, diagnostics[i].Content())
			}
		case schema.DiagnosisLevelInfo:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Info, diagnostics[i].Content())
			}
		case schema.DiagnosisLevelDebug:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Debug, diagnostics[i].Content())
			}
		case schema.DiagnosisLevelTrace:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Trace, diagnostics[i].Content())
			}
		case schema.DiagnosisLevelFatal:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Info, diagnostics[i].Content())
			}
		}
	}
	return err
}

func PrintDiagnostic(diagnostics []*schema.Diagnostic) error {
	InitLogger()
	var err error
	for i := range diagnostics {
		switch diagnostics[i].Level() {
		case schema.DiagnosisLevelError:
			_, f, l, ok := runtime.Caller(1)
			if ok {
				if defaultLogger != nil {
					defaultLogger.Log(hclog.Error, "%s %s:%d", diagnostics[i].Content(), f, l)
				}
			}
			PrintCustomizeLn(ErrorColor, diagnostics[i].Content())
			err = errors.New(diagnostics[i].Content())
		case schema.DiagnosisLevelWarn:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Warn, diagnostics[i].Content())
			}
			PrintWarningLn(diagnostics[i].Content())
		case schema.DiagnosisLevelInfo:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Info, diagnostics[i].Content())
			}
			PrintInfoLn(diagnostics[i].Content())
		case schema.DiagnosisLevelDebug:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Debug, diagnostics[i].Content())
			}
			PrintSuccessLn(diagnostics[i].Content())
		case schema.DiagnosisLevelTrace:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Trace, diagnostics[i].Content())
			}
			PrintInfoLn(diagnostics[i].Content())
		case schema.DiagnosisLevelFatal:
			if defaultLogger != nil {
				defaultLogger.Log(hclog.Info, diagnostics[i].Content())
			}
			PrintErrorLn(diagnostics[i].Content())
		}
	}
	return err
}
