package eventsv1

import (
	"encoding/json"

	"github.com/datatrails/go-datatrails-common-api-gen/attribute/v2/attribute"
	"github.com/zeebo/bencode"
)

/**
 * serializes datatatrails v1 events.
 *
 * Event v1 serialization is the following:
 *
 *  1. take the attributes and trails fields from the v1 event
 *  2. json marshal the attributes and trails
 *  3. bencode the json marshaled attributes and trails
 */

// SerializableEvent is the event
// containing only fields that will be serialized.
type SerializableEvent struct {
	Attributes map[string]*attribute.Attribute `json:"attributes"`
	Trails     []string                        `json:"trails"`
}

// SerializeEvent serializes a v1 event
func SerializeEvent(attributes map[string]*attribute.Attribute, trails []string) ([]byte, error) {

	// 1. take attributes and trails from v1 event
	serializableEvent := SerializableEvent{
		Attributes: attributes,
		Trails:     trails,
	}

	// 2. json marshal the attributes and trails
	jsonSerializedEvent, err := json.Marshal(&serializableEvent)
	if err != nil {
		return nil, err
	}

	// 3. bencode the json marshaled attributes and trails
	//
	// NOTE: this gives a consistent ordering to the attributes that json doesn't give.
	bencodeSerializedEvent, err := bencode.EncodeBytes(jsonSerializedEvent)
	if err != nil {
		return nil, err
	}

	return bencodeSerializedEvent, nil
}
