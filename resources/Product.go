package resources

import (
	"log"
	"net/http"

	"github.com/midoblgsm/go-products/utils"
)

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      Rating `json:"rating"`
}

func CreateProduct(name string, description string) Product {
	return Product{Id: name, Name: name, Description: description}
}

func (p *Product) RateProduct(text string, stars int) error {
	return nil
}

func (p *Product) GetRating(ratingURL string) error {
	ratingsURL := utils.FormatURL(ratingURL, p.Id)
	httpClient := &http.Client{}
	response, err := utils.HttpExecute(httpClient, "GET", ratingsURL, nil)
	if err != nil {
		log.Printf(err.Error())
	}

	err = utils.UnmarshalResponse(response, &p.Rating)
	if err != nil {
		return err
	}

	return nil
}

type Rating struct {
	Id    string `json:"id"`
	Stars int    `json:"stars"`
	Color string `json:"color"`
}
