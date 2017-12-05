package main

import (
	"fmt"
	"bytes"
	"log"
	"encoding/gob"
)

// TXOutput represents a transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

type TXOutputs struct {
	Outputs []TXOutput
}

// Lock signs the output
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	fmt.Printf("publicHash size: %d \n", len(pubKeyHash))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	//pubKeyHash = pubKeyHash[1:30]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}


//Seralize
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}


//DeserializeOutputs
func DeserializeOutputs(data[] byte) TXOutputs{
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}
	return outputs
}

