package cmd

import (
	"fmt"
	"github.com/ducmeit1/zmp3/pkg"
	"github.com/ducmeit1/zmp3/pkg/common"
	"github.com/ducmeit1/zmp3/pkg/zingmp3"
	"github.com/spf13/cobra"
	"strings"
)

var config *pkg.Config

var downloadSongCmd = &cobra.Command{
	Use:   "song",
	Short: "Download MP3 song from Zing MP3",
	Long:  "Download MP3 song from Zing MP3",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !pkg.IsConfigFileExist() {
			err := pkg.WriteDefaultConfig()
			if err != nil {
				return err
			}
		}

		cfg, err := pkg.ReadConfigFile()
		if err != nil {
			return err
		}

		err = cfg.IsValidConfig()
		if err != nil {
			return err
		}

		config = cfg
		config.CreateDownloadFolderIfNotExist()
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return downloadSong()
	},
}

var downloadVideoCmd = &cobra.Command{
	Use:   "video",
	Short: "Download MP3 video from Zing MP3",
	Long:  "Download MP3 video from Zing MP3",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !pkg.IsConfigFileExist() {
			err := pkg.WriteDefaultConfig()
			if err != nil {
				return err
			}
		}

		cfg, err := pkg.ReadConfigFile()
		if err != nil {
			return err
		}

		err = cfg.IsValidConfig()
		if err != nil {
			return err
		}

		config = cfg
		config.CreateDownloadFolderIfNotExist()
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return downloadVideo()
	},
}

func downloadSong() error {
	url, err := common.PromptString("Url")
	if err != nil {
		return err
	}

	if !strings.Contains(url, zingmp3.Song) {
		return common.InvalidSongUrl
	}

	song, err := zingmp3.GetDownloadLinks(url)
	if err != nil {
		return err
	}

	downloadObj := &zingmp3.DownloadObject{
		Title: song.Title,
		Artist: strings.ReplaceAll(song.Artist, " ", ""),
		Type: "mp3",
	}

	switch config.Mp3Quality {
	case zingmp3.Normal:
		downloadObj.DownloadUrl = song.Source.Audio.Num128.Download
		break
	case zingmp3.VIP:
		downloadObj.DownloadUrl = song.Source.Audio.Num320.Download
		if downloadObj.DownloadUrl == "" {
			fmt.Println("This song is not supported 320kbps quality, 128kbps will be downloaded instead")
			downloadObj.DownloadUrl = song.Source.Audio.Num128.Download
		}
		break
	}

	err = zingmp3.Download(downloadObj, config.GetDownloadFolder())
	if err != nil {
		return err
	}

	return nil
}

func downloadVideo() error {
	url, err := common.PromptString("Url")
	if err != nil {
		return err
	}

	if !strings.Contains(url, zingmp3.VideoClip) {
		return common.InvalidVideoUrl
	}

	video, err := zingmp3.GetDownloadLinks(url)
	if err != nil {
		return err
	}

	downloadObj := &zingmp3.DownloadObject{
		Title: video.Title,
		Artist: strings.ReplaceAll(video.Artist, " ", ""),
		Type: "mp4",
	}

	switch config.Mp4Quality {
	case zingmp3.SD_360:
		downloadObj.DownloadUrl = video.Source.Video.Num360.Download
		break
	case zingmp3.SD_480:
		downloadObj.DownloadUrl = video.Source.Video.Num480.Download
		break
	case zingmp3.HD_720:
		downloadObj.DownloadUrl = video.Source.Video.Num720.Download
		break
	case zingmp3.FULL_HD_1080:
		downloadObj.DownloadUrl = video.Source.Video.Num1080.Download
		break
	}

	if downloadObj.DownloadUrl == "" {
		fmt.Printf("Video with quality %dp is not exist, default quality 360p will be downloaded instead", config.Mp4Quality)
		downloadObj.DownloadUrl = video.Source.Video.Num360.Download
	}

	err = zingmp3.Download(downloadObj, config.GetDownloadFolder())
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(downloadSongCmd)
	rootCmd.AddCommand(downloadVideoCmd)
}
