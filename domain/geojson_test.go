package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeatureCollection_Validate(t *testing.T) {
	var featureCollection FeatureCollection

	t.Run("empty feature collection", func(t *testing.T) {
		err := featureCollection.Validate()
		assert.Equal(t, ErrNoFeatures, err)
	})

	featureCollection.Type = FeatureCollectionString
	featureCollection.Features = append(
		featureCollection.Features,
		&Feature{
			ID:         1,
			Type:       FeatureString,
			Properties: map[string]interface{}{},
		},
	)

	t.Run("featureCollection - valid", func(t *testing.T) {
		err := featureCollection.Validate()
		assert.NoError(t, err)
	})
}

func TestFeatureCollection_String(t *testing.T) {
	featureCollection := FeatureCollection{}

	t.Run("feature collection string - empty", func(t *testing.T) {
		str, err := featureCollection.String()
		assert.NoError(t, err)
		assert.Equal(t, "{\"type\":\"\",\"features\":null}", str)

	})

}
