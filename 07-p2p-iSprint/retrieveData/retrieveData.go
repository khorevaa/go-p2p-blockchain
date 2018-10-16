package main

import (
	"log"
	"os"
	"encoding/gob"
	"runtime"
	"github.com/davecgh/go-spew/spew"

)

///// GLOBAL VARIABLES

// Raw Material Transaction (Type 1)
type RawMaterialTransaction struct {
	SerialNo int
	ProductCode string
	ProductName string
	ProductBatchNo string
	Quantity int
	RawMaterial []RawMaterial
}
type RawMaterial struct {
	RawMaterialBatchNo string
	RawMaterialsID string
	RawMaterialName string
	RawMaterialQuantity float32
	RawMaterialMeasurementUnit string
}

// Delivery Transaction (Type 2)
type DeliveryTransaction struct {
	SerialNo int
	RecGenerator string
	ShipmentID string
	Timestamp string	
	Longitude string
	Latitude string
	ShippedFromCompID string
	ShippedToCompID string
	LocationID string
	DeliveryStatus string 	
	DeliveryType string
	Product []Product
	Document Document
}
type Product struct {
	ProductCode string
	ProductName string
	ProductBatch []ProductBatch
}
type ProductBatch struct {
	ProductBatchNo string
	ProductBatchQuantity int
}
type Document struct {
	DocumentURL string
	DocumentType string
	DocumentHash string
	DocumentSign string
}

// Block represents each 'item' in the blockchain
type Block struct {
	Index		int
	Timestamp	string
	TxnType		int
	TxnPayload	interface{}
	Comment		string
	Proposer	string
	PrevHash 	string
	ThisHash	string
}

// Standard Input takes incoming string from terminal for `comment` in block (Type 0)
type StdInput struct {
	Comment string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

//////// FUNCTIONS

func init() { // Idea from https://appliedgo.net/networking/
	log.SetFlags(log.Lshortfile)
	gob.Register(RawMaterialTransaction{})
	gob.Register(DeliveryTransaction{})
	gob.Register(RawMaterial{})
	gob.Register(Product{})
	gob.Register(ProductBatch{})
	gob.Register(map[string]interface{}{})
}

func main() {
	dataFile := "../data5000/blockchain-4.gob"
	log.Println("Loading Blockchain from", dataFile)
	gobCheck(readGob(&Blockchain, dataFile))
	log.Println(Blockchain)
	spew.Dump(Blockchain)
}

func readGob(object interface{}, filePath string) error {
       file, err := os.Open(filePath)
       if err == nil {
              decoder := gob.NewDecoder(file)
              err = decoder.Decode(object)
       }
       file.Close()
       return err
}

func gobCheck(e error) { // Inspired from http://www.robotamer.com/code/go/gotamer/gob.html
    if e != nil {
        _, file, line, _ := runtime.Caller(1)
        log.Println(line, "\t", file, "\n", e)
        os.Exit(1)
    }
}