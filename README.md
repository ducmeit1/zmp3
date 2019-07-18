#DOWNLOAD SONG & VIDEO FROM ZING MP3

zmp3 is a CLI tool helps you download highest quality of music and video on https://mp3.zing.vn for free

- Programing language: Go
- Dependencies:
    - Cobra
    - Viper
    - PromptUI
    - Bar
    - zing-mp3.glitch.me
    
- Quality are supported:
    - MP3:
        - 128 kbps
        - 320 kbps
    - MP4:
        - 360p
        - 480p
        - 720p
        - 1080p

- Usages:
    - Installing:
        ```
        cd $GOPATH/github.com/ducmeit1/zmp3-golang
        go install ./...
        ```
    - Command: *zmp3*
        - Helper: *zmp3 -h*
        - Download Song: *zmp3 song*
        - Download Video: *zmp3 video*
        - Etc...

- Features:
    - Download Song
    - Download Video
    - Set custom configuration (Quality of MP3, MP4, Directory of downloader folder)

*Copyright @ 2019 - ducmeit1*     


