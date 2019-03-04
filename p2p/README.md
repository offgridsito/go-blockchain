**P2P Blockchain**

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

**Screenshots**
![screen shot 2019-03-04 at 10 47 14 am](https://user-images.githubusercontent.com/17554983/53748521-18341500-3e6b-11e9-81d6-cbec14dd9212.png)
![screen shot 2019-03-04 at 10 47 47 am](https://user-images.githubusercontent.com/17554983/53748540-208c5000-3e6b-11e9-94b3-722ff610eb6b.png)
![screen shot 2019-03-04 at 10 48 00 am](https://user-images.githubusercontent.com/17554983/53748557-2aae4e80-3e6b-11e9-88ab-adc7ba135440.png)
![screen shot 2019-03-04 at 10 48 10 am](https://user-images.githubusercontent.com/17554983/53748575-34d04d00-3e6b-11e9-9662-a5b01470305a.png)
![screen shot 2019-03-04 at 10 48 35 am](https://user-images.githubusercontent.com/17554983/53748589-3b5ec480-3e6b-11e9-89c8-525fb92e921e.png)
