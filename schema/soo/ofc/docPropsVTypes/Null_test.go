// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangpy/unioffice/schema/soo/ofc/docPropsVTypes"
)

func TestNullConstructor(t *testing.T) {
	v := docPropsVTypes.NewNull()
	if v == nil {
		t.Errorf("docPropsVTypes.NewNull must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Null should validate: %s", err)
	}
}

func TestNullMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewNull()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewNull()
	xml.Unmarshal(buf, v2)
}
