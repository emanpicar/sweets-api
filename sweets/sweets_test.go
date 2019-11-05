package sweets

import (
	"errors"
	"reflect"
	"testing"

	"github.com/emanpicar/sweets-api/db/entities"
)

type mockDBManager struct{}

func (m mockDBManager) BatchFirstOrCreate(data *[]entities.SweetsCollection) {}
func (m mockDBManager) Insert(data *entities.SweetsCollection) error {
	return nil
}
func (m mockDBManager) UpdateByProductID(pID string, data *entities.SweetsCollection) error {
	return nil
}
func (m mockDBManager) DeleteByProductID(pID string) (string, error) {
	if pID == "000" {
		return "", errors.New("delete failed")
	}

	return "delete success", nil
}
func (m mockDBManager) GetSweetByID(pID string) (*entities.SweetsCollection, error) {
	return &entities.SweetsCollection{}, nil
}
func (m mockDBManager) GetSweetsCollection() *[]entities.SweetsCollection {
	return &[]entities.SweetsCollection{
		entities.SweetsCollection{
			Name:        "CreamIce",
			ImageClosed: "cc.png",
			ImageOpen:   "xx.png",
			Description: "World breaking culinary innovation called CreamIce",
			Story:       "nothing fancy",
			SourcingValues: []entities.SourcingValues{
				entities.SourcingValues{Value: "universe"},
				entities.SourcingValues{Value: "yes"},
			},
			Ingredients: []entities.Ingredients{
				entities.Ingredients{Value: "sugar"},
				entities.Ingredients{Value: "spice"},
				entities.Ingredients{Value: "everything nice"},
			},
			AllergyInfo:           "for everyone",
			DietaryCertifications: "iceice creamporation",
			ProductID:             "9087",
		},
		entities.SweetsCollection{
			Name:        "Nilo energy drink",
			ImageClosed: "nn.png",
			ImageOpen:   "aa.png",
			Description: "Nilo everyday",
			Story:       "life of ni",
			SourcingValues: []entities.SourcingValues{
				entities.SourcingValues{Value: "no"},
				entities.SourcingValues{Value: "thing"},
			},
			Ingredients: []entities.Ingredients{
				entities.Ingredients{Value: "sugar"},
				entities.Ingredients{Value: "sugar"},
				entities.Ingredients{Value: "sugar"},
			},
			AllergyInfo:           "!for everyone",
			DietaryCertifications: "nistle",
			ProductID:             "9070",
		},
	}
}

func Test_sweetHandler_GetAllSweets(t *testing.T) {
	tests := []struct {
		name string
		sw   *sweetHandler
		want *[]SweetsCollection
	}{
		struct {
			name string
			sw   *sweetHandler
			want *[]SweetsCollection
		}{
			name: "Sweets Collection containing multiple data",
			sw:   &sweetHandler{mockDBManager{}},
			want: &[]SweetsCollection{
				SweetsCollection{
					Name:                  "CreamIce",
					ImageClosed:           "cc.png",
					ImageOpen:             "xx.png",
					Description:           "World breaking culinary innovation called CreamIce",
					Story:                 "nothing fancy",
					SourcingValues:        []string{"universe", "yes"},
					Ingredients:           []string{"sugar", "spice", "everything nice"},
					AllergyInfo:           "for everyone",
					DietaryCertifications: "iceice creamporation",
					ProductID:             "9087",
				},
				SweetsCollection{
					Name:                  "Nilo energy drink",
					ImageClosed:           "nn.png",
					ImageOpen:             "aa.png",
					Description:           "Nilo everyday",
					Story:                 "life of ni",
					SourcingValues:        []string{"no", "thing"},
					Ingredients:           []string{"sugar", "sugar", "sugar"},
					AllergyInfo:           "!for everyone",
					DietaryCertifications: "nistle",
					ProductID:             "9070",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sw.GetAllSweets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sweetHandler.GetAllSweets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sweetHandler_DeleteSweet(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name    string
		sw      *sweetHandler
		args    args
		want    string
		wantErr bool
	}{
		struct {
			name    string
			sw      *sweetHandler
			args    args
			want    string
			wantErr bool
		}{
			name:    "Delete success",
			sw:      &sweetHandler{mockDBManager{}},
			args:    args{params: map[string]string{"productId": "999"}},
			want:    "delete success",
			wantErr: false,
		},
		struct {
			name    string
			sw      *sweetHandler
			args    args
			want    string
			wantErr bool
		}{
			name:    "Delete provided error",
			sw:      &sweetHandler{mockDBManager{}},
			args:    args{params: map[string]string{"productId": "000"}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sw.DeleteSweet(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("sweetHandler.DeleteSweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sweetHandler.DeleteSweet() = %v, want %v", got, tt.want)
			}
		})
	}
}
