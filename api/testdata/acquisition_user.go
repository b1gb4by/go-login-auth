package testdata

import (
	"api/domain/model"
)

var AcquisitionUser = map[string]*model.AcquisitionUser{

	"success_1": {
		ID:        1,
		FirstName: "sample",
		LastName:  "sample",
		Email:     "sample@example.com",
		Password:  []byte("$2a$14$dMGDex7fprpYwK1N.MuTFO1U1WQwnvIELXohTwP8I3B/UCcd87Kou"),
	},
}
