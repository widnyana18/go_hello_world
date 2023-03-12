package main

import "reflect"

func checkDataType(data interface{}) {
	var typeDt = reflect.TypeOf(data)
	var value = reflect.ValueOf(data)

	if value.Kind() == reflect.String {
		println("value: ", value.String())
	}

	println("type: ", typeDt)
	println("field: ", typeDt.Field(0).Name)
}

func inspectData(data interface{}) {
	var value = reflect.ValueOf(data)

	println("num field: ", value.NumField())

	for i := 0; i < value.NumField(); i++ {
		if value.Kind() == reflect.Map {
			println("property type: ", value.Type().Field(i).Type)
			println("property name: ", value.Type().Field(i).Name)
			// println("property value: ", value.Field(i))
		}
	}
}
