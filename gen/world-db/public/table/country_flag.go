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

var CountryFlag = newCountryFlagTable("public", "country_flag", "")

type countryFlagTable struct {
	postgres.Table

	// Columns
	Code2   postgres.ColumnString
	Emoji   postgres.ColumnString
	Unicode postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CountryFlagTable struct {
	countryFlagTable

	EXCLUDED countryFlagTable
}

// AS creates new CountryFlagTable with assigned alias
func (a CountryFlagTable) AS(alias string) *CountryFlagTable {
	return newCountryFlagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CountryFlagTable with assigned schema name
func (a CountryFlagTable) FromSchema(schemaName string) *CountryFlagTable {
	return newCountryFlagTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CountryFlagTable with assigned table prefix
func (a CountryFlagTable) WithPrefix(prefix string) *CountryFlagTable {
	return newCountryFlagTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CountryFlagTable with assigned table suffix
func (a CountryFlagTable) WithSuffix(suffix string) *CountryFlagTable {
	return newCountryFlagTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCountryFlagTable(schemaName, tableName, alias string) *CountryFlagTable {
	return &CountryFlagTable{
		countryFlagTable: newCountryFlagTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newCountryFlagTableImpl("", "excluded", ""),
	}
}

func newCountryFlagTableImpl(schemaName, tableName, alias string) countryFlagTable {
	var (
		Code2Column    = postgres.StringColumn("code2")
		EmojiColumn    = postgres.StringColumn("emoji")
		UnicodeColumn  = postgres.StringColumn("unicode")
		allColumns     = postgres.ColumnList{Code2Column, EmojiColumn, UnicodeColumn}
		mutableColumns = postgres.ColumnList{EmojiColumn, UnicodeColumn}
	)

	return countryFlagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Code2:   Code2Column,
		Emoji:   EmojiColumn,
		Unicode: UnicodeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
