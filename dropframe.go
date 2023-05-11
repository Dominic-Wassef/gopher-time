package gophertime

func CountDroppedFrames(minutes int64) int64 {

	notMultTen := minutes - (minutes / 10)

	return notMultTen * 2

}

func CountFramesToDrop(frames int64, rateNum int64) int64 {

	var minutes int64

	var droppedFrames int64

	for frames > 0 {

		frames -= rateNum * 60
		if frames < 0 {
			break
		}

		minutes++

		if minutes%10 > 0 {

			droppedFrames += 2

			frames += 2

		}

	}

	return droppedFrames

}
