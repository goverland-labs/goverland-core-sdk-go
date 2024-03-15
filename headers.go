package goverlandcorewebsdk

import (
	"net/http"
	"strconv"
)

const (
	HeaderTotalCount    = "X-Total-Count"
	HeaderTotalVp       = "X-Total-Vp"
	HeaderCurrentOffset = "X-Offset"
	HeaderLimit         = "X-Limit"
)

func GetOffsetFromHeaders(headers http.Header) int {
	return getNumberFromHeaders(headers, HeaderCurrentOffset)
}

func GetLimitFromHeaders(headers http.Header) int {
	return getNumberFromHeaders(headers, HeaderLimit)
}

func GetTotalCntFromHeaders(headers http.Header) int {
	return getNumberFromHeaders(headers, HeaderTotalCount)
}

func GetTotalVpFromHeaders(headers http.Header) float32 {
	data, ok := headers[HeaderTotalVp]
	if !ok || len(data) != 1 {
		return 0
	}
	vp, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return 0
	}

	return float32(vp)
}

func getNumberFromHeaders(headers http.Header, name string) int {
	data, ok := headers[name]
	if !ok || len(data) != 1 {
		return 0
	}

	number, err := strconv.Atoi(data[0])
	if err != nil {
		return 0
	}

	return number
}
