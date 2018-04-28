package main

type Database struct {
	Id              string `json:"id"`
	ResourceId      string `json:"_rid"`
	Timestamp       int    `json:"_ts"`
	SelfLink        string `json:"_self"`
	ETag            string `json:"_etag"`
	CollectionsLink string `json:"_colls"`
	UsersLink       string `json:"_users"`
}
