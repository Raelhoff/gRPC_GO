package configs

import (
    "fmt"
    "github.com/hyperledger/fabric-sdk-go/pkg/gateway"
    "os"
    "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
)

func ReadWallet() * string   {
    
      w, err := gateway.NewFileSystemWallet("./wallets")
     if err != nil {
                fmt.Printf("Failed to create wallet: %s\n", err)
                os.Exit(1)
     }

     if !w.Exists("Admin") {
                fmt.Println("Failed to get Admin from wallet")
                os.Exit(1)
     }

    fmt.Println("Connected to MongoDB")
    return w
}


//Client instance
var wallet * string = ReadWallet()
