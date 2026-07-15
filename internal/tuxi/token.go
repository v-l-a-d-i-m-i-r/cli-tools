package tuxi

import "strings"

type tabTokenOpts struct {
	text   string
	bg     string
	fg     string
	weight string
}

func tabToken(t tabTokenOpts) string {
	var parts []string

	if t.bg != "" {
		parts = append(parts, "bg="+t.bg)
	}

	if t.fg != "" {
		parts = append(parts, "fg="+t.fg)
	}

	if t.weight != "" {
		parts = append(parts, t.weight)
	}

	partsStr := strings.Join(parts, ", ")

	if len(partsStr) == 0 {
		return ""
	}

	return "#[" + partsStr + "]" + t.text + "#[default]"
}
