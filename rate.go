package gophertime

import (
	"fmt"
	"time"
)

var (
	Rate_23_976 = Rate{24, true, 24000, 1001}

	Rate_24 = Rate{24, false, 24, 1}

	Rate_23_98 = Rate{24, false, 24000, 1001}
	Rate_30    = Rate{30, false, 30, 1}
	Rate_29_97 = Rate{30, true, 30000, 1001}
	Rate_60    = Rate{60, false, 60, 1}
	Rate_59_94 = Rate{60, false, 60000, 1001}
)

type Rate struct {
	Num                      int64
	DropFrame                bool
	TemporalNum, TemporalDen int64
}

func (r *Rate) String() string {
	return fmt.Sprintf("%d", r.Num)
}

func (r *Rate) PlaybackFrameDuration() time.Duration {
	return time.Second * time.Duration(r.TemporalDen) / time.Duration(r.TemporalNum)
}
