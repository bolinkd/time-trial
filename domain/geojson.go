package domain

import (
	"github.com/gin-gonic/gin/json"
)

type GeometryType string
type FeatureCollectionType string
type FeatureType string

// The geometry types supported by GeoJSON 1.0
const (
	GeometryPointAndRadius GeometryType = "PointAndRadius"
	GeometryMultiPolygon   GeometryType = "MultiPolygon"
)

const (
	FeatureString FeatureType = "Feature"
)

const (
	FeatureCollectionString FeatureCollectionType = "FeatureCollection"
)

type FeatureCollection struct {
	Type        FeatureCollectionType `json:"type"`
	BoundingBox []float64             `json:"bbox,omitempty"`
	Features    []*Feature            `json:"features"`
}

func (collection FeatureCollection) Validate() error {
	if collection.Features == nil || len(collection.Features) == 0 {
		return ErrNoFeatures
	}
	//note: all properties checks are done at a higher level
	return nil
}

type Feature struct {
	ID          interface{}            `json:"id,omitempty"`
	Type        FeatureType            `json:"type"`
	BoundingBox []float64              `json:"bbox,omitempty"`
	Geometry    *Geometry              `json:"geometry,omitempty"`
	Properties  map[string]interface{} `json:"properties"`
}

type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Geometry struct {
	Type            GeometryType           `json:"type"`
	BoundingBox     []float64              `json:"bbox,omitempty"`
	Point           Point                  `json:"point,omitempty"`
	Radius          float64                `json:"radius,omitempty"`
	MultiPoint      [][]float64            `json:"-"`
	LineString      [][]float64            `json:"-"`
	MultiLineString [][][]float64          `json:"-"`
	Polygon         [][][]float64          `json:"-"`
	MultiPolygon    [][][][]float64        `json:"-"`
	Geometries      []*Geometry            `json:"-"`
	CRS             map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

func (collection FeatureCollection) String() (string, error) {
	bytes, err := json.Marshal(collection)
	return string(bytes), err
}
