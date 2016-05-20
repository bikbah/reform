package models

// generated with github.com/AlekSi/reform

import (
	"fmt"
	"strings"

	"github.com/AlekSi/reform"
	"github.com/AlekSi/reform/parse"
)

type personTable struct {
	s parse.StructInfo
	z []interface{}
}

// Name returns a view or table name in SQL database (people).
func (v *personTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *personTable) Columns() []string {
	return []string{"id", "name", "email", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *personTable) NewStruct() reform.Struct {
	return new(Person)
}

// NewRecord makes a new record for that table.
func (v *personTable) NewRecord() reform.Record {
	return new(Person)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *personTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// PersonTable represents people view or table in SQL database.
var PersonTable = &personTable{
	s: parse.StructInfo{Type: "Person", SQLName: "people", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Name", Type: "string", Column: "name"}, {Name: "Email", Type: "*string", Column: "email"}, {Name: "CreatedAt", Type: "time.Time", Column: "created_at"}, {Name: "UpdatedAt", Type: "*time.Time", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(Person).Values(),
}

// String returns a string representation of this struct or record.
func (s Person) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "Email: " + reform.Inspect(s.Email, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Person) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.Email,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Person) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.Email,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Person) View() reform.View {
	return PersonTable
}

// Table returns Table object for that record.
func (s *Person) Table() reform.Table {
	return PersonTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Person) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Person) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Person) HasPK() bool {
	return s.ID != PersonTable.z[PersonTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Person) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = PersonTable
	_ reform.Struct = new(Person)
	_ reform.Table  = PersonTable
	_ reform.Record = new(Person)
	_ fmt.Stringer  = new(Person)
)

type projectTable struct {
	s parse.StructInfo
	z []interface{}
}

// Name returns a view or table name in SQL database (projects).
func (v *projectTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *projectTable) Columns() []string {
	return []string{"name", "id", "start", "end"}
}

// NewStruct makes a new struct for that view or table.
func (v *projectTable) NewStruct() reform.Struct {
	return new(Project)
}

// NewRecord makes a new record for that table.
func (v *projectTable) NewRecord() reform.Record {
	return new(Project)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *projectTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ProjectTable represents projects view or table in SQL database.
var ProjectTable = &projectTable{
	s: parse.StructInfo{Type: "Project", SQLName: "projects", Fields: []parse.FieldInfo{{Name: "Name", Type: "string", Column: "name"}, {Name: "ID", Type: "string", Column: "id"}, {Name: "Start", Type: "time.Time", Column: "start"}, {Name: "End", Type: "*time.Time", Column: "end"}}, PKFieldIndex: 1},
	z: new(Project).Values(),
}

// String returns a string representation of this struct or record.
func (s Project) String() string {
	res := make([]string, 4)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "ID: " + reform.Inspect(s.ID, true)
	res[2] = "Start: " + reform.Inspect(s.Start, true)
	res[3] = "End: " + reform.Inspect(s.End, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Project) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.ID,
		s.Start,
		s.End,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Project) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.ID,
		&s.Start,
		&s.End,
	}
}

// View returns View object for that struct.
func (s *Project) View() reform.View {
	return ProjectTable
}

// Table returns Table object for that record.
func (s *Project) Table() reform.Table {
	return ProjectTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Project) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Project) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Project) HasPK() bool {
	return s.ID != ProjectTable.z[ProjectTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Project) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = string(i64)
	} else {
		s.ID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = ProjectTable
	_ reform.Struct = new(Project)
	_ reform.Table  = ProjectTable
	_ reform.Record = new(Project)
	_ fmt.Stringer  = new(Project)
)

type personProjectView struct {
	s parse.StructInfo
	z []interface{}
}

// Name returns a view or table name in SQL database (person_project).
func (v *personProjectView) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *personProjectView) Columns() []string {
	return []string{"person_id", "project_id"}
}

// NewStruct makes a new struct for that view or table.
func (v *personProjectView) NewStruct() reform.Struct {
	return new(PersonProject)
}

// PersonProjectView represents person_project view or table in SQL database.
var PersonProjectView = &personProjectView{
	s: parse.StructInfo{Type: "PersonProject", SQLName: "person_project", Fields: []parse.FieldInfo{{Name: "PersonID", Type: "int32", Column: "person_id"}, {Name: "ProjectID", Type: "string", Column: "project_id"}}, PKFieldIndex: -1},
	z: new(PersonProject).Values(),
}

// String returns a string representation of this struct or record.
func (s PersonProject) String() string {
	res := make([]string, 2)
	res[0] = "PersonID: " + reform.Inspect(s.PersonID, true)
	res[1] = "ProjectID: " + reform.Inspect(s.ProjectID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *PersonProject) Values() []interface{} {
	return []interface{}{
		s.PersonID,
		s.ProjectID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *PersonProject) Pointers() []interface{} {
	return []interface{}{
		&s.PersonID,
		&s.ProjectID,
	}
}

// View returns View object for that struct.
func (s *PersonProject) View() reform.View {
	return PersonProjectView
}

// check interfaces
var (
	_ reform.View   = PersonProjectView
	_ reform.Struct = new(PersonProject)
	_ fmt.Stringer  = new(PersonProject)
)

func init() {
	parse.AssertUpToDate(&PersonTable.s, new(Person))
	parse.AssertUpToDate(&ProjectTable.s, new(Project))
	parse.AssertUpToDate(&PersonProjectView.s, new(PersonProject))
}
