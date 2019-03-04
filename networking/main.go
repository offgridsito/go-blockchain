package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

// keeps block headers
type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
}

// keeps sequence of blocks
var Blockchain []Block

// handles incoming concurrent Blocks
var bcServer chan []Block
var mutex = &sync.Mutex{}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	//create genesis block
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	httpPort := os.Getenv("PORT")

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}
	log.Printf("HTTP Server Listening on port : %v", httpPort)
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Enter a new BPM:")

	scanner := bufio.NewScanner(conn)

	//take in BPM from stdin and add it to blockchain after validating

	go func(){
		for scanner.Scan(){
			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", err)
				continue
			}
			newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], bpm)
			if err != nil {
				log.Println(err)
				continue
			}
			if confirmValidBlock(newBlock, Blockchain[len(Blockchain) -1]) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "\nEnter a new BPM:")
		}
	}()

	//simulate receiving broadcast

	go func(){
		for {
			time.Sleep(30 * time.Second)
			mutex.Lock()
			output, err := json.Marshal(Blockchain)

			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			io.WriteString(conn, string(output))
		}
	}()

	for range bcServer {
		spew.Dump(Blockchain)
	}
}

//check if block is valid by checking index, and comparing has of prev block
func confirmValidBlock(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}

	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

//check if chain we are tracking is longer than current blockchain
func replaceChain(newBlocks []Block){
	mutex.Lock()
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}

	mutex.Unlock()
}

//SHA256 hashing
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//create a new block
func generateBlock(prevBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
