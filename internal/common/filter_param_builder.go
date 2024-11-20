package common

import (
	"fmt"
	"reflect"
	"strings"
)

type FilterParams struct {
}

// BuildFilterConditions
// Dynamically builds the WHERE clause and arguments based on non-nil struct fields.
func (f *FilterParams) BuildFilterConditions(filters interface{}) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	// Dereference the pointer to get the struct value
	value := reflect.ValueOf(filters)
	if value.Kind() == reflect.Ptr {
		value = value.Elem() // Dereference if it's a pointer
	}
	typeOfS := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typeOfS.Field(i)
		fieldName := fieldType.Name

		// Skip if the field is nil or has zero value
		if !isZeroValue(field) {
			columnName := strings.ToLower(fieldName)
			conditions = append(
				conditions,
				fmt.Sprintf("%s ILIKE $%d", columnName, len(args)+1), // Correctly format the placeholder
			)
			args = append(args, field.Interface())
		}
	}

	// Join all conditions with AND
	filterConditions := strings.Join(conditions, " AND ")
	if len(filterConditions) > 0 {
		filterConditions = " AND " + filterConditions
	}

	return filterConditions, args
}

// isZeroValue checks whether the given reflect.Value is a zero value for its type.
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.Len() == 0
	case reflect.Struct:
		// For struct, compare with a zero value struct of the same type
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	default:
		// Compare other types directly with their zero value
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
}
