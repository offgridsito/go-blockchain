#P2P Blockchain

**Getting Started**

Youll need to set up `gx` package manager. 
Assuming you have go installed:

- `go get -d github.com/libp2p/go-libp2p/...`
- navigate to cloned directory from above
- `make`
- `make deps`


Once thats set up. In 1st terminal `go run main.go -secio -l 10000`

Follow instructions in 1st terminal and copy paste the given comman into 2nd terminal e.g. 
`go run main.go -l 10001 -d /ip4/127.0.0.1/tcp/10000/ipfs/QmZ8NayvdXc2U2A1cwh9qGaHK7uxXXVrZQEYwDqbfFydfj -secio`

Follow the instructions in your 2nd terminal and copy and paste the given command into your 3rd terminal e.g. 
`go run main.go -l 10002 -d /ip4/127.0.0.1/tcp/10001/ipfs/QmRAj9JJVKRJmWHbDKzvzKDVVFPWxuWYio3bPym4SgGPgF -secio`

Type a BPM into any of your terminals and watch blockchain be broadcast to all terminals
