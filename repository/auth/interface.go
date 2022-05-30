package auth

type AuthRepositoryInterface interface {
	LoginSeller(email string, password string) (string, error) //login menggunakan pendekatan email dan password
	LoginBuyer(email string, password string) (string, error)  //login menggunakan pendekatan email dan password
}
