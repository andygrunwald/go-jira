// +build generated

package models

import (
	"github.com/ghostsquad/genny/generic"
	"github.com/valyala/fastjson"
)

type GenericType generic.Type

func UnmarshalGenericTypeFromObject(o *fastjson.Object) (GenericType, error) {
	if o == nil {
		return GenericType{}, nil
	}

	return UnmarshalGenericType(o.String())
}
