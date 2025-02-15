## Rocket.Chat Client for Go

### Installation
```bash
go get github.com/davesavic/rocketchat
```

### Usage
```go
package main

import (
    "fmt"
    "github.com/davesavic/rocketchat"
    "context"
)

func main() {
    client := rocketchat.NewClient("http://localhost:3000", "authtoken", "userid")

    // Create new user
    user, err := client.Users().Create(context.Background(), &rocketchat.CreateUser{})
}
```
