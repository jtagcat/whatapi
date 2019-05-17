package whatapi

import (
	"fmt"
	"html"
	"os"
	"strconv"
)

type ArtistGroupStruct struct {
	GroupID              int                   `json:"groupId"`
	GroupYear            int                   `json:"groupYear"`
	GroupRecordLabel     string                `json:"groupRecordLabel"`
	GroupCatalogueNumber string                `json:"groupCatalogueNumber"`
	TagsF                []string              `json:"tags"`
	ReleaseTypeF         int                   `json:"releaseType"`
	GroupVanityHouse     bool                  `json:"groupVanityHouse"`
	HasBookmarked        bool                  `json:"hasBookmarked"`
	Torrent              []ArtistTorrentStruct `json:"torrent"`
	GroupName            string                `json:"groupName"`
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
	NameF                string `json:"name"`
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

func (a Artist) Name() string {
	return html.UnescapeString(a.NameF)
}

func (g ArtistGroupStruct) ID() int {
	return g.GroupID
}

func (g ArtistGroupStruct) Name() string {
	return html.UnescapeString(g.GroupName)
}

func (g ArtistGroupStruct) Artist() string {
	// this is embarrassing, we're embedded in an Artist object
	// we should be able to produce the Artist name
	fmt.Fprint(os.Stderr, "ArtistGroupStruct.Artist() not implemented")
	os.Exit(-1)
	if len(g.ArtistsF) > 0 {
		return g.ArtistsF[0].Name
	}
	for i := 1; i <= 7; i++ {
		s := strconv.Itoa(i)
		if len(g.ExtendedArtists[s]) > 0 {
			return g.ExtendedArtists[s][0].Name
		}
	}
	return " (Unknown Artist) "
}

func (g ArtistGroupStruct) Year() int {
	return g.GroupYear
}

func (g *ArtistGroupStruct) makeArtistsImportance() {
	g.artists = make([]string, 0, 7)
	g.importance = make([]int, 0, 7)
	for i := 1; i <= 7; i++ {
		for _, a := range g.ExtendedArtists[strconv.Itoa(i)] {
			g.artists = append(g.artists, a.Name)
			g.importance = append(g.importance, i)
		}
	}
}

func (g ArtistGroupStruct) Artists() []string {
	if g.artists == nil {
		g.makeArtistsImportance()
	}
	return g.artists
}

func (g ArtistGroupStruct) Importance() []int {
	if g.importance == nil {
		g.makeArtistsImportance()
	}
	return g.importance
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

func (g ArtistGroupStruct) String() string {
	return GroupString(g)
}
