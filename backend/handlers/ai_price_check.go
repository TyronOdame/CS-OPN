package handlers

import (
	"hash/fnv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PriceCheckRequest struct {
	SkinName string `json:"skin_name" binding:"required"`
}

// PriceCheckMock provides a deterministic mocked price response.
func PriceCheckMock(c *gin.Context) {
	var req PriceCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "skin_name is required"})
		return
	}

	// Deterministic pseudo-price from skin name for stable UI demos.
	h := fnv.New32a()
	_, _ = h.Write([]byte(req.SkinName))
	seed := float64(h.Sum32()%15000) / 100.0 // 0.00 - 149.99
	mockPrice := 5.0 + seed                   // 5.00 - 154.99

	c.JSON(http.StatusOK, gin.H{
		"provider":      "mock",
		"skin_name":     req.SkinName,
		"suggested_usd": mockPrice,
		"message":       "Mocked AI price check result",
	})
}

