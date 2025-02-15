package progress

import (
	"fmt"
	"time"

	"github.com/VILJkid/poopy-turtle/constants"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func Run() {
	theme := progressbar.Theme{
		Saucer:        constants.EmojiPoop,
		SaucerHead:    constants.EmojiTurtle,
		SaucerPadding: constants.EmojiWaterWave,
		BarStart:      constants.EmojiLion,
		BarEnd:        constants.EmojiDesertIsland,
	}

	opt := []progressbar.Option{
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetDescription(constants.Description),
		progressbar.OptionSetTheme(theme),
	}

	bar := progressbar.NewOptions(constants.MaxSize, opt...)
	for range bar.GetMax() {
		bar.Add(1)
		time.Sleep(time.Duration(bar.GetMax()) * time.Millisecond)
	}

	fmt.Printf("\n"+constants.Message+"\n", constants.EmojiTurtle)
}
