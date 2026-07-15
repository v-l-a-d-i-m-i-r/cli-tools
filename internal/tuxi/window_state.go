package tuxi

import "encoding/json"

type rawWindowState struct {
	ActiveWindowIndex int    `json:"active_window_index"`
	WindowIndex       int    `json:"window_index"`
	WindowName        string `json:"window_name"`
	WindowRawFlags    string `json:"window_raw_flags"`
}

func parseWindowState(str string) (rawWindowState, error) {
	var state rawWindowState
	err := json.Unmarshal([]byte(str), &state)
	return state, err
}
