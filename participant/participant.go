package participant

import (
	"errors"
	"fmt"
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
}

//TeamDetails - data model
type UpdateResponse struct {
	Status  int    `bson:"status" json:"status"`
	Message string `bson:"message" json:"message"`
}

//GetParticipant - handler to get expenses
func GetParticipant(teamName string) (TeamDetails, error) {
	participantObject := TeamDetails{}
	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return participantObject, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()

	fmt.Println("input teamName : ", teamName)

	// query
	err = session.DB("UBHacking").C("Participant").Find(bson.M{"userName": teamName}).One(&participantObject)
	if err != nil {
		fmt.Println("participantObject > ", participantObject)
		fmt.Println("Unable to find participantObject by ID", err.Error())
		return participantObject, errors.New("Unable to find participantObject by ID " + err.Error())
	}

	return participantObject, err
}

//getAllteamDetails -  obtain all team details
func GetAllTeamDetails()([]TeamDetails,error){
	teams := []TeamDetails{}
	teamsResponse := []TeamDetails{}

	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return teamsResponse, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()
	err = session.DB("UBHacking").C("TeamDetails").Find(nil).All(&teams)

	for _,val := range teams{
		if(!val.IsAdmin){
			teamsResponse = append(teamsResponse,val);
		}
	}
	return teamsResponse,err;

}
////GetParticipant - handler to get expenses
func UpdateTeamDetails(team TeamDetails) error {
	participantObject := TeamDetails{}
	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return errors.New("Mongo connection Error " + err.Error())
	}

	defer session.Close()

	//query to get team details
	err = session.DB("UBHacking").C("TeamDetails").Find(bson.M{"userName": team.UserName}).One(&participantObject)

	//team doesn't exist, create new team
	if err != nil {
		participantObject.Password = "testpwd"
		fmt.Println("Unable to find participantObject by ID", err.Error())
		//return errors.New("Unable to find participantObject by ID " + err.Error())
	}

	//one by one modify team details
	participantObject.UserName = team.UserName
	participantObject.TeamName = team.TeamName
	participantObject.ProjectObjective = team.ProjectObjective
	participantObject.Description = team.Description
	participantObject.TeamLeadName = team.TeamLeadName
	participantObject.TeamPlayers = team.TeamPlayers
	participantObject.SoftwareOrProgrammingLanguageUsed = team.SoftwareOrProgrammingLanguageUsed
	participantObject.HardwareUsed = team.HardwareUsed

	//modify team details with upsert, if doesn't exist create else update
	_, err = session.DB("UBHacking").C("TeamDetails").Upsert(bson.M{"userName": team.UserName}, bson.M{"$set": participantObject})

	//err = session.DB("UBHacking").C("TeamDetails").Upsert(
	//						bson.M{"userName":team.UserName},
	//						bson.M{"$set":
	//							bson.M{
	//								"userName": team.UserName,
	//								"teamName": team.TeamName,
	//								"projectObjective": team.ProjectObjective,
	//								"description": team.Description,
	//								"teamLeadName": team.TeamLeadName,
	//								"teamMembers": team.TeamPlayers,
	//								"softwareOrProgrammingLanguageUsed" : team.SoftwareOrProgrammingLanguageUsed,
	//								"hardwareUsed" : team.HardwareUsed,
	//						},
	//						})
	if err != nil {
		fmt.Println("Unable to find participantObject by ID", err.Error())
		return errors.New("Unable to find participantObject by ID " + err.Error())
	}

	return err
}

//getAllteamDetails -  obtain all team details
func GetAllTeamDetails() ([]TeamDetails, error) {
	teams := []TeamDetails{}
	teamsResponse := []TeamDetails{}

	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return teamsResponse, errors.New("Mongo connection Error " + err.Error())
	}
	defer session.Close()
	err = session.DB("UBHacking").C("TeamDetails").Find(nil).All(&teams)

	for _, val := range teams {
		if !val.IsAdmin {
			teamsResponse = append(teamsResponse, val)
		}
	}
	return teamsResponse, err

}

//Login - Login
func Login(userName, password string) (LoginResponse, error) {
	loginResponse := LoginResponse{}

	session, err := mgo.Dial("127.0.0.1") //todo: change this to AWS mongo URL
	//session, err := mgo.Dial("ec2-54-200-178-6.us-west-2.compute.amazonaws.com") //todo: change this to AWS mongo URL
	//session, err := mgo.Dial("54.200.178.6") //todo: change this to AWS mongo URL
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return loginResponse, errors.New("Mongo connection Error " + err.Error())
	}

	defer session.Close()

	// query for authentication
	err = session.DB("UBHacking").C("TeamDetails").Find(bson.M{"userName": userName, "password": password}).One(&loginResponse)
	if err != nil {
		fmt.Println("Unable to find user", err.Error())
		return loginResponse, errors.New("Unable to find user " + err.Error())
	}

	loginResponse.Status = 200

	return loginResponse, err
}
