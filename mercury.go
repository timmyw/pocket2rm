package pocket2rm

import (
	"net/url"
	"fmt"
)

// ArticleDetails contains details about a URL, usually from the
// Mercury API
type ArticleDetails struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	DatePublished string `json:"date_published"`
	LeadImageURL string `json:"lead_image_url"`
	Dek string `json:"dek"`
	URL string `json:"url"`
	Domain string `json:"domain"`
	Excerpt string `json:"excerpt"`
	WordCount int `json:"word_count"`
	Direction string `json:"direction"`
	TotalPages int `json:"total_pages"`
	RenderedPages int `json:"rendered_pages"`
}

// GetArticleDetails will use the Mercury API to retrieve details of
// the article, using the provided URL
func (p *Pocket2RM) GetArticleDetails(articleURL string, details *ArticleDetails) error {

	headers := map[string]string {
		"x-api-key" : p.Config.GetString("mercury_key"),
	}
	params := url.Values{
		"url": {articleURL},
	}
	path := fmt.Sprintf("/parser?%s", params.Encode())
	err := GetJSON(path,
		APIOriginMercury,
		headers,
		details)
	if err != nil {
		return err
	}

	return nil
}

