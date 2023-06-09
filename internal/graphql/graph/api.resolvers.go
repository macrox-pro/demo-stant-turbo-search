package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/legion-zver/vss-brain-search/internal/graphql/graph/model"
	"github.com/legion-zver/vss-brain-search/internal/search"
	"github.com/mitchellh/mapstructure"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, query string, where *model.SearchWhereInput, useNlp *bool) (*model.SearchResponse, error) {
	var searchWhere *search.Where
	if where != nil {
		searchWhere = &search.Where{
			Active:  where.Active,
			Service: where.Service,
		}
	}
	result, nlpResult, err := r.SearchEngine.Search(ctx, query, searchWhere, useNlp)
	if err != nil {
		return nil, err
	}
	resp := &model.SearchResponse{
		Documents: make([]*model.IndexObject, len(result.Hits)),
		Metadata: &model.SearchResponseMetadata{
			Query: query,
		},
	}
	if nlpResult != nil {
		if nlpResult.Intent != nil {
			resp.Metadata.Intent = &model.SearchIntent{
				Name:       nlpResult.Intent.Name,
				Confidence: float64(nlpResult.Intent.Confidence),
			}
		}
		if len(nlpResult.Entities) > 0 {
			resp.Metadata.Entities = make([]*model.SearchEntity, len(nlpResult.Entities))
			for i, entity := range nlpResult.Entities {
				searchEntity := &model.SearchEntity{
					Start: int(entity.Start),
					Value: entity.Value,
					Type:  entity.Type,
					End:   int(entity.End),
				}
				if len(entity.NormalValue) > 0 {
					searchEntity.NormalValue = &entity.NormalValue
				}
				resp.Metadata.Entities[i] = searchEntity
			}
		}
	}
	for i, hit := range result.Hits {
		obj := &model.IndexObject{
			ID:    hit.ID,
			Score: hit.Score,
		}
		_ = mapstructure.WeakDecode(hit.Fields, obj)
		switch obj.Service {
		case "premier.one":
			var url string
			if obj.Slug != nil {
				url = fmt.Sprintf("https://premier.one/show/%s", *obj.Slug)
			} else {
				url = fmt.Sprintf("https://premier.one/show/%s", obj.ID)
			}
			obj.URL = &url
		}
		resp.Documents[i] = obj
	}
	return resp, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
