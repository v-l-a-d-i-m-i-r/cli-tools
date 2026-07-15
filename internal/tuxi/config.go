package tuxi

import "encoding/json"

type rawConfig struct {
	ActiveWindowTitleBg   string `json:"active_window_title_bg"`
	ActiveWindowTitleFg   string `json:"active_window_title_fg"`
	InactiveWindowTitleBg string `json:"inactive_window_title_bg"`
	InactiveWindowTitleFg string `json:"inactive_window_title_fg"`
	MinWindowStatusWidth  int    `json:"min_window_status_width"`
	SeparatorChar         string `json:"window_status_separator_char"`
}

func parseConfig(str string) (rawConfig, error) {
	var config rawConfig
	err := json.Unmarshal([]byte(str), &config)
	return config, err
}
