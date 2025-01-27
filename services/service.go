package services

import "net/http"

// Service defines a generic interface for integrations
type Service interface {
    HandleRequest(w http.ResponseWriter, r *http.Request) // Process incoming HTTP requests
    GetName() string                                      // Returns the name of the service
}
