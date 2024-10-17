package pg

import "time"

type Attendance struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	LectureID int       `json:"lecture_id"`
	Date      time.Time `json:"date"`
}

type CreateAttendanceDTO struct {
	StudentID int `json:"student_id"`
	LectureID int `json:"lecture_id"`
}

type UpdateAttendanceDTO struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	LectureID int `json:"lecture_id"`
}

type DeleteAttendanceDTO struct {
	ID int `json:"id"`
}
