package assets

type AssetType string
type FileType string

const (
	Avatar AssetType = "avatar"
	Cover  AssetType = "cover"

	Catalog AssetType = "content"
)

const (
	File  FileType = "files"
	Image FileType = "images"
)
