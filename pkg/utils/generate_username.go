package utils

// GenerateUsernameFromEmail generate username using email
// email : using email
// ref : reference email type, ex: @gmail.com
func GenerateUsernameFromEmail(email string, ref ...string) string  {
	reference := "@gmail.com"
	if len(ref) > 0 {
		reference = ref[0]
	}
	return email[:len(email)-len(reference)]
}