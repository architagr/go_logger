package logger

import (
	"bytes"
	"io"
)


type loggerTesting struct {
	name, logCategory     string
	message, formatSample string
	formatReplacer        []interface{}
	utc                   bool
	buf                   io.Writer
}

var testCases = []loggerTesting{
    {
        name:           "test longFile name",
        formatSample:   "%s|%s:%s|%s|%s",
        formatReplacer: []interface{}{Llongfile, Ldate, Ltime, LType, Lmsg},
        utc:            true,
        buf:            new(bytes.Buffer),
        logCategory:    "Debug",
        message:        "hi",
    },
    {
        name:           "test shortFile name",
        formatSample:   "%s|%s:%s|%s|%s",
        formatReplacer: []interface{}{Lshortfile, Ldate, Ltime, LType, Lmsg},
        utc:            true,
        buf:            new(bytes.Buffer),
        logCategory:    "Debug",
        message:        "hi",
    },
    {
        name:           "test logger format with constant string",
        formatSample:   "Test|%s|%s:%s|%s|%s",
        formatReplacer: []interface{}{Lshortfile, Ldate, Ltime, LType, Lmsg},
        utc:            true,
        buf:            new(bytes.Buffer),
        logCategory:    "Debug",
        message:        "hi",
    },
    {
        name:           "test logger with utc false",
        formatSample:   "Test|%s|%s:%s|%s|%s",
        formatReplacer: []interface{}{Lshortfile, Ldate, Ltime, LType, Lmsg},
        utc:            false,
        buf:            new(bytes.Buffer),
        logCategory:    "Debug",
        message:        "hi",
    },
}