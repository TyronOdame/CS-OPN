package seed

import (
	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"log"

	"github.com/google/uuid"
)

// SeedCases adds test cases to the database
func SeedCases() {
	// check if cases already exist
	var count int64
	database.DB.Model(&models.Case{}).Count(&count)
	if count > 0 {
		log.Println("📦 Cases already seeded, skipping...")
		return 
	}


	cases := []models.Case{
        {
            ID:          uuid.New(),
            Name:        "Chroma Case",
            Description: "Contains colorful weapon skins from the Chroma Collection",
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRfX0DbRvCiwMbQVg8kdFAYsba0KQVv1fjGYTxD4eO0hoaEgbmmZ-KBx28Gu8N1i-qRotzw0FW3rxc4fSmtc5jXIFVoNQnSr1ftk-i6hpHu6pTImHI27Cl04XuOmhCy1xpPPuZxxavJH7RMaw",
            Price:       2.50,
        },
        {
            ID:          uuid.New(),
            Name:        "Gamma Case",
            Description: "Features skins from the Gamma Collection with vibrant colors",
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulReQ0HdQOqohZ-CBBVqJFdFubSaIAVv1fjGYTxD4eO0hoaEgbmmYrjVlWJT7cB0i-qX89il0ALi-kBqMWj3LIZGI1Q8ZA7V_FG-ku6915Dp6pTImHI27Sx0sHqLzBGzhgYfcKUx0r9M7f61",
            Price:       3.00,
        },
        {
            ID:          uuid.New(),
            Name:        "Spectrum Case",
            Description: "Rainbow-themed weapon finishes from the Spectrum Collection",
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulReQ0HdUuqohJzHTlVzJEhWv7SaKQZ53P3BTzRO5dvsw9PZlaPwYb3XlWFS_sAlijj89NWi3le1-Ec_ZWjxLNCcJlJvMl2Dq1i3we7tgpDu753MyHRnuXIgtyzVmRbk0k0ZZ-Rxxavs-lE9MA",
            Price:       2.75,
        },
        {
            ID:          uuid.New(),
            Name:        "Prisma Case",
            Description: "Exclusive weapon skins with prismatic designs",
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulReQ0HdUuqohJzHTlVzJEhWubSaKQZ53P3BTzxD4N6zkYGJmOXLPr7Yml4FsMAij7qXp9_wiVHn-xBkZzqmd46LMlhpAXSG_Vy-x-q50JTvu5XMmiBh6HMq4y3eyRWw0k8abLJxxavRJSj50A",
            Price:       4.00,
        },
        {
            ID:          uuid.New(),
            Name:        "Danger Zone Case",
            Description: "High-risk, high-reward skins from the Danger Zone",
            ImageURL:    "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulReQ0HdUuqohJzHTlVzJEhWv7SaKQZu1PfAYjxD4N6zkYGJmOXLPr7Yml4FsJNz2rvEpNil0VHm_0ZkN2D1d4LEJw45YFrSq1C6366x0sKQS5oA",
            Price:       5.00,
        },
    }

	for _, c := range cases {
		if err := database.DB.Create(&c).Error; err != nil {
			log.Fatalf("❌ Failed to seed cases %s: %v", c.Name, err)
		} else {
			log.Printf("✅ Seeded case: %s ($%.2f)", c.Name, c.Price)
		}
	}

	log.Println("📦 Cases seeding completed.")
}