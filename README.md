# UBHackathonManagement

# Use case Diagram
<img src="https://s3-us-west-2.amazonaws.com/ubhacking/Screen+Shot+2017-11-04+at+2.34.40+AM.png" width="700" height="500"/>

# Live Demo
http://s3-us-west-2.amazonaws.com/website.frontend/UBHMS_FrontEnd/registration.html


# APIs List
###### 1. Login (persona : team , admin)


	URL : http://54.200.178.6:8889/login?userName=admin&password=adminpwd
```javascript
Type : GET
Success Resp:
{
    "status": 200,
    "isAdmin": true,
    "typeAPI": 1
}
Failure Response:
{
    "status": 403,
    "isAdmin": false,
    "typeAPI": 1
}
```



###### 2. Update Project Info (participant)


	URL : http://54.200.178.6:8889/team/update
```javascript
Type : POST

Body :

{
	"userName": "team4",
	"password": "team4pwd",
	"isAdmin": false,
	"teamName": "new team 4",
	"projectObjective": "new Objective 4",
	"description": "new description team 4",
	"teamLeadName": "sourabh",
	"teamMembers": [
		{
			"fName" : "john",
			"lName" : "doe",
			"universityName" : "MIT",
			"year" : "3rd",
			"city" : "CA",
			"major" : "EE",
			"degree" : "BS",
			"phoneNumber" : "+1 645 34 34 32",
			"email" : "dsads@aol.com"
		}
	],
	"softwareOrProgrammingLanguageUsed": ["c","c++"],
	"hardwareUsed": ["verilog"]
}

Success Resp: Create New Team
{
    "status": 200,
    "message": "Successfully Create New Team",
    "typeAPI": 2.1
}

Success Resp: Update Existing Team
{
    "status": 200,
    "message": "Successfully Updated Team Details",
    "typeAPI": 2.2
}
```


###### 3. Get Team By Name


	URL : http://54.200.178.6:8889/getTeamByName?teamName=team2

```javascript
Type : GET
Success Response
 {
    "status": 200,
    "team": {
        "userName": "team2",
        "password": "team2pwd",
        "isAdmin": false,
        "teamName": "team2",
        "projectObjective": "Uber for Dogs",
        "description": "Uber for Dogs",
        "teamLeadName": "Bob",
        "teamMembers": [
            {
                "fName": "Alice",
                "lName": "Bob",
                "universityName": "UCLA",
                "year": "2st Year",
                "city": "Los Angeles",
                "major": "CS",
                "degree": "MS",
                "phoneNumber": "+928374",
                "email": "anandr@buffalo.edu"
            }
        ],
        "softwareOrProgrammingLanguageUsed": [
            "C",
            "C++",
            "js"
        ],
        "hardwareUsed": [
            "respberry pi",
            "google glass"
        ]
    },
    "typeAPI": 4
}
Add Comment
```


  ###### 4. Get All Project Derails (admin)


	URL : http://54.200.178.6:8889/getAllTeams

```javascript
Type : GET
{
    "status": 200,
    "team": [
        {
            "userName": "team1",
            "password": "team1pwd",
            "isAdmin": false,
            "teamName": "team1",
            "projectObjective": "Hackathon Management",
            "description": "Hackathon Management",
            "teamLeadName": "sharif",
            "teamMembers": [
                {
                    "fName": "John",
                    "lName": "Cena",
                    "universityName": "UB",
                    "year": "1st Year",
                    "city": "Buffalo",
                    "major": "CS",
                    "degree": "MS",
                    "phoneNumber": "+928374",
                    "email": "sourabhb@buffalo.edu"
                }
            ],
            "softwareOrProgrammingLanguageUsed": [
                "golang",
                "mongo",
                "js"
            ],
            "hardwareUsed": []
        },
        {
            "userName": "team2",
            "password": "team2pwd",
            "isAdmin": false,
            "teamName": "team2",
            "projectObjective": "Uber for Dogs",
            "description": "Uber for Dogs",
            "teamLeadName": "Rambo",
            "teamMembers": [
                {
                    "fName": "Bob",
                    "lName": "Fisher",
                    "universityName": "UCLA",
                    "year": "2st Year",
                    "city": "Los Angeles",
                    "major": "CS",
                    "degree": "MS",
                    "phoneNumber": "+928374",
                    "email": "johns@buffalo.edu"
                }
            ],
            "softwareOrProgrammingLanguageUsed": [
                "C",
                "C++",
                "js"
            ],
            "hardwareUsed": [
                "respberry pi",
                "google glass"
            ]
        }
    ],
    "typeAPI": 3
}
```


## Programming Language & Technologies Used
```python
golang, JS, HTML, mongoDB, AWS S3, AWS EC2,
```