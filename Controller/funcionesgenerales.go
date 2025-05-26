package Controller

import (
	"database/sql"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Option struct {
	Value string
	Label string
}

func GetOpcionesPregunta(db *sql.DB, Query string) ([]Option, error) {
	rows, err := db.Query(Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opciones []Option
	for rows.Next() {
		var value, label string
		err := rows.Scan(&value, &label)
		if err != nil {
			return nil, err
		}
		opciones = append(opciones, Option{Value: value, Label: label})
	}

	return opciones, nil
}

func QuitarTildesYMayusculas(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, input)
	return strings.ToUpper(result)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) && r != 'ñ' && r != 'Ñ' // Mn: nonspacing marks
}
