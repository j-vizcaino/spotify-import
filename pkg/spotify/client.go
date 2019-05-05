package spotify

import "github.com/zmb3/spotify"

func NewClient() (spotify.Client, error) {
	auth := spotify.NewAuthenticator(
		RedirectUrl.String(),
		spotify.ScopeUserLibraryModify,
		spotify.ScopeUserLibraryRead,
	)

	token, err := getAuthToken(auth)
	if err != nil {
		return spotify.Client{}, err
	}

	return auth.NewClient(token), nil
}
