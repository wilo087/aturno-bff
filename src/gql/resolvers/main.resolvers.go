package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/wilo087/aturno-bff/src/database/model"
	"github.com/wilo087/aturno-bff/src/gql/generated"
)

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var id int = 1
	var ij *int = &id
	user := []*model.User{
		{
			ID:        ij,
			Name:      "Wilowayne",
			LastName:  "De La Cruz",
			Email:     "wilo087@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
	}
	return user, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
