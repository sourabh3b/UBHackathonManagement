package main

import (
	"fmt"
	"github.com/unrolled/render"
	"net/http"
	"github.com/UBHackathonManagement/participant"
)

//GetTeamByName - handler to get team details by name
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
	render := render.New()
	render.JSON(w, http.StatusOK, nil)
	return
}

func main() {
	fmt.Println("Started UB Hackathon Management....")
	http.HandleFunc("/test", TestRoute)
	http.ListenAndServe(":8889", nil)
}
