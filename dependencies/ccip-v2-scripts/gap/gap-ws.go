package gap

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
)

func waitForWebSocket(url string, timeout time.Duration) error {
	logger.Info("Waiting for websocket service to be available", "url", url)
	start := time.Now()
	for {
		if time.Since(start) > timeout {
			return fmt.Errorf("timeout waiting for WebSocket server: %s", url)
		}
		err := tryConnecting(url)
		if err == nil {
			logger.Debug("connection successful")

			return nil
		}

		time.Sleep(10 * time.Second)
	}
}

func tryConnecting(url string) error {
	headers, err := getHeaders()
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Debug("dialing websocket endpoint", "url", url, "headers", headers)

	conn, resp, err := websocket.DefaultDialer.Dial(url, headers)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		logger.Info("WebSocket server not available",
			"error", err.Error(),
		)
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)

			logger.Info("Error details: ",
				"url", url,
				"statusCode", resp.StatusCode,
				"body", body,
				"headers", resp.Header,
			)
		}
	}
	defer conn.Close()

	return err
}

//nolint:gosec
func TestWSConnectionViaGAP(env config.DevspaceEnv) {
	logger.Info("Testing connection to WebSocket via GAP")

	// WebSocket server URL (example: "ws://localhost:8080/ws")
	// serverURL := "wss://crib-local-geth-1337-ws.main.stage.cldev.sh"
	serverURL := fmt.Sprintf("wss://gap-%s-geth-1337-ws.public.main.stage.cldev.sh", env.Namespace)

	// Ignore TLS handshake errors
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{}

	err := waitForWebSocket(serverURL, 10*time.Minute)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Connect to the WebSocket server
	headers, err := getHeaders()
	if err != nil {
		logger.Error(err.Error())
	}
	conn, resp, err := websocket.DefaultDialer.Dial(serverURL, headers)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		logger.Error("Failed to connect to WebSocket server", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to WebSocket server")

	// Example message to send
	message := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`

	// Send the message to the WebSocket server
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		logger.Error("Failed to send message:", "error", err)
		os.Exit(1)
	}

	fmt.Println("Message sent to server")

	// Wait for a response from the server
	messageType, response, err := conn.ReadMessage()
	if err != nil {
		logger.Error("Failed to read message: %v", "error", err)
		os.Exit(1)
	}

	if messageType == websocket.TextMessage {
		fmt.Printf("Received response from server: %s\n", string(response))
	} else {
		fmt.Println("Received non-text message from server")
	}
}

func getHeaders() (map[string][]string, error) {
	headers := make(map[string][]string)

	// token := os.Getenv("GH_JWT_TOKEN")
	token, err := FetchJWTTokenForGAP(context.Background())
	if err != nil {
		return headers, fmt.Errorf("failed to fetch JWT token: %v", err)
	}
	if token == "" {
		return headers, fmt.Errorf("token is empty")
	}
	headers["x-authorization-github-jwt"] = []string{fmt.Sprintf("Bearer %s", token)}
	return headers, nil
}
