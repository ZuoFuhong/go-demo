package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type SimpleSession struct {
	Username string
	TTL      int64
}
