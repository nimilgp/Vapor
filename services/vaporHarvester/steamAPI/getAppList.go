package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const GetAppListEndpoint = "http://api.steampowered.com/ISteamApps/GetAppList/v2/"

// App represents a single game entry.
type App struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}

// AppList contains the array of all App structs.
type AppList struct {
	Apps []App `json:"apps"`
}

// Response is the top-level structure of the entire JSON payload.
type Response struct {
	Applist AppList `json:"applist"`
}

func GetAppList() {
	fmt.Println("Querying Steam Web API for the complete app list...")
	resp, err := http.Get(GetAppListEndpoint)
	if err != nil {
		log.Fatalf("Failed to get app list: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK status code: %d", resp.StatusCode)
	}

	// 2. Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 3. Unmarshal the JSON into the defined Response struct
	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Printf("Successfully fetched %d total Steam apps. Displaying all:\n\n", len(data.Applist.Apps))

	// 4. Iterate and print the required fields (Appid and Name)
	for _, app := range data.Applist.Apps {
		fmt.Printf("App ID: %d, Name: %s\n", app.Appid, app.Name)
	}
}
