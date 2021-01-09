package offers

import (
	"log"

	database "github.com/advenjourney/api/internal/pkg/db/mysql"
	"github.com/advenjourney/api/internal/users"
)

// #1
type Offer struct {
	ID            string
	Title         string
	Location      string
	Description   string
	TitleImageURL string
	User          *users.User
}

//#2
func (offer Offer) Save() int64 {
	//#3
	stmt, err := database.DB.Prepare("INSERT INTO Offers(Title,Location,Description,TitleImageURL, UserID) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(offer.Title, offer.Location, offer.Description, offer.TitleImageURL)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")

	return id
}

func GetAll() []Offer {
	stmt, err := database.DB.Prepare("select O.id, O.title, O.location, O.description, O.titleimageurl, O.UserID, U.Username from Offers O inner join Users U on O.UserID = U.ID")

	if err != nil {
		// TODO: handle this more gracefully
		log.Panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		// TODO: handle this more gracefully
		log.Panic(err)
	}
	defer rows.Close()

	var offers []Offer
	var id, username string
	for rows.Next() {
		var offer Offer
		err := rows.Scan(&offer.ID, &offer.Title, &offer.Location, &offer.Description, &offer.TitleImageURL, &id, &username)
		if err != nil {
			// TODO: should not panic here yet
			log.Panic(err)
		}
		offer.User = &users.User{
			ID:       id,
			Username: username,
		}
		offers = append(offers, offer)
	}
	if err = rows.Err(); err != nil {
		// TODO: needs proper error handling // logging to not exit on fatal error
		log.Panic(err)
	}

	return offers
}
