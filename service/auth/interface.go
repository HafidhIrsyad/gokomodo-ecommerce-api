package auth

type AuthServiceInterface interface {
	LoginSeller(email string, password string) (string, error) //login menggunakan pendekatan email dan password
	LoginBuyer(email string, password string) (string, error)  //login menggunakan pendekatan email dan password
}
