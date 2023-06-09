package gophertime_test

import (
	"testing"

	gophertime "github.com/dominic-wassef/gopher-time"
)

func TestGophertime_FrameToString_DF(t *testing.T) {
	cases := map[int64]string{
		2878: "00:02:00;02",
	}
	for f, tcode := range cases {
		tc := gophertime.FromFrame(f, gophertime.Rate_23_976)
		if str := tc.String(); str != tcode {
			t.Errorf("Frame %d should be equivalent to gophertime %s. Got %s\n", f, tcode, str)
		} else {
			t.Logf("Success, frame %d equals gophertime %s\n", f, tcode)
		}
	}
}

func TestGophertime_Identity_DF(t *testing.T) {
	cases := []string{
		"00:02:00;02",
		"00:00:00;00",
		"00:00:59;23",
		"00:01:00;02",
		"00:03:59;23",
		"00:04:00;02",
		"00:01:59;23",
		"00:09:59;23",
		"00:10:00;00",
	}
	for _, tcode := range cases {
		tc, _ := gophertime.Parse(tcode, gophertime.Rate_23_976)
		if str := tc.String(); str != tcode {
			t.Errorf("gophertime %s became %s during parsing and printing\n", tcode, str)
		} else {
			t.Logf("Success, identity valid for %s\n", tcode)
		}
	}
}

func TestGophertime_AddOne_DF(t *testing.T) {
	sequences := map[string]string{
		"00:00:59;23": "00:01:00;02",
		"00:03:59;23": "00:04:00;02",
		"00:01:59;23": "00:02:00;02",
		"00:09:59;23": "00:10:00;00",
	}
	for fromTC, toTC := range sequences {
		tc, _ := gophertime.Parse(fromTC, gophertime.Rate_23_976)
		next := tc.Add(gophertime.Frame(1))
		if str := next.String(); str != toTC {
			t.Errorf("Expected %s => %s, got %s\n", fromTC, toTC, str)
		} else {
			t.Logf("Success, got %s => %s\n", fromTC, toTC)
		}
	}
}

func bruteForceAdd1(c gophertime.Components) gophertime.Components {
	c.Frames++
	if c.Frames >= 24 {
		c.Frames -= 24
		c.Seconds++
		if c.Seconds >= 60 {
			c.Seconds -= 60
			c.Minutes++
			if c.Minutes >= 60 {
				c.Minutes -= 60
				c.Hours++
			}
		}
	}
	return c
}

func TestGophertimeSequenceNDF(t *testing.T) {
	startgophertimes := map[string]int{
		"00:00:00:00": 100000,
		"03:59:59:00": 100000,
	}
	for startgophertimeStr, iterations := range startgophertimes {
		prevTc, _ := gophertime.Parse(startgophertimeStr, gophertime.Rate_24)
		prevComp := prevTc.Components()

		for i := 0; i < iterations; i++ {
			tc := prevTc.Add(gophertime.Frame(1))
			comp := tc.Components()
			expectedComp := bruteForceAdd1(prevComp)
			if !comp.Equals(expectedComp) {
				t.Errorf("Add 1 frame, skipped from %s to %s\n", prevTc.String(), tc.String())
			}
			prevTc = tc
			prevComp = comp
		}
	}
}

func bruteForceAdd1_DF(c gophertime.Components) gophertime.Components {
	c.Frames++
	if c.Frames >= 24 {
		c.Frames -= 24
		c.Seconds++
		if c.Seconds >= 60 {
			c.Seconds -= 60
			c.Minutes++
			if c.Minutes >= 60 {
				c.Minutes -= 60
				c.Hours++
			}
		}
	}
	if (c.Minutes%10 > 0) && (c.Seconds == 0) && (c.Frames == 0 || c.Frames == 1) {
		c.Frames = 2
	}
	return c
}

func TestGophertimeSequenceDF(t *testing.T) {
	startgophertimes := map[string]int{
		"00:00:00:00": 100000,
		"03:59:59:00": 100000,
		"01:05:59:23": 100000,
	}
	for startgophertimeStr, iterations := range startgophertimes {
		prevTc, _ := gophertime.Parse(startgophertimeStr, gophertime.Rate_23_976)
		prevComp := prevTc.Components()

		for i := 0; i < iterations; i++ {
			tc := prevTc.Add(gophertime.Frame(1))
			comp := tc.Components()
			expectedComp := bruteForceAdd1_DF(prevComp)
			if !comp.Equals(expectedComp) {
				t.Errorf("Add 1 frame, skipped from %s to %s\n", prevTc.String(), tc.String())
			}
			prevTc = tc
			prevComp = comp
		}
	}
}
