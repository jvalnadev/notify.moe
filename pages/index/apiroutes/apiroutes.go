package apiroutes

import (
	"strings"

	"github.com/animenotifier/notify.moe/pages/post"
	"github.com/animenotifier/notify.moe/pages/sse"
	"github.com/animenotifier/notify.moe/pages/thread"

	"github.com/aerogo/aero"

	"github.com/aerogo/layout"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/pages/animeimport"
	"github.com/animenotifier/notify.moe/pages/apiview"
	"github.com/animenotifier/notify.moe/pages/apiview/apidocs"
	"github.com/animenotifier/notify.moe/pages/character"
	"github.com/animenotifier/notify.moe/pages/editor/jobs"
	"github.com/animenotifier/notify.moe/pages/me"
	"github.com/animenotifier/notify.moe/pages/notifications"
	"github.com/animenotifier/notify.moe/pages/popular"
	"github.com/animenotifier/notify.moe/pages/soundtrack"
	"github.com/animenotifier/notify.moe/pages/upload"
)

// Register registers the page routes.
func Register(l *layout.Layout, app *aero.Application) {
	// API pages
	l.Page("/api", apiview.Get)

	for name := range arn.DB.Types() {
		l.Page("/api/"+strings.ToLower(name), apidocs.ByType(name))
	}

	// API
	app.Get("/api/me", me.Get)
	app.Get("/api/popular/anime/titles/:count", popular.AnimeTitles)
	app.Get("/api/test/notification", notifications.Test)
	app.Get("/api/count/notifications/unseen", notifications.CountUnseen)
	app.Get("/api/mark/notifications/seen", notifications.MarkNotificationsAsSeen)
	app.Get("/api/user/:id/notifications/latest", notifications.Latest)
	app.Get("/api/random/soundtrack", soundtrack.Random)
	app.Get("/api/next/soundtrack", soundtrack.Next)
	app.Get("/api/character/:id/ranking", character.Ranking)

	// Live updates
	app.Get("/api/sse/events", sse.Events)

	// Thread
	app.Get("/api/thread/:id/reply/ui", thread.ReplyUI)

	// Post
	app.Get("/api/post/:id/reply/ui", post.ReplyUI)

	// SoundTrack
	app.Post("/api/soundtrack/:id/download", soundtrack.Download)

	// Upload
	app.Post("/api/upload/avatar", upload.Avatar)
	app.Post("/api/upload/cover", upload.Cover)
	app.Post("/api/upload/anime/:id/image", upload.AnimeImage)
	app.Post("/api/upload/character/:id/image", upload.CharacterImage)
	app.Post("/api/upload/group/:id/image", upload.GroupImage)
	app.Post("/api/upload/amv/:id/file", upload.AMVFile)

	// Import anime
	app.Post("/api/import/kitsu/anime/:id", animeimport.Kitsu)
	app.Post("/api/delete/kitsu/anime/:id", animeimport.DeleteKitsu)

	// Jobs
	app.Post("/api/job/:job/start", jobs.Start)
}
