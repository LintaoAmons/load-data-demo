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

var Country = newCountryTable("public", "country", "")

type countryTable struct {
	postgres.Table

	// Columns
	Code           postgres.ColumnString
	Name           postgres.ColumnString
	Continent      postgres.ColumnString
	Region         postgres.ColumnString
	SurfaceArea    postgres.ColumnFloat
	IndepYear      postgres.ColumnInteger
	Population     postgres.ColumnInteger
	LifeExpectancy postgres.ColumnFloat
	Gnp            postgres.ColumnFloat
	GnpOld         postgres.ColumnFloat
	LocalName      postgres.ColumnString
	GovernmentForm postgres.ColumnString
	HeadOfState    postgres.ColumnString
	Capital        postgres.ColumnInteger
	Code2          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CountryTable struct {
	countryTable

	EXCLUDED countryTable
}

// AS creates new CountryTable with assigned alias
func (a CountryTable) AS(alias string) *CountryTable {
	return newCountryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CountryTable with assigned schema name
func (a CountryTable) FromSchema(schemaName string) *CountryTable {
	return newCountryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CountryTable with assigned table prefix
func (a CountryTable) WithPrefix(prefix string) *CountryTable {
	return newCountryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CountryTable with assigned table suffix
func (a CountryTable) WithSuffix(suffix string) *CountryTable {
	return newCountryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCountryTable(schemaName, tableName, alias string) *CountryTable {
	return &CountryTable{
		countryTable: newCountryTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newCountryTableImpl("", "excluded", ""),
	}
}

func newCountryTableImpl(schemaName, tableName, alias string) countryTable {
	var (
		CodeColumn           = postgres.StringColumn("code")
		NameColumn           = postgres.StringColumn("name")
		ContinentColumn      = postgres.StringColumn("continent")
		RegionColumn         = postgres.StringColumn("region")
		SurfaceAreaColumn    = postgres.FloatColumn("surface_area")
		IndepYearColumn      = postgres.IntegerColumn("indep_year")
		PopulationColumn     = postgres.IntegerColumn("population")
		LifeExpectancyColumn = postgres.FloatColumn("life_expectancy")
		GnpColumn            = postgres.FloatColumn("gnp")
		GnpOldColumn         = postgres.FloatColumn("gnp_old")
		LocalNameColumn      = postgres.StringColumn("local_name")
		GovernmentFormColumn = postgres.StringColumn("government_form")
		HeadOfStateColumn    = postgres.StringColumn("head_of_state")
		CapitalColumn        = postgres.IntegerColumn("capital")
		Code2Column          = postgres.StringColumn("code2")
		allColumns           = postgres.ColumnList{CodeColumn, NameColumn, ContinentColumn, RegionColumn, SurfaceAreaColumn, IndepYearColumn, PopulationColumn, LifeExpectancyColumn, GnpColumn, GnpOldColumn, LocalNameColumn, GovernmentFormColumn, HeadOfStateColumn, CapitalColumn, Code2Column}
		mutableColumns       = postgres.ColumnList{NameColumn, ContinentColumn, RegionColumn, SurfaceAreaColumn, IndepYearColumn, PopulationColumn, LifeExpectancyColumn, GnpColumn, GnpOldColumn, LocalNameColumn, GovernmentFormColumn, HeadOfStateColumn, CapitalColumn, Code2Column}
	)

	return countryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Code:           CodeColumn,
		Name:           NameColumn,
		Continent:      ContinentColumn,
		Region:         RegionColumn,
		SurfaceArea:    SurfaceAreaColumn,
		IndepYear:      IndepYearColumn,
		Population:     PopulationColumn,
		LifeExpectancy: LifeExpectancyColumn,
		Gnp:            GnpColumn,
		GnpOld:         GnpOldColumn,
		LocalName:      LocalNameColumn,
		GovernmentForm: GovernmentFormColumn,
		HeadOfState:    HeadOfStateColumn,
		Capital:        CapitalColumn,
		Code2:          Code2Column,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}