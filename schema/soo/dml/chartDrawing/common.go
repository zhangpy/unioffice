// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import "github.com/zhangpy/unioffice"

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ShapeNonVisual", NewCT_ShapeNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Shape", NewCT_Shape)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ConnectorNonVisual", NewCT_ConnectorNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Connector", NewCT_Connector)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Picture", NewCT_Picture)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrameNonVisual", NewCT_GraphicFrameNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrame", NewCT_GraphicFrame)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShapeNonVisual", NewCT_GroupShapeNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShape", NewCT_GroupShape)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Marker", NewCT_Marker)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_RelSizeAnchor", NewCT_RelSizeAnchor)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_AbsSizeAnchor", NewCT_AbsSizeAnchor)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Drawing", NewCT_Drawing)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_ObjectChoices", NewEG_ObjectChoices)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_Anchor", NewEG_Anchor)
}
