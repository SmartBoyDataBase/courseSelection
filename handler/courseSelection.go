package handler

import (
	"courseSelection/infrastructure"
	"courseSelection/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	teachCourseIdStr := r.URL.Query().Get("teach_course_id")
	teachCourseId, err := strconv.ParseUint(teachCourseIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	rows, err := infrastructure.DB.Query(`
	SELECT student_id, regular_grade, exam_grade, final_grade
	FROM courseselection
	WHERE teachcourse_id=$1;
	`, teachCourseId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var result []model.CourseSelection
	for rows.Next() {
		current := model.CourseSelection{
			TeachCourseId: teachCourseId,
		}
		err = rows.Scan(&current.StudentId, &current.RegularGrade, &current.ExamGrade, &current.FinalGrade)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		result = append(result, current)
	}
	var body []byte
	if len(result) != 0 {
		body, _ = json.Marshal(result)
	} else {
		body = []byte("[]")
	}
	_, _ = w.Write(body)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var all []model.CourseSelection
	var err error
	teachcourseIdStr := r.URL.Query().Get("teachcourse_id")
	if teachcourseIdStr != "" {
		teachcourseId, _ := strconv.ParseUint(teachcourseIdStr, 10, 64)
		all, err = model.FetchByTeachCourseId(teachcourseId)
	} else {
		all, err = model.All()
	}
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
