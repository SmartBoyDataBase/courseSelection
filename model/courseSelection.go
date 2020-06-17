package model

import (
	"courseSelection/infrastructure"
)

type CourseSelection struct {
	StudentId     uint64 `json:"student_id"`
	TeachCourseId uint64 `json:"teach_course_id"`
	RegularGrade  *uint8 `json:"regular_grade,omitempty"`
	ExamGrade     *uint8 `json:"exam_grade,omitempty"`
	FinalGrade    *uint8 `json:"final_grade,omitempty"`
}

func Create(CourseSelection CourseSelection) (CourseSelection, error) {
	_, err := infrastructure.DB.Exec(`
	INSERT INTO CourseSelection(student_id, teachcourse_id, regular_grade, exam_grade, final_grade)
	VALUES ($1, $2, $3, $4, $5);`,
		CourseSelection.StudentId, CourseSelection.TeachCourseId,
		CourseSelection.RegularGrade, CourseSelection.ExamGrade,
		CourseSelection.FinalGrade)
	return CourseSelection, err
}

func Put(selection CourseSelection) (CourseSelection, error) {
	_, err := infrastructure.DB.Exec(`
	UPDATE courseselection
	SET regular_grade=$3,
		exam_grade=$4,
		final_grade=$5
	WHERE teachcourse_id=$1 AND student_id=$2;
	`, selection.TeachCourseId, selection.StudentId, selection.RegularGrade, selection.ExamGrade, selection.FinalGrade)
	return selection, err
}

func All() ([]CourseSelection, error) {
	rows, err := infrastructure.DB.Query(`
	SELECT student_id, teachcourse_id, regular_grade, exam_grade, final_grade
	FROM CourseSelection;
	`)
	if err != nil {
		return nil, err
	}
	var CourseSelections []CourseSelection
	for rows.Next() {
		var CourseSelection CourseSelection
		err := rows.Scan(&CourseSelection.StudentId,
			&CourseSelection.TeachCourseId,
			&CourseSelection.RegularGrade,
			&CourseSelection.ExamGrade,
			&CourseSelection.FinalGrade)
		if err != nil {
			return CourseSelections, err
		}
		CourseSelections = append(CourseSelections, CourseSelection)
	}
	return CourseSelections, nil
}

func FetchByTeachCourseId(id uint64) ([]CourseSelection, error) {
	rows, err := infrastructure.DB.Query(`
	SELECT student_id, teachcourse_id, regular_grade, exam_grade, final_grade
	FROM CourseSelection
	WHERE teachcourse_id=$1;
	`, id)
	if err != nil {
		return nil, err
	}
	var CourseSelections []CourseSelection
	for rows.Next() {
		var CourseSelection CourseSelection
		err := rows.Scan(&CourseSelection.StudentId,
			&CourseSelection.TeachCourseId,
			&CourseSelection.RegularGrade,
			&CourseSelection.ExamGrade,
			&CourseSelection.FinalGrade)
		if err != nil {
			return CourseSelections, err
		}
		CourseSelections = append(CourseSelections, CourseSelection)
	}
	return CourseSelections, nil
}
