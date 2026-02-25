package seed

import (
	"log"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/models"
)

// SeedCaseSkins seeds case contents used by open-case logic.
// Note: despite the name, runtime open flow reads from case_contents.
func SeedCaseSkins() {
	// Check if case contents are already seeded.
	var contentCount int64
	database.DB.Model(&models.CaseContent{}).Count(&contentCount)
	if contentCount > 0 {
		log.Println("📦 Case contents already seeded, skipping...")
		return
	}

	// Get all cases.
	var cases []models.Case
	if err := database.DB.Find(&cases).Error; err != nil {
		log.Printf("❌ Failed to load cases for case_contents seeding: %v", err)
		return
	}
	if len(cases) == 0 {
		log.Println("⚠️  No cases found, skipping case_contents seeding.")
		return
	}

	// Get skins by rarity buckets.
	var commonSkins, uncommonSkins, rareSkins, veryRareSkins, extremeRareSkins, legendarySkins []models.Skin
	database.DB.Where("rarity IN ?", []string{"Consumer Grade", "Industrial Grade"}).Find(&commonSkins)
	database.DB.Where("rarity = ?", "Mil-Spec").Find(&uncommonSkins)
	database.DB.Where("rarity = ?", "Restricted").Find(&rareSkins)
	database.DB.Where("rarity = ?", "Classified").Find(&veryRareSkins)
	database.DB.Where("rarity = ?", "Covert").Find(&extremeRareSkins)
	database.DB.Where("rarity = ?", "Rare Special").Find(&legendarySkins)

	for _, c := range cases {
		log.Printf("📦 Seeding case contents for case: %s", c.Name)

		// Total drop chances per bucket (sum ~= 1.0).
		seedCaseContentsForCase(c, commonSkins, 0.80)
		seedCaseContentsForCase(c, uncommonSkins, 0.15)
		seedCaseContentsForCase(c, rareSkins, 0.03)
		seedCaseContentsForCase(c, veryRareSkins, 0.015)
		seedCaseContentsForCase(c, extremeRareSkins, 0.004)
		seedCaseContentsForCase(c, legendarySkins, 0.001)
	}

	log.Println("📦 Case contents seeding complete!")
}

func seedCaseContentsForCase(caseItem models.Case, skins []models.Skin, totalDropChance float64) {
	if len(skins) == 0 || totalDropChance <= 0 {
		return
	}

	perSkinChance := totalDropChance / float64(len(skins))

	for _, skin := range skins {
		content := models.CaseContent{
			CaseID:     caseItem.ID,
			SkinID:     skin.ID,
			DropChance: perSkinChance,
		}
		if err := database.DB.Create(&content).Error; err != nil {
			log.Printf("❌ Failed to create case_content for case=%s skin=%s: %v", caseItem.Name, skin.Name, err)
		}
	}
}