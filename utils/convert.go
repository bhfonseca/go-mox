package utils

func SecondsToDays(seconds int) int {
	return seconds / 86400
}

func SecondsToHours(seconds int) int {
	return seconds / 3600
}

func SecondsToMinutes(seconds int) int {
	return seconds / 60
}

func BytesToGigabytes(bytes uint64) float64 {
	const bytesPerGB = 1024 * 1024 * 1024
	return float64(bytes) / float64(bytesPerGB)
}
