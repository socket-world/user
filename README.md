# socket.world user
Welcome!  Here lies the `user` service for socket.world -- this service is responsible for maintaining all User resources in the system.

## Development Guide
This repo uses go modules -- when adding new dependencies, `go.mod` and `go.sum` should be updated.

To do so, run `go mod tidy`, then stage the commit with `git add go.mod && git add go.sum`.

### Docker Build and Run
`docker build -t socketworld/user . && docker run -p 80:8080/tcp socketworld/user`

### Docker Build and Push
`docker build -t socketworld/user . && docker push socketworld/user`
