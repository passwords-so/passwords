package storage

import _ "embed"

type schemaMigration struct {
	version int
	name    string
	sql     string
}

//go:embed migrations/001_initial.sql
var initialSchemaSQL string

var schemaMigrations = []schemaMigration{
	{
		version: 1,
		name:    "initial schema",
		sql:     initialSchemaSQL,
	},
}
