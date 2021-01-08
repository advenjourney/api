package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen/graph/generated"
	"gqlgen/graph/model"
)

func (r *mutationResolver) CreateOffer(ctx context.Context, input model.NewOffer) (*model.Offer, error) {
	var offer model.Offer
	var user model.User
	offer.Description = input.Description
	offer.Location = input.Location
	offer.TitleImageURL = input.TitleImageURL
	offer.Title = input.Title
	user.Name = "test-user"
	offer.User = &user
	return &offer, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Offers(ctx context.Context) ([]*model.Offer, error) {
	var offers []*model.Offer
	dummyOffer := model.Offer{
		Title:         "Coworking in Thalasso Villa del Conde",
		Location:      "Gran Canaria",
		Description:   "Between the ocean and the pool of the fantastic Lopesan Villa del Conde Resort in Meloneras, south Gran Canaria, is the hotel’s spectacular spa. With a refined, linear design and ultramodern interiors the Corallium Thalasso Villa del Conde offers fantasy experiences in its Ocean View Suites. These private suites offer a solarium and seawater pools and are the perfect setting for an Oceanic Paradise experience; an hour-long jacuzzi with spectacular sea views and French Moët & Chandon champagne.",
		TitleImageURL: "https://www.hellocanaryislands.com/sites/default/files/resource/thalasso_spa_coralium-gran_canaria_0.jpg",
		User:          &model.User{Name: "admin"},
	}
	offers = append(offers, &dummyOffer)
	return offers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
