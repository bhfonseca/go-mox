package utils

import (
	"encoding/json"
	"strconv"
)

func SecondsToDays(seconds int) int {
	return seconds / 86400
}

func BytesToGigabytes(bytes json.Number) json.Number {
	const bytesPerGB = 1024 * 1024 * 1024
	floatBytes, err := bytes.Float64()
	if err != nil {
		return "0"
	}
	gb := floatBytes / float64(bytesPerGB)
	return json.Number(strconv.FormatFloat(gb, 'f', 2, 64))
}
