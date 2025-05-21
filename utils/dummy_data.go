package utils

import "my-first-server/models"

var Books = []models.Book{
	{
		ID: "1",
		Title: "Yaa Qaatay Burcadkeygii",
		Author: &models.Author{
			FirstName: "Shaafi Abdi",
			LastName:  "Abdi",
		},
	},
	{
		ID: "2",
		Title: "Dabin",
		Author: &models.Author{
			FirstName: "Ahmed",
			LastName:  "Mohamed",
		},
	},
	{
		ID: "3",
		Title: "Ninka Ugu Taajirsan Baabil",
		Author: &models.Author{
			FirstName: "Abyan",
			LastName:  "Ahmed",
		},
	},
}
