package matcher

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"../search"
)

type (
	//item defines the fields associated with the item tag in the res document
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	//Image defines the fields associated with the image tag in the rss document.

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	//channel defines the fields associated with the channel tag in the rss document.

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubData        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"WebMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	//rssDocument defines the fields associated with the rss document
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

//rssMAtcher implements the Matcher interface
type rssMatcher struct{}

//init registers the matcher with the program

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

//retrieves performs a HTTP request for thr rss feed and decodes

func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}
	//retrieve the rss feed from document from the web
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	//close the response once we return from the function
	defer resp.Body.Close()

	//Check the status code for a 200 so we know we have received a proper response.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)

	}
	//Decode the rss feed document into our struct type.
	//We dont need to check the errors, the caller can do this
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)

	return &document, err

}

//Search looks at the document for the specified search term
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", feed.Type, feed.Name, feed.URI)

	//retrieve the data to search

	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		//check the title for the search term
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)

		if err != nil {
			return nil, err
		}

		// if we found a match save the result
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		//check the description for the search term
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)

		if err != nil {
			return nil, err
		}
		//if we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
