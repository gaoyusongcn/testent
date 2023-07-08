package ent

// 使用命令行参数生成 Schema 文件
// --feature sql/versioned-migration 	用版本迁移
// --idtype uint64 						默认的表 ID 用 uint64
// --target 							第一个位置指定通过 schema 生成的代码目录 第二个位置指定 schema 目录
//--go:generate go run -x -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --idtype uint64 --target ./gen ./schema

// 使用 entc 代码生成 Schema 文件
//go:generate go run -mod=mod entc.go
