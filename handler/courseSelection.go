package handler

import (
	"courseSelection/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var toCreate model.CourseSelection
	_ = json.Unmarshal(body, &toCreate)
	result, err := model.Create(toCreate)
	if err != nil {
		log.Println("Create courseSelection failed")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		log.Println("CourseSelection ", result.TeachCourseId, ',', result.StudentId, "created")
	}
	response, err := json.Marshal(result)
	_, _ = w.Write(response)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postHandler(w, r)
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	all, err := model.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var body []byte
	if len(all) != 0 {
		body, _ = json.Marshal(all)
	} else {
		body = []byte("[]")
	}
	_, _ = w.Write(body)
}
