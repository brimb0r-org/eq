package excel

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"os"
)

type ExcelConfig struct {
	Apikey      string `yaml:"apiKey"`
	AccessToken string `yaml:"apiToken"`
}

func (c *ExcelConfig) getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := c.tokenFromSecret()
	if err != nil {
		fmt.Println(err)
		tok = c.getTokenFromWeb(config)
		c.saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func (c *ExcelConfig) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func (c *ExcelConfig) tokenFromSecret() (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.Unmarshal([]byte(c.AccessToken), tok)
	return tok, err
}

// Saves a token to a file path.
func (c *ExcelConfig) saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (c *ExcelConfig) NewClient() *sheets.Service {
	ctx := context.Background()
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON([]byte(c.Apikey), "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := c.getClient(config)

	clientSrv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return clientSrv
}
