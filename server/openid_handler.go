package server

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"regexp"
)

const (
	idKey      = "id"
	successKey = "success"
	messageKey = "message"
)

// OpenID handler structure.
type openIdHandler struct {
	redirectUrl url.URL
	regex       *regexp.Regexp
}

// http.Handler.ServeHTTP
func (h *openIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := url.Values{}

	id, err := h.parse(r.URL)

	if err != nil {
		logrus.WithError(err).Error("unable to get steam id")

		values.Add(successKey, "false")
		values.Add(messageKey, "Unable to get steam ID.")
	} else {
		values.Add(idKey, id)
		values.Add(successKey, "true")
		values.Add(messageKey, "OK.")
	}

	u := h.redirectUrl
	u.RawQuery = values.Encode()

	http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
}

func (h *openIdHandler) parse(u *url.URL) (string, error) {
	if u == nil {
		return "", errors.New("unable to parse nil url")
	}

	claimedId := u.Query().Get("openid.claimed_id")

	if len(claimedId) == 0 {
		return "", errors.New("claimed id is missing")
	}

	matches := h.regex.FindStringSubmatch(claimedId)

	if len(matches) != 2 {
		return "", errors.New("unable to parse steam id")
	}

	id := matches[len(matches)-1]

	return id, nil
}

// Creates OpenID handler with the provided parameters.
func NewOpenIdHandler(redirectUrl url.URL) http.Handler {
	regex := regexp.MustCompile("/openid/id/(.+)$")

	return &openIdHandler{
		redirectUrl: redirectUrl,
		regex:       regex,
	}
}
