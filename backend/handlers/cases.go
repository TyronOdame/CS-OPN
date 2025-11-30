package handlers

import (
    "math/rand"
    "net/http"
    "time"

    "github.com/TyronOdame/CS-OPN/backend/database"
    "github.com/TyronOdame/CS-OPN/backend/middleware"
    "github.com/TyronOdame/CS-OPN/backend/models"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// GetAllCases returns all active cases
func GetAllCases(c *gin.Context) {
    var cases []models.Case

    // Only get active cases
    if err := database.DB.Where("is_active = ?", true).Find(&cases).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch cases",
        })
        return
    }

    // Convert to JSON
    var response []map[string]interface{}
    for _, cs := range cases {
        response = append(response, cs.ToJSON())
    }

    c.JSON(http.StatusOK, gin.H{
        "cases": response,
    })
}

// GetCaseByID returns a case with all possible skins
func GetCaseByID(c *gin.Context) {
    caseID := c.Param("id")  // ✅ Fixed: Added :=

    // Validate UUID
    if _, err := uuid.Parse(caseID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid case ID",
        })
        return
    }
        
    var caseItem models.Case
    if err := database.DB.First(&caseItem, "id = ? AND is_active = ?", caseID, true).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Case not found",
        })
        return
    }

    // get all skins in this case with drop chances
    var contents []models.CaseContent
    if err := database.DB.Preload("Skin").Where("case_id = ?", caseID).Find(&contents).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{  // Changed from NotFound
            "error": "Failed to fetch case contents",
        })
        return
    }

    // build response with case info and skins
    var skins []map[string]interface{}
    for _, content := range contents {  // ✅ Fixed: Removed space
        skinData := content.Skin.ToJSON()
        skinData["drop_chance"] = content.DropChance
        skinData["drop_percentage"] = content.GetDropPercentage()
        skins = append(skins, skinData)
    }

    response := caseItem.ToJSON()
    response["skins"] = skins

    c.JSON(http.StatusOK, response)  // ✅ Fixed: StatusOK (capital K)
}

// OpenCase handles case opening logic
func OpenCase(c *gin.Context) {
    // Get user ID from JWT  // ✅ Fixed: Comment now says "user ID"
    userID, err := middleware.GetUserID(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Unauthorized",
        })
        return
    }

    // Get case ID from URL
    caseID := c.Param("id")
    parsedCaseID, err := uuid.Parse(caseID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid case ID",
        })
        return
    }

    // Start transaction
    tx := database.DB.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Get case details
    var caseItem models.Case
    if err := tx.First(&caseItem, "id = ?", parsedCaseID).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Case not found",
        })
        return
    }

    // Check if case is active
    if !caseItem.IsActive {
        tx.Rollback()
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "This case is no longer available",
        })
        return
    }

    // Get user and check balance
    var user models.User
    if err := tx.First(&user, "id = ?", userID).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch user",
        })
        return
    }

    // Check if user has enough Case Bucks
    if user.Casebucks < caseItem.Price {  // ✅ Fixed: Changed Balance to Casebucks
        tx.Rollback()
        c.JSON(http.StatusBadRequest, gin.H{
            "error":           "Insufficient Case Bucks",  // ✅ Better error message
            "required":        caseItem.Price,
            "current_balance": user.Casebucks,
        })
        return
    }

    // Get all possible skins in the case with drop chances
    var contents []models.CaseContent
    if err := tx.Preload("Skin").Where("case_id = ?", parsedCaseID).Find(&contents).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch case contents",
        })
        return
    }

    // Select random skin based on drop chances
    selectedContent := selectRandomSkin(contents)  // ✅ Fixed: Removed ', err'
    if selectedContent == nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to select skin",
        })
        return
    }

    // Generate random float for skins 
    randomFloat := generateFloat()

    // Calculate value based on the float 
    skin := selectedContent.Skin
    floatMultiplier := 1.0 - randomFloat // better float will result in higher value
    skinValue := skin.MinValue + (skin.MaxValue-skin.MinValue)*floatMultiplier

    // Deduct Case Bucks from user 
    balanceBefore := user.Casebucks
    user.Casebucks -= caseItem.Price
    if err := tx.Save(&user).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to update balance",
        })
        return 
    }

    // Create transaction record for case opening
    transaction := models.Transaction{
        UserID:        userID,
        Type:          models.TransactionTypeCaseOpen,
        Amount:        -caseItem.Price,
        BalanceBefore: balanceBefore,
        BalanceAfter:  user.Casebucks,
        Description:   "Opened " + caseItem.Name,
        ReferenceID:   &parsedCaseID,
    }
    if err := tx.Create(&transaction).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create transaction",
        })
        return
    }

    // Add skin to user's inventory
    inventory := models.Inventory{
        UserID:       userID,
        SkinID:       skin.ID,
        Float:        randomFloat,
        AcquiredFrom: caseItem.Name,
        Value:        skinValue,
        IsSold:       false,
    }
    if err := tx.Create(&inventory).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to add skin to inventory",
        })
        return
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to complete case opening",
        })
        return
    }

    // Build response
    c.JSON(http.StatusOK, gin.H{
        "message":        "Case opened successfully!",
        "case":           caseItem.ToJSON(),
        "skin":           skin.ToJSON(),
        "float":          randomFloat,
        "condition":      inventory.GetCondition(),
        "value":          skinValue,
        "new_balance":    user.Casebucks,
        "inventory_id":   inventory.ID,
        "transaction_id": transaction.ID,
    })
}

// selectRandomSkin uses weighted random selection based on drop chances
func selectRandomSkin(contents []models.CaseContent) *models.CaseContent {
    if len(contents) == 0 {
        return nil
    }

    // Calculate total drop chance (should sum to ~1.0)
    var totalChance float64
    for _, content := range contents {
        totalChance += content.DropChance
    }

    // Generate random number between 0 and total chance
    rand.Seed(time.Now().UnixNano())
    randomValue := rand.Float64() * totalChance

    // Select skin based on weighted probability
    var cumulativeChance float64
    for i := range contents {
        cumulativeChance += contents[i].DropChance
        if randomValue <= cumulativeChance {
            return &contents[i]
        }
    }

    // Fallback to last skin (shouldn't happen)
    return &contents[len(contents)-1]
}

// generateFloat generates a random float value between 0.0 and 1.0
func generateFloat() float64 {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()
}