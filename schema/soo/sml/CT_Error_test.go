// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangpy/unioffice/schema/soo/sml"
)

func TestCT_ErrorConstructor(t *testing.T) {
	v := sml.NewCT_Error()
	if v == nil {
		t.Errorf("sml.NewCT_Error must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Error should validate: %s", err)
	}
}

func TestCT_ErrorMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Error()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Error()
	xml.Unmarshal(buf, v2)
}
