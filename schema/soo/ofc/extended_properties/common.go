// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package extended_properties

import "github.com/zhangpy/unioffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_Properties", NewCT_Properties)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorVariant", NewCT_VectorVariant)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorLpstr", NewCT_VectorLpstr)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_DigSigBlob", NewCT_DigSigBlob)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "Properties", NewProperties)
}
