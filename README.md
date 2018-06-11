# go-google-wifi
Golang wrapper for the Google Wifi API

So far, only the status endpoint has been implemented.

## Usage

```go
import (
	"fmt"
	"log"

	"github.com/timewasted/linode/dns"
)

// print the WAN ip addreess
func main() {
	status, err := gw_api.GetStatus()
	if err != nil {
		log.Fatalln(err)
	}

fmt.Println(status.Wan.LocalIpAddress)

```


