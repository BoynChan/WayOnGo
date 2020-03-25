package schema

import (
	"7-days-gdorm/gdorm/dialect"
	"go/ast"
	"reflect"
)

// Author:Boyn
// Date:2020/3/23

type Field struct {
	Name string
	Type string
	Tag  string
}

// 被映射的对象
// 包含model是对象本身,name是表名,fields与fieldNames记录所有的字段和列名
// fieldMap记录字段名和field的映射关系
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 我们使用的是对dest传入先取值,再取其指针,所以传入的参数应为结构体对应的指针
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 判断字段p是否可导出
		// 如果是可导出的,那就将其写入schema中
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("gdorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, field.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
