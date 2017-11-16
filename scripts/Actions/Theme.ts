import { AnimeNotifier } from "../AnimeNotifier"

let light = {}
let dark = {
	"hue": "45",
	"saturation": "100%",

	"text-color": "hsl(0, 0%, 90%)",
	"bg-color": "hsl(0, 0%, 24%)",
	"link-color": "hsl(var(--hue), var(--saturation), 66%)",
	"link-hover-color": "hsl(var(--hue), var(--saturation), 76%)",
	"link-hover-text-shadow": "0 0 8px hsla(var(--hue), var(--saturation), 66%, 0.5)",
	"ui-background": "hsla(0, 0%, 8%, 0.3)",
	"sidebar-background": "rgba(0, 0, 0, 0.2)",
	"sidebar-opaque-background": "var(--ui-background)",
	"table-row-hover-background": "hsla(0, 0%, 100%, 0.01)",

	"theme-white": "var(--bg-color)",
	"theme-black": "var(--text-color)",

	"main-color": "var(--link-color)",
	"link-active-color": "var(--link-hover-color)",
	"button-hover-color": "var(--link-hover-color)",
	"button-hover-background": "hsla(0, 0%, 12%, 0.5)",
	"tab-background": "hsla(0, 0%, 0%, 0.1)",
	"tab-hover-background": "hsla(0, 0%, 0%, 0.2)",
	"tab-active-color": "hsl(0, 0%, 95%)",
	"tab-active-background": "hsla(0, 0%, 2%, 0.5)",
	"loading-anim-color": "var(--link-color)",
	"anime-alternative-title-color": "hsla(0, 0%, 100%, 0.5)",
	"anime-list-item-name-color": "var(--text-color)",

	"post-like-color": "var(--link-color)",
	"post-unlike-color": "var(--link-color)",
	"post-permalink-color": "var(--link-color)"
}

// Light theme
export function lightTheme() {
	console.log("Light")

	let root = document.documentElement

	for(let property in light) {
		if(!light.hasOwnProperty(property)) {
			continue
		}

		root.style.setProperty(`--${property}`, light[property])
	}
}

// Dark theme
export function darkTheme() {
	console.log("Dark")

	let root = document.documentElement

	for(let property in dark) {
		if(!dark.hasOwnProperty(property)) {
			continue
		}

		light[property] = root.style.getPropertyValue(`--${property}`)
		root.style.setProperty(`--${property}`, dark[property])
	}
}