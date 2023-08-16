package gonjalla

import (
	"encoding/json"
)

// Server struct contains data returned by api calls that deal with server state
type Server struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	ID          string   `json:"id"`
	Status      string   `json:"status"`
	Os          string   `json:"os"`
	Expiry      string   `json:"expiry"`
	Autorenew   bool     `json:"autorenew"`
	SSHKey      string   `json:"ssh_key"`
	Ips         []string `json:"ips"`
	ReverseName string   `json:"reverse_name"`
	OsState     string   `json:"os_state"`
}

// ListServers returns a listing of all servers for a given account
func ListServers(token string) ([]Server, error) {
	params := map[string]interface{}{}

	data, err := Request(token, "list-servers", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Servers []Server `json:"servers"`
	}

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Servers, nil
}

// ListServerImages returns a listing of the avaliable server images
func ListServerImages(token string) ([]string, error) {
	params := map[string]interface{}{}

	data, err := Request(token, "list-server-images", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Images []string `json:"images"`
	}

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Images, nil
}

// ListServerTypes returns a listing of the avaliable server types
func ListServerTypes(token string) ([]string, error) {
	params := map[string]interface{}{}

	data, err := Request(token, "list-server-types", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Types []string `json:"types"`
	}

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Types, nil
}

// StopServer stops a server from running.  Server data will not be destroyed.
func StopServer(token string, id string) (Server, error) {
	params := map[string]interface{}{
		"id": id,
	}

	var server Server

	data, err := Request(token, "stop-server", params)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}

// StartServer starts a server
func StartServer(token string, id string) (Server, error) {
	params := map[string]interface{}{
		"id": id,
	}

	var server Server

	data, err := Request(token, "start-server", params)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}

// RestartServer restarts a server
func RestartServer(token string, id string) (Server, error) {
	params := map[string]interface{}{
		"id": id,
	}

	var server Server

	data, err := Request(token, "restart-server", params)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}

// ResetServer resets a server with new server settings.  Server data WILL be destroyed
func ResetServer(token string, id string, os string, publicKey string, instanceType string) (Server, error) {
	params := map[string]interface{}{
		"id": id,
		"os": os,
		"ssh_key": publicKey,
		"type": instanceType,
	}

	var server Server

	data, err := Request(token, "reset-server", params)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}

// AddServer creates a new server.
func AddServer(token string, name string, instanceType string, os string, publicKey string, months int) (Server, error) {
	params := map[string]interface{}{
		"name": name,
		"type": instanceType,
		"os": os,
		"ssh_key": publicKey,
		"months": months,
	}

	var server Server

	data, err := Request(token, "add-server", params)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}

// RemoveServer removes a server. Server data WILL be destroyed.
func RemoveServer(token string, id string) (Server, error) {
	params := map[string]interface{}{
		"id": id,
	}

	var server Server

	data, err := Request(token, "remove-server", params)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}
