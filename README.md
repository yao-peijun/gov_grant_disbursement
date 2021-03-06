# Backend Technical Assessment : Government Grant Disbursement APIs

## Content
1. Set-up
- Database
- Host HTTP Server

2. Create APIs 
- /api/household
- /api/household/:householdID/familyMember

3. Query APIs
- /api/household/:householdID
- /api/household/all
- /api/grants?household=size&totalIncome=income

4. Delete APIs
- /api/household/:householdID
- /api/familyMember/:familyMemberID

## Set-up

### Database (MySQL)
1. import tables: (1) household and (2) familyMember to MySQL
2. files: household_household.sql and household_familyMember.sql @ folder: project_folder_path/set_up

### Host HTTP Server

#### Primary method: Run using executable binary file: gov_grant_disbursement.exe
1. Change DB connection settings @ project_folder_path/conf/app.conf. Default settings are as follow:
- mysqlurls = 127.0.0.1:3306
- mysqldb = household
- mysqluser = root
- mysqlpass = root
2. Git Bash @ root of project folder
3. command: ./gov_grant_disbursement.exe
4. http server will run @ ip address: "http://127.0.0.1:8080" or "http://localhost:8080" in default

Remarks:
- host http server using another ip address or port by changing "httpaddr" or "httpport" in project_folder_path/conf/app.conf
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
        "householdID": 9,
        "householdType": "hdb",
        "familyMembers": [
            {
                "familyMemberID": 16,
                "name": "mother2",
                "gender": "F",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 50000,
                "DOB": "1988-01-31"
            },
            {
                "familyMemberID": 17,
                "name": "father2",
                "gender": "M",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 50000,
                "DOB": "1988-01-31"
            },
            {
                "familyMemberID": 18,
                "name": "child",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "student",
                "annualIncome": 0,
                "DOB": "2005-02-28"
            },
            {
                "familyMemberID": 19,
                "name": "kid",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "student",
                "annualIncome": 0,
                "DOB": "2016-02-28"
            }
        ]
    }
]
```

### 2. Show household

#### Endpoint: /api/household/all

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
                "DOB": "2003-02-28"
            },
            {
                "familyMemberID": 9,
                "name": "mother",
                "gender": "F",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 50000,
                "DOB": "2000-01-31"
            },
            {
                "familyMemberID": 10,
                "name": "father",
                "gender": "F",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 50000,
                "DOB": "2000-01-31"
            }
        ]
    },
    {
        "householdID": 8,
        "householdType": "hdb",
        "familyMembers": null
    },
    {
        "householdID": 9,
        "householdType": "hdb",
        "familyMembers": [
            {
                "familyMemberID": 16,
                "name": "mother2",
                "gender": "F",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 50000,
                "DOB": "1988-01-31"
            },
            {
                "familyMemberID": 17,
                "name": "father2",
                "gender": "M",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 50000,
                "DOB": "1988-01-31"
            },
            {
                "familyMemberID": 18,
                "name": "child",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "student",
                "annualIncome": 0,
                "DOB": "2005-02-28"
            },
            {
                "familyMemberID": 19,
                "name": "kid",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "student",
                "annualIncome": 0,
                "DOB": "2016-02-28"
            }
        ]
    },
    {
        "householdID": 10,
        "householdType": "hdb",
        "familyMembers": [
            {
                "familyMemberID": 24,
                "name": "elderly",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "unemployed",
                "annualIncome": 0,
                "DOB": "1945-01-31"
            }
        ]
    },
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
                "DOB": "1990-01-31"
            },
            {
                "familyMemberID": 23,
                "name": "john",
                "gender": "F",
                "maritalStatus": "married",
                "occupationType": "employed",
                "annualIncome": 100000,
                "DOB": "1988-01-31"
            },
            {
                "familyMemberID": 25,
                "name": "child",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "student",
                "annualIncome": 0,
                "DOB": "2005-02-28"
            },
            {
                "familyMemberID": 29,
                "name": "elderly2",
                "gender": "M",
                "maritalStatus": "single",
                "occupationType": "unemployed",
                "annualIncome": 0,
                "DOB": "1971-01-31"
            }
        ]
    }
]
```

### 3. Search for households and recipients of grant disbursement

#### Endpoint: /api/grants?household=size&totalIncome=income

#### Assumptions
- household display should met requirement of "household size" and "total income" provided and filtered based on the criteria provided
- child should be age of > 5 and less than the age 16 and 18 for Student Encouragement Bonus and Family Togetherness Scheme respectively

#### sameple output
Note: size=5, total_income=200000
```json
{
    "studentEncouragementBonus": [
        {
            "householdID": 9,
            "householdType": "hdb",
            "familyMembers": [
                {
                    "familyMemberID": 18,
                    "name": "child",
                    "gender": "M",
                    "maritalStatus": "single",
                    "occupationType": "student",
                    "annualIncome": 0,
                    "DOB": "2005-02-28"
                }
            ]
        }
    ],
    "familyTogethernes": [
        {
            "householdID": 6,
            "householdType": "landed",
            "familyMembers": [
                {
                    "familyMemberID": 9,
                    "name": "mother",
                    "gender": "F",
                    "maritalStatus": "married",
                    "occupationType": "employed",
                    "annualIncome": 50000,
                    "DOB": "2000-01-31"
                },
                {
                    "familyMemberID": 10,
                    "name": "father",
                    "gender": "F",
                    "maritalStatus": "married",
                    "occupationType": "employed",
                    "annualIncome": 50000,
                    "DOB": "2000-01-31"
                }
            ]
        },
        {
            "householdID": 9,
            "householdType": "hdb",
            "familyMembers": [
                {
                    "familyMemberID": 16,
                    "name": "mother2",
                    "gender": "F",
                    "maritalStatus": "married",
                    "occupationType": "employed",
                    "annualIncome": 50000,
                    "DOB": "1988-01-31"
                },
                {
                    "familyMemberID": 17,
                    "name": "father2",
                    "gender": "M",
                    "maritalStatus": "married",
                    "occupationType": "employed",
                    "annualIncome": 50000,
                    "DOB": "1988-01-31"
                }
            ]
        },
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
                    "DOB": "1990-01-31"
                },
                {
                    "familyMemberID": 23,
                    "name": "john",
                    "gender": "F",
                    "maritalStatus": "married",
                    "occupationType": "employed",
                    "annualIncome": 100000,
                    "DOB": "1988-01-31"
                }
            ]
        }
    ],
    "elderBonus": [
        {
            "householdID": 10,
            "householdType": "hdb",
            "familyMembers": [
                {
                    "familyMemberID": 24,
                    "name": "elderly",
                    "gender": "M",
                    "maritalStatus": "single",
                    "occupationType": "unemployed",
                    "annualIncome": 0,
                    "DOB": "1945-01-31"
                }
            ]
        },
        {
            "householdID": 11,
            "householdType": "condominium",
            "familyMembers": [
                {
                    "familyMemberID": 29,
                    "name": "elderly2",
                    "gender": "M",
                    "maritalStatus": "single",
                    "occupationType": "unemployed",
                    "annualIncome": 0,
                    "DOB": "1971-01-31"
                }
            ]
        }
    ],
    "babySunshunGrant": [
        {
            "householdID": 9,
            "householdType": "hdb",
            "familyMembers": [
                {
                    "familyMemberID": 19,
                    "name": "kid",
                    "gender": "M",
                    "maritalStatus": "single",
                    "occupationType": "student",
                    "annualIncome": 0,
                    "DOB": "2016-02-28"
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
            "householdID": 9,
            "householdType": "hdb",
            "totalAnnualIncome": 100000
        },
        {
            "householdID": 10,
            "householdType": "hdb",
            "totalAnnualIncome": 0
        }
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