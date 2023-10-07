package main

import "fmt"

func main() {

	authUsername := &Auth{}

	authWithUserAndEmail := &EmailAuth{
		iAuth: authUsername,
	}

	authWithEmailAndPhone := &PhoneAuth{
		iAuth: authWithUserAndEmail,
	}

	result := authWithEmailAndPhone.signIn()
	if result {
		fmt.Println("Sign in successfully")
	} else {
		fmt.Println("Sign in failed")
	}
}
