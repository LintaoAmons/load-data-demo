//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type CountryLanguage struct {
	CountryCode string `sql:"primary_key"`
	Language    string `sql:"primary_key"`
	IsOfficial  bool
	Percentage  float32
}
