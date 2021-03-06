package domain

import "github.com/pajarraco93/graphql-test/pkg/library/domain/entities"

type Repository interface {
	CreateGroup(entities.Group) error
	CreateAlbum(entities.Album) error
	CreateSong(entities.Song) error

	AllGroups() ([]entities.Group, error)
	AllAlbums() ([]entities.Album, error)
	AllSongs() ([]entities.Song, error)

	GetGroupsByIDs([]int) ([]entities.Group, error)
	GetAlbumsByIDs([]int) ([]entities.Album, error)

	GetAlbumsByGroupID(int) ([]entities.Album, error)
	GetSongsByAlbumID(int) ([]entities.Song, error)
}

type InfoRepo interface {
	GetGroupInfo(entities.Group) (string, error)
	GetAlbumInfo(entities.Album) (string, error)
	GetSongInfo(entities.Song) (string, error)
}
