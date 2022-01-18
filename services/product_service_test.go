package services

// testcase
/*
- Products added to the catalog more than 1 year ago are reservable
*/
import (
	"go-testify/models"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ProductTestifyMock struct {
	mock.Mock
}

func (p ProductTestifyMock) IsProductReservable(id int) (models.Product, error) {
	// user is args here ..
	args := p.Called(id)
	// get returns args at specific index
	result := args.Get(0)
	return result.(models.Product), args.Error(1)
}

func TestProduct(t *testing.T) {
	assertions := require.New(t)

	// Register test mocks
	productMock := ProductTestifyMock{}
	productMock.On("IsProductReservable", 1).Return(models.Product{
		Id:          1,
		Name:        "Product-1",
		Description: "Product created 2 years ago",
		CreatedAt:   time.Now().AddDate(-2, 0, 0),
	}, nil)
	productMock.On("IsProductReservable", 2).Return(models.Product{
		Id:          2,
		Name:        "Product-2",
		Description: "Product created 1 years ago",
		CreatedAt:   time.Now().AddDate(-1, 0, 0),
	}, nil)
	productMock.On("IsProductReservable", 3).Return(models.Product{
		Id:          3,
		Name:        "Product-3",
		Description: "Product recently created",
		CreatedAt:   time.Now(),
	}, nil)

	testDataSet := map[int]bool{
		1: true,
		2: true,
		3: false,
	}

	productServiceImpl := NewProductService(productMock)

	for productId, expectedResult := range testDataSet {
		reservable, err := productServiceImpl.IsProductReservable(productId)
		assertions.Equalf(expectedResult, reservable, "Got wrong reservable info for product id %v")
		assertions.NoErrorf(err, "Failed to check if product %v is reservable: %s", productId, err)

	}
}
