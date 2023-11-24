package router

import (
	"AdHub/internal/pkg/entities"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (mr *PublicRouter) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	Token := r.URL.Query().Get("id")
	if Token == "" {
		http.Error(w, "Ad token is missing", http.StatusBadRequest)
		return
	}

	adID, err := mr.ULink.GetAdId(Token)

	mr.ULink.ULinkRemove(&entities.ULink{
		Token: Token,
		AdId:  adID,
	})
	token := uuid.New().String()

	mr.ULink.ULinkCreate(&entities.ULink{
		Token: token,
		AdId:  adID,
	})

	ad, err := mr.Ad.AdRead(adID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting ad URL: %v", err), http.StatusInternalServerError)
		return
	}
	website := "http://" + ad.Website_link

	http.Redirect(w, r, website, http.StatusSeeOther)
}
