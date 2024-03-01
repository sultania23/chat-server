package hash

type PasswordHasher interface {
	Hash(password string) string
}
