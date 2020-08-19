package schema

import (
	"amiorm/dialect"
	"testing"
)

type User struct {
	Name string `amiorm:"primary key"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("mysql")

func Test_parse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	//fmt.Printf("parsed schema %v", json)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "primary key" {
		t.Fatal("fail to parse primary")
	}
}
