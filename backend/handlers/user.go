package handlers

import (
	"net/http"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"github.com/TyronOdame/CS-OPN/backend/middleware"
	"github.com/gin-gonic/gin"
)

// GetProfile returns the current user profile 