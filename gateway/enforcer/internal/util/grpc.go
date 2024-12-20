package util

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

// CreateGRPCConnection creates a gRPC connection using the provided context, host, port, and TLS configuration.
// It also sets up keepalive parameters for the connection to ensure that the connection remains alive even if no data is being sent.
// It returns a gRPC connection object and an error if the connection fails.
func CreateGRPCConnection(ctx context.Context, host, port string, tlsConfig *tls.Config) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	
	kacp := keepalive.ClientParameters{
		Time:                300 * time.Second,
		PermitWithoutStream: true,             
	}

	dialOptions := []grpc.DialOption{
		grpc.WithKeepaliveParams(kacp),
	}

	dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	conn, err := grpc.NewClient(
		address,
		dialOptions...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %v", err)
	}
	return conn, nil
}

// CreateGRPCConnectionWithRetry attempts to create a gRPC connection with retry logic.
// If the connection fails, it retries based on the provided maxRetries and retryInterval.
// If successful, it returns the gRPC connection; otherwise, it returns an error.
func CreateGRPCConnectionWithRetry(ctx context.Context, host, port string, tlsConfig *tls.Config, maxRetries int, retryInterval time.Duration) (*grpc.ClientConn, error) {
	for retries := 0; maxRetries == -1 || retries < maxRetries; retries++ {
		conn, err := CreateGRPCConnection(ctx, host, port, tlsConfig)
		if err == nil {
			return conn, nil
		}
		time.Sleep(retryInterval)
	}
	return nil, fmt.Errorf("failed to create gRPC connection after %d retries", maxRetries)
}

// CreateGRPCConnectionWithRetryAndPanic is similar to CreateGRPCConnectionWithRetry, but it panics if the connection cannot be established 
// after the specified number of retries. It is typically used when the application cannot proceed without the connection.
func CreateGRPCConnectionWithRetryAndPanic(ctx context.Context, host, port string, tlsConfig *tls.Config, maxRetries int, retryInterval time.Duration) *grpc.ClientConn {
	conn, err := CreateGRPCConnectionWithRetry(ctx, host, port, tlsConfig, maxRetries, retryInterval)
	if err != nil {
		panic(err)
	}
	return conn
}

// CreateGRPCServer creates a new gRPC server using the provided public and private key paths to load the TLS credentials.
// It returns the created gRPC server or an error if the credentials cannot be loaded.
func CreateGRPCServer(publicKeyPath, privateKeyPath string) (*grpc.Server, error) {
	// Load TLS credentials
	cert, err := LoadCertificates(publicKeyPath, privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load TLS credentials: %v", err)
	}

	creds := credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{cert}})
	// Create and return a new gRPC server with the loaded credentials
	return grpc.NewServer(grpc.Creds(creds)), nil
}