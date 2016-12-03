package workers

import (
	"log"

	"github.com/8tomat8/GoRepost/task"
	"github.com/yanple/vk_api"
)

// Handler for Task
func vk(t *task.Task) {
	// time.Sleep(15000 * time.Millisecond)
	var groupIDs []string
	api := &vk_api.Api{}
	api.AccessToken = "070db029e63209872170a2e83a13700126347f86aac0cf0e3c01c4b2b482c2c8c8023ef81d3cacaacf484"

	for _, d := range t.Destinations {
		if d.Social == "vk" {
			groupIDs = d.GroupIDs
		}
	}
	params := make(map[string]string)
	params["owner_id"] = groupIDs[0]
	params["from_group"] = "1"
	params["message"] = t.Message

	for _, a := range t.Attachments {
		// TODO: Add mechanism of processing attachments
		if a.Type == "photo" {
			photoLoader(a)
		}
	}

	strResp, err := api.Request("wall.post", params)
	if err != nil {
		panic(err)
	}
	if strResp != "" {
		log.Println(strResp)
	}
}

func photoLoader(l *task.Attachment) {
	return
}
