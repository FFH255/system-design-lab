package redis

type UID string

type Student struct {
	UID        UID    `json:"uid"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type CreateStudentDTO struct {
	Student
}

type UpdateStudentDTO struct {
	Student
}
