### Initialization:

1. We initialize project with the command:

```sh
go mod init github.com/prashant1k99/Go-RSS-Scraper
```

Here usually the Github repo URL is added.

2. Then add `main.go` file which will be the entrypoint of the application.

3. Use command `go build` to build a binary for the project.

4. Add `.env` file to add Env Variables.

5. In order to import env variables in Go project, we can use the project: `github.com/joho/godotenv`

```sh
go get github.com/joho/godotenv
```

6. Once the project is installed run:

```sh
go mod vendor
```

So that it can be accessed in our project and output binary.

7. Then in the `mian.go` use it like following:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Server is running on port: ", PORT)
}
```
