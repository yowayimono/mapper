package mapper

import "reflect"

func Map(source, target interface{}) {
	sourceValue := reflect.ValueOf(source).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	mapping := getFieldMapping(targetValue.Type())

	for sourceField, targetField := range mapping {
		if sourceValueField := sourceValue.FieldByName(sourceField); sourceValueField.IsValid() && sourceValueField.CanSet() {
			targetValueField := targetValue.FieldByName(targetField)
			if targetValueField.IsValid() && targetValueField.CanSet() {
				targetValueField.Set(sourceValueField)
			}
		}
	}
}

func getFieldMapping(targetType reflect.Type) map[string]string {
	mapping := make(map[string]string)

	for i := 0; i < targetType.NumField(); i++ {
		targetField := targetType.Field(i)
		targetFieldName := targetField.Name
		sourceFieldName := targetField.Tag.Get("mapper")

		if sourceFieldName != "" {
			mapping[sourceFieldName] = targetFieldName
		} else {
			mapping[targetFieldName] = targetFieldName
		}
	}

	return mapping
}
