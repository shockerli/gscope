package main

var tplPkg = `
// Code generated by github.com/shockerli/gormer/sql-builder-gen; DO NOT EDIT

// Package {{.pkg}} SQL build helper functions for field
package {{.pkg}}

import (
{{- range $v := .import}}
	{{$v}}
{{- end}}
)
`

var tplFunc = map[string]string{
	EQ: `
// Eq WHERE {{.Column}} = ?
func (i {{.Name}}) Eq(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} = ?", v[0])
	}
}
`,
	NEQ: `
// Neq WHERE {{.Column}} <> ?
func (i {{.Name}}) Neq(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} <> ?", v[0])
	}
}
`,
	GT: `
// Gt WHERE {{.Column}} > ?
func (i {{.Name}}) Gt(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} > ?", v[0])
	}
}
`,
	GTE: `
// Gte WHERE {{.Column}} >= ?
func (i {{.Name}}) Gte(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} >= ?", v[0])
	}
}
`,
	LT: `
// Lt WHERE {{.Column}} < ?
func (i {{.Name}}) Lt(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} < ?", v[0])
	}
}
`,
	LTE: `
// Lte WHERE {{.Column}} <= ?
func (i {{.Name}}) Lte(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} <= ?", v[0])
	}
}
`,
	IN: `
// In WHERE {{.Column}} IN (?)
func (i {{.Name}}) In(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} IN (?)", v)
	}
}
`,
	NOTIN: `
// NotIn WHERE {{.Column}} NOT IN (?)
func (i {{.Name}}) NotIn(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} NOT IN (?)", v)
	}
}
`,
	BETWEEN: `
// Between WHERE {{.Column}} BETWEEN ? AND ?
func (i {{.Name}}) Between(a, b {{.Type}}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} BETWEEN ? AND ?", a, b)
	}
}
`,
	LIKE: `
// Like WHERE {{.Column}} LIKE '%?%'
func (i {{.Name}}) Like(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} LIKE '%?%'", v[0])
	}
}
`,
	RLIKE: `
// RLike WHERE {{.Column}} LIKE '?%'
func (i {{.Name}}) RLike(v ...{{.Type}}) func(db *gorm.DB) *gorm.DB {
	if len(v) == 0 {
		v = append(v, i.{{.Field}})
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.Column}} LIKE '?%'", v[0])
	}
}
`,
}

var tplExtend = map[string]string{
	TIME: `
// Datetime return datetime
func (i {{.Name}}) Datetime() string {
	return time.Unix(i.{{.Field}}, 0).Format("2006-01-02 15:04:05")
}

// Date return date
func (i {{.Name}}) Date() string {
	return time.Unix(i.{{.Field}}, 0).Format("2006-01-02")
}

// DateInt return date with int
func (i {{.Name}}) DateInt() int {
	v, _ := strconv.Atoi(time.Unix(i.CT, 0).Format("20060102"))
	return v
}

// Time return time.Time
func (i {{.Name}}) Time() time.Time {
	return time.Unix(i.{{.Field}}, 0)
}

// String return format datetime
func (i {{.Name}}) String() string {
	return i.Datetime()
}
`,
}
