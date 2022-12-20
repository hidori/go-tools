package runtime

import (
	"runtime"
	"strings"
)

type CallerInfo struct {
	OK         bool
	FileName   string
	LineNumber int
	FuncName   string
}

var callerInfoFileNamePrefixToTrim = ""

func SetCallerInfoFileNamePrefixToTrim(prefix string) string {
	prev := callerInfoFileNamePrefixToTrim
	callerInfoFileNamePrefixToTrim = prefix

	return prev
}

func GetCallerInfo() CallerInfo {
	const skip = 3

	return GetCallerInfoWithSkip(skip)
}

func GetCallerInfoWithSkip(skip int) CallerInfo {
	return GetCallerInfoWithSkipAndTrimFileName(skip, func(fileName string) string {
		if callerInfoFileNamePrefixToTrim != "" {
			fileName = strings.TrimPrefix(strings.TrimPrefix(fileName, callerInfoFileNamePrefixToTrim), "/")
		}

		return fileName
	})
}

func GetCallerInfoWithSkipAndTrimFileName(skip int, trim func(fileName string) string) CallerInfo {
	pc, fileName, lineNumber, ok := runtime.Caller(skip)

	if ok && trim != nil {
		fileName = trim(fileName)
	}

	return CallerInfo{
		OK:         ok,
		FileName:   fileName,
		LineNumber: lineNumber,
		FuncName:   runtime.FuncForPC(pc).Name(),
	}
}
