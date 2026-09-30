package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/renatopgn/Go/packages"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(db *pgxpool.Pool) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/api/users", GetHandler(db))
	r.Post("/api/users", PostHandler(db))
	r.Get("/api/users/{id}", GetIDHandler(db))
	r.Put("/api/users/{id}", PutIDHandler(db))
	r.Delete("/api/users/{id}", DeleteHandler(db))

	return r
}

func GetHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := packages.FindAll(db)
		if err != nil {
			SendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		SendJSON(w, Response{Data: user}, http.StatusOK)
	}

}

func PostHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser packages.User
		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			slog.Error("failed to decode body", "error", err)
			SendJSON(w, Response{Error: "fail to save body"}, http.StatusBadRequest)
			return
		}

		if newUser.FirstName == "" || newUser.LastName == "" || newUser.Biography == "" {
			SendJSON(w, Response{Error: "invalid body"}, http.StatusBadRequest)
			return
		}

		if len(newUser.FirstName) < 2 || len(newUser.FirstName) > 20 {
			SendJSON(w, Response{Error: "invalid body, first name"}, http.StatusBadRequest)
			return
		}
		if len(newUser.LastName) < 2 || len(newUser.LastName) > 20 {
			SendJSON(w, Response{Error: "invalid body, last name"}, http.StatusBadRequest)
			return
		}
		if len(newUser.Biography) < 20 || len(newUser.Biography) > 450 {
			SendJSON(w, Response{Error: "invalid body, biography"}, http.StatusBadRequest)
			return
		}

		createdUser, err := packages.Insert(db, newUser)
		if err != nil {
			SendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		SendJSON(w, Response{Data: createdUser}, http.StatusCreated)
	}

}

func GetIDHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		SearchID, err := uuid.Parse(idStr)
		if err != nil {
			SendJSON(w, Response{Error: "invlid body"}, http.StatusBadRequest)
			return
		}

		ID := packages.Id(SearchID)

		user, err := packages.FindById(db, ID)
		if err != nil {
			slog.Error("erro pra mandar pra funcao", "erro", err)
			SendJSON(w, Response{Error: "User not found"}, http.StatusNotFound)
			return
		}

		SendJSON(w, Response{Data: user}, http.StatusOK)

	}

}

func PutIDHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		SearchID, err := uuid.Parse(idStr)
		if err != nil {
			SendJSON(w, Response{Error: "invlid body"}, http.StatusBadRequest)
			return
		}

		ID := packages.Id(SearchID)

		var updatesUser packages.User
		if err := json.NewDecoder(r.Body).Decode(&updatesUser); err != nil {
			SendJSON(w, Response{Error: "fail to save body"}, http.StatusBadRequest)
			return
		}

		if updatesUser.FirstName == "" || updatesUser.LastName == "" || updatesUser.Biography == "" {
			SendJSON(w, Response{Error: "invalid body"}, http.StatusBadRequest)
			return
		}

		if len(updatesUser.FirstName) < 2 || len(updatesUser.FirstName) > 20 {
			SendJSON(w, Response{Error: "invalid body, first name"}, http.StatusBadRequest)
			return
		}
		if len(updatesUser.LastName) < 2 || len(updatesUser.LastName) > 20 {
			SendJSON(w, Response{Error: "invalid body, last name"}, http.StatusBadRequest)
			return
		}
		if len(updatesUser.Biography) < 20 || len(updatesUser.Biography) > 450 {
			SendJSON(w, Response{Error: "invalid body, biography"}, http.StatusBadRequest)
			return
		}

		UserUpdate, err := packages.Update(ID, updatesUser, db)
		if err != nil {
			SendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		SendJSON(w, Response{Data: UserUpdate}, http.StatusOK)
	}
}

func DeleteHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		SearchID, err := uuid.Parse(idStr)
		if err != nil {
			SendJSON(w, Response{Error: "invlid body"}, http.StatusBadRequest)
			return
		}

		ID := packages.Id(SearchID)

		removedUser, err := packages.Delete(db, ID)

		if err != nil {
			SendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		SendJSON(w, Response{Data: removedUser}, http.StatusOK)

	}

}

func SendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)

		SendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}
