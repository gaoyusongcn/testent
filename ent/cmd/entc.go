//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {

	opts := []entc.Option{}

	err := entc.Generate("../schema", &gen.Config{
		// 设置生成的 Schema 自增 ID 类型为 Uint64
		IDType: &field.TypeInfo{Type: field.TypeUint64},
		// 设置生成代码存放的路径 DOT EDIT!
		Target:   "../gen",
		Package:  "testent/ent/gen",
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...)

	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
