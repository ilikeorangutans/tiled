package tmx

import (
	"encoding/xml"
)

// TMX map structure
type Map struct {
	XMLName     xml.Name `xml:"map"`
	Version     string   `xml:"version,attr"`
	Orientation string   `xml:"orientation,attr"`
	Width       int      `xml:"width,attr"`
	Height      int      `xml:"height,attr"`
	Tilewidth   int      `xml:"tilewidth,attr"`
	Tileheight  int      `xml:"tileheight,attr"`

	Tileset []Tileset `xml:"tileset"`
	Layer   []Layer   `xml:"layer"`
}

type Tileset struct {
	XMLName    xml.Name `xml:"tileset"`
	Firstgid   int      `xml:"firstgid,attr"`
	Source     string   `xml:"source,attr"`
	Name       string   `xml:"name,attr"`
	Tilewidth  int      `xml:"tilewidth,attr"`
	Tileheight int      `xml:"tileheight,attr"`
	Spacing    int      `xml:"spacing,attr"`
	Margin     int      `xml:"margin,attr"`
	Image      Image    `xml:"image"`
}

// Describes an image in a tileset.
type Image struct {
	XMLName xml.Name `xml:"image"`
	Format  string   `xml:"format,attr"` // Format of the embedded image
	Source  string   `xml:"source,attr"` // Source path
	Trans   string   `xml:"trans,attr"`  // Transparent color for this image
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

type Layer struct {
	XMLName xml.Name  `xml:"layer"`
	Name    string    `xml:"name,attr"`
	Data    LayerData `xml:"data"`
}

type LayerData struct {
	XMLName     xml.Name `xml:"data"`
	Encoding    string   `xml:"encoding,attr"`    // Encoding of the data, optional. Can be base64 or csv
	Compression string   `xml:"compression,attr"` // Optional, compression algorithm used for the contents of the data element. Can be gzip or zlib.
	Tiles       []Tile   `xml:"tile"`             // The actual tiles, only available if data is stored in XML.
	RawData     []byte   `xml:",innerxml"`        // Holds the raw contents of data, used to decode different encodings.
}

type Tile struct {
	XMLName xml.Name `xml:"tile"`
	Gid     int      `xml:"gid,attr"`
}
