package documentdb

import "fmt"

// Database represents a Cosmos DB Database
type Database struct {
	ID              string `json:"id"`
	ResourceID      string `json:"_rid"`
	Timestamp       int    `json:"_ts"`
	SelfLink        string `json:"_self"`
	ETag            string `json:"_etag"`
	CollectionsLink string `json:"_colls"`
	UsersLink       string `json:"_users"`
}

func (db *Database) String() string {
	return fmt.Sprintf("{ID: %v, ResourceID: %v, Timestamp: %v, SelfLink: %v, ETag: %v, CollectionsLink: %v, UsersLink: %v}", db.ID, db.ResourceID, db.Timestamp, db.SelfLink, db.ETag, db.CollectionsLink, db.UsersLink)
}
