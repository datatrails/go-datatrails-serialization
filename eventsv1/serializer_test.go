package eventsv1

import (
	"testing"

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
func TestSerializeEventFromProtoConsistency(t *testing.T) {
	type event struct {
		attributes map[string]any
		trails     []string
		event_type string
	}
	tests := []struct {
		name              string
		event1            event
		event2            event
		sameSerialization bool
		hasEventType      bool
	}{
		{
			name: "same attributes and trails in same order",
			event1: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			sameSerialization: true,
		},
		{
			name: "same attributes and trails in same order and event type",
			event1: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
				event_type: "receipt",
			},
			event2: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
				event_type: "receipt",
			},
			sameSerialization: true,
			hasEventType:      true,
		},
		{
			name: "same attributes and trails, attributes in different order",
			event1: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]any{
					"vanilla extract": "1 tsp",
					"flour":           "500g",
					"milk":            "300ml",
					"eggs":            "2",
					"sugar":           "250g",
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
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
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
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]any{
					"flour":               "500g",
					"sugar":               "250g",
					"eggs":                "2",
					"milk":                "300ml",
					"vanilla extract":     "1 tsp",
					"bicarbonate of soda": "2 tsp",
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
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
				},
				trails: []string{
					"cake",
					"viccy sponge",
				},
			},
			event2: event{
				attributes: map[string]any{
					"flour":           "500g",
					"sugar":           "250g",
					"eggs":            "2",
					"milk":            "300ml",
					"vanilla extract": "1 tsp",
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

			serializableEvent1 := SerializableEvent{
				Attributes: test.event1.attributes,
				Trails:     test.event1.trails,
				EventType:  test.event1.event_type,
			}

			serializableEvent2 := SerializableEvent{
				Attributes: test.event2.attributes,
				Trails:     test.event2.trails,
				EventType:  test.event2.event_type,
			}

			serializedEvent1, err := serializableEvent1.Serialize()
			require.Nil(t, err)

			serializedEvent2, err := serializableEvent2.Serialize()
			require.Nil(t, err)

			// check that two events serilaize to the same value
			// or not depending on test data
			if test.sameSerialization {
				assert.Equal(t, serializedEvent1, serializedEvent2)
			} else {
				assert.NotEqual(t, serializedEvent1, serializedEvent2)
			}

			// check that event_type got serialized or ommitted if not present
			// in the initila data
			if test.hasEventType {
				assert.Contains(t, string(serializedEvent1), "event_type")
				assert.Contains(t, string(serializedEvent2), "event_type")
			} else {
				assert.NotContains(t, string(serializedEvent1), "event_type")
				assert.NotContains(t, string(serializedEvent2), "event_type")
			}

		})
	}
}

// TestSerializableEvent_Serialize tests:
//
// 1. an event with all types of attributes [string|list|dict] can be serialized without error.
func TestSerializableEvent_Serialize(t *testing.T) {
	type args struct {
		attributes map[string]any
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
				attributes: map[string]any{
					"flour": "500g",
					"method": []map[string]string{
						{"1": "put flour sugar into mixing bowl"},
						{"2": "put in eggs and mix"},
						{"3": "put in milk and mix"},
					},
					"baking time": map[string]string{
						"oven time": "30 mins",
					},
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

			serializableEvent := SerializableEvent{
				Attributes: test.args.attributes,
				Trails:     test.args.trails,
			}

			actual, err := serializableEvent.Serialize()

			require.NotNil(t, actual)
			assert.Equal(t, test.err, err)

		})
	}
}
