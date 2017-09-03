// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_PhoneticRun struct {
	// Base Text Start Index
	SbAttr uint32
	// Base Text End Index
	EbAttr uint32
	// Text
	T string
}

func NewCT_PhoneticRun() *CT_PhoneticRun {
	ret := &CT_PhoneticRun{}
	return ret
}

func (m *CT_PhoneticRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sb"},
		Value: fmt.Sprintf("%v", m.SbAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "eb"},
		Value: fmt.Sprintf("%v", m.EbAttr)})
	e.EncodeToken(start)
	set := xml.StartElement{Name: xml.Name{Local: "x:t"}}
	gooxml.AddPreserveSpaceAttr(&set, m.T)
	e.EncodeElement(m.T, set)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PhoneticRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sb" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SbAttr = uint32(parsed)
		}
		if attr.Name.Local == "eb" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.EbAttr = uint32(parsed)
		}
	}
lCT_PhoneticRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "t":
				if err := d.DecodeElement(&m.T, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PhoneticRun %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PhoneticRun
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PhoneticRun and its children
func (m *CT_PhoneticRun) Validate() error {
	return m.ValidateWithPath("CT_PhoneticRun")
}

// ValidateWithPath validates the CT_PhoneticRun and its children, prefixing error messages with path
func (m *CT_PhoneticRun) ValidateWithPath(path string) error {
	return nil
}
