package model

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type BinaryUUID uuid.UUID

func (b BinaryUUID) String() string {
	return uuid.UUID(b).String()
}

// StringToBinaryUUID -> parse string to BinaryUUID
func StringToBinaryUUID(s string) (BinaryUUID, error) {
	id, err := uuid.Parse(s)
	return BinaryUUID(id), err
}

func (b BinaryUUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(b)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (b *BinaryUUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*b = BinaryUUID(s)
	return err
}

func (b *BinaryUUID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	data, err := uuid.FromBytes(bytes)
	*b = BinaryUUID(data)
	return err
}

func (b BinaryUUID) Value() (driver.Value, error) {
	return uuid.UUID(b).MarshalBinary()
}
