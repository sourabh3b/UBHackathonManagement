package participant

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TeamDetails - data model
type TeamDetails struct {
	UserName                          string   `bson:"userName" json:"userName"`
	Password                          string   `bson:"password" json:"password"`
	IsAdmin                           bool     `bson:"isAdmin" json:"isAdmin"`
	TeamName                          string   `bson:"teamName" json:"teamName"`
	ProjectObjective                  string   `bson:"projectObjective" json:"projectObjective"`
	Description                       string   `bson:"description" json:"description"`
	TeamLeadName                      string   `bson:"teamLeadName" json:"teamLeadName"`
	TeamPlayers                       []Member `bson:"teamMembers" json:"teamMembers"`
	SoftwareOrProgrammingLanguageUsed []string `bson:"softwareOrProgrammingLanguageUsed" json:"softwareOrProgrammingLanguageUsed"`
	HardwareUsed                      []string `bson:"hardwareUsed" json:"hardwareUsed"`
}

type Member struct {
	FName          string `bson:"fName" json:"fName"`
	LName          string `bson:"lName" json:"lName"`
	UniversityName string `bson:"universityName" json:"universityName"`
	Year           string `bson:"year" json:"year"`
	City           string `bson:"city" json:"city"`
	Major          string `bson:"major" json:"major"`   //1st year, 2nd year, drop down
	Degree         string `bson:"degree" json:"degree"` //BS, MS //drop down
	PhoneNumber    string `bson:"phoneNumber" json:"phoneNumber"`
	Email          string `bson:"email" json:"email"`
}

//TeamDetails - data model
type LoginResponse struct {
	Status  int  `bson:"status" json:"status"`
	IsAdmin bool `bson:"isAdmin" json:"isAdmin"`
	TypeAPI int  `bson:"typeAPI" json:"typeAPI"`
}

//TeamDetails - data model
type UpdateResponse struct {
	Status  int     `bson:"status" json:"status"`
	Message string  `bson:"message" json:"message"`
	TypeAPI float64 `bson:"typeAPI" json:"typeAPI"`
}

//GetTeamResponse - data model
type GetTeamResponse struct {
	Status  int         `bson:"status" json:"status"`
	Team    TeamDetails `bson:"team" json:"team"`
	TypeAPI int         `bson:"typeAPI" json:"typeAPI"`
}

//GetAllTeamsResponse - data model
type GetAllTeamsResponse struct {
	Status  int           `bson:"status" json:"status"`
	Team    []TeamDetails `bson:"team" json:"team"`
	TypeAPI int           `bson:"typeAPI" json:"typeAPI"`
}

//GetParticipant - handler to get expenses
func GetTeamByName(teamName string) (TeamDetails, error) {
	teamObj := TeamDetails{}
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Error("Mongo error", err.Error())
		return teamObj, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()

	// query
	err = session.DB("UBHacking").C("TeamDetails").Find(bson.M{"userName": teamName}).One(&teamObj)
	if err != nil {
		fmt.Println("Unable to find participantObject by ID", err.Error())
		return teamObj, errors.New("Unable to find participantObject by ID " + err.Error())
	}
	return teamObj, err
}

////GetParticipant - handler to get expenses
func UpdateTeamDetails(team TeamDetails) (error, bool) {

	var isNew bool

	participantObject := TeamDetails{}
	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	if err != nil {
		log.Error("Mongo error", err.Error())
		return errors.New("Mongo connection Error " + err.Error()), false
	}

	defer session.Close()

	//query to get team details
	err = session.DB("UBHacking").C("TeamDetails").Find(bson.M{"userName": team.UserName}).One(&participantObject)

	//participantObject.Password = enteredPassword
	//team doesn't exist, create new team
	if err != nil {
		isNew = true
		log.Info("Unable to find participantObject by ID", err.Error())
		//return errors.New("Unable to find participantObject by ID " + err.Error())
	}

	//one by one modify team details
	participantObject.UserName = team.UserName
	participantObject.Password = team.Password
	participantObject.TeamName = team.TeamName
	participantObject.ProjectObjective = team.ProjectObjective
	participantObject.Description = team.Description
	participantObject.TeamLeadName = team.TeamLeadName
	participantObject.TeamPlayers = team.TeamPlayers
	participantObject.SoftwareOrProgrammingLanguageUsed = team.SoftwareOrProgrammingLanguageUsed
	participantObject.HardwareUsed = team.HardwareUsed

	//modify team details with upsert, if doesn't exist create else update
	_, err = session.DB("UBHacking").C("TeamDetails").Upsert(bson.M{"userName": team.UserName}, bson.M{"$set": participantObject})
	if err != nil {
		log.Error("Unable to find participantObject by ID", err.Error())
		return errors.New("Unable to find participantObject by ID " + err.Error()), false
	}

	return err, isNew
}

//getAllteamDetails -  obtain all team details
func GetAllTeamDetails() (GetAllTeamsResponse, error) {
	response := GetAllTeamsResponse{}

	teams := []TeamDetails{}
	teamsResponse := []TeamDetails{}

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Error("Mongo error ", err.Error())
		return response, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()
	err = session.DB("UBHacking").C("TeamDetails").Find(nil).All(&teams)

	for _, val := range teams {
		if !val.IsAdmin {
			teamsResponse = append(teamsResponse, val)
		}
	}
	response.TypeAPI = 3
	response.Team = teamsResponse
	return response, err

}

//Login - Login
func Login(userName, password string) (LoginResponse, error) {
	loginResponse := LoginResponse{}

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Error("Mongo error ", err.Error())
		return loginResponse, errors.New("Mongo connection Error " + err.Error())
	}

	defer session.Close()

	// query for authentication
	err = session.DB("UBHacking").C("TeamDetails").Find(bson.M{"userName": userName, "password": password}).One(&loginResponse)
	if err != nil {
		log.Error("Unable to find User ", err.Error())
		return loginResponse, errors.New("Unable to find user " + err.Error())
	}

	loginResponse.Status = 200
	loginResponse.TypeAPI = 1

	return loginResponse, err
}
