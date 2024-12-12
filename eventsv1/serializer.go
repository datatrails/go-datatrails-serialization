package eventsv1

import (
	"encoding/json"

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
	Attributes map[string]any `json:"attributes"`
	Trails     []string       `json:"trails"`
}

// SerializeEventFromJson serializes a v1 event from datatrails eventv1 api respone
func SerializeEventFromJson(eventJson []byte) ([]byte, error) {

	serializableEvent := SerializableEvent{}

	// json marshal the attributes and trails
	err := json.Unmarshal(eventJson, &serializableEvent)
	if err != nil {
		return nil, err
	}

	return serializableEvent.Serialize()

}

// Serialize the event
func (se *SerializableEvent) Serialize() ([]byte, error) {

	// json marshal the attributes and trails
	jsonSerializedEvent, err := json.Marshal(&se)
	if err != nil {
		return nil, err
	}

	// bencode the json marshaled attributes and trails
	//
	// NOTE: this gives a consistent ordering to the attributes that json doesn't give.
	bencodeSerializedEvent, err := bencode.EncodeBytes(jsonSerializedEvent)
	if err != nil {
		return nil, err
	}

	return bencodeSerializedEvent, nil

}
