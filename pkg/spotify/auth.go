package spotify

import (
	"encoding/base64"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

var (
	RedirectUrl *url.URL
)

func init() {
	var err error
	RedirectUrl, err = url.Parse("http://localhost:8080/auth")
	if err != nil {
		panic(err)
	}
}

func getAuthToken(auth spotify.Authenticator) (*oauth2.Token, error) {
	tokenCh := make(chan *oauth2.Token)
	defer close(tokenCh)
	errCh := make(chan error)
	defer close(errCh)

	randomState, err := getRandomState()
	if err != nil {
		return nil, err
	}

	// HTTP server to handle OAuth
	authHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				errCh <- err
			}
		}()

		token, err := auth.Token(randomState, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			return
		}
		if st := r.FormValue("state"); st != randomState {
			http.NotFound(w, r)
			err = fmt.Errorf("State mismatch: %s != %s\n", st, randomState)
			return
		}
		tokenCh <- token
		fmt.Fprintln(w, "Authentication complete. This window can be safely closed.")
	})

	authServer := httptest.NewUnstartedServer(authHandler)
	authServer.Listener, err = net.Listen("tcp", RedirectUrl.Host)
	if err != nil {
		return nil, err
	}
	go authServer.Start()
	defer authServer.Close()

	webURL := auth.AuthURL(randomState)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", webURL)

	// wait for auth to complete
	select {
	case token := <-tokenCh:
		return token, nil
	case err := <-errCh:
		return nil, err
	case <-time.After(time.Minute):
		return nil, fmt.Errorf("no credentials provided after one minute")
	}
}

func getRandomState() (string, error) {
	const stateLength = 16
	buf := make([]byte, stateLength)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf), nil
}
