package whatapi

import (
	"regexp"
	"strconv"
	"strings"
)

type Torrent struct {
	Group   GroupType   `json:"group"`
	Torrent TorrentType `json:"torrent"`
}

type GroupType struct {
	WikiBody        string `json:"wikiBody"`
	WikiImage       string `json:"wikiImage"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Year            int    `json:"year"`
	RecordLabel     string `json:"recordLabel"`
	CatalogueNumber string `json:"catalogueNumber"`
	ReleaseType     int    `json:"releaseType"`
	CategoryID      int    `json:"caregoryId"`
	CategoryName    string `json:"categoryName"`
	Time            string `json:"time"`
	VanityHouse     bool   `json:"vanityHouse"`
	MusicInfo       struct {
		Composers []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"composers"`
		DJ []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"dj"`
		Artists []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		With []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"with"`
		Conductor []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"conductor"`
		RemixedBy []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"remixedBy"`
		Producer []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"producer"`
	} `json:"musicInfo"`
	Tags []string `json:"tags"`
}

type TorrentType struct {
	ID                      int    `json:"id"`
	Media                   string `json:"media"`
	Format                  string `json:"format"`
	Encoding                string `json:"encoding"`
	Remastered              bool   `json:"remastered"`
	RemasterYear            int    `json:"remasterYear"`
	RemasterTitle           string `json:"remasterTitle"`
	RemasterRecordLabel     string `json:"remasterRecordLabel"`
	RemasterCatalogueNumber string `json:"remasterCatalogueNumber"`
	Scene                   bool   `json:"scene"`
	HasLog                  bool   `json:"hasLog"`
	HasCue                  bool   `json:"hasCue"`
	LogScore                int    `json:"logScore"`
	FileCount               int    `json:"fileCount"`
	Size                    int    `json:"size"`
	Seeders                 int    `json:"seeders"`
	Leechers                int    `json:"leechers"`
	Snatched                int    `json:"snatched"`
	FreeTorrent             bool   `json:"freeTorrent"`
	Time                    string `json:"time"`
	Description             string `json:"description"`
	FileList                string `json:"fileList"`
	FilePath                string `json:"filePath"`
	UserID                  int    `json:"userID"`
	Username                string `json:"username"`
}

// FileStruct represents what we know about the files in a torrent
type FileStruct struct {
	Name string
	Size int64
}

// Files returns a slice of FileStruts for a torrent
func (t TorrentType) Files() ([]FileStruct, error) {
	f := []FileStruct{}
	re := regexp.MustCompile("(.*){{{(.*)}}}")
	for _, s := range strings.Split(t.FileList, "|||") {
		m := re.FindStringSubmatch(s)
		i, err := strconv.ParseInt(m[2], 10, 64)
		if err != nil {
			return nil, err
		}
		f = append(f, FileStruct{m[1], i})
	}
	return f, nil
}
