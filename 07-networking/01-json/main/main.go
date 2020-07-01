package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "Some name",
	"type": "deposit",
	"amount": 1000000.3
}
`

type Request struct {
	Login  string  `01-json:"user"`
	Type   string  `01-json:"type"`
	Amount float64 `01-json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data) // Simulation of a file/socket

	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", req)

	prevBalance := 8500000.0 // balance from "database"
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
}
