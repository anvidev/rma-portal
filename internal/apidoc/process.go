package apidoc

import (
	"reflect"
	"strings"
)

// structToSlice converts a struct into a map of body fields. It handles nested structs and slices recursively.
func structToSlice(i any) []BodyField {
	reflectValue := reflect.ValueOf(i)

	if reflectValue.Kind() == reflect.Pointer {
		// if pointer, dereference it
		reflectValue = reflectValue.Elem()
	}

	if reflectValue.Kind() != reflect.Struct {
		return nil
	}

	reflectType := reflectValue.Type()
	fieldCount := reflectValue.NumField()
	var result []BodyField

	for i := range fieldCount {
		field := reflectValue.Field(i)
		fieldType := reflectType.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		description := fieldType.Tag.Get("description")
		validateTag := fieldType.Tag.Get("validate")
		apidocTag := fieldType.Tag.Get("apidoc")

		if !fieldType.IsExported() || jsonTag == "-" || strings.Contains(apidocTag, "ignore") {
			continue
		}

		fieldName := fieldType.Name
		fieldKind := field.Kind()

		jsonName := fieldName
		if jsonTag != "" {
			jsonName = strings.Split(jsonTag, ",")[0]
		}

		validation := parseValidationRules(validateTag)

		bodyFieldType := fieldType.Type.String()

		if strings.Contains(bodyFieldType, ".") {
			bodyFieldType = strings.Split(bodyFieldType, ".")[1]
		}

		bodyField := BodyField{
			Name:        jsonName,
			Type:        bodyFieldType,
			Required:    strings.Contains(validateTag, "required"),
			Description: description,
			Validation:  validation,
		}

		if fieldKind == reflect.Struct || (fieldKind == reflect.Pointer && field.Elem().Kind() == reflect.Struct) {
			bodyField.Fields = structToSlice(field.Interface())
		} else if fieldKind == reflect.Slice {
			sliceType := fieldType.Type.Elem()
			if strings.Contains(sliceType.String(), ".") {
				bodyField.Type = "[]" + strings.Split(sliceType.String(), ".")[1]
			} else {
				bodyField.Type = "[]" + sliceType.String()
			}

			if sliceType.Kind() == reflect.Struct || (sliceType.Kind() == reflect.Pointer && sliceType.Elem().Kind() == reflect.Struct) {
				bodyField.Fields = structToSlice(reflect.New(sliceType).Interface())
			}
		}
		result = append(result, bodyField)
	}

	return result
}

func parseValidationRules(validateTag string) map[string]string {
	rules := make(map[string]string)
	if validateTag == "" {
		return rules
	}

	for _, rule := range strings.Split(validateTag, ",") {
		rule := strings.TrimSpace(rule)
		parts := strings.SplitN(rule, "=", 2)
		if len(parts) == 2 {
			rules[parts[0]] = parts[1]
		}
	}

	return rules
}
