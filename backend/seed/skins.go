package seed

import (
    "log"

    "github.com/TyronOdame/CS-OPN/backend/database"
    "github.com/TyronOdame/CS-OPN/backend/models"
    "github.com/google/uuid"
)

// SeedSkins adds test skins to the database
func SeedSkins() {
    // Check if skins already exist
    var count int64
    database.DB.Model(&models.Skin{}).Count(&count)
    if count > 0 {
        log.Println("🎨 Skins already seeded, skipping...")
        return
    }

    skins := []models.Skin{
        // COMMON SKINS ($0.50 - $2.00)
        {
            ID:         uuid.New(),
            Name:       "P250 | Valence",
            WeaponType: "Pistol",
            Rarity:     "Consumer Grade",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopujwezhjxszYI2gS09-5lpKKqPrxN7LEmyVQ7MEpiLuSrYmnjQO3-UdsZWD7cNeTc1Q9ZlyC_gW4kOfxxcjrT92cGQ",
            MinValue:   0.50,
            MaxValue:   2.00,
        },
        {
            ID:         uuid.New(),
            Name:       "MAC-10 | Fade",
            WeaponType: "SMG",
            Rarity:     "Industrial Grade",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou7umeldf0Ob3fDxBvYyJgIGbmOXgDKnBmX5D18h0i-yVo9-m3gfnrRJpZjz3cIXEJgI3Zw3Z_lbswb3shpa8uZvKyXpmvSYk7C3ZyRLjgU5SLrs40cUNcFo",
            MinValue:   1.00,
            MaxValue:   3.50,
        },

        // UNCOMMON SKINS ($3.00 - $8.00)
        {
            ID:         uuid.New(),
            Name:       "AK-47 | Blue Laminate",
            WeaponType: "Rifle",
            Rarity:     "Mil-Spec",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot7HxfDhjxszJemkV08-jhIGZmP_gfb7UlWJQ-vp9g-7J4cL33Qzh-kI6ZTv0JNLGdgM6YlnW_FW7lebxxcjrrvON_0A",
            MinValue:   3.00,
            MaxValue:   7.00,
        },
        {
            ID:         uuid.New(),
            Name:       "M4A4 | Radiation Hazard",
            WeaponType: "Rifle",
            Rarity:     "Mil-Spec",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwT09K5g4-Kgvr1P7rDqWZU7Mxkh6fA9Nimig3nqEs5Nj3xIYfHdlM8Z1rZ_FK-wL281pDq7pnPyCdgvT5iuyjUNDsC-g",
            MinValue:   5.00,
            MaxValue:   8.00,
        },

        // RARE SKINS ($10.00 - $25.00)
        {
            ID:         uuid.New(),
            Name:       "AWP | Graphite",
            WeaponType: "Sniper Rifle",
            Rarity:     "Restricted",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJD_9W7m5a0mvLwOq7c2GdQ-sJ0teXI8oThxgzs-kdvYDr6LYaSdgM_ZQzT-Vm-wOu-hJXu7p_ImCBguyUl5CqJzRGy1B9FcKUx0hDvdZpG",
            MinValue:   10.00,
            MaxValue:   20.00,
        },
        {
            ID:         uuid.New(),
            Name:       "Desert Eagle | Hypnotic",
            WeaponType: "Pistol",
            Rarity:     "Restricted",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgposr-kLAtl7PLZTjlH_9mkgIWKkPvxDLDEm2JS4Mp1mOjG-LP5gVO8v11uaz2gd9fBdABqMl7U-le3k-7r1sXvvcjNyCdhvnUm7SyOmBPj00webbI",
            MinValue:   12.00,
            MaxValue:   25.00,
        },

        // VERY RARE SKINS ($30.00 - $75.00)
        {
            ID:         uuid.New(),
            Name:       "AK-47 | Redline",
            WeaponType: "Rifle",
            Rarity:     "Classified",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot7HxfDhjxszJemkV09-5lpKKqPrxN7LEm1Rd6dd2j6eQ9N2t2wK3-EprZW-mddeTcQBrNV_Y-VS7weq91p-1tZvOmiY3viUrsHmPnxe2hBwabekx0uveFwuEfB5IMG4",
            MinValue:   30.00,
            MaxValue:   60.00,
        },
        {
            ID:         uuid.New(),
            Name:       "M4A1-S | Hyper Beast",
            WeaponType: "Rifle",
            Rarity:     "Classified",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhz2v_Nfz5H_uO1gb-Gw_alIITSgn1u-_p9g-7J4cKkiQWx-hE4YGuncYTEd1VqM13Y_lW4we7xxcjr7Qoc4Fo",
            MinValue:   40.00,
            MaxValue:   75.00,
        },

        // EXTREMELY RARE SKINS ($100.00 - $300.00)
        {
            ID:         uuid.New(),
            Name:       "AWP | Asiimov",
            WeaponType: "Sniper Rifle",
            Rarity:     "Covert",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJU7c6kgoWYkuHxPYTTl29u_hVwibqX84-tigHl_0FpYmigcI6WIQU6aAnX-gDqw-fphJ65vsvAn3YwsyYktWGdwUIMvIhqhg",
            MinValue:   100.00,
            MaxValue:   200.00,
        },
        {
            ID:         uuid.New(),
            Name:       "Karambit | Fade",
            WeaponType: "Knife",
            Rarity:     "Covert",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovbSsLQJf2PLacDBA5ciJlY20kfbkI7PYhG5u4MBwnPCPpdSs3lC3qUA-Ymn1II-VdQc4ZFDV-Vi2xO3ohZXuu5vMnCNh6CEm5nrYnUbk1gYMMLIAkLu0RA",
            MinValue:   200.00,
            MaxValue:   500.00,
        },

        // LEGENDARY SKINS ($500.00+)
        {
            ID:         uuid.New(),
            Name:       "M4A4 | Howl",
            WeaponType: "Rifle",
            Rarity:     "Rare Special",
            ImageURL:   "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwT09S5g4yCmfDLP7LWnn8f7cF02eiQ8Nr3jQbi_UQ4ZG-mJoOUIQQ-YQmF_Vjqx-7mgpC06pTIzSdmuyEjsy3D30vggEVZOGo",
            MinValue:   500.00,
            MaxValue:   1500.00,
        },
    }

    for _, s := range skins {
        if err := database.DB.Create(&s).Error; err != nil {
            log.Printf("❌ Failed to seed skin %s: %v", s.Name, err)
        } else {
            log.Printf("✅ Seeded skin: %s (%s) - $%.2f-$%.2f", s.Name, s.Rarity, s.MinValue, s.MaxValue)
        }
    }

    log.Println("🎨 Skins seeding complete!")
}