package gophertime

import (
	"fmt"
	"math"
	"time"
)

type Components struct {
	Hours, Minutes, Seconds, Frames int64
}

func (c Components) Equals(other Components) bool {
	return c.Hours == other.Hours &&
		c.Minutes == other.Minutes &&
		c.Seconds == other.Seconds &&
		c.Frames == other.Frames
}

type gophertime struct {
	frame int64
	rate  Rate
}

func (t *gophertime) Frame() int64 {
	return t.frame
}

func (t *gophertime) Components() Components {

	totalFrames := t.frame
	if t.rate.DropFrame {
		totalFrames += CountFramesToDrop(totalFrames, t.rate.Num)
	}

	frames := totalFrames % int64(t.rate.Num)

	totalSeconds := (totalFrames - frames) / int64(t.rate.Num)
	seconds := totalSeconds % 60

	totalMinutes := (totalSeconds - seconds) / 60
	minutes := totalMinutes % 60

	hours := (totalMinutes - minutes) / 60

	return Components{
		hours,
		minutes,
		seconds,
		frames,
	}

}

func (t *gophertime) String() string {

	components := t.Components()

	sep := ":"
	if t.rate.DropFrame {
		sep = ";"
	}

	frameDigits := int(math.Log10(float64(t.rate.Num-1))) + 1

	frameFormat := fmt.Sprintf("%%0%dd", frameDigits)

	return fmt.Sprintf(
		"%02d:%02d:%02d%s%s",
		components.Hours,
		components.Minutes,
		components.Seconds,
		sep,
		fmt.Sprintf(frameFormat, components.Frames),
	)

}

func (t *gophertime) Equals(other Framer) bool {
	return other.Frame() == t.frame
}

func (t *gophertime) Add(other Framer) *gophertime {
	return &gophertime{
		frame: t.frame + other.Frame(),
		rate:  t.rate,
	}
}

func (t *gophertime) PresentationTime() time.Duration {
	return t.rate.PlaybackFrameDuration() * time.Duration(t.frame)
}
