package eventsv1

import (
	"testing"

	"github.com/datatrails/go-datatrails-common-api-gen/attribute/v2/attribute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSerializeEventConsistency tests:
//
// 1. an event with the same attributes and trails in the same order, serializes to the same bytes.
// 2. an event with the same trails and same attributes in a different order, serializes to the same bytes.
// 3. an event with the same attributes and same trails in a different order, serializes to different bytes.
// 4. an event with the same trails and different attributes serializes to different bytes.
// 5. an event with the same attributes and different trails, serializes to different bytes.
func TestSerializeEventConsistency(t *testing.T) {
	type event struct {
		attributes map[string]*attribute.Attribute
		trails     []string
	}
	tests := []struct {
		name              string
		event1            event
		event2            event
		sameSerialization bool
	}{
		{
			name: "same attributes and trails in same order",
			event1: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			sameSerialization: true,
		},
		{
			name: "same attributes and trails, attributes in different order",
			event1: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]*attribute.Attribute{
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
					"flour":           attribute.NewStringAttribute("500g"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"eggs":            attribute.NewStringAttribute("2"),
					"sugar":           attribute.NewStringAttribute("250g"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			sameSerialization: true,
		},
		{
			name: "same attributes and trails, trails in different order",
			event1: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"viccy sponge",
					"cake",
				},
			},
			sameSerialization: false,
		},
		{
			name: "same trails and different attributes in same order",
			event1: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]*attribute.Attribute{
					"flour":               attribute.NewStringAttribute("500g"),
					"sugar":               attribute.NewStringAttribute("250g"),
					"eggs":                attribute.NewStringAttribute("2"),
					"milk":                attribute.NewStringAttribute("300ml"),
					"vanilla extract":     attribute.NewStringAttribute("1 tsp"),
					"bicarbonate of soda": attribute.NewStringAttribute("2 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			sameSerialization: false,
		},
		{
			name: "same attributes and different trails in same order",
			event1: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]*attribute.Attribute{
					"flour":           attribute.NewStringAttribute("500g"),
					"sugar":           attribute.NewStringAttribute("250g"),
					"eggs":            attribute.NewStringAttribute("2"),
					"milk":            attribute.NewStringAttribute("300ml"),
					"vanilla extract": attribute.NewStringAttribute("1 tsp"),
				},
				trails: []string{
					"cake",
					"dessert",
				},
			},
			sameSerialization: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			serializedEvent1, err := SerializeEvent(test.event1.attributes, test.event1.trails)
			require.Nil(t, err)

			serializedEvent2, err := SerializeEvent(test.event2.attributes, test.event2.trails)
			require.Nil(t, err)

			if test.sameSerialization {
				assert.Equal(t, serializedEvent1, serializedEvent2)
			} else {
				assert.NotEqual(t, serializedEvent1, serializedEvent2)
			}

		})
	}
}

// TestSerializeEvent tests:
//
// 1. an event with all types of attributes [string|list|dict] can be serialized without error.
func TestSerializeEvent(t *testing.T) {
	type args struct {
		attributes map[string]*attribute.Attribute
		trails     []string
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "all attribute types no error",
			args: args{
				attributes: map[string]*attribute.Attribute{
					"flour": attribute.NewStringAttribute("500g"),
					"method": attribute.NewListAttribute([]map[string]string{
						{"1": "put flour sugar into mixing bowl"},
						{"2": "put in eggs and mix"},
						{"3": "put in milk and mix"},
					}),
					"baking time": attribute.NewDictAttribute(map[string]string{
						"oven time": "30 mins",
					}),
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			err: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actual, err := SerializeEvent(test.args.attributes, test.args.trails)
			require.NotNil(t, actual)

			assert.Equal(t, test.err, err)

		})
	}
}
