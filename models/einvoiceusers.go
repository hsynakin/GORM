package models

import (
	"encoding/xml"
	"time"
)

// EInvoiceUsers ...
// xml files geting values struct
type EInvoiceUsers struct {
	XMLName           xml.Name  `json:"-"`
	Identifier        string    `json:"identifier" gorm:"primary_key"`
	Alias             string    `json:"alias; not null"`
	Title             string    `json:"title; not null"`
	Type              string    `json:"type"`
	FirstCreationTime time.Time `json:"firstCreationTime"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}

// Users ...
// Users xml struct
type Users struct {
	XMLName xml.Name        `xml:"UserList"`
	Users   []EInvoiceUsers `xml:"User"`
}
