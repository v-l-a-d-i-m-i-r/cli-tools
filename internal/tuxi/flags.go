package tuxi

type windowFlag string

const (
	activeWindowFlag   windowFlag = "*"
	lastWindowFlag     windowFlag = "-"
	activityWindowFlag windowFlag = "#"
	bellWindowFlag     windowFlag = "!"
	silenceWindowFlag  windowFlag = "~"
	markedWindowFlag   windowFlag = "M"
	zoomedWindowFlag   windowFlag = "Z"
)

type windowFlags struct {
	Active   bool
	Last     bool
	Activity bool
	Bell     bool
	Silence  bool
	Marked   bool
	Zoomed   bool
}

func parseWindowFlags(raw string) windowFlags {
	var f windowFlags

	for _, r := range raw {
		switch windowFlag(r) {
		case activeWindowFlag:
			f.Active = true
		case lastWindowFlag:
			f.Last = true
		case activityWindowFlag:
			f.Activity = true
		case bellWindowFlag:
			f.Bell = true
		case silenceWindowFlag:
			f.Silence = true
		case markedWindowFlag:
			f.Marked = true
		case zoomedWindowFlag:
			f.Zoomed = true
		}
	}

	return f
}
