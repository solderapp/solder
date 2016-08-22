package model

// ClientPacks is simply a collection of client pack structs.
type ClientPacks []*ClientPack

// ClientPack represents a client pack model definition.
type ClientPack struct {
	ClientID int     `json:"client_id" sql:"index"`
	Client   *Client `json:"client,omitempty"`
	PackID   int     `json:"pack_id" sql:"index"`
	Pack     *Pack   `json:"pack,omitempty"`
}
