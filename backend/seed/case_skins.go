package seed

import (
    "github.com/TyronOdame/CS-OPN/backend/database"
    "github.com/TyronOdame/CS-OPN/backend/models"
    "log"
)

// SeedCaseSkins links skins to cases with drop rates
func SeedCaseSkins() {
    // Check if already seeded
    var count int64
    database.DB.Model(&models.CaseSkin{}).Count(&count)
    if count > 0 {
        log.Println("🔗 Case-Skin links already seeded, skipping...")
        return
    }

    // Get all cases
    var cases []models.Case
    database.DB.Find(&cases)

    // Get skins by rarity
    var commonSkins, uncommonSkins, rareSkins, veryRareSkins, extremeRareSkins, legendarySkins []models.Skin
    database.DB.Where("rarity IN ?", []string{"Consumer Grade", "Industrial Grade"}).Find(&commonSkins)
    database.DB.Where("rarity = ?", "Mil-Spec").Find(&uncommonSkins)
    database.DB.Where("rarity = ?", "Restricted").Find(&rareSkins)
    database.DB.Where("rarity = ?", "Classified").Find(&veryRareSkins)
    database.DB.Where("rarity = ?", "Covert").Find(&extremeRareSkins)
    database.DB.Where("rarity = ?", "Rare Special").Find(&legendarySkins)

    // Link skins to each case with realistic drop rates
    for _, c := range cases {
        log.Printf("🔗 Linking skins to case: %s", c.Name)

        // Common skins (80% total drop rate)
        for _, skin := range commonSkins {
            database.DB.Create(&models.CaseSkin{
                CaseID:   c.ID,
                SkinID:   skin.ID,
                DropRate: 40.0, // Each common has 40% chance
            })
        }

        // Uncommon skins (15% total drop rate)
        for _, skin := range uncommonSkins {
            database.DB.Create(&models.CaseSkin{
                CaseID:   c.ID,
                SkinID:   skin.ID,
                DropRate: 7.5, // Each uncommon has 7.5% chance
            })
        }

        // Rare skins (3% total drop rate)
        for _, skin := range rareSkins {
            database.DB.Create(&models.CaseSkin{
                CaseID:   c.ID,
                SkinID:   skin.ID,
                DropRate: 1.5, // Each rare has 1.5% chance
            })
        }

        // Very rare skins (1.5% total drop rate)
        for _, skin := range veryRareSkins {
            database.DB.Create(&models.CaseSkin{
                CaseID:   c.ID,
                SkinID:   skin.ID,
                DropRate: 0.75, // Each very rare has 0.75% chance
            })
        }

        // Extremely rare skins (0.4% total drop rate)
        for _, skin := range extremeRareSkins {
            database.DB.Create(&models.CaseSkin{
                CaseID:   c.ID,
                SkinID:   skin.ID,
                DropRate: 0.2, // Each extremely rare has 0.2% chance
            })
        }

        // Legendary skins (0.1% drop rate)
        for _, skin := range legendarySkins {
            database.DB.Create(&models.CaseSkin{
                CaseID:   c.ID,
                SkinID:   skin.ID,
                DropRate: 0.1, // Legendary has 0.1% chance (1 in 1000)
            })
        }
    }

    log.Println("🔗 Case-Skin linking complete!")
}