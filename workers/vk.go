package workers

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/8tomat8/GoRepost/task"
	"github.com/golang/glog"
	"github.com/tidwall/gjson"
)

// Handler for Task
func vk(t *task.Task) {
	groups, _ := t.Destinations["vk"]

	params := make(map[string]string)
	params["attachments"] = ""
	params["message"] = t.Message

	for _, g := range *groups {
		if len(g.AccessKey) != 85 { // 85 - Size of Vk standard Access key
			glog.Error(g.Id + "s access_key not valid! Skipping group.")
			continue
		}

		if g.FromGroup {
			params["from_group"] = "1"
		} else {
			params["from_group"] = "0"
		}
		params["owner_id"] = "-" + g.Id
		for _, a := range t.Attachments {
			switch a.Type {
			case "photo":
				photo, err := photoLoader(a, g)
				if err != nil {
					glog.Error(err)
				} else {
					addToAttachments(params, photo)
				}
			case "link":
				if len(a.Link) != 0 {
					addToAttachments(params, a.Link)
				}
			default:
				glog.Error(a.Type + " not supported yet =(")
			}
		}
		strResp, err := callAPI("wall.post", params, &g.AccessKey)
		if err != nil {
			panic(err)
		}
		if strResp != "" {
			glog.Info("VK API response for " + g.Id + ": " + strResp)
		}
	}
}

func callAPI(method string, params map[string]string, token *string) (string, error) {
	u, err := url.Parse("https://api.vk.com/method/" + method)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("access_token", *token)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	result := string(data)
	return result, nil
}

func photoLoader(a *task.Attachment, g *task.Group) (string, error) {
	params := make(map[string]string)
	params["group_id"] = g.Id

	var b bytes.Buffer

	jsonResp, err := callAPI("photos.getWallUploadServer", params, &g.AccessKey)
	if err != nil {
		return "", errors.New("Failed to load" + a.Type + " " + a.Link)
	}

	uploadURL := gjson.Get(jsonResp, "response.upload_url")

	resp, err := http.Get(a.Link)
	defer resp.Body.Close()

	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("photo", "image.jpeg")
	if _, err = io.Copy(fw, resp.Body); err != nil {
		return "", errors.New("Can't read photo from " + a.Link)
	}

	w.Close()

	req, err := http.NewRequest("POST", uploadURL.String(), &b)
	if err != nil {
		return "", errors.New("Can't create Request for " + a.Link)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(io.LimitReader(res.Body, 2<<19))
	if err != nil {
		return "", errors.New("Can't get response from VK on " + a.Link)
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("Can't upload photo " + a.Link)
	}

	strResp := string(body)
	params["group_id"] = strings.Replace(params["group_id"], "-", "", -1)
	params["photo"] = gjson.Get(strResp, "photo").String()
	params["server"] = gjson.Get(strResp, "server").String()
	params["hash"] = gjson.Get(strResp, "hash").String()

	jsonResp, err = callAPI("photos.saveWallPhoto", params, &g.AccessKey)
	if err != nil {
		return "", errors.New("Can't get response from VK on saving " + a.Link)
	}
	return gjson.Get(jsonResp, "response.#.id").Array()[0].String(), nil
}

func addToAttachments(params map[string]string, newElement string) {
	if len([]rune(params["attachments"])) != 0 {
		params["attachments"] = params["attachments"] + ","
	}
	params["attachments"] = params["attachments"] + newElement
}
