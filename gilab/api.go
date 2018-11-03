package gilab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//GetRepos call gitlab api for get repository list
func GetRepos(host, accessToken string, page int) ([]Repoinfo, error) {

	url := fmt.Sprintf("%s/api/v4/projects?private_token=%s&per_page=100&page=%d", host, accessToken, page)
	bd, err := urlLoader(url)

	resp := []Repoinfo{}
	if err := json.Unmarshal(bd, &resp); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return resp, err
}

//GetBranches call gitlab api for get repository branches
func GetBranches(host, accessToken string, id int) ([]BranchInfo, error) {

	url := fmt.Sprintf("%s/api/v4/projects/%d/repository/branches?private_token=%s", host, id, accessToken)
	bd, err := urlLoader(url)

	resp := []BranchInfo{}
	if err := json.Unmarshal(bd, &resp); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return resp, err
}

func urlLoader(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return contents, nil
}
