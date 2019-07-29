//
// ginrequestid
//
// Set an UUID4 string as Request ID into response headers ("X-Request-Id") and
// expose that value as "RequestId" in order to use it inside the application for logging
// or propagation to other systems.
//
package ginrequestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestID == "" {
			uuid4 := uuid.Must(uuid.NewRandom())
			requestID = uuid4.String()
		}

		// Expose it for use in the application
		c.Set("RequestId", requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
