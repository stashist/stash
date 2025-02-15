package api

import (
	"context"

	"github.com/stashapp/stash/pkg/api/urlbuilders"
	"github.com/stashapp/stash/pkg/models"
)

func (r *studioResolver) Name(ctx context.Context, obj *models.Studio) (string, error) {
	if obj.Name.Valid {
		return obj.Name.String, nil
	}
	panic("null name") // TODO make name required
}

func (r *studioResolver) URL(ctx context.Context, obj *models.Studio) (*string, error) {
	if obj.URL.Valid {
		return &obj.URL.String, nil
	}
	return nil, nil
}

func (r *studioResolver) ImagePath(ctx context.Context, obj *models.Studio) (*string, error) {
	baseURL, _ := ctx.Value(BaseURLCtxKey).(string)
	imagePath := urlbuilders.NewStudioURLBuilder(baseURL, obj).GetStudioImageURL()

	var hasImage bool
	if err := r.withReadTxn(ctx, func(repo models.ReaderRepository) error {
		var err error
		hasImage, err = repo.Studio().HasImage(obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	// indicate that image is missing by setting default query param to true
	if !hasImage {
		imagePath = imagePath + "?default=true"
	}

	return &imagePath, nil
}

func (r *studioResolver) SceneCount(ctx context.Context, obj *models.Studio) (ret *int, err error) {
	var res int
	if err := r.withReadTxn(ctx, func(repo models.ReaderRepository) error {
		res, err = repo.Scene().CountByStudioID(obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return &res, err
}

func (r *studioResolver) ParentStudio(ctx context.Context, obj *models.Studio) (ret *models.Studio, err error) {
	if !obj.ParentID.Valid {
		return nil, nil
	}

	if err := r.withReadTxn(ctx, func(repo models.ReaderRepository) error {
		ret, err = repo.Studio().Find(int(obj.ParentID.Int64))
		return err
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *studioResolver) ChildStudios(ctx context.Context, obj *models.Studio) (ret []*models.Studio, err error) {
	if err := r.withReadTxn(ctx, func(repo models.ReaderRepository) error {
		ret, err = repo.Studio().FindChildren(obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *studioResolver) StashIds(ctx context.Context, obj *models.Studio) (ret []*models.StashID, err error) {
	if err := r.withReadTxn(ctx, func(repo models.ReaderRepository) error {
		ret, err = repo.Studio().GetStashIDs(obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return ret, nil
}
