package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating user references")
	defer arn.Node.Close()

	arn.DB.Clear("NickToUser")
	arn.DB.Clear("EmailToUser")
	arn.DB.Clear("GoogleToUser")
	arn.DB.Clear("FacebookToUser")

	// Iterate over the stream
	count := 0

	users, err := arn.AllUsers()
	arn.PanicOnError(err)

	// Make the most recently seen users overwrite the references of older accounts.
	// Some people accidentally created 2 accounts and in case of overlaps we want
	// to keep the references for the account that is more likely used.
	arn.SortUsersLastSeenLast(users)

	// Write new references
	for _, user := range users {
		count++
		println(count, user.Nick)

		user.ForceSetNick(user.Nick)

		if user.Email != "" {
			user.SetEmail(user.Email)
		}

		if user.Accounts.Google.ID != "" {
			user.ConnectGoogle(user.Accounts.Google.ID)
		}

		if user.Accounts.Facebook.ID != "" {
			user.ConnectFacebook(user.Accounts.Facebook.ID)
		}

		user.Save()
	}

	color.Green("Finished.")
}
