package gonfman

import (
	"testing"

	"github.com/axkit/gonfig"
)

// type MockDB struct{}

// func (MockDB) Query(string, ...interface{}) (*sql.Rows, error) {
// 	return &MockRows{
// 		rows: []Param{
// 			{ID: "gopath", DataTypeID: "string", RawValue: "/home/go/src"},
// 			{ID: "goroot", DataTypeID: "string", RawValue: "/home/go"},
// 			{ID: "goos", DataTypeID: "string", RawValue: "amd64"}},
// 	}, nil

// }

// type MockRows struct {
// 	curpos int
// 	rows   []Param
// }

// func (MockRows) Close() error {
// 	return nil
// }
// func (m *MockRows) Next() bool {
// 	return m.curpos < len(m.rows)
// }

// func (m *MockRows) Scan(dest ...interface{}) error {
// 	*(dest[0]).(*string) = m.rows[m.curpos].ID
// 	dest[1] = m.rows[m.curpos].SectionID
// 	dest[2] = m.rows[m.curpos].PositionOrder
// 	dest[3] = m.rows[m.curpos].Name
// 	*dest[4].(*string) = m.rows[m.curpos].DataTypeID
// 	dest[5] = m.rows[m.curpos].ControlID
// 	*dest[6].(*string) = m.rows[m.curpos].RawValue
// 	dest[7] = m.rows[m.curpos].IsReadonly
// 	dest[8] = m.rows[m.curpos].IsNullable
// 	m.curpos++
// 	return nil
// }

// func (MockRows) Err() error {
// 	return nil
// }

func TestConfigManager_ApplyTo(t *testing.T) {

	cfg := gonfig.New()

	mdb := &MockDB{}

	cm := New(mdb)
	if err := cm.readParams(); err != nil {
		t.Error(err)
	}

	if err := cm.ApplyTo(cfg, false); err != nil {
		t.Error(err)
	}

	if cfg.IsExist("gopath") == false {
		t.Error("no gopath var")
	}
}
