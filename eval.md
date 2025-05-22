Project Path: /home/idk/go/src/github.com/go-i2p/go-meta-listener

Source Tree:

```
go-meta-listener
├── go.sum
├── example
│   └── main.go
├── eval.md
├── mirror
│   ├── listener.go
│   ├── metaproxy
│   │   ├── main.go
│   │   └── README.md
│   ├── README.md
│   └── header.go
├── metalistener.go
├── go.mod
├── LICENSE
├── README.md
└── Makefile

```

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/go.sum`:

```````sum
github.com/cretz/bine v0.2.0 h1:8GiDRGlTgz+o8H9DSnsl+5MeBK4HsExxgl6WgzOCuZo=
github.com/cretz/bine v0.2.0/go.mod h1:WU4o9QR9wWp8AVKtTM1XD5vUHkEqnf2vVSo6dBqbetI=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/go-i2p/i2pkeys v0.0.0-20241108200332-e4f5ccdff8c4/go.mod h1:m5TlHjPZrU5KbTd7Lr+I2rljyC6aJ88HdkeMQXV0U0E=
github.com/go-i2p/i2pkeys v0.33.10-0.20241113193422-e10de5e60708 h1:Tiy9IBwi21maNpK74yCdHursJJMkyH7w87tX1nXGWzg=
github.com/go-i2p/i2pkeys v0.33.10-0.20241113193422-e10de5e60708/go.mod h1:m5TlHjPZrU5KbTd7Lr+I2rljyC6aJ88HdkeMQXV0U0E=
github.com/go-i2p/onramp v0.33.92 h1:Dk3A0SGpdEw829rSjW2LqN8o16pUvuhiN0vn36z7Gpc=
github.com/go-i2p/onramp v0.33.92/go.mod h1:5sfB8H2xk05gAS2K7XAUZ7ekOfwGJu3tWF0fqdXzJG4=
github.com/go-i2p/sam3 v0.33.9 h1:3a+gunx75DFc6jxloUZTAVJbdP6736VU1dy2i7I9fKA=
github.com/go-i2p/sam3 v0.33.9/go.mod h1:oDuV145l5XWKKafeE4igJHTDpPwA0Yloz9nyKKh92eo=
github.com/opd-ai/wileedot v0.0.0-20241217172720-521d4175e624 h1:FXCTQV93+31Yj46zpYbd41es+EYgT7qi4RK6KSVrGQM=
github.com/opd-ai/wileedot v0.0.0-20241217172720-521d4175e624/go.mod h1:ftKSvvGC9FnxZeuL3B4MB6q/DOzVSV0kET08YUyDwbM=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.9.3 h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=
github.com/sirupsen/logrus v1.9.3/go.mod h1:naHLuLoDiP4jHNo9R0sCBMtWGeIprob74mVsIT4qYEQ=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.4 h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=
github.com/stretchr/testify v1.8.4/go.mod h1:sz/lmYIOXD/1dqDmKjjqLyZ2RngseejIcXlSw2iwfAo=
golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a/go.mod h1:P+XmwS30IXTQdn5tA2iutPOUgjI07+tq3H3K9MVA1s8=
golang.org/x/crypto v0.31.0 h1:ihbySMvVjLAeSH1IbfcRTkD/iNscyz8rGzjF/E5hV6U=
golang.org/x/crypto v0.31.0/go.mod h1:kDsLvtWBEx7MV9tJOj9bnXsPbxwJQ6csT/x4KIN4Ssk=
golang.org/x/net v0.0.0-20210226172049-e18ecbb05110/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
golang.org/x/net v0.0.0-20210525063256-abc453219eb5/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/net v0.31.0 h1:68CPQngjLL0r2AlUKiSxtQFKvzRVbnzLwMUn5SzcLHo=
golang.org/x/net v0.31.0/go.mod h1:P4fl1q7dY2hnZFxEk4pPSkDHF+QqjitcnDjUQyMM+pM=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.27.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/sys v0.28.0 h1:Fksou7UEQUWlKvIdsqzJmUmCX3cZuD2+P3XyyzwMhlA=
golang.org/x/sys v0.28.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.21.0 h1:zyQAAkrwaneQ066sspRyJaG9VNi/YJ1NfzcGB3hZ/qo=
golang.org/x/text v0.21.0/go.mod h1:4IBbMaMmOPCJ8SecivzSH54+73PCFmPWxNTLm+vZkEQ=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/example/main.go`:

```````go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-i2p/go-meta-listener"
)

func main() {
	// Create a new meta listener
	metaListener := meta.NewMetaListener()
	defer metaListener.Close()

	// Create and add TCP listener
	tcpListener, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatalf("Failed to create TCP listener: %v", err)
	}
	if err := metaListener.AddListener("tcp", tcpListener); err != nil {
		log.Fatalf("Failed to add TCP listener: %v", err)
	}
	log.Println("Added TCP listener on 127.0.0.1:8082")

	// Create and add a Unix socket listener (on Unix systems)
	socketPath := "/tmp/example.sock"
	os.Remove(socketPath) // Clean up from previous runs
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Printf("Failed to create Unix socket listener: %v", err)
	} else {
		if err := metaListener.AddListener("unix", unixListener); err != nil {
			log.Printf("Failed to add Unix socket listener: %v", err)
		} else {
			log.Println("Added Unix socket listener on", socketPath)
		}
	}
	log.Println("Starting http server...")

	// Create a simple HTTP server using the meta listener
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from MetaListener! You connected via: %s\n", r.Proto)
		}),
	}
	log.Println("Server is ready to accept connections...")

	// Handle server shutdown gracefully
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server starting, listening on multiple transports")
		if err := server.Serve(metaListener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down the HTTP server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// Wait for all listener goroutines to exit
	if err := metaListener.WaitForShutdown(ctx); err != nil {
		log.Printf("Timed out waiting for listener shutdown: %v", err)
	}

	log.Println("Server stopped")
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/eval.md`:

```````md
Project Path: /home/idk/go/src/github.com/go-i2p/go-meta-listener

Source Tree:

```
go-meta-listener
├── go.sum
├── example
│   └── main.go
├── mirror
│   ├── listener.go
│   ├── metaproxy
│   │   ├── main.go
│   │   └── README.md
│   ├── README.md
│   └── header.go
├── metalistener.go
├── go.mod
├── LICENSE
├── README.md
└── Makefile

```

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/go.sum`:

```````sum
github.com/cretz/bine v0.2.0 h1:8GiDRGlTgz+o8H9DSnsl+5MeBK4HsExxgl6WgzOCuZo=
github.com/cretz/bine v0.2.0/go.mod h1:WU4o9QR9wWp8AVKtTM1XD5vUHkEqnf2vVSo6dBqbetI=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/go-i2p/i2pkeys v0.0.0-20241108200332-e4f5ccdff8c4/go.mod h1:m5TlHjPZrU5KbTd7Lr+I2rljyC6aJ88HdkeMQXV0U0E=
github.com/go-i2p/i2pkeys v0.33.10-0.20241113193422-e10de5e60708 h1:Tiy9IBwi21maNpK74yCdHursJJMkyH7w87tX1nXGWzg=
github.com/go-i2p/i2pkeys v0.33.10-0.20241113193422-e10de5e60708/go.mod h1:m5TlHjPZrU5KbTd7Lr+I2rljyC6aJ88HdkeMQXV0U0E=
github.com/go-i2p/onramp v0.33.92 h1:Dk3A0SGpdEw829rSjW2LqN8o16pUvuhiN0vn36z7Gpc=
github.com/go-i2p/onramp v0.33.92/go.mod h1:5sfB8H2xk05gAS2K7XAUZ7ekOfwGJu3tWF0fqdXzJG4=
github.com/go-i2p/sam3 v0.33.9 h1:3a+gunx75DFc6jxloUZTAVJbdP6736VU1dy2i7I9fKA=
github.com/go-i2p/sam3 v0.33.9/go.mod h1:oDuV145l5XWKKafeE4igJHTDpPwA0Yloz9nyKKh92eo=
github.com/opd-ai/wileedot v0.0.0-20241217172720-521d4175e624 h1:FXCTQV93+31Yj46zpYbd41es+EYgT7qi4RK6KSVrGQM=
github.com/opd-ai/wileedot v0.0.0-20241217172720-521d4175e624/go.mod h1:ftKSvvGC9FnxZeuL3B4MB6q/DOzVSV0kET08YUyDwbM=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.9.3 h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=
github.com/sirupsen/logrus v1.9.3/go.mod h1:naHLuLoDiP4jHNo9R0sCBMtWGeIprob74mVsIT4qYEQ=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.4 h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=
github.com/stretchr/testify v1.8.4/go.mod h1:sz/lmYIOXD/1dqDmKjjqLyZ2RngseejIcXlSw2iwfAo=
golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a/go.mod h1:P+XmwS30IXTQdn5tA2iutPOUgjI07+tq3H3K9MVA1s8=
golang.org/x/crypto v0.31.0 h1:ihbySMvVjLAeSH1IbfcRTkD/iNscyz8rGzjF/E5hV6U=
golang.org/x/crypto v0.31.0/go.mod h1:kDsLvtWBEx7MV9tJOj9bnXsPbxwJQ6csT/x4KIN4Ssk=
golang.org/x/net v0.0.0-20210226172049-e18ecbb05110/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
golang.org/x/net v0.0.0-20210525063256-abc453219eb5/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/net v0.31.0 h1:68CPQngjLL0r2AlUKiSxtQFKvzRVbnzLwMUn5SzcLHo=
golang.org/x/net v0.31.0/go.mod h1:P4fl1q7dY2hnZFxEk4pPSkDHF+QqjitcnDjUQyMM+pM=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.27.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/sys v0.28.0 h1:Fksou7UEQUWlKvIdsqzJmUmCX3cZuD2+P3XyyzwMhlA=
golang.org/x/sys v0.28.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.21.0 h1:zyQAAkrwaneQ066sspRyJaG9VNi/YJ1NfzcGB3hZ/qo=
golang.org/x/text v0.21.0/go.mod h1:4IBbMaMmOPCJ8SecivzSH54+73PCFmPWxNTLm+vZkEQ=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/example/main.go`:

```````go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-i2p/go-meta-listener"
)

func main() {
	// Create a new meta listener
	metaListener := meta.NewMetaListener()
	defer metaListener.Close()

	// Create and add TCP listener
	tcpListener, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatalf("Failed to create TCP listener: %v", err)
	}
	if err := metaListener.AddListener("tcp", tcpListener); err != nil {
		log.Fatalf("Failed to add TCP listener: %v", err)
	}
	log.Println("Added TCP listener on 127.0.0.1:8082")

	// Create and add a Unix socket listener (on Unix systems)
	socketPath := "/tmp/example.sock"
	os.Remove(socketPath) // Clean up from previous runs
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Printf("Failed to create Unix socket listener: %v", err)
	} else {
		if err := metaListener.AddListener("unix", unixListener); err != nil {
			log.Printf("Failed to add Unix socket listener: %v", err)
		} else {
			log.Println("Added Unix socket listener on", socketPath)
		}
	}
	log.Println("Starting http server...")

	// Create a simple HTTP server using the meta listener
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from MetaListener! You connected via: %s\n", r.Proto)
		}),
	}
	log.Println("Server is ready to accept connections...")

	// Handle server shutdown gracefully
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server starting, listening on multiple transports")
		if err := server.Serve(metaListener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down the HTTP server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// Wait for all listener goroutines to exit
	if err := metaListener.WaitForShutdown(ctx); err != nil {
		log.Printf("Timed out waiting for listener shutdown: %v", err)
	}

	log.Println("Server stopped")
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/listener.go`:

```````go
package mirror

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/go-i2p/go-meta-listener"
	"github.com/go-i2p/onramp"

	wileedot "github.com/opd-ai/wileedot"
)

type Mirror struct {
	*meta.MetaListener
	Onions  map[string]*onramp.Onion
	Garlics map[string]*onramp.Garlic
}

var _ net.Listener = &Mirror{}

func (m *Mirror) Close() error {
	log.Println("Closing Mirror")
	if err := m.MetaListener.Close(); err != nil {
		log.Println("Error closing MetaListener:", err)
	} else {
		log.Println("MetaListener closed")
	}
	for _, onion := range m.Onions {
		if err := onion.Close(); err != nil {
			log.Println("Error closing Onion:", err)
		} else {
			log.Println("Onion closed")
		}
	}
	for _, garlic := range m.Garlics {
		if err := garlic.Close(); err != nil {
			log.Println("Error closing Garlic:", err)
		} else {
			log.Println("Garlic closed")
		}
	}
	log.Println("Mirror closed")
	return nil
}

func NewMirror(name string) (*Mirror, error) {
	log.Println("Creating new Mirror")
	inner := meta.NewMetaListener()
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "")
	if name == "" {
		name = "mirror"
	}
	log.Printf("Creating new MetaListener with name: '%s'\n", name)
	onion, err := onramp.NewOnion("metalistener-" + name)
	if err != nil {
		return nil, err
	}
	log.Println("Created new Onion manager")
	garlic, err := onramp.NewGarlic("metalistener-"+name, "127.0.0.1:7656", onramp.OPT_WIDE)
	if err != nil {
		return nil, err
	}
	log.Println("Created new Garlic manager")
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		port = "3000"
	}
	onions := make(map[string]*onramp.Onion)
	garlics := make(map[string]*onramp.Garlic)
	onions[port] = onion
	garlics[port] = garlic
	ml := &Mirror{
		MetaListener: inner,
		Onions:       onions,
		Garlics:      garlics,
	}
	log.Printf("Mirror created with name: '%s' and port: '%s', '%s'\n", name, port, ml.MetaListener.Addr().String())
	return ml, nil
}

func (ml Mirror) Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
	log.Println("Starting Mirror Listener")
	log.Printf("Actual args: name: '%s' addr: '%s' certDir: '%s' hiddenTls: '%t'\n", name, addr, certdir, hiddenTls)
	// get the port:
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		// check if host is an IP address
		if net.ParseIP(name) == nil {
			// host = "127.0.0.1"
		}
		port = "3000"
	}
	if strings.HasSuffix(port, "22") {
		log.Println("Port ends with 22, setting hiddenTls to false")
		log.Println("This is a workaround for the fact that the default port for SSH is 22")
		log.Println("This is so self-configuring SSH servers can be used without TLS, which would make connecting to them wierd")
		hiddenTls = false
	}
	localAddr := net.JoinHostPort("127.0.0.1", port)
	// Listen on plain HTTP
	tcpListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return nil, err
	}
	if err := ml.AddListener(port, tcpListener); err != nil {
		return nil, err
	}
	log.Printf("HTTP Local listener added http://%s\n", tcpListener.Addr())
	log.Println("Checking for existing onion and garlic listeners")
	listenerId := fmt.Sprintf("metalistener-%s-%s", name, port)
	log.Println("Listener ID:", listenerId)
	// Check if onion and garlic listeners already exist
	if ml.Onions[port] == nil {
		// make a new onion listener
		// and add it to the map
		log.Println("Creating new onion listener")
		onion, err := onramp.NewOnion(listenerId)
		if err != nil {
			return nil, err
		}
		log.Println("Onion listener created for port", port)
		ml.Onions[port] = onion
	}
	if ml.Garlics[port] == nil {
		// make a new garlic listener
		// and add it to the map
		log.Println("Creating new garlic listener")
		garlic, err := onramp.NewGarlic(listenerId, "127.0.0.1:7656", onramp.OPT_WIDE)
		if err != nil {
			return nil, err
		}
		log.Println("Garlic listener created for port", port)
		ml.Garlics[port] = garlic
	}
	if hiddenTls {
		// make sure an onion and a garlic listener exist at ml.Onions[port] and ml.Garlics[port]
		// and listen on them, check existence first
		onionListener, err := ml.Onions[port].ListenTLS()
		if err != nil {
			return nil, err
		}
		oid := fmt.Sprintf("onion-%s", onionListener.Addr().String())
		if err := ml.AddListener(oid, onionListener); err != nil {
			return nil, err
		}
		log.Printf("OnionTLS listener added https://%s\n", onionListener.Addr())
		garlicListener, err := ml.Garlics[port].ListenTLS()
		if err != nil {
			return nil, err
		}
		gid := fmt.Sprintf("garlic-%s", garlicListener.Addr().String())
		if err := ml.AddListener(gid, garlicListener); err != nil {
			return nil, err
		}
		log.Printf("GarlicTLS listener added https://%s\n", garlicListener.Addr())
	} else {
		onionListener, err := ml.Onions[port].Listen()
		if err != nil {
			return nil, err
		}
		oid := fmt.Sprintf("onion-%s", onionListener.Addr().String())
		if err := ml.AddListener(oid, onionListener); err != nil {
			return nil, err
		}
		log.Printf("Onion listener added http://%s\n", onionListener.Addr())
		garlicListener, err := ml.Garlics[port].Listen()
		if err != nil {
			return nil, err
		}
		gid := fmt.Sprintf("garlic-%s", garlicListener.Addr().String())
		if err := ml.AddListener(gid, garlicListener); err != nil {
			return nil, err
		}
		log.Printf("Garlic listener added http://%s\n", garlicListener.Addr())
	}
	if addr != "" {
		cfg := wileedot.Config{
			Domain:         name,
			AllowedDomains: []string{name},
			CertDir:        certdir,
			Email:          addr,
		}
		tlsListener, err := wileedot.New(cfg)
		if err != nil {
			return nil, err
		}
		tid := fmt.Sprintf("tls-%s", tlsListener.Addr().String())
		if err := ml.AddListener(tid, tlsListener); err != nil {
			return nil, err
		}
		log.Printf("TLS listener added https://%s\n", tlsListener.Addr())
	}
	return &ml, nil
}

// Listen creates a new Mirror instance and sets up listeners for TLS, Onion, and Garlic.
// It returns the Mirror instance and any error encountered during setup.
// name is the domain name used for the TLS listener, required for Let's Encrypt.
// addr is the email address used for Let's Encrypt registration.
// It is recommended to use a valid email address for production use.
func Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
	ml, err := NewMirror(name)
	if err != nil {
		return nil, err
	}
	return ml.Listen(name, addr, certdir, hiddenTls)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/metaproxy/main.go`:

```````go
package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	"github.com/go-i2p/go-meta-listener/mirror"
)

// main function sets up a meta listener that forwards connections to a specified host and port.
// It listens for incoming connections and forwards them to the specified destination.
func main() {
	host := flag.String("host", "localhost", "Host to forward connections to")
	port := flag.Int("port", 8080, "Port to forward connections to")
	listenPort := flag.Int("listen-port", 3002, "Port to listen for incoming connections")
	domain := flag.String("domain", "i2pgit.org", "Domain name for TLS listener")
	email := flag.String("email", "", "Email address for Let's Encrypt registration")
	certDir := flag.String("certdir", "./certs", "Directory for storing certificates")
	hiddenTls := flag.Bool("hidden-tls", false, "Enable hidden TLS")
	flag.Parse()
	addr := net.JoinHostPort(*domain, fmt.Sprintf("%d", *listenPort))
	// Create a new meta listener
	metaListener, err := mirror.Listen(addr, *email, *certDir, *hiddenTls)
	if err != nil {
		panic(err)
	}
	defer metaListener.Close()
	// forward all connections recieved on the meta listener to a local host:port
	for {
		conn, err := metaListener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			localConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
			if err != nil {
				panic(err)
			}
			defer localConn.Close()
			go io.Copy(localConn, conn)
			io.Copy(conn, localConn)
		}()
	}
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/metaproxy/README.md`:

```````md
# metaproxy - Connection Forwarder

A simple utility that forwards connections from a meta listener to a specified host and port. Part of the `go-i2p/go-meta-listener` toolset.
Automatically forwards a local service to a TLS Service, an I2P Eepsite, and a Tor Onion service at the same time.

## Installation

To install the metaproxy utility, use:

```bash
go install github.com/go-i2p/go-meta-listener/mirror/metaproxy@latest
```

## Usage

```bash
metaproxy [options]
```

### Options

- `-host`: Host to forward connections to (default: "localhost")
- `-port`: Port to forward connections to (default: 8080)
- `-domain`: Domain name for TLS listener (default: "i2pgit.org")
- `-email`: Email address for Let's Encrypt registration (default: "example@example.com")
- `-certdir`: Directory for storing certificates (default: "./certs")
- `-hidden-tls`: Enable hidden TLS (default: false)

## Description

metaproxy creates a meta listener that can accept connections from multiple transport types and forwards them to a specified destination (host:port).
It supports TLS with automatic certificate management through Let's Encrypt, I2P EepSites, and Tor Onion Services.

## Examples

Forward connections to a local web server:
```bash
metaproxy -host localhost -port 3000
```

Forward connections with custom TLS settings:
```bash
metaproxy -domain yourdomain.com -email you@example.com -certdir /etc/certs -port 8443
```

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/README.md`:

```````md
# Mirror Listener

A network listener implementation that simultaneously listens on clearnet (TLS), Tor onion services, and I2P garlic services.

## Overview

Mirror Listener is a wrapper around the [go-meta-listener](https://github.com/go-i2p/go-meta-listener) package that provides a simplified interface for setting up multi-protocol listeners. It automatically configures:

- TLS-secured clearnet connections with Let's Encrypt certificates
- Tor onion service endpoints
- I2P garlic service endpoints

This allows you to run a single service that's accessible through multiple network layers and protocols.

## Installation

```bash
go get github.com/go-i2p/go-meta-listener/mirror
```

## Usage

```go
import (
    "github.com/go-i2p/go-meta-listener/mirror"
    "net/http"
)

func main() {
    // Create a multi-protocol listener
    listener, err := mirror.Listen(
        "yourdomain.com",        // Domain name for TLS
        "your.email@example.com", // Email for Let's Encrypt
        "./certs",               // Certificate directory
        false,                   // Enable/disable TLS on hidden services
    )
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    
    // Use with standard library
    http.Serve(listener, yourHandler)
}
```

## Configuration Options

- **Domain Name**: Required for TLS certificate issuance through Let's Encrypt
- **Email Address**: Used for Let's Encrypt registration
- **Certificate Directory**: Where TLS certificates will be stored
- **Hidden TLS**: When set to true, enables TLS for Tor and I2P services as well

## Example: Connection Forwarding

See the [example directory](./example) for a complete example of using Mirror Listener to forward connections to a local service.
```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/header.go`:

```````go
package mirror

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// AddHeaders adds headers to the connection.
// It takes a net.Conn and a map of headers as input.
// It only adds headers if the connection is an HTTP connection.
// It returns a net.Conn with the headers added.
func AddHeaders(conn net.Conn, headers map[string]string) net.Conn {
	// Create a buffer to store the original request
	var buf bytes.Buffer
	teeReader := io.TeeReader(conn, &buf)

	// Try to read the request, but also save it to our buffer
	req, err := http.ReadRequest(bufio.NewReader(teeReader))
	if err != nil {
		// Not an HTTP request or couldn't parse, return original connection
		return conn
	}

	// Add our headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Create a pipe to connect our modified request with the output
	pr, pw := io.Pipe()

	// Write the modified request to one end of the pipe
	go func() {
		req.Write(pw)
		// Then copy the rest of the original connection
		io.Copy(pw, conn)
		pw.Close()
	}()

	// Return a ReadWriter that reads from our pipe and writes to the original connection
	return &readWriteConn{
		Reader: pr,
		Writer: conn,
		conn:   conn,
	}
}

// readWriteConn implements net.Conn
type readWriteConn struct {
	io.Reader
	io.Writer
	conn net.Conn
}

// Implement the rest of net.Conn interface by delegating to the original connection
func (rwc *readWriteConn) Close() error                       { return rwc.conn.Close() }
func (rwc *readWriteConn) LocalAddr() net.Addr                { return rwc.conn.LocalAddr() }
func (rwc *readWriteConn) RemoteAddr() net.Addr               { return rwc.conn.RemoteAddr() }
func (rwc *readWriteConn) SetDeadline(t time.Time) error      { return rwc.conn.SetDeadline(t) }
func (rwc *readWriteConn) SetReadDeadline(t time.Time) error  { return rwc.conn.SetReadDeadline(t) }
func (rwc *readWriteConn) SetWriteDeadline(t time.Time) error { return rwc.conn.SetWriteDeadline(t) }

// Accept accepts a connection from the listener.
// It takes a net.Listener as input and returns a net.Conn with the headers added.
// It is used to accept connections from the meta listener and add headers to them.
func (ml *Mirror) Accept() (net.Conn, error) {
	// Accept a connection from the listener
	conn, err := ml.MetaListener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return nil, err
	}

	// Check if the connection is a TLS connection
	if tlsConn, ok := conn.(*tls.Conn); ok {
		// If it is a TLS connection, perform the handshake
		if err := tlsConn.Handshake(); err != nil {
			log.Println("Error performing TLS handshake:", err)
			return nil, err
		}
		// If the handshake is successful, get the underlying connection
		//conn = tlsConn.NetConn()
	}

	host := map[string]string{
		"Host":              conn.LocalAddr().String(),
		"X-Forwarded-For":   conn.RemoteAddr().String(),
		"X-Forwarded-Proto": "http",
	}

	// Add headers to the connection
	conn = AddHeaders(conn, host)

	return conn, nil
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/metalistener.go`:

```````go
package meta

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	// ErrListenerClosed is returned when attempting to accept on a closed listener
	ErrListenerClosed = errors.New("listener is closed")
	// ErrNoListeners is returned when the meta listener has no active listeners
	ErrNoListeners = errors.New("no active listeners")
)

// MetaListener implements the net.Listener interface and manages multiple
// underlying network listeners as a unified interface.
type MetaListener struct {
	// listeners is a map of registered listeners with their unique identifiers
	listeners map[string]net.Listener
	// listenerWg tracks active listener goroutines for graceful shutdown
	listenerWg sync.WaitGroup
	// connCh is used to receive connections from all managed listeners
	connCh chan ConnResult
	// closeCh signals all goroutines to stop
	closeCh chan struct{}
	// isClosed indicates whether the meta listener has been closed
	isClosed bool
	// mu protects concurrent access to the listener's state
	mu sync.RWMutex
}

// ConnResult represents a connection received from a listener
type ConnResult struct {
	net.Conn
	src string // source listener ID
}

// NewMetaListener creates a new MetaListener instance ready to manage multiple listeners.
func NewMetaListener() *MetaListener {
	return &MetaListener{
		listeners: make(map[string]net.Listener),
		connCh:    make(chan ConnResult, 100), // Larger buffer for high connection volume
		closeCh:   make(chan struct{}),
	}
}

// AddListener adds a new listener with the specified ID.
// Returns an error if a listener with the same ID already exists or if the
// provided listener is nil.
func (ml *MetaListener) AddListener(id string, listener net.Listener) error {
	if listener == nil {
		return errors.New("cannot add nil listener")
	}

	ml.mu.Lock()
	defer ml.mu.Unlock()

	if ml.isClosed {
		return ErrListenerClosed
	}

	if _, exists := ml.listeners[id]; exists {
		return fmt.Errorf("listener with ID '%s' already exists", id)
	}

	ml.listeners[id] = listener

	// Start a goroutine to handle connections from this listener
	ml.listenerWg.Add(1)
	go ml.handleListener(id, listener)

	return nil
}

// RemoveListener stops and removes the listener with the specified ID.
// Returns an error if no listener with that ID exists.
func (ml *MetaListener) RemoveListener(id string) error {
	ml.mu.Lock()
	defer ml.mu.Unlock()

	listener, exists := ml.listeners[id]
	if !exists {
		return fmt.Errorf("no listener with ID '%s' exists", id)
	}

	// Close the specific listener
	err := listener.Close()
	delete(ml.listeners, id)

	return err
}

// handleListener runs in a separate goroutine for each added listener
// and forwards accepted connections to the connCh channel.
func (ml *MetaListener) handleListener(id string, listener net.Listener) {
	defer func() {
		log.Printf("Listener goroutine for %s exiting", id)
		ml.listenerWg.Done()
	}()

	for {
		// First check if the MetaListener is closed
		select {
		case <-ml.closeCh:
			log.Printf("MetaListener closed, stopping %s listener", id)
			return
		default:
		}

		// Set a deadline for Accept to prevent blocking indefinitely
		if deadline, ok := listener.(interface{ SetDeadline(time.Time) error }); ok {
			deadline.SetDeadline(time.Now().Add(1 * time.Second))
		}

		conn, err := listener.Accept()
		if err != nil {
			// Check if this is a timeout error (which we expect due to our deadline)
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			// Check if this is any other temporary error
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Printf("Temporary error in %s listener: %v, retrying in 100ms", id, err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			log.Printf("Permanent error in %s listener: %v, stopping", id, err)
			ml.mu.Lock()
			delete(ml.listeners, id)
			ml.mu.Unlock()
			return
		}

		// If we reach here, we have a valid connection
		log.Printf("Listener %s accepted connection from %s", id, conn.RemoteAddr())

		// Try to forward the connection, but don't block indefinitely
		select {
		case ml.connCh <- ConnResult{Conn: conn, src: id}:
			log.Printf("Connection from %s successfully forwarded via %s", conn.RemoteAddr(), id)
		case <-ml.closeCh:
			log.Printf("MetaListener closing while forwarding connection, closing connection")
			conn.Close()
			return
		case <-time.After(5 * time.Second):
			// If we can't forward within 5 seconds, something is seriously wrong
			log.Printf("WARNING: Connection forwarding timed out, closing connection from %s", conn.RemoteAddr())
			conn.Close()
		}
	}
}

// Accept implements the net.Listener Accept method.
// It waits for and returns the next connection from any of the managed listeners.
func (ml *MetaListener) Accept() (net.Conn, error) {
	for {
		ml.mu.RLock()
		if ml.isClosed {
			ml.mu.RUnlock()
			return nil, ErrListenerClosed
		}

		if len(ml.listeners) == 0 {
			ml.mu.RUnlock()
			return nil, ErrNoListeners
		}
		ml.mu.RUnlock()

		// Wait for either a connection or close signal
		select {
		case result, ok := <-ml.connCh:
			if !ok {
				return nil, ErrListenerClosed
			}
			log.Printf("Accept returning connection from %s via %s",
				result.RemoteAddr(), result.src)
			return result.Conn, nil
		case <-ml.closeCh:
			return nil, ErrListenerClosed
		}
	}
}

// Close implements the net.Listener Close method.
// It closes all managed listeners and releases resources.
func (ml *MetaListener) Close() error {
	ml.mu.Lock()

	if ml.isClosed {
		ml.mu.Unlock()
		return nil
	}

	log.Printf("Closing MetaListener with %d listeners", len(ml.listeners))
	ml.isClosed = true

	// Signal all goroutines to stop
	close(ml.closeCh)

	// Close all listeners
	var errs []error
	for id, listener := range ml.listeners {
		if err := listener.Close(); err != nil {
			log.Printf("Error closing %s listener: %v", id, err)
			errs = append(errs, err)
		}
	}

	ml.mu.Unlock()

	// Wait for all listener goroutines to exit
	ml.listenerWg.Wait()
	log.Printf("All listener goroutines have exited")

	// Return combined errors if any
	if len(errs) > 0 {
		return fmt.Errorf("errors closing listeners: %v", errs)
	}

	return nil
}

// Addr implements the net.Listener Addr method.
// It returns a MetaAddr representing all managed listeners.
func (ml *MetaListener) Addr() net.Addr {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	addresses := make([]net.Addr, 0, len(ml.listeners))
	for _, listener := range ml.listeners {
		addresses = append(addresses, listener.Addr())
	}

	return &MetaAddr{addresses: addresses}
}

// ListenerIDs returns the IDs of all active listeners.
func (ml *MetaListener) ListenerIDs() []string {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	ids := make([]string, 0, len(ml.listeners))
	for id := range ml.listeners {
		ids = append(ids, id)
	}

	return ids
}

// Count returns the number of active listeners.
func (ml *MetaListener) Count() int {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	return len(ml.listeners)
}

// WaitForShutdown blocks until all listener goroutines have exited.
// This is useful for ensuring clean shutdown in server applications.
func (ml *MetaListener) WaitForShutdown(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		ml.listenerWg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// MetaAddr implements the net.Addr interface for a meta listener.
type MetaAddr struct {
	addresses []net.Addr
}

// Network returns the name of the network.
func (ma *MetaAddr) Network() string {
	return "meta"
}

// String returns a string representation of all managed addresses.
func (ma *MetaAddr) String() string {
	if len(ma.addresses) == 0 {
		return "meta(empty)"
	}

	result := "meta("
	for i, addr := range ma.addresses {
		if i > 0 {
			result += ", "
		}
		result += addr.String()
	}
	result += ")"

	return result
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/go.mod`:

```````mod
module github.com/go-i2p/go-meta-listener

go 1.23.5

require (
	github.com/go-i2p/onramp v0.33.92
	github.com/opd-ai/wileedot v0.0.0-20241217172720-521d4175e624
)

require (
	github.com/cretz/bine v0.2.0 // indirect
	github.com/go-i2p/i2pkeys v0.33.10-0.20241113193422-e10de5e60708 // indirect
	github.com/go-i2p/sam3 v0.33.9 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/LICENSE`:

```````
MIT License

Copyright (c) 2025 I2P For Go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/README.md`:

```````md
# go-meta-listener

A Go package that implements a unified network listener interface capable of simultaneously handling connections from multiple underlying transport protocols.

[![Go Reference](https://pkg.go.dev/badge/github.com/go-i2p/go-meta-listener.svg)](https://pkg.go.dev/github.com/go-i2p/go-meta-listener)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

`go-meta-listener` provides a "meta listener" implementation that:
- Manages multiple network listeners through a single interface
- Supports any `net.Listener` implementation (TCP, Unix sockets, TLS, etc.)
- Handles connections and errors from all managed listeners
- Enables graceful shutdown across all listeners

The package also includes a specialized `mirror` implementation for multi-protocol network services supporting:
- TLS-secured clearnet connections
- Tor onion services
- I2P garlic services

## Installation

```bash
# Install core package
go get github.com/go-i2p/go-meta-listener

# For multi-protocol mirror functionality
go get github.com/go-i2p/go-meta-listener/mirror
```

## Basic Usage

```go
package main

import (
    "log"
    "net"
    "net/http"

    "github.com/go-i2p/go-meta-listener"
)

func main() {
    // Create a new meta listener
    metaListener := meta.NewMetaListener()
    defer metaListener.Close()

    // Add a TCP listener
    tcpListener, _ := net.Listen("tcp", ":8080")
    metaListener.AddListener("tcp", tcpListener)

    // Add a TLS listener
    tlsListener, _ := tls.Listen("tcp", ":8443", tlsConfig)
    metaListener.AddListener("tls", tlsListener)

    // Use with standard http server
    http.Serve(metaListener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from any protocol!"))
    }))
}
```

## Mirror Functionality

The `mirror` package provides a simpler interface for creating services available on clearnet, Tor, and I2P simultaneously:

```go
import "github.com/go-i2p/go-meta-listener/mirror"

// Create a multi-protocol listener
listener, err := mirror.Listen(
    "yourdomain.com",        // Domain name for TLS
    "your.email@example.com", // Email for Let's Encrypt
    "./certs",               // Certificate directory
    false                    // Enable/disable TLS on hidden services
)
defer listener.Close()

// Use with standard library
http.Serve(listener, yourHandler)
```

## Examples

See the [example directory](./example) for complete HTTP server examples and the [mirror/metaproxy directory](./mirror/metaproxy) for multi-protocol connection forwarding.

## License

MIT License - Copyright (c) 2025 I2P For Go

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/Makefile`:

```````
fmt:
	find . -name '*.go' -exec gofumpt -s -extra -w {} \;
```````
```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/listener.go`:

```````go
package mirror

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/go-i2p/go-meta-listener"
	"github.com/go-i2p/onramp"

	wileedot "github.com/opd-ai/wileedot"
)

type Mirror struct {
	*meta.MetaListener
	Onions  map[string]*onramp.Onion
	Garlics map[string]*onramp.Garlic
}

var _ net.Listener = &Mirror{}

func (m *Mirror) Close() error {
	log.Println("Closing Mirror")
	if err := m.MetaListener.Close(); err != nil {
		log.Println("Error closing MetaListener:", err)
	} else {
		log.Println("MetaListener closed")
	}
	for _, onion := range m.Onions {
		if err := onion.Close(); err != nil {
			log.Println("Error closing Onion:", err)
		} else {
			log.Println("Onion closed")
		}
	}
	for _, garlic := range m.Garlics {
		if err := garlic.Close(); err != nil {
			log.Println("Error closing Garlic:", err)
		} else {
			log.Println("Garlic closed")
		}
	}
	log.Println("Mirror closed")
	return nil
}

func NewMirror(name string) (*Mirror, error) {
	log.Println("Creating new Mirror")
	inner := meta.NewMetaListener()
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "")
	if name == "" {
		name = "mirror"
	}
	log.Printf("Creating new MetaListener with name: '%s'\n", name)
	onion, err := onramp.NewOnion("metalistener-" + name)
	if err != nil {
		return nil, err
	}
	log.Println("Created new Onion manager")
	garlic, err := onramp.NewGarlic("metalistener-"+name, "127.0.0.1:7656", onramp.OPT_WIDE)
	if err != nil {
		return nil, err
	}
	log.Println("Created new Garlic manager")
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		port = "3000"
	}
	onions := make(map[string]*onramp.Onion)
	garlics := make(map[string]*onramp.Garlic)
	onions[port] = onion
	garlics[port] = garlic
	ml := &Mirror{
		MetaListener: inner,
		Onions:       onions,
		Garlics:      garlics,
	}
	log.Printf("Mirror created with name: '%s' and port: '%s', '%s'\n", name, port, ml.MetaListener.Addr().String())
	return ml, nil
}

func (ml Mirror) Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
	log.Println("Starting Mirror Listener")
	log.Printf("Actual args: name: '%s' addr: '%s' certDir: '%s' hiddenTls: '%t'\n", name, addr, certdir, hiddenTls)
	// get the port:
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		// check if host is an IP address
		if net.ParseIP(name) == nil {
			// host = "127.0.0.1"
		}
		port = "3000"
	}
	if strings.HasSuffix(port, "22") {
		log.Println("Port ends with 22, setting hiddenTls to false")
		log.Println("This is a workaround for the fact that the default port for SSH is 22")
		log.Println("This is so self-configuring SSH servers can be used without TLS, which would make connecting to them wierd")
		hiddenTls = false
	}
	localAddr := net.JoinHostPort("127.0.0.1", port)
	// Listen on plain HTTP
	tcpListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return nil, err
	}
	if err := ml.AddListener(port, tcpListener); err != nil {
		return nil, err
	}
	log.Printf("HTTP Local listener added http://%s\n", tcpListener.Addr())
	log.Println("Checking for existing onion and garlic listeners")
	listenerId := fmt.Sprintf("metalistener-%s-%s", name, port)
	log.Println("Listener ID:", listenerId)
	// Check if onion and garlic listeners already exist
	if ml.Onions[port] == nil {
		// make a new onion listener
		// and add it to the map
		log.Println("Creating new onion listener")
		onion, err := onramp.NewOnion(listenerId)
		if err != nil {
			return nil, err
		}
		log.Println("Onion listener created for port", port)
		ml.Onions[port] = onion
	}
	if ml.Garlics[port] == nil {
		// make a new garlic listener
		// and add it to the map
		log.Println("Creating new garlic listener")
		garlic, err := onramp.NewGarlic(listenerId, "127.0.0.1:7656", onramp.OPT_WIDE)
		if err != nil {
			return nil, err
		}
		log.Println("Garlic listener created for port", port)
		ml.Garlics[port] = garlic
	}
	if hiddenTls {
		// make sure an onion and a garlic listener exist at ml.Onions[port] and ml.Garlics[port]
		// and listen on them, check existence first
		onionListener, err := ml.Onions[port].ListenTLS()
		if err != nil {
			return nil, err
		}
		oid := fmt.Sprintf("onion-%s", onionListener.Addr().String())
		if err := ml.AddListener(oid, onionListener); err != nil {
			return nil, err
		}
		log.Printf("OnionTLS listener added https://%s\n", onionListener.Addr())
		garlicListener, err := ml.Garlics[port].ListenTLS()
		if err != nil {
			return nil, err
		}
		gid := fmt.Sprintf("garlic-%s", garlicListener.Addr().String())
		if err := ml.AddListener(gid, garlicListener); err != nil {
			return nil, err
		}
		log.Printf("GarlicTLS listener added https://%s\n", garlicListener.Addr())
	} else {
		onionListener, err := ml.Onions[port].Listen()
		if err != nil {
			return nil, err
		}
		oid := fmt.Sprintf("onion-%s", onionListener.Addr().String())
		if err := ml.AddListener(oid, onionListener); err != nil {
			return nil, err
		}
		log.Printf("Onion listener added http://%s\n", onionListener.Addr())
		garlicListener, err := ml.Garlics[port].Listen()
		if err != nil {
			return nil, err
		}
		gid := fmt.Sprintf("garlic-%s", garlicListener.Addr().String())
		if err := ml.AddListener(gid, garlicListener); err != nil {
			return nil, err
		}
		log.Printf("Garlic listener added http://%s\n", garlicListener.Addr())
	}
	if addr != "" {
		cfg := wileedot.Config{
			Domain:         name,
			AllowedDomains: []string{name},
			CertDir:        certdir,
			Email:          addr,
		}
		tlsListener, err := wileedot.New(cfg)
		if err != nil {
			return nil, err
		}
		tid := fmt.Sprintf("tls-%s", tlsListener.Addr().String())
		if err := ml.AddListener(tid, tlsListener); err != nil {
			return nil, err
		}
		log.Printf("TLS listener added https://%s\n", tlsListener.Addr())
	}
	return &ml, nil
}

// Listen creates a new Mirror instance and sets up listeners for TLS, Onion, and Garlic.
// It returns the Mirror instance and any error encountered during setup.
// name is the domain name used for the TLS listener, required for Let's Encrypt.
// addr is the email address used for Let's Encrypt registration.
// It is recommended to use a valid email address for production use.
func Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
	ml, err := NewMirror(name)
	if err != nil {
		return nil, err
	}
	return ml.Listen(name, addr, certdir, hiddenTls)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/metaproxy/main.go`:

```````go
package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	"github.com/go-i2p/go-meta-listener/mirror"
)

// main function sets up a meta listener that forwards connections to a specified host and port.
// It listens for incoming connections and forwards them to the specified destination.
func main() {
	host := flag.String("host", "localhost", "Host to forward connections to")
	port := flag.Int("port", 8080, "Port to forward connections to")
	listenPort := flag.Int("listen-port", 3002, "Port to listen for incoming connections")
	domain := flag.String("domain", "i2pgit.org", "Domain name for TLS listener")
	email := flag.String("email", "", "Email address for Let's Encrypt registration")
	certDir := flag.String("certdir", "./certs", "Directory for storing certificates")
	hiddenTls := flag.Bool("hidden-tls", false, "Enable hidden TLS")
	flag.Parse()
	addr := net.JoinHostPort(*domain, fmt.Sprintf("%d", *listenPort))
	// Create a new meta listener
	metaListener, err := mirror.Listen(addr, *email, *certDir, *hiddenTls)
	if err != nil {
		panic(err)
	}
	defer metaListener.Close()
	// forward all connections recieved on the meta listener to a local host:port
	for {
		conn, err := metaListener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			localConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
			if err != nil {
				panic(err)
			}
			defer localConn.Close()
			go io.Copy(localConn, conn)
			io.Copy(conn, localConn)
		}()
	}
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/metaproxy/README.md`:

```````md
# metaproxy - Connection Forwarder

A simple utility that forwards connections from a meta listener to a specified host and port. Part of the `go-i2p/go-meta-listener` toolset.
Automatically forwards a local service to a TLS Service, an I2P Eepsite, and a Tor Onion service at the same time.

## Installation

To install the metaproxy utility, use:

```bash
go install github.com/go-i2p/go-meta-listener/mirror/metaproxy@latest
```

## Usage

```bash
metaproxy [options]
```

### Options

- `-host`: Host to forward connections to (default: "localhost")
- `-port`: Port to forward connections to (default: 8080)
- `-domain`: Domain name for TLS listener (default: "i2pgit.org")
- `-email`: Email address for Let's Encrypt registration (default: "example@example.com")
- `-certdir`: Directory for storing certificates (default: "./certs")
- `-hidden-tls`: Enable hidden TLS (default: false)

## Description

metaproxy creates a meta listener that can accept connections from multiple transport types and forwards them to a specified destination (host:port).
It supports TLS with automatic certificate management through Let's Encrypt, I2P EepSites, and Tor Onion Services.

## Examples

Forward connections to a local web server:
```bash
metaproxy -host localhost -port 3000
```

Forward connections with custom TLS settings:
```bash
metaproxy -domain yourdomain.com -email you@example.com -certdir /etc/certs -port 8443
```

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/README.md`:

```````md
# Mirror Listener

A network listener implementation that simultaneously listens on clearnet (TLS), Tor onion services, and I2P garlic services.

## Overview

Mirror Listener is a wrapper around the [go-meta-listener](https://github.com/go-i2p/go-meta-listener) package that provides a simplified interface for setting up multi-protocol listeners. It automatically configures:

- TLS-secured clearnet connections with Let's Encrypt certificates
- Tor onion service endpoints
- I2P garlic service endpoints

This allows you to run a single service that's accessible through multiple network layers and protocols.

## Installation

```bash
go get github.com/go-i2p/go-meta-listener/mirror
```

## Usage

```go
import (
    "github.com/go-i2p/go-meta-listener/mirror"
    "net/http"
)

func main() {
    // Create a multi-protocol listener
    listener, err := mirror.Listen(
        "yourdomain.com",        // Domain name for TLS
        "your.email@example.com", // Email for Let's Encrypt
        "./certs",               // Certificate directory
        false,                   // Enable/disable TLS on hidden services
    )
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    
    // Use with standard library
    http.Serve(listener, yourHandler)
}
```

## Configuration Options

- **Domain Name**: Required for TLS certificate issuance through Let's Encrypt
- **Email Address**: Used for Let's Encrypt registration
- **Certificate Directory**: Where TLS certificates will be stored
- **Hidden TLS**: When set to true, enables TLS for Tor and I2P services as well

## Example: Connection Forwarding

See the [example directory](./example) for a complete example of using Mirror Listener to forward connections to a local service.
```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/mirror/header.go`:

```````go
package mirror

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// AddHeaders adds headers to the connection.
// It takes a net.Conn and a map of headers as input.
// It only adds headers if the connection is an HTTP connection.
// It returns a net.Conn with the headers added.
func AddHeaders(conn net.Conn, headers map[string]string) net.Conn {
	// Create a buffer to store the original request
	var buf bytes.Buffer
	teeReader := io.TeeReader(conn, &buf)

	// Try to read the request, but also save it to our buffer
	req, err := http.ReadRequest(bufio.NewReader(teeReader))
	if err != nil {
		// Not an HTTP request or couldn't parse, return original connection
		return conn
	}

	// Add our headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Create a pipe to connect our modified request with the output
	pr, pw := io.Pipe()

	// Write the modified request to one end of the pipe
	go func() {
		req.Write(pw)
		// Then copy the rest of the original connection
		io.Copy(pw, conn)
		pw.Close()
	}()

	// Return a ReadWriter that reads from our pipe and writes to the original connection
	return &readWriteConn{
		Reader: pr,
		Writer: conn,
		conn:   conn,
	}
}

// readWriteConn implements net.Conn
type readWriteConn struct {
	io.Reader
	io.Writer
	conn net.Conn
}

// Implement the rest of net.Conn interface by delegating to the original connection
func (rwc *readWriteConn) Close() error                       { return rwc.conn.Close() }
func (rwc *readWriteConn) LocalAddr() net.Addr                { return rwc.conn.LocalAddr() }
func (rwc *readWriteConn) RemoteAddr() net.Addr               { return rwc.conn.RemoteAddr() }
func (rwc *readWriteConn) SetDeadline(t time.Time) error      { return rwc.conn.SetDeadline(t) }
func (rwc *readWriteConn) SetReadDeadline(t time.Time) error  { return rwc.conn.SetReadDeadline(t) }
func (rwc *readWriteConn) SetWriteDeadline(t time.Time) error { return rwc.conn.SetWriteDeadline(t) }

// Accept accepts a connection from the listener.
// It takes a net.Listener as input and returns a net.Conn with the headers added.
// It is used to accept connections from the meta listener and add headers to them.
func (ml *Mirror) Accept() (net.Conn, error) {
	// Accept a connection from the listener
	conn, err := ml.MetaListener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return nil, err
	}

	// Check if the connection is a TLS connection
	if tlsConn, ok := conn.(*tls.Conn); ok {
		// If it is a TLS connection, perform the handshake
		if err := tlsConn.Handshake(); err != nil {
			log.Println("Error performing TLS handshake:", err)
			return nil, err
		}
		// If the handshake is successful, get the underlying connection
		//conn = tlsConn.NetConn()
	}

	host := map[string]string{
		"Host":              conn.LocalAddr().String(),
		"X-Forwarded-For":   conn.RemoteAddr().String(),
		"X-Forwarded-Proto": "http",
	}

	// Add headers to the connection
	conn = AddHeaders(conn, host)

	return conn, nil
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/metalistener.go`:

```````go
package meta

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	// ErrListenerClosed is returned when attempting to accept on a closed listener
	ErrListenerClosed = errors.New("listener is closed")
	// ErrNoListeners is returned when the meta listener has no active listeners
	ErrNoListeners = errors.New("no active listeners")
)

// MetaListener implements the net.Listener interface and manages multiple
// underlying network listeners as a unified interface.
type MetaListener struct {
	// listeners is a map of registered listeners with their unique identifiers
	listeners map[string]net.Listener
	// listenerWg tracks active listener goroutines for graceful shutdown
	listenerWg sync.WaitGroup
	// connCh is used to receive connections from all managed listeners
	connCh chan ConnResult
	// closeCh signals all goroutines to stop
	closeCh chan struct{}
	// isClosed indicates whether the meta listener has been closed
	isClosed bool
	// mu protects concurrent access to the listener's state
	mu sync.RWMutex
}

// ConnResult represents a connection received from a listener
type ConnResult struct {
	net.Conn
	src string // source listener ID
}

// NewMetaListener creates a new MetaListener instance ready to manage multiple listeners.
func NewMetaListener() *MetaListener {
	return &MetaListener{
		listeners: make(map[string]net.Listener),
		connCh:    make(chan ConnResult, 100), // Larger buffer for high connection volume
		closeCh:   make(chan struct{}),
	}
}

// AddListener adds a new listener with the specified ID.
// Returns an error if a listener with the same ID already exists or if the
// provided listener is nil.
func (ml *MetaListener) AddListener(id string, listener net.Listener) error {
	if listener == nil {
		return errors.New("cannot add nil listener")
	}

	ml.mu.Lock()
	defer ml.mu.Unlock()

	if ml.isClosed {
		return ErrListenerClosed
	}

	if _, exists := ml.listeners[id]; exists {
		return fmt.Errorf("listener with ID '%s' already exists", id)
	}

	ml.listeners[id] = listener

	// Start a goroutine to handle connections from this listener
	ml.listenerWg.Add(1)
	go ml.handleListener(id, listener)

	return nil
}

// RemoveListener stops and removes the listener with the specified ID.
// Returns an error if no listener with that ID exists.
func (ml *MetaListener) RemoveListener(id string) error {
	ml.mu.Lock()
	defer ml.mu.Unlock()

	listener, exists := ml.listeners[id]
	if !exists {
		return fmt.Errorf("no listener with ID '%s' exists", id)
	}

	// Close the specific listener
	err := listener.Close()
	delete(ml.listeners, id)

	return err
}

// handleListener runs in a separate goroutine for each added listener
// and forwards accepted connections to the connCh channel.
func (ml *MetaListener) handleListener(id string, listener net.Listener) {
	defer func() {
		log.Printf("Listener goroutine for %s exiting", id)
		ml.listenerWg.Done()
	}()

	for {
		// First check if the MetaListener is closed
		select {
		case <-ml.closeCh:
			log.Printf("MetaListener closed, stopping %s listener", id)
			return
		default:
		}

		// Set a deadline for Accept to prevent blocking indefinitely
		if deadline, ok := listener.(interface{ SetDeadline(time.Time) error }); ok {
			deadline.SetDeadline(time.Now().Add(1 * time.Second))
		}

		conn, err := listener.Accept()
		if err != nil {
			// Check if this is a timeout error (which we expect due to our deadline)
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			// Check if this is any other temporary error
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Printf("Temporary error in %s listener: %v, retrying in 100ms", id, err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			log.Printf("Permanent error in %s listener: %v, stopping", id, err)
			ml.mu.Lock()
			delete(ml.listeners, id)
			ml.mu.Unlock()
			return
		}

		// If we reach here, we have a valid connection
		log.Printf("Listener %s accepted connection from %s", id, conn.RemoteAddr())

		// Try to forward the connection, but don't block indefinitely
		select {
		case ml.connCh <- ConnResult{Conn: conn, src: id}:
			log.Printf("Connection from %s successfully forwarded via %s", conn.RemoteAddr(), id)
		case <-ml.closeCh:
			log.Printf("MetaListener closing while forwarding connection, closing connection")
			conn.Close()
			return
		case <-time.After(5 * time.Second):
			// If we can't forward within 5 seconds, something is seriously wrong
			log.Printf("WARNING: Connection forwarding timed out, closing connection from %s", conn.RemoteAddr())
			conn.Close()
		}
	}
}

// Accept implements the net.Listener Accept method.
// It waits for and returns the next connection from any of the managed listeners.
func (ml *MetaListener) Accept() (net.Conn, error) {
	for {
		ml.mu.RLock()
		if ml.isClosed {
			ml.mu.RUnlock()
			return nil, ErrListenerClosed
		}

		if len(ml.listeners) == 0 {
			ml.mu.RUnlock()
			return nil, ErrNoListeners
		}
		ml.mu.RUnlock()

		// Wait for either a connection or close signal
		select {
		case result, ok := <-ml.connCh:
			if !ok {
				return nil, ErrListenerClosed
			}
			log.Printf("Accept returning connection from %s via %s",
				result.RemoteAddr(), result.src)
			return result.Conn, nil
		case <-ml.closeCh:
			return nil, ErrListenerClosed
		}
	}
}

// Close implements the net.Listener Close method.
// It closes all managed listeners and releases resources.
func (ml *MetaListener) Close() error {
	ml.mu.Lock()

	if ml.isClosed {
		ml.mu.Unlock()
		return nil
	}

	log.Printf("Closing MetaListener with %d listeners", len(ml.listeners))
	ml.isClosed = true

	// Signal all goroutines to stop
	close(ml.closeCh)

	// Close all listeners
	var errs []error
	for id, listener := range ml.listeners {
		if err := listener.Close(); err != nil {
			log.Printf("Error closing %s listener: %v", id, err)
			errs = append(errs, err)
		}
	}

	ml.mu.Unlock()

	// Wait for all listener goroutines to exit
	ml.listenerWg.Wait()
	log.Printf("All listener goroutines have exited")

	// Return combined errors if any
	if len(errs) > 0 {
		return fmt.Errorf("errors closing listeners: %v", errs)
	}

	return nil
}

// Addr implements the net.Listener Addr method.
// It returns a MetaAddr representing all managed listeners.
func (ml *MetaListener) Addr() net.Addr {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	addresses := make([]net.Addr, 0, len(ml.listeners))
	for _, listener := range ml.listeners {
		addresses = append(addresses, listener.Addr())
	}

	return &MetaAddr{addresses: addresses}
}

// ListenerIDs returns the IDs of all active listeners.
func (ml *MetaListener) ListenerIDs() []string {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	ids := make([]string, 0, len(ml.listeners))
	for id := range ml.listeners {
		ids = append(ids, id)
	}

	return ids
}

// Count returns the number of active listeners.
func (ml *MetaListener) Count() int {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	return len(ml.listeners)
}

// WaitForShutdown blocks until all listener goroutines have exited.
// This is useful for ensuring clean shutdown in server applications.
func (ml *MetaListener) WaitForShutdown(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		ml.listenerWg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// MetaAddr implements the net.Addr interface for a meta listener.
type MetaAddr struct {
	addresses []net.Addr
}

// Network returns the name of the network.
func (ma *MetaAddr) Network() string {
	return "meta"
}

// String returns a string representation of all managed addresses.
func (ma *MetaAddr) String() string {
	if len(ma.addresses) == 0 {
		return "meta(empty)"
	}

	result := "meta("
	for i, addr := range ma.addresses {
		if i > 0 {
			result += ", "
		}
		result += addr.String()
	}
	result += ")"

	return result
}

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/go.mod`:

```````mod
module github.com/go-i2p/go-meta-listener

go 1.23.5

require (
	github.com/go-i2p/onramp v0.33.92
	github.com/opd-ai/wileedot v0.0.0-20241217172720-521d4175e624
)

require (
	github.com/cretz/bine v0.2.0 // indirect
	github.com/go-i2p/i2pkeys v0.33.10-0.20241113193422-e10de5e60708 // indirect
	github.com/go-i2p/sam3 v0.33.9 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/LICENSE`:

```````
MIT License

Copyright (c) 2025 I2P For Go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/README.md`:

```````md
# go-meta-listener

A Go package that implements a unified network listener interface capable of simultaneously handling connections from multiple underlying transport protocols.

[![Go Reference](https://pkg.go.dev/badge/github.com/go-i2p/go-meta-listener.svg)](https://pkg.go.dev/github.com/go-i2p/go-meta-listener)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

`go-meta-listener` provides a "meta listener" implementation that:
- Manages multiple network listeners through a single interface
- Supports any `net.Listener` implementation (TCP, Unix sockets, TLS, etc.)
- Handles connections and errors from all managed listeners
- Enables graceful shutdown across all listeners

The package also includes a specialized `mirror` implementation for multi-protocol network services supporting:
- TLS-secured clearnet connections
- Tor onion services
- I2P garlic services

## Installation

```bash
# Install core package
go get github.com/go-i2p/go-meta-listener

# For multi-protocol mirror functionality
go get github.com/go-i2p/go-meta-listener/mirror
```

## Basic Usage

```go
package main

import (
    "log"
    "net"
    "net/http"

    "github.com/go-i2p/go-meta-listener"
)

func main() {
    // Create a new meta listener
    metaListener := meta.NewMetaListener()
    defer metaListener.Close()

    // Add a TCP listener
    tcpListener, _ := net.Listen("tcp", ":8080")
    metaListener.AddListener("tcp", tcpListener)

    // Add a TLS listener
    tlsListener, _ := tls.Listen("tcp", ":8443", tlsConfig)
    metaListener.AddListener("tls", tlsListener)

    // Use with standard http server
    http.Serve(metaListener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from any protocol!"))
    }))
}
```

## Mirror Functionality

The `mirror` package provides a simpler interface for creating services available on clearnet, Tor, and I2P simultaneously:

```go
import "github.com/go-i2p/go-meta-listener/mirror"

// Create a multi-protocol listener
listener, err := mirror.Listen(
    "yourdomain.com",        // Domain name for TLS
    "your.email@example.com", // Email for Let's Encrypt
    "./certs",               // Certificate directory
    false                    // Enable/disable TLS on hidden services
)
defer listener.Close()

// Use with standard library
http.Serve(listener, yourHandler)
```

## Examples

See the [example directory](./example) for complete HTTP server examples and the [mirror/metaproxy directory](./mirror/metaproxy) for multi-protocol connection forwarding.

## License

MIT License - Copyright (c) 2025 I2P For Go

```````

`/home/idk/go/src/github.com/go-i2p/go-meta-listener/Makefile`:

```````
fmt:
	find . -name '*.go' -exec gofumpt -s -extra -w {} \;
```````