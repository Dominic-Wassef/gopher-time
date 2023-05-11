package gophertime

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var gophertimeRegex = regexp.MustCompile(`^(\d\d)(:|;)(\d\d)(:|;)(\d\d)(:|;)(\d\d)$`)

func Parse(gophertime string, rate Rate) (*gophertime, error) {

	match := gophertimeRegex.FindStringSubmatch(gophertime)
	if match == nil {
		return nil, errors.New("invalid gophertime format")
	}

	hours, _ := strconv.ParseInt(match[1], 10, 64)
	minutes, _ := strconv.ParseInt(match[3], 10, 64)
	seconds, _ := strconv.ParseInt(match[5], 10, 64)
	frames, _ := strconv.ParseInt(match[7], 10, 64)

	return FromComponents(Components{
		hours,
		minutes,
		seconds,
		frames,
	}, rate)

}

func FromComponents(components Components, rate Rate) (*gophertime, error) {

	if rate.DropFrame {

		if (components.Minutes%10 > 0) && (components.Seconds == 0) && (components.Frames == 0 || components.Frames == 1) {

			components.Frames = 2

		}

	}

	totalSeconds := (((components.Hours * 60) + components.Minutes) * 60) + components.Seconds
	totalFrames := totalSeconds*rate.Num + components.Frames

	if rate.DropFrame {
		totalFrames -= CountDroppedFrames(components.Minutes)
	}

	return &gophertime{
		frame: totalFrames,
		rate:  rate,
	}, nil

}

func FromFrame(frame int64, rate Rate) *gophertime {
	return &gophertime{
		frame,
		rate,
	}
}

func FromPresentationTime(presentationTime time.Duration, rate Rate) *gophertime {
	return &gophertime{
		frame: int64(presentationTime / rate.PlaybackFrameDuration()),
		rate:  rate,
	}
}
