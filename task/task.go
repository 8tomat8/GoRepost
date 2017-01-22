package task

// Task - main struct from incoming message
type Task struct {
	Id           string             `json:"id"`
	CallBackUrl  string             `json:"call_back_url"`
	Destinations map[string]*Groups `json:"destinations"`
	Message      string             `json:"message"`
	Attachments  []*Attachment      `json:"attachments"`
}

// Tasks - List of Task. Shortcut
type Tasks []Task

// Attachment - struct for incoming attachments
type Attachment struct {
	Type string `json:"type"`
	Link string `json:"link"`
}

type Group struct {
	Id        string    `json:"id"`
	AccessKey string    `json:"access_key"`
	FromGroup bool      `json:"from_group"`
	Status    string    `json:"status"`
}

type Groups []*Group

func NewTask() *Task {
	return &Task{
		Destinations:make(map[string]*Groups),
		Attachments: make([]*Attachment, 10),
	}
}
