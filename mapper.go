package main

import "reflect"

func MapFields(source, target interface{}) {
	sourceValue := reflect.ValueOf(source).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceField := sourceValue.Type().Field(i)
		targetField := targetValue.FieldByName(sourceField.Name)

		if targetField.IsValid() && targetField.CanSet() {
			if sourceField.Type.Kind() == reflect.Struct && targetField.Type().Kind() == reflect.Struct {
				MapFields(sourceValue.Field(i).Addr().Interface(), targetField.Addr().Interface())
			} else {
				targetField.Set(sourceValue.Field(i))
			}
		}
	}
}
