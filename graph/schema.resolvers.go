package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/advenjourney/api/graph/generated"
	"github.com/advenjourney/api/graph/model"
	"github.com/advenjourney/api/internal/auth"
	"github.com/advenjourney/api/internal/offers"
	"github.com/advenjourney/api/internal/users"
	"github.com/advenjourney/api/pkg/jwt"
)

func (r *mutationResolver) CreateOffer(ctx context.Context, input model.NewOffer) (*model.Offer, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Offer{}, fmt.Errorf("access denied")
	}

	var offer offers.Offer
	offer.Description = input.Description
	offer.Location = input.Location
	offer.TitleImageURL = input.TitleImageURL
	offer.Title = input.Title
	offer.User = user
	offerID := offer.Save()
	grahpqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}

	return &model.Offer{ID: strconv.FormatInt(offerID, 10), Title: offer.Title, Location: offer.Location, Description: offer.Description, TitleImageURL: offer.TitleImageURL, User: grahpqlUser}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Offers(ctx context.Context) ([]*model.Offer, error) {
	var resultOffers []*model.Offer
	dbOffers := offers.GetAll()
	for _, offer := range dbOffers {
		// grahpqlUser := &model.User{
		//	 ID:   link.User.ID,
		//	 Name: link.User.Username,
		// }
		resultOffers = append(resultOffers, &model.Offer{ID: offer.ID, Title: offer.Title, Location: offer.Location, Description: offer.Description, TitleImageURL: offer.TitleImageURL})
	}

	return resultOffers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
