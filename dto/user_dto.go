package dto

type SignUpRequest struct {
	Email    string `json:"email" example:"risywanazriel@gmail.com"`
	Password string `json:"password" example:"123456"`
}

// SignInRequest represents the payload for the sign in request
type SignInRequest struct {
    Email    string `json:"email" example:"risywanazriel@gmail.com"`
    Password string `json:"password" example:"123456"`
}

type SignInResponse struct {
    User struct {
        AccessToken string `json:"access_token"`
        TokenType   string `json:"token_type"`
        ExpiresIn   int    `json:"expires_in"`
        RefreshToken string `json:"refresh_token"`
        User struct {
            ID                  string                 `json:"id"`
            Aud                 string                 `json:"aud"`
            Role                string                 `json:"role"`
            Email               string                 `json:"email"`
            InvitedAt           string                 `json:"invited_at"`
            ConfirmedAt         string                 `json:"confirmed_at"`
            ConfirmationSentAt  string                 `json:"confirmation_sent_at"`
            AppMetadata         map[string]interface{} `json:"app_metadata"`
            UserMetadata        UserMetadata           `json:"user_metadata"`
            CreatedAt           string                 `json:"created_at"`
            UpdatedAt           string                 `json:"updated_at"`
        } `json:"user"`
        ProviderToken         string `json:"provider_token"`
        ProviderRefreshToken  string `json:"provider_refresh_token"`
    } `json:"user"`
}

// SignUpResponse represents the structure of the response for the Sign Up endpoint
type SignUpResponse struct {
    User struct {
        ID                  string                 `json:"id"`
        Aud                 string                 `json:"aud"`
        Role                string                 `json:"role"`
        Email               string                 `json:"email"`
        InvitedAt           string                 `json:"invited_at"`
        ConfirmedAt         string                 `json:"confirmed_at"`
        ConfirmationSentAt  string                 `json:"confirmation_sent_at"`
        AppMetadata         map[string]interface{} `json:"app_metadata"`
        UserMetadata        UserMetadata           `json:"user_metadata"`
        CreatedAt           string                 `json:"created_at"`
        UpdatedAt           string                 `json:"updated_at"`
    } `json:"user"`
}

// UserMetadata represents the user metadata in the response
type UserMetadata struct {
    Email         string `json:"email"`
    EmailVerified bool   `json:"email_verified"`
    PhoneVerified bool   `json:"phone_verified"`
    Sub           string `json:"sub"`
}