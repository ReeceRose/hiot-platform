package types

//IAuthService is an interface which provides method signatures for a auth service
type IAuthService interface {
	Validate(string) StandardResponse
	Issue() StandardResponse
}
