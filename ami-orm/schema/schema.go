package schema

import (
	"amiorm/dialect"
	"go/ast"
	"reflect"
)

//type User struct {
//	Name string `amiorm:"PRIMARY KEY"`
//	Age  int
//}

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	// 对象
	Model interface{}
	// 表名
	Name string
	// 字段Fields
	Fields []*Field
	// FieldNames 所有字段名（列名)
	FieldNames []string
	// 字段名字和Field的映射
	fieldMap map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 将对象解析为Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// Valueof 返回指向dest的指针
	// Indirect 返回类型
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 非匿名以及 首字母大写
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				// TODO go反射
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}

			if v, ok := p.Tag.Lookup("amiorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
