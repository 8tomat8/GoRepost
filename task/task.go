package task

// Task - main struct from incoming message
type Task struct {
	Destinations []*Destination `json:"destinations"`
	Message      string         `json:"message"`
	Attachments  []*Attachment  `json:"attachments"`
}

// Tasks - List of Task. Shortcut
type Tasks []Task

// Attachment - struct for incoming attachments
type Attachment struct {
	Type string `json:"type"`
	Link string `json:"link"`
}

// Destination - struct to specify socials and groups in them
type Destination struct {
	Social   string   `json:"social"`
	GroupIDs []string `json:"group_ids"`
}
