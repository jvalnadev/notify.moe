package soundtrack

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Lyrics of a soundtrack.
func Lyrics(ctx *aero.Context) string {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	openGraph := getOpenGraph(ctx, track)

	if track.Lyrics.Native != "" {
		openGraph.Tags["og:description"] = utils.CutLongDescription(track.Lyrics.Native)
	}

	if track.Lyrics.Romaji != "" {
		openGraph.Tags["og:description"] = utils.CutLongDescription(track.Lyrics.Romaji)
	}

	ctx.Data = openGraph
	return ctx.HTML(components.SoundTrackLyricsPage(track, user))
}
