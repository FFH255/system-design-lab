package neo4j

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
