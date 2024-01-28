package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	name, repos, err := githubInfo("wrzolekLukasz")
	fmt.Println(name, repos, err)
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	resp, err := http.Get("https://api.github.com/users/" + login)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("error: %s", resp.Status)
	}
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.Public_Repos, nil
}

type Reply struct {
	Name         string
	Public_Repos int `json:"public_repos"` // !! json tag for the field name mapping -> public_repos being mapped to the specific field from json
}

/*  JSON in go
true/false -> bool
string -> string
number -> float64
null -> nil
array -> []interface{}
object -> map[string]interface{}, struct
*/
