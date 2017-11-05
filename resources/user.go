package resources

type User struct {
	username string
	email string
	password_hash string
}

// UserResource is the REST layer to the User domain
type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

