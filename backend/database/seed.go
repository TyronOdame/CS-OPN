package database

import (
    "log"

    "github.com/TyronOdame/CS-OPN/backend/models"
)

// SeedDatabase populates the database with sample cases and skins
func SeedDatabase() error {
    log.Println("🌱 Seeding database with sample data...")

    // Check if data already exists
    var caseCount int64
    DB.Model(&models.Case{}).Count(&caseCount)
    if caseCount > 0 {
        log.Println("⏭️  Database already seeded, skipping...")
        return nil
    }

    // Create skins first
    skins, err := createSkins()
    if err != nil {
        return err
    }

    // Create cases
    cases, err := createCases()
    if err != nil {
        return err
    }

    // Link skins to cases with drop rates
    if err := linkCaseContents(cases, skins); err != nil {
        return err
    }

    log.Println("✅ Database seeded successfully!")
    return nil
}

// createSkins creates sample skins across all rarities
func createSkins() (map[string]models.Skin, error) {
    // Define skins with their properties
    skinDefinitions := map[string]models.Skin{
        // Consumer Grade (Common - 79.92%)
        "tec9_groundwater": {
            Name:        "Tec-9 | Groundwater",
            WeaponType:  "Pistol",
            Rarity:      "Consumer Grade",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpoor-mcjhhwszcdD4b09-3moS0mvLwOq7cqWdQ-sJ0teXI8oThxlKx-RdrZW6lI4CWJwBqNQnU_FXswO7rgMO96p_KzHVnunR25SrZzAv3308",
            MinValue:    0.50,
            MaxValue:    1.00,
            Description: "It has been painted using a hydrogrpahic pattern of a topographical map.",
        },
        "p250_mint_kimono": {
            Name:        "P250 | Mint Kimono",
            WeaponType:  "Pistol",
            Rarity:      "Consumer Grade",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopujwezhhwszYI2gS09-5lpKKqPrxN7LEmyVQ7MEpiLuSrYmnjQPkrRE-ZzqmJoORcVdtaQ3U-AXswbzngcPq7czKzHVlvSkm5H3D30vgtY9ZMhA",
            MinValue:    0.50,
            MaxValue:    1.00,
            Description: "It has been decorated with a mint-colored pattern of traditional Japanese designs.",
        },
        // Industrial Grade (Uncommon - 15.98%)
        "mac10_fade": {
            Name:        "MAC-10 | Fade",
            WeaponType:  "SMG",
            Rarity:      "Industrial Grade",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou7umeldf0Ob3fDxBvYyJgYWKkPvxDLfYkWNFppwp2L6QrI6m3wK2-hFsYT2lINCWdgQ8aArX_FK8krq6hcK86pTPn3A1s3Ug5nzazBG1gBwYO7Zsh_ePRF_VTerFBg",
            MinValue:    2.00,
            MaxValue:    4.00,
            Description: "It has been painted by airbrushing transparent paints that fade together over a chrome base coat.",
        },
        "ssg08_acid_fade": {
            Name:        "SSG 08 | Acid Fade",
            WeaponType:  "Sniper",
            Rarity:      "Industrial Grade",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopamie19f0Ob3Yi5FvISJmYWPnvb4J4Tdn2xZ_Isli7CZ8I2j3lCw-xI-ZWihd4WWcg85YV_T-1HowO3v1MC8tZTAz3F9-n51W_pXP2Q",
            MinValue:    2.00,
            MaxValue:    4.00,
            Description: "It has been painted by airbrushing transparent paints that fade together over a chrome base coat.",
        },
        // Mil-Spec (Restricted - 3.2%)
        "m4a1s_hyper_beast": {
            Name:        "M4A1-S | Hyper Beast",
            WeaponType:  "Rifle",
            Rarity:      "Mil-Spec",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhz2v_Nfz5H_uO1gb-Gw_alIITZk2pH8fp8j-jE_Jn4xlC9vh5yYzv2IYTBdgBqYAnZ_la6wL_mgpDu6oOJlyW-N5hc3A",
            MinValue:    5.00,
            MaxValue:    10.00,
            Description: "It has been custom painted with a beastly creature in psychedelic colors.",
        },
        "glock18_water_elemental": {
            Name:        "Glock-18 | Water Elemental",
            WeaponType:  "Pistol",
            Rarity:      "Mil-Spec",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgposbaqKAxf0Ob3djFN79fnzL-ckvbnNrfummpD78A_3OqXo9ug2AHnqRU-Y2_7I4DGIAU7Yw7S-1K7krjxxcjr4pUKfw",
            MinValue:    5.00,
            MaxValue:    10.00,
            Description: "It has been custom painted with a depiction of a water spirit.",
        },
        // Restricted (Purple - 0.64%)
        "ak47_redline": {
            Name:        "AK-47 | Redline",
            WeaponType:  "Rifle",
            Rarity:      "Restricted",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot7HxfDhjxszJemkV09-5lpKKqPv9NLPF2G5V-vp9g-7J4bP5iUazrl1lZDzwJtfAdFU2aFqB_VTswuzm05a_6Z6dySBluyEg-z-DyN-tCSAD",
            MinValue:    15.00,
            MaxValue:    30.00,
            Description: "It has been painted using a carbon fiber hydrographic over a red and black base coat.",
        },
        "aug_chameleon": {
            Name:        "AUG | Chameleon",
            WeaponType:  "Rifle",
            Rarity:      "Restricted",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot6-iFAR17PLfYQJD_9W7m5a0mvLwOq7c2GtXu8Ag3e2Wodz22lDg_kJrYmr1ItDHdlI6aQrU_lC3kOjxxcjrTvRbpGA",
            MinValue:    12.00,
            MaxValue:    25.00,
            Description: "It has been custom painted with a multicolored pattern.",
        },
        // Classified (Pink - 0.32%)
        "m4a4_desolate_space": {
            Name:        "M4A4 | Desolate Space",
            WeaponType:  "Rifle",
            Rarity:      "Classified",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJQJD_9W7m5a0n_L1J6_um25V4dB8xO2WrI2t2VCx-UduYjz3JoWVdQA7N1vT_QK5wejxxcjr-kZmQrA",
            MinValue:    30.00,
            MaxValue:    60.00,
            Description: "It has been custom painted with a cosmic design.",
        },
        "p90_trigon": {
            Name:        "P90 | Trigon",
            WeaponType:  "SMG",
            Rarity:      "Classified",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopuP1FAR17P7YKAJA4867kpKOqPv9NLPF2G5V-sB02-qUrN-s3gS2_0NuYGj3doCdcVU9ZQzS-VLowuq9gpO67s6dzSA3s3Fx7WGdwULxXE8d1A",
            MinValue:    25.00,
            MaxValue:    50.00,
            Description: "It has been painted with a trigon pattern.",
        },
        // Covert (Red - 0.64%)
        "awp_asiimov": {
            Name:        "AWP | Asiimov",
            WeaponType:  "Sniper",
            Rarity:      "Covert",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJU7c-ikZKSqPv9NLPF2GpTu8Ag2r-Zp9z32lLh_0FvZ2-lI4WddwI3ZV2C_FC-x-fp1p-4vp7KzCY37yEl4mGdwUIo7A80lQ",
            MinValue:    80.00,
            MaxValue:    150.00,
            Description: "It has been custom painted with a sci-fi design.",
        },
        "desert_eagle_blaze": {
            Name:        "Desert Eagle | Blaze",
            WeaponType:  "Pistol",
            Rarity:      "Covert",
            Float:       0.5,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgposr-kLAtl7PLZTjlH_9mkgIWKkPvLPr7Vn35cppEo27-Q8N-t3wW3_UdsZ2_0IYOcIFI3N13Z-wO6wOq9hMC46ZvPyyRh6CQ8pSGK2P3giBM",
            MinValue:    70.00,
            MaxValue:    130.00,
            Description: "It has been anodized in a flame pattern.",
        },
        // Rare Special (Gold - Knives - 0.26%)
        "karambit_fade": {
            Name:        "★ Karambit | Fade",
            WeaponType:  "Knife",
            Rarity:      "Rare Special",
            Float:       0.01,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovbSsLQJf2PLacDBA5ciJlY20k_jkI7fUhFRd4cJ5nqeQrdSl21Hm-hdoYGv7cI6Rdw47YlyDqADoxO3ngpLovJzAznJnuykq-z-DyB0S6bvY",
            MinValue:    800.00,
            MaxValue:    2000.00,
            Description: "It has been painted by airbrushing transparent paints that fade together over a chrome base coat.",
        },
        "butterfly_slaughter": {
            Name:        "★ Butterfly Knife | Slaughter",
            WeaponType:  "Knife",
            Rarity:      "Rare Special",
            Float:       0.01,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovbSsLQJf0ebcZThQ6tCvq4GGqPL6IITdn2xZ_Isli7jD9I2j2lGx-RVkMGnwLI-dcFU7YFvU_Fa8yOy-hJ-76YOJlyUIg41AoA",
            MinValue:    700.00,
            MaxValue:    1800.00,
            Description: "It has been painted in a zebra-stripe pattern with aluminum and chrome paints with various reflectivities.",
        },
        "m9_bayonet_doppler": {
            Name:        "★ M9 Bayonet | Doppler",
            WeaponType:  "Knife",
            Rarity:      "Rare Special",
            Float:       0.01,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovbSsLQJf3qr3czxb49KzgL-DjsjwN6vdk1Rd4cJ5nqfA89ul2lDsqBBoMWygIIKUIw46YFDR_VK4wO3v1p7quZvIziMwuCEm-z-DyGhpZX7D",
            MinValue:    600.00,
            MaxValue:    1500.00,
            Description: "It has been painted with black and silver metallic paints using a marbleizing medium, then candy coated.",
        },
        "bayonet_tiger_tooth": {
            Name:        "★ Bayonet | Tiger Tooth",
            WeaponType:  "Knife",
            Rarity:      "Rare Special",
            Float:       0.01,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovbSsLQJf3qr3czxb49KzgL-DjsjwN6vdk1Rd4cJ5ntbN9J7yjRrg-RE4MGv7I4TBcAJrZAzS-FDtyejv05e46Z7Jn3Nk6yQ8pSGKrUP1J1w",
            MinValue:    500.00,
            MaxValue:    1200.00,
            Description: "It has been painted in a striped pattern.",
        },
    }

    // Now create each skin and store it with its generated ID
    skins := make(map[string]models.Skin)
    for key, skinDef := range skinDefinitions {
        skin := skinDef // Copy the definition
        if err := DB.Create(&skin).Error; err != nil {
            return nil, err
        }
        skins[key] = skin // Store with generated ID
    }

    log.Printf("✅ Created %d skins\n", len(skins))
    return skins, nil
}

// createCases creates sample cases
func createCases() (map[string]models.Case, error) {
    // Define cases with their properties
    caseDefinitions := map[string]models.Case{
        "chroma": {
            Name:        "Chroma Case",
            Price:       2.50,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DAQ1h3LAVbv6mxFABs3OXNYgJR_Nm1nYGHnuTgDKnCmGpa7cdlmdbN_Iv9nBri-xZqMWqndYKXJw85ZwyC-FHrxOjmjcfv6pXJm2wj5HdzFbcCcw",
            Description: "The Chroma Case contains the Chroma Collection and was released as part of the January 8, 2015 update. This case features community-designed weapon finishes from the Chroma Collection and introduces rare special items - knives with Chroma finishes.",
            IsActive:    true,
        },
        "gamma": {
            Name:        "Gamma Case",
            Price:       3.00,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DAQ1h3LAVbv6mxFABs3OXNYgJR_Nm1nYGHnuTgDLfYkWNFppUk3riXo96njA3g_UJoaz-lIo-QcVc8Z1-F-APqx-y6gJe-7MzOzHY1siYi5WGdwULaISkPuw",
            Description: "The Gamma Case contains 17 community-designed weapon finishes and the all-new Gamma Finishes for knives. A portion of the proceeds from this case goes to the weapon finish designers.",
            IsActive:    true,
        },
        "revolution": {
            Name:        "Revolution Case",
            Price:       5.00,
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DAQ1h3LAVbv6mxFABs3OXNYgJR_Nm1nYGHnuTgDLbQhH9u5cRjiOXI_Iv9nBqxqEFlMGuhII_DIQNrZw7Q_Fe5wb_nm5W8ot2XzhK8xQjg",
            Description: "The Revolution Case contains the Revolution Collection and was released in February 2023. Features popular community designs and introduces new knife finishes.",
            IsActive:    true,
        },
    }

    // Now create each case and store it with its generated ID
    cases := make(map[string]models.Case)
    for key, caseDef := range caseDefinitions {
        caseItem := caseDef // Copy the definition
        if err := DB.Create(&caseItem).Error; err != nil {
            return nil, err
        }
        cases[key] = caseItem // Store with generated ID
    }

    log.Printf("✅ Created %d cases\n", len(cases))
    return cases, nil
}

// linkCaseContents links skins to cases with CS2-realistic drop rates
func linkCaseContents(cases map[string]models.Case, skins map[string]models.Skin) error {
    // Chroma Case contents
    chromaContents := []struct {
        skinKey string
        drop    float64
    }{
        {"tec9_groundwater", 0.1598},       // Consumer Grade
        {"p250_mint_kimono", 0.1598},       // Consumer Grade
        {"mac10_fade", 0.0799},             // Industrial Grade
        {"ssg08_acid_fade", 0.0799},        // Industrial Grade
        {"m4a1s_hyper_beast", 0.032},       // Mil-Spec
        {"glock18_water_elemental", 0.032}, // Mil-Spec
        {"ak47_redline", 0.0064},           // Restricted
        {"aug_chameleon", 0.0064},          // Restricted
        {"m4a4_desolate_space", 0.0032},    // Classified
        {"p90_trigon", 0.0032},             // Classified
        {"awp_asiimov", 0.0026},            // Covert
        {"karambit_fade", 0.00065},         // Rare Special (Knife)
        {"butterfly_slaughter", 0.00065},   // Rare Special (Knife)
    }

    for _, content := range chromaContents {
        cc := models.CaseContent{
            CaseID:     cases["chroma"].ID,
            SkinID:     skins[content.skinKey].ID,
            DropChance: content.drop,
        }
        if err := DB.Create(&cc).Error; err != nil {
            return err
        }
    }

    // Gamma Case contents
    gammaContents := []struct {
        skinKey string
        drop    float64
    }{
        {"p250_mint_kimono", 0.1598},
        {"ssg08_acid_fade", 0.1598},
        {"glock18_water_elemental", 0.032},
        {"aug_chameleon", 0.0064},
        {"p90_trigon", 0.0032},
        {"desert_eagle_blaze", 0.0026},
        {"m9_bayonet_doppler", 0.00065},
        {"bayonet_tiger_tooth", 0.00065},
    }

    for _, content := range gammaContents {
        cc := models.CaseContent{
            CaseID:     cases["gamma"].ID,
            SkinID:     skins[content.skinKey].ID,
            DropChance: content.drop,
        }
        if err := DB.Create(&cc).Error; err != nil {
            return err
        }
    }

    // Revolution Case contents
    revolutionContents := []struct {
        skinKey string
        drop    float64
    }{
        {"tec9_groundwater", 0.1598},
        {"mac10_fade", 0.0799},
        {"m4a1s_hyper_beast", 0.032},
        {"ak47_redline", 0.0064},
        {"m4a4_desolate_space", 0.0032},
        {"awp_asiimov", 0.0026},
        {"karambit_fade", 0.00065},
    }

    for _, content := range revolutionContents {
        cc := models.CaseContent{
            CaseID:     cases["revolution"].ID,
            SkinID:     skins[content.skinKey].ID,
            DropChance: content.drop,
        }
        if err := DB.Create(&cc).Error; err != nil {
            return err
        }
    }

    log.Println("✅ Linked skins to cases with drop rates")
    return nil
}