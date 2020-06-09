package internal

import (
	"fmt"

	aero "github.com/aerospike/aerospike-client-go"
)

type DummyAero struct {
	client *aero.Client
}

func NewDummyAero(host string, port int) (*DummyAero, error) {
	client, err := connectToAero(host, port)
	if err != nil {
		return nil, err
	}

	return &DummyAero{
		client: client,
	}, nil
}

func connectToAero(hostname string, port int) (*aero.Client, error) {
	// define a client to connect to
	client, err := aero.NewClient(hostname, port)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (do *DummyAero) TestConnection() error {
	key, err := aero.NewKey("test", "aerospike", "key")
	if err != nil {
		return err
	}

	// define some bins with data
	bins := aero.BinMap{
		"bin1": 42,
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 2009},
	}

	// write the bins
	err = do.client.Put(nil, key, bins)
	if err != nil {
		return err
	}

	// read it back!
	rec, err := do.client.Get(nil, key)
	if err != nil {
		return err
	}
	if rec == nil {
		return fmt.Errorf("rec for %s is nil", key)
	}

	// delete the key, and check if key exists
	existed, err := do.client.Delete(nil, key)
	if err != nil {
		return err
	}
	fmt.Printf("Record existed before delete? %v\n", existed)

	return nil
}
