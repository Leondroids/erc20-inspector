package app

import (
	"encoding/json"
	"net/http"
)

type TokenResponse struct {
	ClientVersion   string `json:"clientVersion"`
	ProtocolVersion int64  `json:"protocolVersion"`
	NetworkID       string `json:"networkId"`
}

func (it *TokenHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	info, err := it.processTokenResponse("")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(info)
}

func (it *TokenHandler) processTokenResponse(tokenID string) (*TokenResponse, error) {
	return nil, nil
}
