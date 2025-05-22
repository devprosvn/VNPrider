// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package ledger

import "encoding/json"

// TxInput references a previous transaction output.
type TxInput struct {
	From   string `json:"from"`
	Amount uint64 `json:"amount"`
}

// TxOutput defines a recipient and amount.
type TxOutput struct {
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

// Transaction groups inputs and outputs.
type Transaction struct {
	Inputs  []TxInput  `json:"inputs"`
	Outputs []TxOutput `json:"outputs"`
	Nonce   uint64     `json:"nonce"`
}

// Serialize returns the JSON representation of the transaction.
func (tx *Transaction) Serialize() ([]byte, error) {
	return json.Marshal(tx)
}
