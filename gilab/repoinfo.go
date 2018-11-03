package gilab

//Repoinfo repository information struct
type Repoinfo struct {
	ID                int    `json:"id"`
	PathWithNamespace string `json:"path_with_namespace"`
	WebURL            string `json:"web_url"`
	Branch            string `json:"default_branch"`
	Links             struct {
		RepoBranches string `json:"repo_branches"`
	} `json:"_links"`
}

//BranchInfo repository branch info
type BranchInfo struct {
	Name   string `json:"name"`
	Commit struct {
		ID string `json:"id"`
	} `json:"commit"`
}
