//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var City = newCityTable("public", "city", "")

type cityTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	Name        postgres.ColumnString
	CountryCode postgres.ColumnString
	District    postgres.ColumnString
	Population  postgres.ColumnInteger
	LocalName   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CityTable struct {
	cityTable

	EXCLUDED cityTable
}

// AS creates new CityTable with assigned alias
func (a CityTable) AS(alias string) *CityTable {
	return newCityTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CityTable with assigned schema name
func (a CityTable) FromSchema(schemaName string) *CityTable {
	return newCityTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CityTable with assigned table prefix
func (a CityTable) WithPrefix(prefix string) *CityTable {
	return newCityTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CityTable with assigned table suffix
func (a CityTable) WithSuffix(suffix string) *CityTable {
	return newCityTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCityTable(schemaName, tableName, alias string) *CityTable {
	return &CityTable{
		cityTable: newCityTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newCityTableImpl("", "excluded", ""),
	}
}

func newCityTableImpl(schemaName, tableName, alias string) cityTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		NameColumn        = postgres.StringColumn("name")
		CountryCodeColumn = postgres.StringColumn("country_code")
		DistrictColumn    = postgres.StringColumn("district")
		PopulationColumn  = postgres.IntegerColumn("population")
		LocalNameColumn   = postgres.StringColumn("local_name")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, CountryCodeColumn, DistrictColumn, PopulationColumn, LocalNameColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, CountryCodeColumn, DistrictColumn, PopulationColumn, LocalNameColumn}
	)

	return cityTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		CountryCode: CountryCodeColumn,
		District:    DistrictColumn,
		Population:  PopulationColumn,
		LocalName:   LocalNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}