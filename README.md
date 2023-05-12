# Web Dev Intro with Go

## Section 1
- [Go standard library](https://pkg.go.dev/std)
- [Go doc - a great place to find open source packages](https://pkg.go.dev/)
- [jsDeliver - a helpful content deliver network which hosts a lot of open source javascript and CSS packages](https://www.jsdelivr.com/)
- [CDN - another great CDN with externally hosted javascript and CSS resources](https://cdnjs.com/)
- [The Mozilla Dev Network's documentation for Javascript](https://pkg.go.dev/](https://developer.mozilla.org/en-US/docs/Web/javascript)
- [The Mozilla Dev Network's documentation for CSS](https://pkg.go.dev/](https://developer.mozilla.org/en-US/docs/Web/javascript](https://developer.mozilla.org/en-US/docs/Web/CSS)

## Section 2

Create a new module:
`` go mod init github.com/username/repo ``

Install a package:
`` go get github.com/username/repo ``

Run a go file:
`` go run main.go ``

Build a go file:
`` go build main.go ``

Build a go file and run it:
`` go build main.go && ./main ``

Run all tests in the current directory:
`` go test ``

Run all tests in the current directory and subdirectories:
`` go test ./... ``

Coverage in CLI:
`` go test -cover ``

Run all tests in the current directory and subdirectories with coverage:
`` go test ./... -cover ``

Run tests and get detailed results in CLI:
`` go test -v ``

Coverage in html format:
`` go test -coverprofile=coverage.out && go tool cover -html=coverage.out ``
