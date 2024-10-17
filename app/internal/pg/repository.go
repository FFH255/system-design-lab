package pg

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	client *pgx.Conn
}

func NewRepository(
	client *pgx.Conn,
) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Create(ctx context.Context, dto CreateAttendanceDTO) (*Attendance, error) {
	query := `
        INSERT INTO student_attendance (student_id, lecture_id, attendance_date)
        VALUES ($1, $2, current_date)
        RETURNING id, student_id, lecture_id, attendance_date`

	attendance := &Attendance{}
	err := r.client.QueryRow(ctx, query, dto.StudentID, dto.LectureID).Scan(
		&attendance.ID,
		&attendance.StudentID,
		&attendance.LectureID,
		&attendance.Date,
	)

	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *Repository) Update(ctx context.Context, dto UpdateAttendanceDTO) (*Attendance, error) {
	query := `
        UPDATE student_attendance
        SET student_id = $1, lecture_id = $2
        WHERE id = $3
        RETURNING id, student_id, lecture_id, attendance_date`

	attendance := &Attendance{}
	err := r.client.QueryRow(ctx, query, dto.StudentID, dto.LectureID, dto.ID).Scan(
		&attendance.ID,
		&attendance.StudentID,
		&attendance.LectureID,
		&attendance.Date,
	)

	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *Repository) Delete(ctx context.Context, dto DeleteAttendanceDTO) error {
	query := `DELETE FROM student_attendance WHERE id = $1`

	_, err := r.client.Exec(ctx, query, dto.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Attendance, error) {
	query := `SELECT id, student_id, lecture_id, attendance_date FROM student_attendance`

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []Attendance
	for rows.Next() {
		var attendance Attendance
		if err := rows.Scan(
			&attendance.ID,
			&attendance.StudentID,
			&attendance.LectureID,
			&attendance.Date,
		); err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}
