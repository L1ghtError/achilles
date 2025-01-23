package m3u8_parser

import (
	"ahls_srvi/internal/strutil"
	"strings"
)

const EXT_CUE_OUT = "#EXT-X-CUE-OUT"
const EXT_CUE_IN = "#EXT-X-CUE-IN"
const EXT_INF = "#EXTINF"
const DURATION = "DURATION="
const COLON = ":"

func CountTotalINF(pl string) float32 {
	lines := strings.Split(pl, "\n")
	var accum float32 = 0
	for _, line := range lines {
		v := getTagDuration(line, EXT_INF)
		if v > 0 {
			accum += v
		}
	}
	return accum
}

func getTagDuration(p string, prefix string) float32 {
	if !strings.HasPrefix(p, prefix) {
		return -1
	}

	start := strings.Index(p, DURATION)
	if start < len(prefix) {
		start = strings.Index(p, COLON)
		if start < len(prefix) {
			return -1
		}
	}
	return strutil.ExtractNumber(p[start:])
}

func ModifyPlaylist(pl string, adSegs string) (string, bool) {
	lines := strings.Split(pl, "\n")
	var max float32 = 0
	ad_start, ad_end := -1, -1
	for i, line := range lines {
		dur := getTagDuration(line, EXT_CUE_OUT)
		if max < dur {
			max = dur
			ad_start = i
			ad_end = -1
		}
		if ad_end == -1 && strings.HasPrefix(line, EXT_CUE_IN) {
			ad_end = i
		}
	}
	if ad_start == -1 || ad_end == -1 {
		return "", false
	}
	segs := strings.Split(adSegs, "\n")
	for i := 0; i < len(segs) && ad_start+i < ad_end-1; i++ {
		lines[ad_start+i+1] = segs[i]
	}
	_ = ad_start
	return strings.Join(lines, "\n"), true
}
