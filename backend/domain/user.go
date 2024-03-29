package domain

type User struct {
	Id                  int    `db:"id" json:"id,omitempty"`
	Username            string `db:"username" json:"username,omitempty"`
	Email               string `db:"email" json:"email,omitempty"`
	EmailVerified       bool   `db:"email_verified" json:"emailVerified"`
	Password            string `db:"password" json:"password,omitempty"`
	NotificationOptions `json:"notificationOptions"`
}

type NotificationOptions struct {
	Releases        bool `db:"notify_releases" json:"notifyReleases"`
	Recommendations bool `db:"notify_recommendations" json:"notifyRecommendations"`
}
