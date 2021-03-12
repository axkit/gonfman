package gonfman

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/axkit/gonfig"
)

// Param describes param stored in SQL database. Param's value is stored as string
// in RawValue field.
type Param struct {
	ID                string    `json:"id"`
	SectionID         string    `json:"sectionID"`
	PositionOrder     int       `json:"positionOrder"`
	Name              string    `json:"name"`
	DataTypeID        string    `json:"dataTypeID"`
	ControlID         *string   `json:"controlID"`
	RawValue          string    `json:"rawValue"`
	IsReadonly        bool      `json:"isReadonly"`
	IsNullable        bool      `json:"isNullable"`
	UpdatedAt         time.Time `json:"-"`
	UpdateFingerPrint int       `json:"-"`
}

// Section describes a hierarchy of configuration parameters.
type Section struct {
	ID            string `json:"id"`
	ParentID      string `json:"parentID"`
	PositionOrder int    `json:"positionOrder"`
	Name          string `json:"name"`
}

// Control specify ui control to be used for param's value modification.
type Control struct {
	ID                       string  `json:"id"`
	ValidationFunction       *string `json:"validateFunction"`
	FailedValidationResponse *string `json:"failedValidationResponse"`
}

var (
	// TableNameParam holds default table name used as params/values storage.
	TableNameParam = "config_params"

	// TableNameSection holds default table name holding heirarchy of parameters.
	TableNameSection = "config_sections"

	// TableNameControl holds default table name holding ui control names.
	TableNameControl = "config_controls"

	// Mapping holds mapping between gonfig kinds and params.data_type_id.
	Mapping = map[string]gonfig.AKind{
		"int":    gonfig.AInt,
		"bool":   gonfig.ABool,
		"string": gonfig.AString,
		"float":  gonfig.AFloat,
	}
)

// ErrUnsupportedDataType indicates invalid element in Mapping.
var ErrUnsupportedDataType = errors.New("column data_type_id has unknown value")

// ConfigManager implements logic or reading application parameters
// from the sql database.
type ConfigManager struct {
	db     *sql.DB
	params struct {
		list []Param
	}

	sections struct {
		list []Section
	}

	controls struct {
		list []Control
	}
}

// New returns ConfigManager.
func New(db *sql.DB) *ConfigManager {
	return &ConfigManager{db: db}
}

// Init caches rows from config_* tables.
func (s *ConfigManager) Init(ctx context.Context) error {
	if err := s.readSections(); err != nil {
		return err
	}

	if err := s.readControls(); err != nil {
		return err
	}

	if err := s.readParams(); err != nil {
		return err
	}

	return nil
}

// Start
func (s *ConfigManager) Start(ctx context.Context) error {
	return nil
}

func (cm *ConfigManager) readSections() error {

	qry := `select id, parent_id, position_order, name from ` + TableNameSection

	rows, err := cm.db.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()

	var (
		s        Section
		parentID *string
	)

	for rows.Next() {
		if err := rows.Scan(
			&s.ID,
			&parentID,
			&s.PositionOrder,
			&s.Name,
		); err != nil {
			return err
		}
		if parentID == nil {
			s.ParentID = ""
		} else {
			s.ParentID = *parentID
		}
		cm.sections.list = append(cm.sections.list, s)
	}
	return rows.Err()
}

func (cm *ConfigManager) readControls() error {

	qry := `select id, validation_function, failed_validation_response from ` + TableNameControl

	rows, err := cm.db.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()

	var c Control

	for rows.Next() {
		if err := rows.Scan(
			&c.ID,
			&c.ValidationFunction,
			&c.FailedValidationResponse,
		); err != nil {
			return err
		}
		cm.controls.list = append(cm.controls.list, c)
	}
	return rows.Err()
}

func (cm *ConfigManager) readParams() error {

	qry := `select id, section_id, position_order, name, data_type_id, control_id, raw_value, is_readonly, is_nullable from ` + TableNameParam

	rows, err := cm.db.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()

	var (
		p         Param
		sectionID *string
	)

	for rows.Next() {
		if err := rows.Scan(
			&p.ID,
			&sectionID,
			&p.PositionOrder,
			&p.Name,
			&p.DataTypeID,
			&p.ControlID,
			&p.RawValue,
			&p.IsReadonly,
			&p.IsNullable,
		); err != nil {
			return err
		}
		if sectionID == nil {
			p.SectionID = ""
		} else {
			p.SectionID = *sectionID
		}
		cm.params.list = append(cm.params.list, p)
	}
	return rows.Err()
}

// ApplyTo copies all rows from table TableNameParam.
func (cm *ConfigManager) ApplyTo(g gonfig.Configer, ow bool) error {
	return cm.applyTo(g, ow)
}

func (cm *ConfigManager) applyTo(g gonfig.Configer, ow bool) error {

	for _, p := range cm.params.list {

		ak, ok := Mapping[p.DataTypeID]
		if !ok {
			return ErrUnsupportedDataType
		}

		if !ow && g.IsExist(p.ID) {
			continue
		}

		err := g.MustParam(p.ID, ak).Parse(p.RawValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cm *ConfigManager) Controls() []Control {
	res := make([]Control, len(cm.controls.list))
	copy(res, cm.controls.list)
	return res
}

func (cm *ConfigManager) Sections() []Section {
	res := make([]Section, len(cm.sections.list))
	copy(res, cm.sections.list)
	return res
}

func (cm *ConfigManager) Params() []Param {
	res := make([]Param, len(cm.params.list))
	copy(res, cm.params.list)
	return res
}

func (cm *ConfigManager) UpdateParams(m map[string]string) error {
	return nil
}
