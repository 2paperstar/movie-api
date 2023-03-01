package utils

func IdToPosterUrl(identifier string) string {
	const CDN_BASE = "https://boomerangdb.nyc3.cdn.digitaloceanspaces.com/"
	return CDN_BASE + "posters/webp/" + identifier + ".webp"
}
