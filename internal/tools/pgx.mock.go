package tools

import (
	"context"
	"errors"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// MockPGXPool represents a mock implementation of *pgxpool.Pool
type MockPGXPool struct {
	QueryRowFunc func(ctx context.Context, sql string, args ...any) pgx.Row
}

func (m *MockPGXPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(ctx, sql, args...)
	}
	return nil
}

// MockPGXRows represents a mock implementation of pgx.Rows
type MockPGXRows struct {
	data     [][]interface{}
	columns  []string
	rowIndex int
	closed   bool
	errValue error
}

func NewMockPGXRows(columns []string, data [][]interface{}) *MockPGXRows {
	return &MockPGXRows{
		columns:  columns,
		data:     data,
		rowIndex: -1,
	}
}

func (m *MockPGXRows) Close() {
	m.closed = true
}
func (m *MockPGXRows) Err() error {
	return m.errValue
}
func (m *MockPGXRows) CommandTag() pgconn.CommandTag {
	// Implement if needed
	return pgconn.CommandTag{}
}
func (m *MockPGXRows) FieldDescriptions() []pgconn.FieldDescription {
	// Implement if needed
	return nil
}
func (m *MockPGXRows) Next() bool {
	m.rowIndex++
	return m.rowIndex < len(m.data)
}
func (m *MockPGXRows) Scan(dest ...interface{}) error {
	if m.rowIndex >= len(m.data) {
		return pgx.ErrNoRows
	}
	rowData := m.data[m.rowIndex]

	if len(rowData) != len(dest) {
		return errors.New("number of columns doesn't match number of destinations")
	}

	for i, val := range rowData {
		destVal := reflect.ValueOf(dest[i])
		if destVal.Kind() != reflect.Ptr {
			return errors.New("destination is not a pointer")
		}

		if reflect.TypeOf(val).AssignableTo(destVal.Elem().Type()) {
			destVal.Elem().Set(reflect.ValueOf(val))
		} else {
			return errors.New("value type does not match destination type: " + reflect.TypeOf(val).String() + " vs " + destVal.Elem().Type().String())
		}
	}

	return nil
}
func (m *MockPGXRows) Values() ([]interface{}, error) {
	// Implement if needed
	return nil, nil
}
func (m *MockPGXRows) RawValues() [][]byte {
	// Implement if needed
	return nil
}

// MockPGXRow represents a mock implementation of pgx.Row
type MockPGXRow struct {
	columns []string
	values  []interface{}
	index   int
}

func NewMockPGXRow(columns []string, values []interface{}) *MockPGXRow {
	return &MockPGXRow{
		columns: columns,
		values:  values,
		index:   -1,
	}
}

func (m *MockPGXRow) Scan(dest ...interface{}) error {
	if m.index >= len(m.values)-1 {
		return pgx.ErrNoRows
	}
	m.index++

	if len(m.values) != len(dest) {
		return errors.New("number of columns doesn't match number of destinations")
	}

	for i, val := range m.values {
		destVal := reflect.ValueOf(dest[i])
		if destVal.Kind() != reflect.Ptr {
			return errors.New("destination is not a pointer")
		}

		if reflect.TypeOf(val).AssignableTo(destVal.Elem().Type()) {
			destVal.Elem().Set(reflect.ValueOf(val))
		} else {
			return errors.New("value type does not match destination type: " + reflect.TypeOf(val).String() + " vs " + destVal.Elem().Type().String())
		}
	}

	return nil
}
