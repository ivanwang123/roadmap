package model

import (
	"fmt"
	"strings"
	"time"

	LinkPreview "github.com/Junzki/link-preview"
)

type Checkpoint struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Instructions string     `json:"instructions"`
	Status       StatusType `json:"status"`
	Links        Links      `json:"links"`
	RoadmapID    int        `json:"roadmap"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type StatusType Status

type Links []*Link
type Link struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (l *Links) Scan(src interface{}) error {
	linkStr := fmt.Sprintf("%v", src)
	linkArr := strings.Split(linkStr[1:len(linkStr)-1], ",")
	links := make([]*Link, len(linkArr))

	for i, link := range linkArr {
		preview, err := LinkPreview.Preview(link, nil)
		if err != nil {
			links[i] = &Link{
				URL:         link,
				Title:       "",
				Description: "",
				// TODO: Add broken link image
				Image: "",
			}
		} else {
			links[i] = &Link{
				URL:         link,
				Title:       strings.TrimSpace(preview.Title),
				Description: strings.TrimSpace(preview.Description),
				Image:       preview.ImageURL,
			}
		}
	}

	*l = Links(links)
	return nil
}

// TODO: Remove
// func (l Link) Value() (driver.Value, error) {
// 	fmt.Println("VALUE", l)
// 	return driver.Value("{}"), nil
// }
