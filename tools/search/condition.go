package search

import "strings"

type Condition interface {
	SetWhere(k string, v []interface{})
	SetOr(k string, v []interface{})
	SetOrder(k string)
	SetJoinOn(t, on string) Condition
	WithAliasMap(aliasMap map[string]string) Condition
	GetAliasMap() map[string]string
}

type GormCondition struct {
	GormPublic
	Join []*GormJoin
}

type GormPublic struct {
	Where    map[string][]interface{}
	Order    []string
	Or       map[string][]interface{}
	AliasMap map[string]string
}

type GormJoin struct {
	Type   string
	JoinOn string
	GormPublic
}

func (e *GormJoin) SetJoinOn(t, on string) Condition {
	return nil
}

func (e *GormPublic) SetWhere(k string, v []interface{}) {
	if e.Where == nil {
		e.Where = make(map[string][]interface{})
	}
	e.Where[k] = v
}

func (e *GormPublic) SetOr(k string, v []interface{}) {
	if e.Or == nil {
		e.Or = make(map[string][]interface{})
	}
	e.Or[k] = v
}

func (e *GormPublic) SetOrder(k string) {
	if e.Order == nil {
		e.Order = make([]string, 0)
	}
	e.Order = append(e.Order, k)
}

func (e *GormCondition) SetJoinOn(t, on string) Condition {
	if e.Join == nil {
		e.Join = make([]*GormJoin, 0)
	}
	join := &GormJoin{
		Type:       t,
		JoinOn:     on,
		GormPublic: GormPublic{},
	}
	e.Join = append(e.Join, join)
	return join
}

func (c *GormCondition) WithAliasMap(aliasMap map[string]string) Condition {
	c.AliasMap = aliasMap
	return c
}
func (c *GormJoin) WithAliasMap(aliasMap map[string]string) Condition {
	c.AliasMap = aliasMap
	return c
}
func (c *GormCondition) GetAliasMap() map[string]string {
	return c.AliasMap
}
func (c *GormJoin) GetAliasMap() map[string]string {
	return c.AliasMap
}

type resolveSearchTag struct {
	Type      string
	Column    string
	Table     string
	On        []string
	Join      string
	JoinAlias string
}

// makeTag 解析search的tag标签
func makeTag(tag string) *resolveSearchTag {
	r := &resolveSearchTag{}
	tags := strings.Split(tag, ";")
	var ts []string
	for _, t := range tags {
		ts = strings.Split(t, ":")
		if len(ts) == 0 {
			continue
		}
		switch ts[0] {
		case "type":
			if len(ts) > 1 {
				r.Type = ts[1]
			}
		case "column":
			if len(ts) > 1 {
				r.Column = ts[1]
			}
		case "table":
			if len(ts) > 1 {
				r.Table = ts[1]
			}
		case "on":
			if len(ts) > 1 {
				r.On = ts[1:]
			}
		case "join":
			if len(ts) > 1 {
				r.Join = ts[1]
			}
		case "alias":
			if len(ts) > 1 {
				r.JoinAlias = ts[1]
			}
		}
	}
	return r
}
