package domain

type User struct {
	Id            int    `db:"id"`
	Username      string `db:"username"`
	Email         string `db:"email"`
	EmailVerified bool   `db:"email_verified"`
	Password      string `db:"password"`
	NotificationOptions
}

type NotificationOptions struct {
	Releases        bool `db:"notify_releases"`
	Recommendations bool `db:"notify_recommendations"`
}
