package types

import (
	"bytes"
	"encoding/binary"
)

type TxInput struct {
	TxID        []byte
	OutputIndex uint32
	Signature   []byte
	PubKey      []byte
}

type TxOutput struct {
	Value        uint64
	PubKeyHash   []byte
	ScriptPubKey []byte
}

type Transaction struct {
	Version   uint32
	Timestamp int64
	Inputs    []TxInput
	Outputs   []TxOutput
	Payload   []byte
}

// Serialize encodes the transaction deterministically.
func (tx *Transaction) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, tx.Version)
	binary.Write(buf, binary.BigEndian, tx.Timestamp)
	binary.Write(buf, binary.BigEndian, uint32(len(tx.Inputs)))
	for _, in := range tx.Inputs {
		binary.Write(buf, binary.BigEndian, uint32(len(in.TxID)))
		buf.Write(in.TxID)
		binary.Write(buf, binary.BigEndian, in.OutputIndex)
		binary.Write(buf, binary.BigEndian, uint32(len(in.Signature)))
		buf.Write(in.Signature)
		binary.Write(buf, binary.BigEndian, uint32(len(in.PubKey)))
		buf.Write(in.PubKey)
	}
	binary.Write(buf, binary.BigEndian, uint32(len(tx.Outputs)))
	for _, out := range tx.Outputs {
		binary.Write(buf, binary.BigEndian, out.Value)
		binary.Write(buf, binary.BigEndian, uint32(len(out.PubKeyHash)))
		buf.Write(out.PubKeyHash)
		binary.Write(buf, binary.BigEndian, uint32(len(out.ScriptPubKey)))
		buf.Write(out.ScriptPubKey)
	}
	binary.Write(buf, binary.BigEndian, uint32(len(tx.Payload)))
	buf.Write(tx.Payload)
	return buf.Bytes()
}
