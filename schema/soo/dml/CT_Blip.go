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
	"fmt"

	"github.com/zhangpy/unioffice"
)

type CT_Blip struct {
	CstateAttr ST_BlipCompression
	Choice     []*CT_BlipChoice
	ExtLst     *CT_OfficeArtExtensionList
	EmbedAttr  *string
	LinkAttr   *string
}

func NewCT_Blip() *CT_Blip {
	ret := &CT_Blip{}
	return ret
}

func (m *CT_Blip) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CstateAttr != ST_BlipCompressionUnset {
		attr, err := m.CstateAttr.MarshalXMLAttr(xml.Name{Local: "cstate"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EmbedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:embed"},
			Value: fmt.Sprintf("%v", *m.EmbedAttr)})
	}
	if m.LinkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:link"},
			Value: fmt.Sprintf("%v", *m.LinkAttr)})
	}
	e.EncodeToken(start)
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Blip) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "embed" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EmbedAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "link" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LinkAttr = &parsed
			continue
		}
		if attr.Name.Local == "cstate" {
			m.CstateAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Blip:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaBiLevel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaBiLevel"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaBiLevel, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaCeiling"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaCeiling"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaCeiling, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaFloor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaFloor"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaFloor, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaInv"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaInv"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaInv, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaMod"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaMod"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaMod, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaModFix"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaModFix"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaModFix, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaRepl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaRepl"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.AlphaRepl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "biLevel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "biLevel"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.BiLevel, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blur"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blur"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Blur, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrChange"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "clrChange"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.ClrChange, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrRepl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "clrRepl"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.ClrRepl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "duotone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "duotone"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Duotone, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillOverlay"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillOverlay"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.FillOverlay, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "grayscl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "grayscl"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Grayscl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hsl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "hsl"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Hsl, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lum"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Lum, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tint"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tint"}:
				tmp := NewCT_BlipChoice()
				if err := d.DecodeElement(&tmp.Tint, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Blip %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Blip
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Blip and its children
func (m *CT_Blip) Validate() error {
	return m.ValidateWithPath("CT_Blip")
}

// ValidateWithPath validates the CT_Blip and its children, prefixing error messages with path
func (m *CT_Blip) ValidateWithPath(path string) error {
	if err := m.CstateAttr.ValidateWithPath(path + "/CstateAttr"); err != nil {
		return err
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
