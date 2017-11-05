package main

import (
	"encoding/json"
	"fmt"
	"github.com/UBHackathonManagement/participant"
	"github.com/unrolled/render"
	"net/http"
)

//GetTeamByName - handler to get team details by names
func GetTeamByName(w http.ResponseWriter, r *http.Request) {
	render := render.New()

	teamName := r.URL.Query().Get("teamName")

	expenseObject, err := participant.GetParticipant(teamName)
	if err != nil {
		fmt.Println("Cannot get Team Name ", err.Error())
		render.JSON(w, http.StatusBadGateway, "Team Details")
		return
	}
	render.JSON(w, http.StatusOK, expenseObject)
	return
}

//TestRoute - test route
func TestRoute(w http.ResponseWriter, r *http.Request) {
	//render := render.New()
	fmt.Fprint(w, "Hello World !")
	//render.JSON(w, http.StatusOK, nil)
	return
}

//todo : use this later for login input in POST request
type User struct {
	UserName string `bson:"userName" json:"userName"`
	Password string `bson:"password" json:"password"`
}

//LoginHandler - LoginHandler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//userObject := User{}
	render := render.New()
	userName := r.URL.Query().Get("userName")
	password := r.URL.Query().Get("password")

	resp, _ := participant.Login(userName, password)

	badResponse := participant.LoginResponse{}
	badResponse.IsAdmin = false
	badResponse.Status = 403

	if resp.Status == 200 {
		render.JSON(w, http.StatusOK, resp)
	} else {
		render.JSON(w, http.StatusForbidden, badResponse)
	}
}

func GetTeamDetailsHandler(w http.ResponseWriter,r *http.Request){
	render := render.New()
	teamsResponse,err :=participant.GetAllTeamDetails();
	if(err != nil){
		render.JSON(w,http.StatusInternalServerError,teamsResponse);
	}else{
		render.JSON(w,http.StatusOK,teamsResponse);
	}

}
//UpdateTeamDetails - UpdateTeamDetails
func UpdateTeamDetails(w http.ResponseWriter, r *http.Request) {

	render := render.New()
	team := participant.TeamDetails{}

	//decoding the request into team, so that it can be used to save the team details
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {

	}

	err = participant.UpdateTeamDetails(team)

	updateResponse := participant.UpdateResponse{}

	//if team doesnot exist, create a new team
	if err != nil {
		updateResponse.Status = 403
		updateResponse.Message = "Cannot Update Team Details"
		render.JSON(w, http.StatusOK, updateResponse)
		return
	} else {
		//else update existing team
		updateResponse.Status = 200
		updateResponse.Message = "Successfully Updated Team Details"
		render.JSON(w, http.StatusOK, updateResponse)
		return
	}
}

func GetTeamDetailsHandler(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	teamsResponse, err := participant.GetAllTeamDetails()
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, teamsResponse)
	} else {
		render.JSON(w, http.StatusOK, teamsResponse)
	}

}
func main() {
	fmt.Println("Started UB Hackathon Management....")
	http.HandleFunc("/test", TestRoute)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/getTeamDetails", GetTeamDetailsHandler)
	http.HandleFunc("/team/update", UpdateTeamDetails)
	http.HandleFunc("/getTeamDetails", GetTeamDetailsHandler)
	http.ListenAndServe(":8889", nil)
}
