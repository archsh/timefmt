package timefmt

import (
    "testing"
    "time"
)

var test_cases map[string]string = map[string]string{
    "%Y-%m-%dT%H:%M:%S":     "2016-09-22T06:04:26",
    "%y-%m-%dT%H:%M:%S":     "16-09-22T06:04:26",
    "%Y-%m-%dT%I:%M:%S":     "2016-09-22T06:04:26",
    "%Y-%m-%dT %p %I:%M:%S": "2016-09-22T AM 06:04:26",
    "%Y-%b-%dT%H:%M:%S":     "2016-Sep-22T06:04:26",
    "%Y-%B-%dT%H:%M:%S":     "2016-September-22T06:04:26",
    "%Y-%b-%dT%H:%-M:%S":    "2016-Sep-22T06:4:26",
    "%c":                    "Thu Sep 22 06:04:26 2016",
    "%x":                    "09/22/16",
    "%X":                    "06:04:26",
    "%Y-%m-%dT%H:%M:%S %z":  "2016-09-22T06:04:26 +0000",
}

func TestStrftime(t *testing.T) {
    var validate = func(tm time.Time, format string, result string) {
        if s, e := Strftime(tm, format); e != nil || s != result {
            t.Errorf("Strftime(/%v/, '%s') should return '%s' but not (%s) (%s)\n", tm, format, result, e, s)
        } else {
            t.Logf("%s: %s => %s", tm, format, s)
        }
    }
    loc, _ := time.LoadLocation("UTC")
    tm := time.Unix(1474524266, 321).In(loc)

    tm.Location()
    validate(tm, "%Y-%m-%dT%H:%M:%S", "2016-09-22T06:04:26")
    validate(tm, "%y-%m-%dT%H:%M:%S", "16-09-22T06:04:26")
    validate(tm, "%Y-%m-%dT%I:%M:%S", "2016-09-22T06:04:26")
    validate(tm, "%Y-%m-%dT %p %I:%M:%S", "2016-09-22T AM 06:04:26")
    validate(tm, "%Y-%b-%dT%H:%M:%S", "2016-Sep-22T06:04:26")
    validate(tm, "%Y-%B-%dT%H:%M:%S", "2016-September-22T06:04:26")
    validate(tm, "%Y-%b-%dT%H:%-M:%S", "2016-Sep-22T06:4:26")
    validate(tm, "%c", "Thu Sep 22 06:04:26 2016")
    validate(tm, "%x", "09/22/16")
    validate(tm, "%X", "06:04:26")
    validate(tm, "%Y-%m-%dT%H:%M:%S %z", "2016-09-22T06:04:26 +0000")
}

func TestStrptime(t *testing.T) {
    var validate = func(val string, format string, result time.Time) {
        if tm, e := Strptime(val, format); e != nil || tm != result {
            t.Errorf("Strptime('%s', '%s') should return /%v/ but not (%v) (%s) \n", val, format, result, tm, e)
            t.Errorf("%v -%v = %v \n", result, tm, result.Sub(tm))
        }
    }
    //loc, _ := time.LoadLocation("Asia/Shanghai")
    //tm := time.Unix(1474524266, 321000).In(loc)
    //validate("2016-Sep-22T14:04:26.000321 Asia/Shanghai", "%Y-%b-%dT%H:%M:%S.%f %Z", tm)
    loc, _ := time.LoadLocation("UTC")
    tm := time.Unix(1474524266, 321000).In(loc)
    validate("2016-Sep-22T06:04:26.000321 UTC", "%Y-%b-%dT%H:%M:%S.%f %Z", tm)
}

func BenchmarkStrftime(b *testing.B) {
    loc, _ := time.LoadLocation("UTC")
    tm := time.Unix(1474524266, 321).In(loc)
    for n := 0; n < b.N; n++ {
        _, _ = Strftime(tm, "%Y-%m-%dT%H:%M:%S %z %Z %p %b %B %a %A")
    }
}

func BenchmarkStrptime(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _, _ = Strptime("2016-Sep-22T06:04:26.000321 UTC", "%Y-%b-%dT%H:%M:%S.%f %Z")
    }
}
