package gonfman

import (
	"sort"
	"testing"

	"github.com/axkit/gonfig"
)

// type Scanner interface {
// 	Scan(dest ...interface{}) error
// }

// type MockDB struct{}

// func (MockDB) Query(string, ...interface{}) (*MockRows, error) {
// 	return &MockRows{
// 		rows: []Param{
// 			{ID: "gopath", DataTypeID: "string", RawValue: "/home/go/src"},
// 			{ID: "goroot", DataTypeID: "string", RawValue: "/home/go"},
// 			{ID: "goos", DataTypeID: "string", RawValue: "amd64"}},
// 	}, nil
// }

// func (MockDB) Ping() error {
// 	return nil
// }
// func (MockDB) Close() error {
// 	return nil
// }
// func (MockDB) Execute(query string, args ...interface{}) error {
// 	return nil
// }

// func (MockDB) QueryRow(query string, args ...interface{}) Scanner {

// 	return nil
// }

// type MockRows struct {
// 	curpos int
// 	rows   []Param
// }

// func (*MockRows) Columns() ([]string, error) {
// 	return nil, nil
// }

// func (*MockRows) Close() error {
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

// func (*MockRows) Err() error {
// 	return nil
// }

// func TestConfigManager_ApplyTo(t *testing.T) {

// 	cfg := gonfig.New()

// 	mdb := &MockDB{}

// 	cm := New(mdb)
// 	if err := cm.readParams(); err != nil {
// 		t.Error(err)
// 	}

// 	if err := cm.ApplyTo(cfg, false); err != nil {
// 		t.Error(err)
// 	}

// 	if cfg.IsExist("gopath") == false {
// 		t.Error("no gopath var")
// 	}
// }
func TestConfigManager_Sections(t *testing.T) {

	cfg := gonfig.New()

	gfm := New(nil)

	gfm.sections.list = append(gfm.sections.list,
		Section{
			ID:   "root1",
			Name: "root1",
		},
		Section{
			ID:   "root2",
			Name: "root2",
		},
		Section{
			ID:       "root21",
			Name:     "root21",
			ParentID: "root2",
		},
		Section{
			ID:       "root22",
			Name:     "root22",
			ParentID: "root2",
		},
		Section{
			ID:       "root221",
			Name:     "root221",
			ParentID: "root22",
		},
		Section{
			ID:   "root3",
			Name: "root3",
		},
	)

	sort.Slice(gfm.sections.list, func(i, j int) bool {
		if gfm.sections.list[i].ParentID == gfm.sections.list[j].ParentID {
			return gfm.sections.list[i].PositionOrder < gfm.sections.list[j].PositionOrder
		}
		return gfm.sections.list[i].PositionOrder < gfm.sections.list[j].PositionOrder
	})
}
