package sql

import (
	"bufio"
	"bytes"
	db_sql "database/sql"
	"fmt"
	"text/template"
	"io/fs"
)

func LoadSchema(db *db_sql.DB, schema_fs fs.FS, table_name string) (string, error) {

	driver := Driver(db)

	fname := fmt.Sprintf("%s.%s.schema", table_name, driver)

	data, err := fs.ReadFile(schema_fs, fname)

	if err != nil {
		return "", fmt.Errorf("Failed to read %s, %w", fname, err)
	}

	t, err := template.New(table_name).Parse(string(data))

	if err != nil {
		return "", fmt.Errorf("Failed to parse %s template, %w", fname, err)
	}

	vars := struct {
		Name string
	}{
		Name: table_name,
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = t.Execute(wr, vars)

	if err != nil {
		return "", fmt.Errorf("Failed to process %s template, %w", fname, err)
	}

	wr.Flush()

	return buf.String(), nil
}
