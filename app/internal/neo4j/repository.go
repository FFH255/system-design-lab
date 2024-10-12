package neo4j

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repository struct {
	driver neo4j.DriverWithContext
}

func NewRepository(driver neo4j.DriverWithContext) *Repository {
	return &Repository{driver: driver}
}

func (r *Repository) CreateStudent(ctx context.Context, student Student) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func() {
		_ = session.Close(ctx)
	}()

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "CREATE (s:Student { id: $id, name: $name })"
		params := map[string]any{
			"id":   student.ID,
			"name": &student.Name,
		}
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})

	return err
}

func (r *Repository) CreateGroup(ctx context.Context, group Group) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func() {
		_ = session.Close(ctx)
	}()

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "CREATE (g:Group { id: $id, name: $name })"
		params := map[string]any{
			"id":   group.ID,
			"name": &group.Name,
		}
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})

	return err
}

func (r *Repository) CreateCourse(ctx context.Context, course Course) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func() {
		_ = session.Close(ctx)
	}()

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "CREATE (c:Course {id: $id, title: $title})"
		params := map[string]interface{}{
			"id":    course.ID,
			"title": course.Title,
		}
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})

	return err
}

func (r *Repository) AddStudentToGroup(ctx context.Context, studentID string, groupID string) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func() {
		_ = session.Close(ctx)
	}()

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (s:Student {id: $studentID}), (g:Group {id: $groupID})
			CREATE (s)-[:BELONGS_TO]->(g)
		`
		params := map[string]interface{}{
			"studentID": studentID,
			"groupID":   groupID,
		}
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})

	return err
}

func (r *Repository) EnrollStudentInCourse(ctx context.Context, studentID, courseID string) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func() {
		_ = session.Close(ctx)
	}()

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (s:Student {id: $studentID}), (c:Course {id: $courseID})
			CREATE (s)-[:ENROLLED_IN]->(c)
		`
		params := map[string]interface{}{
			"studentID": studentID,
			"courseID":  courseID,
		}
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})

	return err
}
