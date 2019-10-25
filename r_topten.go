package whatapi

type TopTenTags []struct {
	Caption string `json:"caption"`
	Tag     string `json:"tag"`
	Limit   int    `json:"limit"`
	Results []struct {
		Name     string `json:"name"`
		Uses     int    `json:"uses"`
		PosVotes int    `json:"posVotes"`
		NegVotes int    `json:"negVotes"`
	} `json:"results"`
}

type TopTenResult struct {
	TorrentID     int      `json:"torrentId"`
	GroupID       int      `json:"groupId"`
	Artist        string   `json:"artist"`
	GroupName     string   `json:"groupName"`
	GroupCategory int      `json:"groupCategory"`
	GroupYear     int      `json:"groupYear"`
	RemasterTitle string   `json:"remasterTitle"`
	Format        string   `json:"format"`
	Encoding      string   `json:"encoding"`
	HasLog        bool     `json:"hasLog"`
	HasCue        bool     `json:"hasCue"`
	HasLogDB      bool     `json:"hasLogDB"`
	LogScore      int      `json:"logsScore"`
	LogChecksum   string   `json:"logChecksum"`
	Media         string   `json:"media"`
	Scene         bool     `json:"scene"`
	Year          int      `json:"year"`
	Tags          []string `json:"tags"`
	Snatched      int      `json:"snatched"`
	Seeders       int      `json:"seeders"`
	Leechers      int      `json:"leechers"`
	Data          int64    `json:"data"`
	Size          int64    `json:"size"`
	WikiImage     string   `json:"wikiImage"`
	ReleaseType   string   `json:"releaseType"`
}

type TopTenTorrents []struct {
	Caption string         `json:"caption"`
	Tag     string         `json:"tag"`
	Limit   int            `json:"limit"`
	Results []TopTenResult `json:"results"`
}

type TopTenUsers []struct {
	Caption string `json:"caption"`
	Tag     string `json:"tag"`
	Limit   int    `json:"limit"`
	Results []struct {
		ID         int     `json:"id"`
		Username   string  `json:"username"`
		Uploaded   float64 `json:"uploaded"`
		UpSpeed    float64 `json:"upSpeed"`
		Downloaded float64 `json:"downloaded"`
		DownSpeed  float64 `json:"downSpeed"`
		NumUploads int     `json:"numUploads"`
		JoinDate   string  `json:"joinDate"`
	} `json:"results"`
}
