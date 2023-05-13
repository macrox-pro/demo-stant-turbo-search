package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/legion-zver/premier-one-bleve-search/internal/graphql/graph/model"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, query string) ([]*model.SearchResultObject, error) {
	resp, err := r.SearchEngine.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	result := make([]*model.SearchResultObject, len(resp.Hits), len(resp.Hits))
	for i, hit := range resp.Hits {
		obj := &model.SearchResultObject{
			ID:    hit.ID,
			Score: hit.Score,
		}
		if mapstructure.Decode(hit.Fields, obj) == nil {
			if obj.Slug != nil {
				url := fmt.Sprintf("https://premier.one/show/%s", *obj.Slug)
				obj.URL = &url
			}
		}
		result[i] = obj
	}
	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
