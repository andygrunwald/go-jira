package models

import (
	"fmt"
	"github.com/valyala/fastjson"
)

var issueJsonParserPool fastjson.ParserPool

// Issue represents a Jira issue.
type Issue struct {
	Expand         string               `json:"expand,omitempty" structs:"expand,omitempty"`
	ID             string               `json:"id,omitempty" structs:"id,omitempty"`
	Self           string               `json:"self,omitempty" structs:"self,omitempty"`
	Key            string               `json:"key,omitempty" structs:"key,omitempty"`
	Fields         *IssueFields         `json:"fields,omitempty" structs:"fields,omitempty"`
	RenderedFields *IssueRenderedFields `json:"renderedFields,omitempty" structs:"renderedFields,omitempty"`
	Changelog      *Changelog           `json:"changelog,omitempty" structs:"changelog,omitempty"`
	Transitions    []Transition         `json:"transitions,omitempty" structs:"transitions,omitempty"`
	Names          map[string]string    `json:"names,omitempty" structs:"names,omitempty"`
}

// Unmarshal unmarshals JSON s to issue.
func (x *Issue) Unmarshal(j string) error {

	p := issueJsonParserPool.Get()
	defer issueJsonParserPool.Put(p)
	v, err := p.Parse(j)

	if err != nil {
		return fmt.Errorf("failed to parse issue json: %w", err)
	}

	x.Expand = string(v.GetStringBytes("expand"))
	x.ID = string(v.GetStringBytes("id"))
	x.Self = string(v.GetStringBytes("self"))
	x.Key = string(v.GetStringBytes("key"))

	{
		y := &IssueFields{}
		o := v.GetObject("fields")
		err := y.UnmarshalFromObj(o)
		if err != nil {
			return err
		}
	}

	{
		y := &IssueRenderedFields{}
		err = y.Unmarshal(v.GetObject("renderedFields").String())
		if err != nil {
			return err
		}
		x.RenderedFields = y
	}

	{
		y := &Changelog{}
		err = y.Unmarshal(v.GetObject("changelog").String())
		if err != nil {
			return err
		}
		x.Changelog = y
	}

	{
		for _, jV := range v.GetArray("transitions") {
			t, err := UnmarshalTransition(jV.String())
			if err != nil {
				return err
			}
			x.Transitions = append(x.Transitions, t)
		}
	}

	{
		y := make(map[string]string)
		n := v.Get("names")
	}

	ms.Foo = append(ms.Foo[:0], v.GetStringBytes("foo")...)
	ms.A = ms.A[:0]
	for _, av := range v.GetArray("a") {
		ms.A = append(ms.A, av.GetInt()
	}
	return nil
}