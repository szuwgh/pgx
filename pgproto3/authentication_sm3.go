package pgproto3

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5/internal/pgio"
)

// AuthenticationMD5Password is a message sent from the backend indicating that an MD5 hashed password is required.
type AuthenticationSM3 struct {
	Salt [4]byte
}

// Backend identifies this message as sendable by the PostgreSQL backend.
func (*AuthenticationSM3) Backend() {}

// Backend identifies this message as an authentication response.
func (*AuthenticationSM3) AuthenticationResponse() {}

// Decode decodes src into dst. src must contain the complete message with the exception of the initial 1 byte message
// type identifier and 4 byte message length.
func (dst *AuthenticationSM3) Decode(src []byte) error {
	if len(src) != 8 {
		return errors.New("bad authentication message size")
	}

	authType := binary.BigEndian.Uint32(src)

	if authType != AuthTypeSM3 {
		return errors.New("bad auth type")
	}

	copy(dst.Salt[:], src[4:8])
	//panic("AuthenticationSM3 Decode is not implemented yet")

	return nil
}

// Encode encodes src into dst. dst will include the 1 byte message type identifier and the 4 byte message length.
func (src *AuthenticationSM3) Encode(dst []byte) ([]byte, error) {
	//panic("AuthenticationSM3 Encode is not implemented yet")
	dst, sp := beginMessage(dst, 'R')
	dst = pgio.AppendUint32(dst, AuthTypeSM3)
	dst = append(dst, src.Salt[:]...)
	return finishMessage(dst, sp)
}

// MarshalJSON implements encoding/json.Marshaler.
func (src AuthenticationSM3) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string
		Salt [4]byte
	}{
		Type: "AuthenticationSM3",
		Salt: src.Salt,
	})
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
func (dst *AuthenticationSM3) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	var msg struct {
		Type string
		Salt [4]byte
	}
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}

	dst.Salt = msg.Salt
	return nil
}
