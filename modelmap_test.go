package main

import (
	"github.com/yowayimono/mapper/mapper"
	"reflect"
	"testing"
)

type SourceStruct struct {
	Username string
	Password string
	Phone    string
}

type TargetStruct struct {
	Username string
	Password string
	Phone    string
}

type NestedStruct struct {
	Field1 string
	Field2 int
}

type SourceStructWithNested struct {
	Username string
	Password string
	Phone    string
	Nested   NestedStruct
}

type TargetStructWithNested struct {
	Username string
	Password string

	Nested NestedStruct
}

func TestMapFields(t *testing.T) {
	source := SourceStructWithNested{
		Username: "john",
		Password: "password123",
		Phone:    "1234567890",
		Nested: NestedStruct{
			Field1: "1234",
			Field2: 18,
		},
	}

	target := TargetStructWithNested{}

	mapper.Maps(&source, &target)

	expected := TargetStructWithNested{
		Username: "john",
		Password: "password123",
		Nested: NestedStruct{
			Field1: "1234",
			Field2: 18,
		},
	}

	if !reflect.DeepEqual(target, expected) {
		t.Errorf("MapFields did not map the fields correctly. Expected: %v, got: %v", expected, target)
	}
}
