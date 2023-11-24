package main

import (
	"encoding/json"
	"fmt"
	
	"net/http"
	"time"

	"github.com/dhruvsolanki0811/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConifg) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}

	decoder:=json.NewDecoder(r.Body)
	param:=parameters{}
	err:=decoder.Decode(&param)
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Error parrsing json %v",err))
		return 
	}
	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: param.Name,
	})
	if err!=nil{
		respondWithError(w, 400, "Couldn't create user")
		return
	}
	
	respondWithJSON(w,200,databaseUserToUser(user))
}