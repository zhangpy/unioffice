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

type EG_LineJoinProperties struct {
	Round *CT_LineJoinRound
	Bevel *CT_LineJoinBevel
	Miter *CT_LineJoinMiterProperties
}

func NewEG_LineJoinProperties() *EG_LineJoinProperties {
	ret := &EG_LineJoinProperties{}
	return ret
}

func (m *EG_LineJoinProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Round != nil {
		seround := xml.StartElement{Name: xml.Name{Local: "a:round"}}
		e.EncodeElement(m.Round, seround)
	}
	if m.Bevel != nil {
		sebevel := xml.StartElement{Name: xml.Name{Local: "a:bevel"}}
		e.EncodeElement(m.Bevel, sebevel)
	}
	if m.Miter != nil {
		semiter := xml.StartElement{Name: xml.Name{Local: "a:miter"}}
		e.EncodeElement(m.Miter, semiter)
	}
	return nil
}

func (m *EG_LineJoinProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_LineJoinProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "round"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "round"}:
				m.Round = NewCT_LineJoinRound()
				if err := d.DecodeElement(m.Round, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bevel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "bevel"}:
				m.Bevel = NewCT_LineJoinBevel()
				if err := d.DecodeElement(m.Bevel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "miter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "miter"}:
				m.Miter = NewCT_LineJoinMiterProperties()
				if err := d.DecodeElement(m.Miter, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_LineJoinProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_LineJoinProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_LineJoinProperties and its children
func (m *EG_LineJoinProperties) Validate() error {
	return m.ValidateWithPath("EG_LineJoinProperties")
}

// ValidateWithPath validates the EG_LineJoinProperties and its children, prefixing error messages with path
func (m *EG_LineJoinProperties) ValidateWithPath(path string) error {
	if m.Round != nil {
		if err := m.Round.ValidateWithPath(path + "/Round"); err != nil {
			return err
		}
	}
	if m.Bevel != nil {
		if err := m.Bevel.ValidateWithPath(path + "/Bevel"); err != nil {
			return err
		}
	}
	if m.Miter != nil {
		if err := m.Miter.ValidateWithPath(path + "/Miter"); err != nil {
			return err
		}
	}
	return nil
}
