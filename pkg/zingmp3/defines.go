package zingmp3

const (
	UpStream  = "https://zing-mp3.glitch.me/?url="
	Song      = "https://zingmp3.vn/bai-hat"
	VideoClip = "https://zingmp3.vn/video-clip"
)

const (
	Normal int64 = 128
	VIP    int64 = 320
)

const (
	SD_360       int64 = 360
	SD_480       int64 = 480
	HD_720       int64 = 720
	FULL_HD_1080 int64 = 1080
)

type SongInfo struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Source struct {
		Video struct {
			Num360 struct {
				View     string `json:"view"`
				Download string `json:"download"`
			} `json:"360"`
			Num480 struct {
				View     string `json:"view"`
				Download string `json:"download"`
			} `json:"480"`
			Num720 struct {
				View     string `json:"view"`
				Download string `json:"download"`
			} `json:"720"`
			Num1080 struct {
				View     string `json:"view"`
				Download string `json:"download"`
			} `json:"1080"`
		} `json:"video"`
		Audio struct {
			Num128 struct {
				View     string `json:"view"`
				Download string `json:"download"`
			} `json:"128"`
			Num320 struct {
				View     string `json:"view"`
				Download string `json:"download"`
			} `json:"320"`
		} `json:"audio"`
	} `json:"source"`
}

type DownloadObject struct {
	Title       string
	Artist      string
	DownloadUrl string
	Type        string
}
