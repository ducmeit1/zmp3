package common

import "errors"

var (
	InvalidZingMp3Url = errors.New("invalid url, the url must be start with https://mp3.zing.vn")
	InvalidSongUrl = errors.New("invalid song url, must be start with https://mp3.zing.vn/bai-hat")
	InvalidVideoUrl = errors.New("invalid video url, must be start with https://mp3.zing.vn/video-clip")
	InvalidMp3Quality = errors.New("invalid mp3 quality, only supported 128 kbps and 320 kbps")
	InvalidMp4Quality = errors.New("invalid mp4 quality, only supported 360p, 480p, 720p, 1080p")
	InvalidDownloadObject = errors.New("download object must be not empty")
	)