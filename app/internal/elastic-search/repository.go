package elastic_search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

type Repository struct {
	client    *elasticsearch.Client
	indexName string
}

func NewRepository(
	client *elasticsearch.Client,
) *Repository {
	return &Repository{
		client:    client,
		indexName: "courses",
	}
}

func (r *Repository) Create(ctx context.Context, course Course) error {
	// Convert the Course struct to JSON
	jsonBody, err := json.Marshal(course)

	if err != nil {
		return fmt.Errorf("error marshaling the course: %s", err)
	}

	// Create a request to index the document
	req := esapi.IndexRequest{
		Index:      r.indexName, // Name of the Elasticsearch index
		DocumentID: course.ID,
		Body:       strings.NewReader(string(jsonBody)),
		Refresh:    "true",
	}

	// Send the request to Elasticsearch
	res, err := req.Do(ctx, r.client)

	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.IsError() {
		return fmt.Errorf("[%s] Error indexing document ID=%s", res.Status(), course.ID)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, course Course) error {
	jsonBody, err := json.Marshal(course)

	if err != nil {
		return fmt.Errorf("error marshaling the course: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      r.indexName,
		DocumentID: course.ID,
		Body:       strings.NewReader(string(jsonBody)),
	}

	res, err := req.Do(ctx, r.client)

	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.IsError() {
		return fmt.Errorf("[%s] Error updating document ID=%s", res.Status(), course.ID)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, courseID string) error {
	req := esapi.DeleteRequest{
		Index:      r.indexName,
		DocumentID: courseID,
	}

	res, err := req.Do(ctx, r.client)

	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.IsError() {
		return fmt.Errorf("[%s] Error deleting document ID=%s", res.Status(), courseID)
	}

	return nil
}

func (r *Repository) Search(ctx context.Context, search string) ([]Course, error) {
	// Create the search query, using the dynamic 'search' string
	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"name": "%s"
			}
		}
	}`, search)

	// Create the SearchRequest
	req := esapi.SearchRequest{
		Index: []string{r.indexName}, // The Elasticsearch index
		Body:  strings.NewReader(query),
	}

	// Send the search request
	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.IsError() {
		return nil, fmt.Errorf("[%s] error in search request", res.Status())
	}

	// Parse the search results
	var results map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	// Extract the search hits
	hits := results["hits"].(map[string]interface{})["hits"].([]interface{})

	// Create a slice to hold the resulting courses
	courses := make([]Course, 0)

	// Iterate through the hits and unmarshal the course data
	for _, hit := range hits {
		hitSource := hit.(map[string]interface{})["_source"]
		course := Course{
			ID:      hitSource.(map[string]interface{})["id"].(string),
			Name:    hitSource.(map[string]interface{})["name"].(string),
			Content: hitSource.(map[string]interface{})["content"].(string),
		}
		courses = append(courses, course)
	}

	return courses, nil
}
