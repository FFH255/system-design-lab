package mongo

type UID string

type Member struct {
	UID      UID    `bson:"uid" json:"uid"`
	FullName string `bson:"full_name" json:"full_name"`
}

type Group struct {
	UID     UID      `bson:"uid" json:"uid"`
	Name    string   `bson:"name" json:"name"`
	Members []Member `bson:"members" json:"members"`
}
