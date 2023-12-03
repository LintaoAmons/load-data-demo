package main

import (
	"fmt"

	"github.com/LintaoAmons/load-data-to-db/gen/world-db/public/model"
	"github.com/LintaoAmons/ltoolbox/functions"
	"github.com/LintaoAmons/ltoolbox/functions/sql"
)

func generateOneCity(id int) string {
	localName := "localName"
	a := model.City{
		ID:          int32(id) + 10000, // to avoid duplicated pk
		Name:        "Name",
		CountryCode: "SGP",
		District:    "District",
		Population:  1234,
		LocalName:   &localName,
	}

	return functions.StructToCsvRow(a, false)
}

func main() {
	sql.GenerateCsvFiles(1000000, 10, sql.RowGeneratorFunc(generateOneCity), "city")
	fmt.Println("Generated")
}
