package main

type repository struct {
	Name        string
	Url         string
	Description string
	Homepage    string
}

type author struct {
	Name  string
	Email string
}

type commit struct {
	Id        string
	Message   string
	Timestamp string
	Url       string
	Author    author
}

type HookData struct {
	Before            string
	After             string
	Ref               string
	UserId            int    `json:user_id`
	UserName          string `json:user_name`
	ProjectId         int    `json:project_id`
	Repository        repository
	Commits           []commit `json:commits`
	TotalCommitsCount int      `json:total_commits_count`
}
