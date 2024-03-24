package types

type GET_USER_BY_EMAIL_INPUT struct{ Email string }
type DELETE_USER_BY_EMAIL_INPUT struct{ Email string }
type UPDATE_USER_BY_EMAIL_INPUT struct {
	Email string
	Field string
	Value string
}

type GET_SHOP_BY_EMAIL_INPUT struct{ Email string }
type DELETE_SHOP_BY_EMAIL_INPUT struct{ Email string }
type UPDATE_SHOP_BY_EMAIL_INPUT struct {
	Email string
	Field string
	Value string
}

type GET_PRODUCT_BY_ID_INPUT struct{ Id int }
type DELETE_PRODUCT_BY_ID_INPUT struct{ Id int }
type UPDATE_PRODUCT_BY_ID_INPUT struct {
	Id    int
	Field string
	Value string
}

type GET_ORDER_BY_ID_INPUT struct{ Id int }
type UPDATE_ORDER_BY_ID_INPUT struct {
	Id    int
	Field string
	Value string
}
