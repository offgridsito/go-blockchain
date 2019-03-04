Code for the **Twitter Blockchain**

Can take the tweets returned from root `main.go` file and run it through **SHA256**
converted online, save returned hash

- Run `go run main.go` 
- Send **POST** request with **Postman** to `localhost:(.env port)` with body `{"TweetHash": (returned hash from SHA256 converted)}`



