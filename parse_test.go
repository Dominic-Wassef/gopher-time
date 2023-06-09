package gophertime_test

import (
	"testing"
	"time"

	gophertime "github.com/dominic-wassef/gopher-time"
)

func TestParse_NDF(t *testing.T) {
	cases := map[string]int64{
		"00:00:00:00": 0,
		"00:00:00:01": 1,
		"00:00:01:01": 61,
		"00:00:11:01": 661,
	}
	for k, f := range cases {
		tc, _ := gophertime.Parse(k, gophertime.Rate_60)
		if frame := tc.Frame(); frame != f {
			t.Errorf("gophertime %s should be equivalent to frame %d. Got %d\n", k, f, frame)
		} else {
			t.Logf("Success, gophertime %s equals frame %d\n", k, f)
		}
	}
}

func TestParse_DF(t *testing.T) {
	cases := map[string]int64{
		"00:02:00;02": 2878,
		"00:01:59;23": 2877,
		"00:01:59;22": 2876,
		"00:03:00;04": 4318,
	}
	for k, f := range cases {
		tc, _ := gophertime.Parse(k, gophertime.Rate_23_976)
		if frame := tc.Frame(); frame != f {
			t.Errorf("gophertime %s should be equivalent to frame %d. Got %d\n", k, f, frame)
		} else {
			t.Logf("Success, gophertime %s equals frame %d\n", k, f)
		}
	}
}

func TestParseInvalidDF(t *testing.T) {
	cases := map[string]string{
		"00:01:59;23": "00:01:59;23",
		"00:02:00;00": "00:02:00;02",
		"00:02:00;01": "00:02:00;02",
		"00:02:00;02": "00:02:00;02",
		"00:02:00;03": "00:02:00;03",
	}
	for k, s := range cases {
		tc, _ := gophertime.Parse(k, gophertime.Rate_23_976)
		if str := tc.String(); str != s {
			t.Errorf("DF gophertime %s should be rounded to gophertime %s. Got %s\n", k, s, str)
		} else {
			t.Logf("Success, DF gophertime %s rounded to %s\n", k, s)
		}
	}
}

func TestParseInvalidDFInNDFRate(t *testing.T) {
	cases := []string{
		"00:01:59:23",
		"00:02:00:00",
		"00:02:00:01",
		"00:02:00:02",
		"00:02:00:03",
	}
	for _, s := range cases {
		tc, _ := gophertime.Parse(s, gophertime.Rate_24)
		if str := tc.String(); str != s {
			t.Errorf("NDF gophertime %s should NOT be rounded. Got %s\n", s, str)
		} else {
			t.Logf("Success, NDF gophertime %s stayed the same\n", s)
		}
	}
}

func TestFromPresentationTime(t *testing.T) {
	cases := map[time.Duration]string{
		time.Minute * 2:             "00:02:00:00",
		time.Minute * 10:            "00:10:00:00",
		time.Second * 6:             "00:00:06:00",
		time.Hour*6 + time.Second/2: "06:00:00:12",
	}
	for pt, s := range cases {
		tc := gophertime.FromPresentationTime(pt, gophertime.Rate_24)
		if str := tc.String(); str != s {
			t.Errorf("Presentation time %s should be gophertime %s. Got %s\n", pt.String(), s, str)
		} else {
			t.Logf("Success, presentation time %s equals gophertime %s\n", pt.String(), s)
		}
	}
}
