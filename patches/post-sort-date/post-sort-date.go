package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Sorting posts by date")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for thread := range arn.StreamThreads() {
		posts := thread.Posts()
		arn.SortPostsLatestLast(posts)
		postIDs := []string{}

		for _, post := range posts {
			postIDs = append(postIDs, post.ID)
		}

		thread.PostIDs = postIDs
		thread.Save()
	}
}
