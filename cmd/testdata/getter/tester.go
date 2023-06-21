package test

type Tester struct {
	field1              string `accessor:"getter"`
	field2              int32  `accessor:"getter:GetSecondField"`
	field3              *bool
	field_in_snake_case *bool `accessor:"getter"`
}
