// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/zhangpy/unioffice"
)

type CT_TableBackgroundStyle struct {
	Fill      *CT_FillProperties
	FillRef   *CT_StyleMatrixReference
	Effect    *CT_EffectProperties
	EffectRef *CT_StyleMatrixReference
}

func NewCT_TableBackgroundStyle() *CT_TableBackgroundStyle {
	ret := &CT_TableBackgroundStyle{}
	return ret
}

func (m *CT_TableBackgroundStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "a:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.FillRef != nil {
		sefillRef := xml.StartElement{Name: xml.Name{Local: "a:fillRef"}}
		e.EncodeElement(m.FillRef, sefillRef)
	}
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "a:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.EffectRef != nil {
		seeffectRef := xml.StartElement{Name: xml.Name{Local: "a:effectRef"}}
		e.EncodeElement(m.EffectRef, seeffectRef)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableBackgroundStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TableBackgroundStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fill"}:
				m.Fill = NewCT_FillProperties()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillRef"}:
				m.FillRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.FillRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effect"}:
				m.Effect = NewCT_EffectProperties()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effectRef"}:
				m.EffectRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.EffectRef, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TableBackgroundStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableBackgroundStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableBackgroundStyle and its children
func (m *CT_TableBackgroundStyle) Validate() error {
	return m.ValidateWithPath("CT_TableBackgroundStyle")
}

// ValidateWithPath validates the CT_TableBackgroundStyle and its children, prefixing error messages with path
func (m *CT_TableBackgroundStyle) ValidateWithPath(path string) error {
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.FillRef != nil {
		if err := m.FillRef.ValidateWithPath(path + "/FillRef"); err != nil {
			return err
		}
	}
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.EffectRef != nil {
		if err := m.EffectRef.ValidateWithPath(path + "/EffectRef"); err != nil {
			return err
		}
	}
	return nil
}
