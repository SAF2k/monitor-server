// main.go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ServerStatus represents the status of a server.
type ServerStatus struct {
	Server   string `json:"server"`
	IsActive bool   `json:"isActive"`
}

var serverStatuses = map[string]bool{
	"Server1": false,
	"Server2": false,
	// Add more servers as needed
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("templates/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"Servers": serverStatuses})
	})

	r.GET("/status", func(c *gin.Context) {
		var statuses []ServerStatus

		for server, isActive := range serverStatuses {
			statuses = append(statuses, ServerStatus{Server: server, IsActive: isActive})
		}

		c.JSON(http.StatusOK, statuses)
	})

	r.POST("/add", func(c *gin.Context) {
		ip := c.PostForm("ip")

		// Add the new server to the status map
		serverStatuses[ip] = false

		// Respond with the updated server list
		c.HTML(http.StatusOK, "index.html", gin.H{"Servers": serverStatuses})
	})

	go func() {
		for {
			time.Sleep(5 * time.Second)

			// Update server statuses (replace this with actual server status checking logic)
			for server := range serverStatuses {
				serverStatuses[server] = isServerActive(server)
			}
		}
	}()

	r.Run(":8080")
}

func isServerActive(server string) bool {
	// Add logic to check the status of the server (e.g., ping, HTTP request, etc.)
	// Return true if the server is active, false otherwise.
	return false
}
