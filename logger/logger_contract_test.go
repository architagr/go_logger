package logger

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestFormatValidator(t *testing.T) {

	t.Run("check if it is validating correct formats", func(test *testing.T) {
		replacersGot, err := validateFormat("logger : %lda : %lty")
		if err != nil {
			test.Fatal(err)
		}
		replacersWant := []string{"logger : ", "%lda", " : ", "%lty"}
		if !reflect.DeepEqual(replacersGot, replacersWant) {
			test.Fatalf("replacers found are not correct as we have got: %+v and we want %+v", replacersGot, replacersWant)
		}
	})
	t.Run("check if it it gives error if format is empty", func(test *testing.T) {
		_, err := validateFormat("")
		if err == nil || err.Error() != "formater is empty" {
			test.Fatal()
		}
	})
	t.Run("check if gives error for not valid replacer", func(test *testing.T) {
		_, err := validateFormat("logger : %ldt : %utc")
		if err == nil || err.Error() != "invalid replacer %ldt" {
			test.Fatal()
		}
	})

	t.Run("check if gives error for no replacer found", func(test *testing.T) {
		_, err := validateFormat("logger")
		if err == nil || err.Error() != "invalid format, no replacer found" {
			test.Fatal()
		}
	})
}

func TestDebugLogging(t *testing.T) {
	for testcaseIndex, testCase := range testCases {
		testCaseName := fmt.Sprintf("test case number %d - ", testcaseIndex+1) + testCase.name
		t.Run(testCaseName, func(t *testing.T) {
			buf := new(bytes.Buffer)
			formatSample := testCase.formatSample
			formatString := fmt.Sprintf(formatSample, testCase.formatReplacer...)
			logger, err := NewLogger(buf, formatString, testCase.utc, 2)
			if err != nil {
				t.Error(err)
			}
			logger.Debug(testCase.message)

			got := buf.String()
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			}
			now := time.Now()
			if testCase.utc{
				now = now.UTC()
			}
			expected := fmt.Sprintf(formatSample+"\n", buildExpectedLog(testCase, now, file, line)...)
			fmt.Printf("testCaseName:%s - expected:[%s]", testCaseName, expected)
			if got != expected {
				t.Errorf("\nexpected:[%s]-\ngot:[%s]-", expected, got)
			}
		})
	}
}

func buildExpectedLog(testCase loggerTesting, t time.Time, file string, line int) []interface{} {
	output := make([]interface{}, 0)
	for _, replacer := range testCase.formatReplacer {
		switch replacer {
		case Lmsg:
			output = append(output, testCase.message)
		case Ldate:
			output = append(output, getDate(t))
		case Ltime:
			output = append(output, getTime(t))
		case Lmicroseconds:
			output = append(output, getMicrosecond(t))
		case Llongfile:
			output = append(output, getLongFile(file, line))
		case Lshortfile:
			output = append(output, getShortFile(file, line))
		case LType:
			output = append(output, testCase.logCategory)
		}
	}
	return output
}
