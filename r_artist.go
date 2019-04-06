package whatapi

import (
	"fmt"
	"strconv"
)

type ArtistGroupStruct struct {
	GroupID              int             `json:"groupId"`
	GroupYear            int             `json:"groupYear"`
	GroupRecordLabel     string          `json:"groupRecordLabel"`
	GroupCatalogueNumber string          `json:"groupCatalogueNumber"`
	TagsF                []string        `json:"tags"`
	ReleaseTypeF         int             `json:"releaseType"`
	GroupVanityHouse     bool            `json:"groupVanityHouse"`
	HasBookmarked        bool            `json:"hasBookmarked"`
	Torrent              []TorrentStruct `json:"torrent"`
	GroupName            string          `json:"groupName"`
	ArtistsF             []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		AliasID int    `json:"aliasid"`
	} `json:"artists"`
	ExtendedArtists map[string][]struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		AliasID int    `json:"aliasid"`
	} `json:"extendedArtists"`
	artists    []string
	importance []int
}

type Artist struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
	HasBookmarked        bool   `json:"hasBookmarked"`
	Image                string `json:"image"`
	Body                 string `json:"body"`
	VanityHouse          bool   `json:"vanityHouse"`
	Tags                 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"tags"`
	SimilarArtists []struct {
		ArtistID  int    `json:"artistId"`
		Name      string `json:"name"`
		Score     int    `json:"score"`
		SimilarID int    `json:"similarId"`
	} `json:"similarArtists"`
	Statistics struct {
		NumGroups   int `json:"numGroups"`
		NumTorrents int `json:"numTorrents"`
		NumSeeders  int `json:"numSeeders"`
		NumLeechers int `json:"numLeechers"`
		NumSnatches int `json:"numSnatches"`
	} `json:"statistics"`
	TorrentGroup []ArtistGroupStruct `json:"torrentgroup"`
	Requests     []struct {
		RequestID  int    `json:"requestId"`
		CategoryID int    `json:"categoryId"`
		Title      string `json:"title"`
		Year       int    `json:"year"`
		TimeAdded  string `json:"timeAdded"`
		Votes      int    `json:"votes"`
		Bounty     int64  `json:"bounty"`
	} `json:"requests"`
}

func (g ArtistGroupStruct) ID() int {
	return g.GroupID
}

func (g ArtistGroupStruct) Name() string {
	return g.GroupName
}

func (g ArtistGroupStruct) Artist() string {
	// this is embarrassing, we're embedded in an Artist object
	// we should be able to produce the Artist name
	return g.ArtistsF[0].Name
}

func (g ArtistGroupStruct) Year() int {
	return g.GroupYear
}

func (g ArtistGroupStruct) WikiImage() (string, error) {
	return "", fmt.Errorf("Artist Group does not contain an image")
}

func (g ArtistGroupStruct) makeArtistsImportance() error {
	g.artists = make([]string, 5)
	g.importance = make([]int, 5)
	for i, e := range g.ExtendedArtists {
		v, err := strconv.Atoi(i)
		if err != nil {
			return err
		}
		for _, a := range e {
			g.artists = append(g.artists, a.Name)
			g.importance = append(g.importance, v)
		}
	}
	return nil
}

func (g ArtistGroupStruct) Artists() ([]string, error) {
	if g.artists == nil {
		if err := g.makeArtistsImportance(); err != nil {
			return []string{}, err
		}
	}
	return g.artists, nil
}

func (g ArtistGroupStruct) Importance() ([]int, error) {
	if g.importance == nil {
		if err := g.makeArtistsImportance(); err != nil {
			return []int{}, err
		}
	}
	return g.importance, nil
}

func (g ArtistGroupStruct) RecordLabel() string {
	return g.GroupRecordLabel
}

func (g ArtistGroupStruct) CatalogueNumber() string {
	return g.GroupCatalogueNumber
}

func (g ArtistGroupStruct) ReleaseType() int {
	return g.ReleaseTypeF
}

func (g ArtistGroupStruct) Tags() []string {
	return g.TagsF
}

func (g ArtistGroupStruct) WikiBody() (string, error) {
	return "", fmt.Errorf("Artist Group does not contain a WikiBody")
}

func (g ArtistGroupStruct) String() string {
	return GroupString(g)
}
