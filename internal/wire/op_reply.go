// Copyright 2021 Baltoro OÜ.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wire

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"

	"github.com/MangoDB-io/MangoDB/internal/bson"
	"github.com/MangoDB-io/MangoDB/internal/types"
	"github.com/MangoDB-io/MangoDB/internal/util/lazyerrors"
)

const maxNumberReturned = 1000

type OpReply struct {
	ResponseFlags  OpReplyFlags
	CursorID       int64
	StartingFrom   int32
	NumberReturned int32
	Documents      []types.Document
}

func (reply *OpReply) msgbody() {}

func (reply *OpReply) readFrom(bufr *bufio.Reader) error {
	if err := binary.Read(bufr, binary.LittleEndian, &reply.ResponseFlags); err != nil {
		return lazyerrors.Errorf("wire.OpReply.ReadFrom (binary.Read): %w", err)
	}
	if err := binary.Read(bufr, binary.LittleEndian, &reply.CursorID); err != nil {
		return lazyerrors.Errorf("wire.OpReply.ReadFrom (binary.Read): %w", err)
	}
	if err := binary.Read(bufr, binary.LittleEndian, &reply.StartingFrom); err != nil {
		return lazyerrors.Errorf("wire.OpReply.ReadFrom (binary.Read): %w", err)
	}
	if err := binary.Read(bufr, binary.LittleEndian, &reply.NumberReturned); err != nil {
		return lazyerrors.Errorf("wire.OpReply.ReadFrom (binary.Read): %w", err)
	}

	if n := reply.NumberReturned; n < 0 || n > maxNumberReturned {
		return lazyerrors.Errorf("wire.OpReply.ReadFrom: invalid NumberReturned %d", n)
	}

	reply.Documents = make([]types.Document, reply.NumberReturned)
	for i := int32(0); i < reply.NumberReturned; i++ {
		var doc bson.Document
		if err := doc.ReadFrom(bufr); err != nil {
			return lazyerrors.Errorf("wire.OpReply.ReadFrom: %w", err)
		}
		reply.Documents[i] = types.MustConvertDocument(&doc)
	}

	return nil
}

func (reply *OpReply) UnmarshalBinary(b []byte) error {
	br := bytes.NewReader(b)
	bufr := bufio.NewReader(br)

	if err := reply.readFrom(bufr); err != nil {
		return lazyerrors.Errorf("wire.OpReply.UnmarshalBinary: %w", err)
	}

	if _, err := bufr.Peek(1); err != io.EOF {
		return lazyerrors.Errorf("unexpected end of the OpReply: %v", err)
	}

	return nil
}

func (reply *OpReply) MarshalBinary() ([]byte, error) {
	if l := len(reply.Documents); int32(l) != reply.NumberReturned {
		return nil, lazyerrors.Errorf("wire.OpReply.MarshalBinary: len(Documents)=%d, NumberReturned=%d", l, reply.NumberReturned)
	}

	var buf bytes.Buffer
	bufw := bufio.NewWriter(&buf)

	if err := binary.Write(bufw, binary.LittleEndian, reply.ResponseFlags); err != nil {
		return nil, lazyerrors.Errorf("wire.OpReply.MarshalBinary (binary.Write): %w", err)
	}
	if err := binary.Write(bufw, binary.LittleEndian, reply.CursorID); err != nil {
		return nil, lazyerrors.Errorf("wire.OpReply.MarshalBinary (binary.Write): %w", err)
	}
	if err := binary.Write(bufw, binary.LittleEndian, reply.StartingFrom); err != nil {
		return nil, lazyerrors.Errorf("wire.OpReply.MarshalBinary (binary.Write): %w", err)
	}
	if err := binary.Write(bufw, binary.LittleEndian, reply.NumberReturned); err != nil {
		return nil, lazyerrors.Errorf("wire.OpReply.UnmarshalBinary (binary.Write): %w", err)
	}

	for _, doc := range reply.Documents {
		if err := bson.MustConvertDocument(doc).WriteTo(bufw); err != nil {
			return nil, lazyerrors.Errorf("wire.OpReply.MarshalBinary: %w", err)
		}
	}

	if err := bufw.Flush(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (reply *OpReply) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ResponseFlags":  reply.ResponseFlags,
		"CursorID":       reply.CursorID,
		"StartingFrom":   reply.StartingFrom,
		"NumberReturned": reply.NumberReturned,
	}

	docs := make([]interface{}, len(reply.Documents))
	for i, d := range reply.Documents {
		docs[i] = bson.MustConvertDocument(d)
	}

	m["Documents"] = docs

	return json.Marshal(m)
}

// check interfaces
var (
	_ MsgBody = (*OpReply)(nil)
)
