package types

type User struct {
	Id          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	UserAddress string
	IsSeller    bool
}

type Shop struct {
	Id              int
	OwnerId         int
	ShopName        string
	Email           string
	PhoneNumber     string
	MapLocation     string
	ShopType        string
	ShopDescription string
}

type Product struct {
	Id                           int
	ProductName                  string
	SellerId                     int
	ShopId                       int
	ProductType                  string
	ProductCondition             string
	Price                        string
	OriginalPurchasedDate        string
	OriginalPurchaisingRecieptNo string
	ProductDescription           string
}

type Order struct {
	Id                  int
	FromMapLocation     string
	ToMapLocation       string
	LastStopMapLocation string
	OrderStatus         string
	PaymentStatus       string
	ProductId           int
	BuyerId             int
	ShopId              int
}
