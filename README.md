# Backend Technical Assessment : Government Grant Disbursement APIs

## Content
1. Set-up
2. Create APIs
3. Query APIs
4. Delete APIs

## Set-up

### Database (MySQL)
1. import tables: (1) household and (2) familyMember to MySQL
2. files: household.sql and familyMember.sql @ folder: project_folder_path/set_up

### APIs Deployment

#### Primary method: Run using executable binary file: gov_grant_disbursement.exe
1. Git Bash @ root of project folder
2. Change DB connection settings @ project_folder_path/conf/app.conf
Default settings are as follow:
- mysqlurls = 127.0.0.1:3306
- mysqldb = household
- mysqluser = root
- mysqlpass = root
3. command: ./gov_grant_disbursement.exe
4. http server will run @ ip address: "http://127.0.0.1:8080" or "http://localhost:8080" in default

Remarks:
- change to another ip address or port by changing "httpaddr" or "httpport" in project_folder_path/conf/app.conf
- gov_grant_disbursement.exe need to be executed with folder: project_folder_path/conf at the same root

#### Alternative method: Compile and Run
1. Place project folder @ $GOPATH/src
2. Git Bash @ root of project folder
2. command to download the dependencies: "go build"
3. command to run: (1) "go run main.go" or (2) "bee run"

Remarks:
- prerequisite: Go installed with $GOPATH set


## Create APIs

### 1. Create Household

#### Endpoint: /api/household

#### input validation
- householdType cannot be empty
- input only accept (1) Landed, (2) Condominium, (3) HDB (Non-case sensitive)

#### body
```json
{
  "householdType": "HDB"
}
```

#### success output
```json
{
  "id": 11,
  "rowAffected": 1
}
```

### 2. Add a family member to household

#### Endpoint: /api/household/:householdID/familyMember

#### input validations
- name, gender, maritalStatus, occupationType, DOB should not be empty
- gender input accept "F" or "M" (Non-case sensitive)
- maritalStatus input accept "Single" or "Married" or "Divorced" or "Windowed" (Non-case sensitive)
- spouse need to be provided if maritalStatus is "Married"
- occupationType input accept "Unemployed" or "Student" or "Employed" (Non-case sensitive)
- annualIncome need to be provided if occupationType is "Employed"
- DOB should be in format DD-MM-YYYY

#### sample body
```json
{
  "name": "Mary",
  "gender": "F",
  "maritalStatus": "married",
  "spouse": "John",
  "occupationType": "employed",
  "annualIncome": 100000,
  "DOB": "12-01-1988"
}
```

#### success output
```json
{
  "id": 22,
  "rowAffected": 1
}
```


## Query APIs

### 1. List households

#### Endpoint: /api/household/:householdID

#### sample output
```json
[
  {
    "householdID": 6,
    "householdType": "landed",
    "familyMembers": [
      {
        "familyMemberID": 8,
        "name": "child",
        "gender": "F",
        "maritalStatus": "single",
        "occupationType": "student",
        "annualIncome": 0,
        "DOB": "2018-01-12"
      },
      {
        "familyMemberID": 9,
        "name": "mother",
        "gender": "F",
        "maritalStatus": "married",
        "occupationType": "employed",
        "annualIncome": 50000,
        "DOB": "2000-01-12"
      },
      {
        "familyMemberID": 10,
        "name": "father",
        "gender": "F",
        "maritalStatus": "married",
        "occupationType": "employed",
        "annualIncome": 50000,
        "DOB": "2000-01-12"
      }
    ]
  },
  {
    "householdID": 10,
    "householdType": "hdb",
    "familyMembers": [
      {
        "familyMemberID": 15,
        "name": "elderly",
        "gender": "M",
        "maritalStatus": "single",
        "occupationType": "unemployed",
        "annualIncome": 0,
        "DOB": "1945-01-31"
      }
    ]
  }
]
```

### 2. Show household

#### Endpoint: /api/household/all

#### output
```json
[
  {
    "householdID": 8,
    "householdType": "hdb",
    "familyMembers": [
      {
        "familyMemberID": 15,
        "name": "student1",
        "gender": "F",
        "maritalStatus": "single",
        "occupationType": "student",
        "annualIncome": 0,
        "DOB": "2010-01-12"
      }
    ]
  }
]
```

### 3. Search for households and recipients of grant disbursement

#### Endpoint: /api/grants?household=<size>&totalIncome=<income>

#### Assumptions
- household display should met requirement of "household size" and "total income" provided and filtered based on the criteria provided
- child should be age of > 5 and less than the age 16 and 18 for Student Encouragement Bonus and Family Togetherness Scheme respectively

#### sameple output
```json
{
    "studentEncouragementBonus": [],
    "familyTogethernes": [],
    "elderBonus": [],
    "babySunshunGrant": [
      {
        "householdID": 11,
        "householdType": "condominium",
        "familyMembers": [
          {
            "familyMemberID": 22,
            "name": "mary",
            "gender": "F",
            "maritalStatus": "married",
            "occupationType": "employed",
            "annualIncome": 100000,
            "DOB": "1950-01-31"
          }
        ]
      }
    ],
    "yoloGstGrant": [
      {
        "householdID": 6,
        "householdType": "landed",
        "totalAnnualIncome": 100000
      },
      {
        "householdID": 8,
        "householdType": "hdb",
        "totalAnnualIncome": 100000
      },
    ]
}
```


## Delete APIs

### 1. Delete Household

#### Endpoint: /api/household/:householdID

Family member with the same householdID will deleted together

#### success output
```json
{
  "rowAffected": 1
}
```

### 2. Delete Family Member

#### Endpoint: /api/familyMember/:familyMemberID

#### success output
```json
{
  "rowAffected": 1
}
```