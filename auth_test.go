package documentdb

import "testing"

func TestMasterKeyTokenProviderGenerateToken(t *testing.T) {
	provider := &MasterKeyTokenProvider{
		Verb:         "GET",
		ResourceType: "dbs",
		ResourceLink: "dbs/ToDoList",
		Date:         "Thu, 27 Apr 2017 00:51:12 GMT",
		Key:          "dsZQi3KtZmCv1ljt3VNWNm7sQUF1y5rJfC6kv5JiwvW0EndXdDku/dkKBp8/ufDToSxLzR4y+O/0H/t4bQtVNw==",
		KeyType:      "master",
		TokenVersion: "1.0",
	}

	expected := "type%3Dmaster%26ver%3D1.0%26sig%3Dc09PEVJrgp2uQRkr934kFbTqhByc7TVr3OHyqlu%2Bc%2Bc%3D"

	signature, err := provider.GenerateToken()

	if err != nil {
		t.Fatal(err)
	}

	if signature != expected {
		t.Fatalf("Signature incorrect:\ngot: %v\nwant: %v\n", signature, expected)
	}
}
