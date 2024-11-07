package log

import (
	"fmt"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
)

func Panic(fileName, msg string, data ...interface{}) {
	withLevel(zerolog.PanicLevel, 3, fileName, msg, data...)
}

func Fatal(fileName, msg string, data ...interface{}) {
	withLevel(zerolog.FatalLevel, 3, fileName, msg, data...)
}

func Error(fileName, msg string, data ...interface{}) {
	withLevel(zerolog.ErrorLevel, 3, fileName, msg, data...)
}

func Warn(fileName, msg string, data ...interface{}) {
	withLevel(zerolog.WarnLevel, 3, fileName, msg, data...)
}

func Info(fileName, msg string, data ...interface{}) {
	withLevel(zerolog.InfoLevel, 3, fileName, msg, data...)
}

func Debug(fileName, msg string, data ...interface{}) {
	withLevel(zerolog.DebugLevel, 3, fileName, msg, data...)
}

func PanicF(fileName, msg string, v ...interface{}) {
	withLevel(zerolog.PanicLevel, 3, fileName, fmt.Sprintf(msg, v...))
}

func FatalF(fileName, msg string, v ...interface{}) {
	withLevel(zerolog.FatalLevel, 3, fileName, fmt.Sprintf(msg, v...))
}

func ErrorF(fileName, msg string, v ...interface{}) {
	withLevel(zerolog.ErrorLevel, 3, fileName, fmt.Sprintf(msg, v...))
}

func WarnF(fileName, msg string, v ...interface{}) {
	withLevel(zerolog.WarnLevel, 3, fileName, fmt.Sprintf(msg, v...))
}

func InfoF(fileName, msg string, v ...interface{}) {
	withLevel(zerolog.InfoLevel, 3, fileName, fmt.Sprintf(msg, v...))
}

func DebugF(fileName, msg string, v ...interface{}) {
	withLevel(zerolog.DebugLevel, 3, fileName, fmt.Sprintf(msg, v...))
}

func LogError(fileName string, err error, msgs ...string) {
	if err == nil {
		return
	}

	msg := err.Error()
	if len(msgs) == 1 {
		msg = msgs[0]
	} else if len(msgs) > 1 {
		msg, _ = sonic.MarshalString(msgs)
	}
	if fileName == FilePanic {
		withLevel(zerolog.PanicLevel, 3, fileName, msg, FormatErrorToJSON(err))
		return
	}

	withLevel(zerolog.ErrorLevel, 3, fileName, msg, FormatErrorToJSON(err))
}

func LogErrorf(fileName string, err error, msg string, v ...interface{}) {
	if err == nil {
		return
	}

	if msg == "" {
		msg = err.Error()
	}
	msg = fmt.Sprintf(msg, v...)
	if fileName == FilePanic {
		withLevel(zerolog.PanicLevel, 3, fileName, msg, FormatErrorToJSON(err))
		return
	}

	withLevel(zerolog.ErrorLevel, 3, fileName, msg, FormatErrorToJSON(err))
}

func FormatErrorToJSON(err error) string {
	str, marshalErr := sonic.MarshalString(eris.ToCustomJSON(err, eris.NewDefaultJSONFormat(eris.FormatOptions{
		InvertOutput: true, // Flag that inverts the error output (wrap errors shown first).
		WithTrace:    true, // Flag that enables stack trace output.
		InvertTrace:  true, // Flag that inverts the stack trace output (top of call stack shown first).
		WithExternal: true, // Flag that enables external error output.
	})))
	if marshalErr != nil {
		return eris.ToString(err, true)
	}
	return str
}

func FormatErrorToText(err error) string {
	upErr := eris.Unpack(err)
	str := ""
	if upErr.ErrExternal != nil {
		str += upErr.ErrExternal.Error() + "\n"
	}
	if upErr.ErrRoot.Msg != "" {
		str += upErr.ErrRoot.Msg + "\n"
	}

	for _, eStack := range upErr.ErrRoot.Stack {
		str += "\t" + eStack.File + ":" + strconv.Itoa(eStack.Line) + " " + eStack.Name + "\n"
	}
	for _, eLink := range upErr.ErrChain {
		str += eLink.Msg + "\n\t" + eLink.Frame.File + ":" + strconv.Itoa(eLink.Frame.Line) + " " + eLink.Frame.Name + "\n"
	}

	return str
}
