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

var CountryLanguage = newCountryLanguageTable("public", "country_language", "")

type countryLanguageTable struct {
	postgres.Table

	// Columns
	CountryCode postgres.ColumnString
	Language    postgres.ColumnString
	IsOfficial  postgres.ColumnBool
	Percentage  postgres.ColumnFloat

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CountryLanguageTable struct {
	countryLanguageTable

	EXCLUDED countryLanguageTable
}

// AS creates new CountryLanguageTable with assigned alias
func (a CountryLanguageTable) AS(alias string) *CountryLanguageTable {
	return newCountryLanguageTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CountryLanguageTable with assigned schema name
func (a CountryLanguageTable) FromSchema(schemaName string) *CountryLanguageTable {
	return newCountryLanguageTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CountryLanguageTable with assigned table prefix
func (a CountryLanguageTable) WithPrefix(prefix string) *CountryLanguageTable {
	return newCountryLanguageTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CountryLanguageTable with assigned table suffix
func (a CountryLanguageTable) WithSuffix(suffix string) *CountryLanguageTable {
	return newCountryLanguageTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCountryLanguageTable(schemaName, tableName, alias string) *CountryLanguageTable {
	return &CountryLanguageTable{
		countryLanguageTable: newCountryLanguageTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newCountryLanguageTableImpl("", "excluded", ""),
	}
}

func newCountryLanguageTableImpl(schemaName, tableName, alias string) countryLanguageTable {
	var (
		CountryCodeColumn = postgres.StringColumn("country_code")
		LanguageColumn    = postgres.StringColumn("language")
		IsOfficialColumn  = postgres.BoolColumn("is_official")
		PercentageColumn  = postgres.FloatColumn("percentage")
		allColumns        = postgres.ColumnList{CountryCodeColumn, LanguageColumn, IsOfficialColumn, PercentageColumn}
		mutableColumns    = postgres.ColumnList{IsOfficialColumn, PercentageColumn}
	)

	return countryLanguageTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CountryCode: CountryCodeColumn,
		Language:    LanguageColumn,
		IsOfficial:  IsOfficialColumn,
		Percentage:  PercentageColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
