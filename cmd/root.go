package cmd

import (
	"fmt"
	"github.com/ducmeit1/simple-cli-at-golang/pkg"
	"github.com/ducmeit1/simple-cli-at-golang/pkg/common"
	"github.com/spf13/cobra"
	"os"
)

var (
	version    bool
	showConfig bool
	setConfig  bool
)

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show current version of CLI")
	rootCmd.Flags().BoolVarP(&showConfig, "show config", "s", false, "show current configuration")
	rootCmd.Flags().BoolVarP(&setConfig, "set config", "c", false, "set new configuration")
}

var rootCmd = &cobra.Command{
	Use:   "zmp3",
	Short: "A simple CLI for Golang use to download Song/Video from Zing Mp3",
	Long:  "A simple CLI for Golang use to download Song/Video from Zing Mp3",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !pkg.IsConfigFileExist() {
			err := pkg.WriteDefaultConfig()
			if err != nil {
				return err
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			return printVersion()
		}
		if showConfig {
			return showCurrentConfig()
		}
		if setConfig {
			return setNewConfig()
		}
		return cmd.Help()
	},
}

func printVersion() error {
	fmt.Println("Version: ", pkg.Version)
	return nil
}

func showCurrentConfig() error {
	cfg, err := pkg.ReadConfigFile()
	if err != nil {
		return err
	}
	err = cfg.IsValidConfig()
	if err != nil {
		return err
	}
	fmt.Printf("MP3 Quality: %d\n"+
		"MP4 Quality: %d\n"+
		"Directory: %s",
		cfg.Mp3Quality, cfg.Mp4Quality, cfg.GetDownloadFolder())

	return nil
}

func setNewConfig() error {
	cfg, err := promptSetNewConfig()
	if err != nil {
		return err
	}

	err = pkg.WriteConfigFile(cfg)
	if err != nil {
		return err
	}
	fmt.Println("Override new configuration with details:")

	fmt.Printf("MP3 Quality: %d\n"+
		"MP4 Quality: %d\n"+
		"Directory: %s",
		cfg.Mp3Quality, cfg.Mp4Quality, cfg.GetDownloadFolder())

	return nil
}

func promptSetNewConfig() (*pkg.Config, error) {
	cfg := &pkg.Config{}
	mp3Quality, err := common.PromptInteger("MP3 Quality")
	if err != nil {
		return nil, err
	}

	if err := pkg.IsValidMP3Quality(mp3Quality); err != nil {
		return nil, err
	}

	cfg.Mp3Quality = mp3Quality

	mp4Quality, err := common.PromptInteger("MP4 Quality")
	if err != nil {
		return nil, err
	}

	if err := pkg.IsValidMP4Quality(mp4Quality); err != nil {
		return nil, err
	}

	cfg.Mp4Quality = mp4Quality

	directory, err := common.PromptString("Directory")
	if err != nil {
		return nil, err
	}
	cfg.Directory = directory

	return cfg, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
