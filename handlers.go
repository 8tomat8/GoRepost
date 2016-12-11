package GoRepost

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/8tomat8/GoRepost/counter"
	"github.com/8tomat8/GoRepost/task"
	"github.com/8tomat8/GoRepost/workers"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func TaskCreate(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 2<<19))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &task); err != nil {
		glog.Error(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			glog.Error(err)
		}
		return
	}
	glog.Info(body)
	go workers.Handler(&task)
}

func TaskStatus(w http.ResponseWriter, r *http.Request) {
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }
}

func TasksList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["taskId"]
	fmt.Fprintln(w, "Todo show:", taskID)
}

func Greating(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "                  .....")
	fmt.Fprintln(w, "                 C C  /")
	fmt.Fprintln(w, "                /<   /     IT'S ALIVE!!!!!!11111")
	fmt.Fprintln(w, " ___ __________/_#__=o     #####################")
	fmt.Fprintln(w, "/(- /(\\_\\________   \\      The number of jobs: "+strconv.FormatUint(counter.GetCounter().GetSize(), 10))
	fmt.Fprintln(w, "\\ ) \\ )_      \\o     \\     #####################")
	fmt.Fprintln(w, "/|\\ /|\\       |'     |       ")
	fmt.Fprintln(w, "              |     _|       ")
	fmt.Fprintln(w, "              /o   __\\       ")
	fmt.Fprintln(w, "             / '     |     ")
	fmt.Fprintln(w, "            / /      |       ")
	fmt.Fprintln(w, "           /_/\\______|       ")
	fmt.Fprintln(w, "          (   _(    <        ")
	fmt.Fprintln(w, "           \\    \\    \\       ")
	fmt.Fprintln(w, "            \\    \\    |    ")
	fmt.Fprintln(w, "             \\____\\____\\   ")
	fmt.Fprintln(w, "             ____\\_\\__\\_\\    ")
}
