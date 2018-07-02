# githuuk
A GitHub webhook receiver written in Go.

This project was originally a fork of [phayes/hookserve](https://github.com/phayes/hookserve), but has been rewritten nearly completely.

```go
import "maunium.net/go/githuuk"

func main() {
	server := githuuk.NewServer()
	server.Port = 8888
	server.Secret = "GitHub webhook secret"
	server.AsyncListenAndServe()

	for rawEvent := range server.Events {
		switch rawEvent.GetType() {
		case githuuk.EventPush:
			evt := rawEvent.(*githuuk.PushEvent)
			fmt.Println(evt.Repository.Owner.Name, evt.Repository.Name, evt.Ref.Name(), evt.HeadCommit.ID)
		}
	}
}
```
