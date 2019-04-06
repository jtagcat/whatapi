package whatapi

type TorrentGroup struct {
	Group   GroupStruct     `json:"group"`
	Torrent []TorrentStruct `json:"torrents"`
}
