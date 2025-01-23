package m3u8_parser

import (
	"math"
	"testing"
)

func TestGetCueDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected float32
	}{
		{"DURATION=12#EXT-X-CUE-OUT: followed by some more text", -1},
		{"#EXT-X-CUE-OUT:DURATION=\"201.467\"", 201.467},
		{"#EXT-X-CUE-OUT:12", 12},
		{"No relevant information here", -1},
		{"", -1},
		{"#EXT-X-CUE-OUT:\"12.12\"", 12.12},
		{"DURATION followed by some more text", -1},
		{"#EXT-X-CUE-OUT:99999.22", 99999.22},
		{"312", -1},
	}
	tolerance := float32(0.05)
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := getTagDuration(test.input, EXT_CUE_OUT)
			if math.Abs(float64(result-test.expected)) > float64(tolerance) {
				t.Errorf("getTagDuration(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestModifyPlaylist(t *testing.T) {
	tests := []struct {
		playlist         string
		adSegments       string
		expectedPlaylist string
		expectedSuccess  bool
	}{
		{`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:8
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:4.379378,
seg0.ts
#EXTINF:6.297956,
seg1.ts
#EXT-X-CUE-OUT:20.000
#EXTINF:3.999744,
blank0.ts
#EXTINF:4.000600,
blank1.ts
#EXTINF:3.999567,
blank2.ts
#EXTINF:4.000144,
blank3.ts
#EXTINF:3.999989,
blank4.ts
#EXT-X-CUE-IN
#EXTINF:3.169833,
seg2.ts
#EXT-X-ENDLIST`,
			`#EXTINF:3.999743,
adseg0.ts
#EXTINF:4.000599,
adseg1.ts
#EXTINF:3.999566,
adseg2.ts
#EXTINF:4.000143,
adseg3.ts
#EXTINF:3.999988,
adseg4.ts`,
			`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:8
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:4.379378,
seg0.ts
#EXTINF:6.297956,
seg1.ts
#EXT-X-CUE-OUT:20.000
#EXTINF:3.999743,
adseg0.ts
#EXTINF:4.000599,
adseg1.ts
#EXTINF:3.999566,
adseg2.ts
#EXTINF:4.000143,
adseg3.ts
#EXTINF:3.999988,
adseg4.ts
#EXT-X-CUE-IN
#EXTINF:3.169833,
seg2.ts
#EXT-X-ENDLIST`, true},
		{`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:8
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:4.379378,
seg0.ts
#EXTINF:6.297956,
seg1.ts
#EXT-X-CUE-OUT:12.000
#EXTINF:3.999744,
blank0.ts
#EXTINF:4.000600,
blank1.ts
#EXTINF:3.999567,
blank2.ts
#EXT-X-CUE-IN
#EXTINF:3.169833,
seg2.ts
#EXTINF:4.389323,
seg3.ts
#EXT-X-CUE-OUT:16.000
#EXTINF:3.999744,
blank0.ts
#EXTINF:4.000600,
blank1.ts
#EXTINF:3.999567,
blank2.ts
#EXTINF:4.000507,
blank3.ts
#EXT-X-CUE-IN
#EXTINF:3.989275,
seg4.ts
#EXT-X-ENDLIST`,
			`#EXTINF:3.999743,
adseg0.ts
#EXTINF:4.000599,
adseg1.ts
#EXTINF:3.999566,
adseg2.ts
#EXTINF:4.000143,
adseg3.ts
#EXTINF:3.999988,
adseg4.ts`,
			`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:8
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:4.379378,
seg0.ts
#EXTINF:6.297956,
seg1.ts
#EXT-X-CUE-OUT:12.000
#EXTINF:3.999744,
blank0.ts
#EXTINF:4.000600,
blank1.ts
#EXTINF:3.999567,
blank2.ts
#EXT-X-CUE-IN
#EXTINF:3.169833,
seg2.ts
#EXTINF:4.389323,
seg3.ts
#EXT-X-CUE-OUT:16.000
#EXTINF:3.999743,
adseg0.ts
#EXTINF:4.000599,
adseg1.ts
#EXTINF:3.999566,
adseg2.ts
#EXTINF:4.000143,
adseg3.ts
#EXT-X-CUE-IN
#EXTINF:3.989275,
seg4.ts
#EXT-X-ENDLIST`, true},
		{"", "", "", false},
	}
	for _, test := range tests {
		t.Run(test.playlist, func(t *testing.T) {
			res, succes := ModifyPlaylist(test.playlist, test.adSegments)
			if res != test.expectedPlaylist || succes != test.expectedSuccess {
				t.Errorf("ModifyPlaylist(%q and %q) = %v; want %v", test.playlist, test.adSegments, res, test.expectedPlaylist)
			}
		})
	}
}
