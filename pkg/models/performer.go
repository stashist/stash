package models

type PerformerReader interface {
	Find(id int) (*Performer, error)
	FindMany(ids []int) ([]*Performer, error)
	FindBySceneID(sceneID int) ([]*Performer, error)
	FindNamesBySceneID(sceneID int) ([]*Performer, error)
	FindByImageID(imageID int) ([]*Performer, error)
	FindByGalleryID(galleryID int) ([]*Performer, error)
	FindByNames(names []string, nocase bool) ([]*Performer, error)
	CountByTagID(tagID int) (int, error)
	Count() (int, error)
	All() ([]*Performer, error)
	Query(performerFilter *PerformerFilterType, findFilter *FindFilterType) ([]*Performer, int, error)
	GetImage(performerID int) ([]byte, error)
	GetStashIDs(performerID int) ([]*StashID, error)
	GetTagIDs(sceneID int) ([]int, error)
}

type PerformerWriter interface {
	Create(newPerformer Performer) (*Performer, error)
	Update(updatedPerformer PerformerPartial) (*Performer, error)
	UpdateFull(updatedPerformer Performer) (*Performer, error)
	Destroy(id int) error
	UpdateImage(performerID int, image []byte) error
	DestroyImage(performerID int) error
	UpdateStashIDs(performerID int, stashIDs []StashID) error
	UpdateTags(sceneID int, tagIDs []int) error
}

type PerformerReaderWriter interface {
	PerformerReader
	PerformerWriter
}
