package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"finishy1995/device-manager/legacy/processor/processorclient"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	// Create a zrpc client configuration
	clientConfig := zrpc.RpcClientConf{
		Endpoints: []string{"localhost:8080"}, // Replace with your service address
		Timeout:   5000,                       // Timeout in milliseconds
	}

	// Initialize the zrpc client
	client := zrpc.MustNewClient(clientConfig)

	// Create a new Processor client
	processorClient := processorclient.NewProcessor(client)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Example: Call GenerateDemoMetadata
	req := &processorclient.GenerateDemoMetadataReq{
		DeviceNumber: 1000,
		DeviceParamNumber: 200,
	}
	resp, err := processorClient.GenerateDemoMetadata(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call GenerateDemoMetadata: %v", err)
	}
	// Print the response
	fmt.Printf("Response from GenerateDemoMetadata: %+v\n", resp)
}
