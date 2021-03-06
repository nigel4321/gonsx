package ipset

import "encoding/xml"

// List - top level <list> element
type List struct {
	XMLName xml.Name `xml:"list"`
	IPSets  []IPSet  `xml:"ipset"`
}

// IPSets top level struct
type IPSets struct {
	IPSets []IPSet `xml:"ipset"`
}

// IPSet object struct
type IPSet struct {
	XMLName            xml.Name `xml:"ipset"`
	ObjectID           string   `xml:"objectId,omitempty"`
	ObjectTypeName     string   `xml:"objectTypeName,omitempty"`
	VsmUUID            string   `xml:"vsmUuid,omitempty"`
	NodeID             string   `xml:"nodeId,omitempty"`
	Revision           int      `xml:"revision,omitempty"`
	TypeName           string   `xml:"type,omitempty>typeName,omitempty"`
	Name               string   `xml:"name,omitempty"`
	Description        string   `xml:"description,omitempty"`
	IsUniversal        bool     `xml:"isUniversal,omitempty"`
	InheritanceAllowed bool     `xml:"inheritanceAllowed,omitempty"`
	Value              string   `xml:"value,omitempty"`
}
