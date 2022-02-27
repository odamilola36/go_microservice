package resthandlers

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"io/ioutil"
	"microservices/api/restutil"
	"microservices/pb"
	"net/http"
	"time"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authHandlers struct {
	auth pb.AuthServiceClient
}

func NewAuthHandlers(auth pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{auth: auth}
}

func (auth *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			restutil.WriteError(w, http.StatusInternalServerError, err)
		}
	}(r.Body)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := new(pb.User)

	err = json.Unmarshal(body, req)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Created = time.Now().Unix()
	req.Updated = time.Now().Unix()
	req.Id = primitive.NewObjectID().Hex()

	res, err := auth.auth.Signup(r.Context(), req)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusCreated, res)

}

func (auth *authHandlers) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			restutil.WriteError(w, http.StatusInternalServerError, err)
		}
	}(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var login Login
	err = json.Unmarshal(body, &login)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	response, err := auth.auth.SignIn(r.Context(), &pb.SigninRequest{
		Email:    login.Email,
		Password: login.Password,
	})
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, response)
}

func (auth *authHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := restutil.AuthRequestWithId(r)
	if err != nil {
		restutil.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			restutil.WriteError(w, http.StatusInternalServerError, err)
		}
	}(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := new(pb.User)
	err = json.Unmarshal(body, req)

	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	req.Id = tokenPayload.UserId

	res, err := auth.auth.Update(r.Context(), req)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusCreated, res)

}

func (auth *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := restutil.AuthRequestWithId(r)
	if err != nil {
		restutil.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	res, err := auth.auth.GetUser(r.Context(), &pb.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		restutil.WriteError(w, http.StatusNotFound, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, res)

}

func (auth *authHandlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := auth.auth.ListUsers(r.Context(), &pb.ListUsersRequest{})

	if err != nil {
		restutil.WriteError(w, http.StatusNotFound, err)
		return
	}

	var users []*pb.User

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusNotFound, err)
			return
		}
		users = append(users, user)
	}
	restutil.WriteAsJson(w, http.StatusOK, users)
}

func (auth *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := restutil.AuthRequestWithId(r)
	if err != nil {
		restutil.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	res, err := auth.auth.DeleteUser(r.Context(), &pb.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		restutil.WriteError(w, http.StatusNotFound, err)
		return
	}
	w.Header().Add("Entity", res.Id)
	restutil.WriteAsJson(w, http.StatusOK, nil)
}
