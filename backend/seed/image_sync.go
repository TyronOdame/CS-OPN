package seed

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/models"
)

const steamAPISImageCatalogURL = "https://api.steamapis.com/image/items/730"

var preferredWearOrder = []string{
	"Factory New",
	"Minimal Wear",
	"Field-Tested",
	"Well-Worn",
	"Battle-Scarred",
}

// SyncImageURLs backfills broken/missing case and skin image URLs.
// It resolves names against SteamApis' CS2 image catalog.
func SyncImageURLs() {
	log.Println("🖼️  Syncing image URLs from SteamApis...")

	imageMap, err := fetchSteamImageCatalog()
	if err != nil {
		log.Printf("⚠️  Image sync skipped (catalog fetch failed): %v", err)
		return
	}

	skinsUpdated := syncSkinImages(imageMap)
	casesUpdated := syncCaseImages(imageMap)

	log.Printf("🖼️  Image sync complete. skins=%d, cases=%d", skinsUpdated, casesUpdated)
}

func fetchSteamImageCatalog() (map[string]string, error) {
	client := &http.Client{Timeout: 12 * time.Second}
	resp, err := client.Get(steamAPISImageCatalogURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var imageMap map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&imageMap); err != nil {
		return nil, err
	}
	return imageMap, nil
}

func syncSkinImages(imageMap map[string]string) int {
	var skins []models.Skin
	if err := database.DB.Find(&skins).Error; err != nil {
		log.Printf("⚠️  Failed loading skins for image sync: %v", err)
		return 0
	}

	updated := 0
	for _, skin := range skins {
		if !shouldReplaceImageURL(skin.ImageURL) {
			continue
		}

		newURL := resolveImageByName(imageMap, skin.Name)
		if newURL == "" {
			continue
		}

		if err := database.DB.Model(&skin).Update("image_url", newURL).Error; err != nil {
			log.Printf("⚠️  Failed updating skin image (%s): %v", skin.Name, err)
			continue
		}
		updated++
	}

	return updated
}

func syncCaseImages(imageMap map[string]string) int {
	var cases []models.Case
	if err := database.DB.Find(&cases).Error; err != nil {
		log.Printf("⚠️  Failed loading cases for image sync: %v", err)
		return 0
	}

	updated := 0
	for _, caseItem := range cases {
		if !shouldReplaceImageURL(caseItem.ImageURL) {
			continue
		}

		newURL := resolveImageByName(imageMap, caseItem.Name)
		if newURL == "" {
			continue
		}

		if err := database.DB.Model(&caseItem).Update("image_url", newURL).Error; err != nil {
			log.Printf("⚠️  Failed updating case image (%s): %v", caseItem.Name, err)
			continue
		}
		updated++
	}

	return updated
}

func shouldReplaceImageURL(url string) bool {
	if strings.TrimSpace(url) == "" {
		return true
	}
	lowerURL := strings.ToLower(url)
	return strings.Contains(lowerURL, "community.cloudflare.steamstatic.com")
}

func resolveImageByName(imageMap map[string]string, name string) string {
	if exact := imageMap[name]; exact != "" {
		return exact
	}

	// Try common wear variants for skins stored without wear suffix.
	for _, wear := range preferredWearOrder {
		candidate := name + " (" + wear + ")"
		if imageMap[candidate] != "" {
			return imageMap[candidate]
		}
	}

	// Fallback: any catalog entry that starts with "<name> (".
	prefix := name + " ("
	for marketHashName, imageURL := range imageMap {
		if strings.HasPrefix(marketHashName, prefix) {
			return imageURL
		}
	}

	return ""
}
