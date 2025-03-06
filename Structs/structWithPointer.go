package main

func (u *User) setEmailUsingPtr(email string) {
	u.Email = email
}
