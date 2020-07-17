package database

import "github.com/hashicorp/go-memdb"

func Schema() *memdb.DBSchema {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"weight_histories": &memdb.TableSchema{
				Name: "weight_histories",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	return schema
}
