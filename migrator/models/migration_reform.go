// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type migrationTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *migrationTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("schema_migrations").
func (v *migrationTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *migrationTableType) Columns() []string {
	return []string{"version", "state", "run_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *migrationTableType) NewStruct() reform.Struct {
	return new(Migration)
}

// NewRecord makes a new record for that table.
func (v *migrationTableType) NewRecord() reform.Record {
	return new(Migration)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *migrationTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// MigrationTable represents schema_migrations view or table in SQL database.
var MigrationTable = &migrationTableType{
	s: parse.StructInfo{Type: "Migration", SQLSchema: "", SQLName: "schema_migrations", Fields: []parse.FieldInfo{{Name: "Version", Type: "string", Column: "version"}, {Name: "State", Type: "string", Column: "state"}, {Name: "RunAt", Type: "time.Time", Column: "run_at"}}, PKFieldIndex: 0},
	z: new(Migration).Values(),
}

// String returns a string representation of this struct or record.
func (s Migration) String() string {
	res := make([]string, 3)
	res[0] = "Version: " + reform.Inspect(s.Version, true)
	res[1] = "State: " + reform.Inspect(s.State, true)
	res[2] = "RunAt: " + reform.Inspect(s.RunAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Migration) Values() []interface{} {
	return []interface{}{
		s.Version,
		s.State,
		s.RunAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Migration) Pointers() []interface{} {
	return []interface{}{
		&s.Version,
		&s.State,
		&s.RunAt,
	}
}

// View returns View object for that struct.
func (s *Migration) View() reform.View {
	return MigrationTable
}

// Table returns Table object for that record.
func (s *Migration) Table() reform.Table {
	return MigrationTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Migration) PKValue() interface{} {
	return s.Version
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Migration) PKPointer() interface{} {
	return &s.Version
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Migration) HasPK() bool {
	return s.Version != MigrationTable.z[MigrationTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Migration) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Version = string(i64)
	} else {
		s.Version = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = MigrationTable
	_ reform.Struct = (*Migration)(nil)
	_ reform.Table  = MigrationTable
	_ reform.Record = (*Migration)(nil)
	_ fmt.Stringer  = (*Migration)(nil)
)

func init() {
	parse.AssertUpToDate(&MigrationTable.s, new(Migration))
}