package fmtrs

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon"
)

func FormatTimeToString(t time.Time, format string) string {
	return fmt.Sprint(carbon.Parse(t.Format(format)).DiffForHumans())
}

func FormatLogString(str, t string, code any) string {
	fmtStr := ""

	if code != 0 {
		fmtStr = fmt.Sprintf("| Code %s-%v: %s. Please see the Athenaeum manual for more information.", t, code, str)
	} else {
		fmtStr = fmt.Sprintf("| %s.", str)
	}

	return fmtStr
}
