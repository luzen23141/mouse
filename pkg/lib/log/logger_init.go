package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// 第一个日期，第二个为文件名
	loggerMap = make(map[string]map[string]*zerolog.Logger)
	logFile   = "./log/%s-%s.log" // 一定要有两个%s 第一个会是文件名，第二个会是日期
)

func getLogger(fileName string) *zerolog.Logger {
	date := time.Now().Format("2006-01-02")
	if _, ok := loggerMap[date]; !ok {
		loggerMap[date] = make(map[string]*zerolog.Logger)
	}

	if len(loggerMap) > 2 {
		mutex := &sync.Mutex{}
		mutex.Lock()

		yesterdayDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		for k := range loggerMap {
			if k == date || k == yesterdayDate {
				continue
			}
			for k1 := range loggerMap[k] {
				delete(loggerMap[k], k1)
			}
			delete(loggerMap, k)
		}
		mutex.Unlock()
	}

	if _, ok := loggerMap[date][fileName]; !ok {
		logger := GetZeroLogger(fmt.Sprintf(logFile, fileName, date))
		mutex := &sync.Mutex{}
		mutex.Lock()
		loggerMap[date][fileName] = &logger
		mutex.Unlock()
	}

	return loggerMap[date][fileName]
}

func GetZeroLogger(name string) zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out: &lumberjack.Logger{
			Filename:   name,
			MaxSize:    100, // MB
			MaxBackups: 10,  // 最多保留3个备份
			MaxAge:     30,  // 天数
			Compress:   true,
		},
		TimeFormat:       "[15:04:05.000]",
		NoColor:          true,
		FormatMessage:    func(i interface{}) string { return i.(string) },
		FormatFieldName:  func(i interface{}) string { return i.(string) + "=" },
		FormatFieldValue: func(i interface{}) string { return i.(string) },
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	return zerolog.New(output).With().Timestamp().Logger()
}

func withLevel(level zerolog.Level, skip int, fileName, msg string, data ...interface{}) {
	if len(data) == 1 && isString(data[0]) {
		msg += " " + data[0].(string)
	} else if len(data) > 0 {
		rawJson, err := sonic.MarshalString(data)
		if err != nil {
			fmt.Println(err)
		}
		msg += " " + rawJson
	}
	msg = getMsgPrefix(skip) + msg
	getLogger(fileName).WithLevel(level).Msg(msg)
}

func isString(s interface{}) bool {
	_, ok := s.(string)
	return ok
}

func getMsgPrefix(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	paths := strings.Split(file, "/")
	l := len(paths)
	if l > 2 {
		file = paths[l-2] + "/" + paths[l-1]
	}
	return file + ":" + strconv.Itoa(line) + " "
}
