> This is open source project and has posted on Medium at: https://medium.com/@ducmeit/make-a-great-downloader-cli-by-cobra-viper-promptui-at-golang-f6408434cafa

zmp3 is a CLI tool helps you download highest quality of music and video on https://mp3.zing.vn for free

- Installing: `go install github.com/ducmeit1/zmp3`
- Downloading at: `https://github.com/ducmeit1/zmp3/releases`
- CLI:
![Demo](png/demo.png)

- Progress bar when downloading
![Progress](png/progress.png)

- Result after downloaded complete
![Result](png/result.png)

- Notice: Use can download binary file `zmp3` to use as a console application
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
        go install github.com/ducmeit1/zmp3
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
- Licensed:
This source code is under MIT license

*Copyright @ 2019 - ducmeit1*     


