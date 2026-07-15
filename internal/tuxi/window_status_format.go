package tuxi

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

const spaceChar = "█"

func newWindowStatusFormatCmd() *cobra.Command {
	var config string
	var windowState string

	cmd := &cobra.Command{
		Use:  "window-status-format",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			log := loggerFromContext(cmd.Context())
			log.Debug("window-status-format request", "config", config, "window_state", windowState)

			tab, err := renderWindowStatus(config, windowState)
			if err != nil {
				log.Error("window-status-format failed", "error", err)
				return err
			}

			log.Debug("window-status-format result", "output", tab)

			if _, err := fmt.Fprintln(cmd.OutOrStdout(), tab); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&config, "config", "", "tuxi config JSON")
	cmd.Flags().StringVar(&windowState, "window-state", "", "tuxi window state JSON")
	_ = cmd.MarkFlagRequired("config")
	_ = cmd.MarkFlagRequired("window-state")

	return cmd
}

func renderWindowStatus(configStr, stateStr string) (string, error) {
	config, err := parseConfig(configStr)
	if err != nil {
		return "", err
	}

	state, err := parseWindowState(stateStr)
	if err != nil {
		return "", err
	}

	flags := parseWindowFlags(state.WindowRawFlags)

	var tab string
	var tabText string

	indexText, indexToken := renderIndex(config, state, flags)
	nameText, nameToken := renderName(config, state, flags)
	flagsText, flagsToken := renderFlags(config, state, flags)
	separatorText, separatorToken := renderSeparator(config, state, flags)

	tabText += indexText
	tabText += nameText
	tabText += flagsText
	tabText += separatorText

	totalChars := utf8.RuneCountInString(tabText)
	spacerLen := config.MinWindowStatusWidth - totalChars
	_, spacerToken := renderSpacer(config, state, flags, spacerLen)

	tab += indexToken
	tab += spacerToken
	tab += nameToken
	tab += flagsToken
	tab += separatorToken

	return tab, nil
}

func renderIndex(
	config rawConfig,
	state rawWindowState,
	flags windowFlags,
) (text string, token string) {
	bg := config.InactiveWindowTitleBg
	fg := config.InactiveWindowTitleFg

	if flags.Active {
		bg = config.ActiveWindowTitleBg
		fg = config.ActiveWindowTitleFg
	}

	text = " " + strconv.Itoa(state.WindowIndex) + ":"
	token = tabToken(tabTokenOpts{
		text: text,
		bg:   bg,
		fg:   fg,
	})

	return text, token
}

func renderName(
	config rawConfig,
	state rawWindowState,
	flags windowFlags,
) (text string, token string) {
	bg := config.InactiveWindowTitleBg
	fg := config.InactiveWindowTitleFg

	if flags.Active {
		bg = config.ActiveWindowTitleBg
		fg = config.ActiveWindowTitleFg
	}

	text = " " + state.WindowName
	token = tabToken(tabTokenOpts{
		text: text,
		bg:   bg,
		fg:   fg,
	})

	return text, token
}

func renderFlags(
	config rawConfig,
	state rawWindowState,
	flags windowFlags,
) (text string, token string) {
	bg := config.InactiveWindowTitleBg
	fg := config.InactiveWindowTitleFg

	if flags.Active {
		bg = config.ActiveWindowTitleBg
		fg = config.ActiveWindowTitleFg
	}

	if len(state.WindowRawFlags) > 0 {
		text += " "
		token += tabToken(tabTokenOpts{
			text:   " ",
			bg:     bg,
			fg:     fg,
			weight: "bold",
		})

		if flags.Active {
			text += string(activeWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(activeWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}

		if flags.Last {
			text += string(lastWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(lastWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}

		if flags.Activity {
			text += string(activityWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(activityWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}

		if flags.Bell {
			text += string(bellWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(bellWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}

		if flags.Silence {
			text += string(silenceWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(silenceWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}

		if flags.Marked {
			text += string(markedWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(markedWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}

		if flags.Zoomed {
			text += string(zoomedWindowFlag)
			token += tabToken(tabTokenOpts{
				text:   string(zoomedWindowFlag),
				bg:     bg,
				fg:     fg,
				weight: "bold",
			})
		}
	}

	return text, token
}

func renderSeparator(
	config rawConfig,
	state rawWindowState,
	flags windowFlags,
) (text string, token string) {
	if flags.Active {
		text = spaceChar
		token = tabToken(tabTokenOpts{
			text: text,
			bg:   config.ActiveWindowTitleFg,
			fg:   config.ActiveWindowTitleBg,
		})

		return text, token
	}

	needsSeparator := state.WindowIndex != state.ActiveWindowIndex-1
	if needsSeparator {
		text = config.SeparatorChar
		token = tabToken(tabTokenOpts{
			text: text,
			bg:   config.InactiveWindowTitleBg,
			fg:   config.ActiveWindowTitleBg,
		})

		return text, token
	}

	if !needsSeparator {
		text = " "
		token = tabToken(tabTokenOpts{
			text: text,
			bg:   config.InactiveWindowTitleBg,
			fg:   config.InactiveWindowTitleFg,
		})
	}

	return text, token
}

func renderSpacer(
	config rawConfig,
	_ rawWindowState,
	flags windowFlags,
	l int,
) (text string, token string) {
	if l <= 0 {
		return "", ""
	}

	bg := config.InactiveWindowTitleFg
	fg := config.InactiveWindowTitleBg

	if flags.Active {
		bg = config.ActiveWindowTitleFg
		fg = config.ActiveWindowTitleBg
	}

	text = strings.Repeat(spaceChar, l)
	token = tabToken(tabTokenOpts{
		text: text,
		bg:   bg,
		fg:   fg,
	})

	return text, token
}
