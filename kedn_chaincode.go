package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// "strconv"
	// "time"
	"errors"
	"strings"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


// // entity student
// type Student struct {
// 	StudentID   string `json:"studentId"`
// 	Name        string `json:"name"`
// 	LeftFinger  string `json:"leftFinger"`
// 	RightFinger string `json:"rightFinger"`
// 	FacePhoto   string `json:"facePhoto"`
// 	States   []string `json:"states"`
// 	RegDate   string `json:"regDate"`
// 	School   string `json:"school"`
// }

// entity student
type Student struct {
	ID   string `json:"id"`
	// StudentID        string `json:"studentId"`
	Name string `json:"name"`
	// PersonDetails  PersonDetails `json:"personDetails"`
	// CurrentGrade string `json:"currentGrade"`
	// Status string `json:"status"`
	School string `json:"school"`
	//ParentsList   []Parents `json:"parentsList"`
	// StatusList   []Status `json:"statusList"`
}

// school account entity
type Account struct {
	ID   string `json:"id"`
	Name   string `json:"name"`
	TotalStudents   int `json:"totalStudents"`
	AssetsIds   []string `json:"assetsIds"`
}

// school details
type School struct {
	ID          string  `json:"id"`
	Name      string  `json:"name"`
	StudentsIds   []string `json:"studentsIds"`
}


// states entity
type State struct {
	StateID   string `json:"stateID"`
	Status   string `json:"status"`
	School   string `json:"school"`
	TimeStamp   string `json:"timeStamp"`
}

// transactions entity
type Transaction struct {
	StudentID   string `json:"studentID"`
	FromSchool   string `json:"fromSchool"`
	ToSchool   string `json:"toSchool"`
	State   string `json:"state"`
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return t.createSchool( stub, args )
}

func (t *SimpleChaincode) createAccount(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// obtain username to create an account
	if len(args) !=1 {
		fmt.Println("Error obtaining username")
		return nil, errors.New("create account accepts only one argument")
	}

	// username := args[0]

	// build an account for the school registrar
	var err error
	var account Account

	fmt.Println("Unmarshling account")
	err = json.Unmarshal([]byte(args[0]), &account)

	if err != nil {
		fmt.Println("Error unmarshling account")
		return nil, errors.New("Invalid account")
	}

	fmt.Println("Attempting to get state of any sexisting account for id: " +account.ID)
	existingByte, err := stub.GetState(account.ID)

	if err == nil {
		// the account already exists
		var school Account
		err = json.Unmarshal(existingByte, &school)

		if err != nil {
			fmt.Println("...")
		}

	}

	return nil, nil
}

func (t *SimpleChaincode) createSchool(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// obtain the username to associate with school
	if len(args) != 2 {
		fmt.Println("Error obtaining username")
		return nil, errors.New("Create school account accepts only one argument")
	}

	username := args[0]
	name := args[1]

	// build an account object for the user
	var studentsIds []string

	var	school = School{ID: username, Name:name, StudentsIds: studentsIds}

	

	fmt.Println("Attempting to get state of any	existing school for: " + school.ID)
	existingSchoolBytes, err := stub.GetState(school.ID)
	if err == nil {
		
		var	existingSchool School
		err = json.Unmarshal(existingSchoolBytes, &existingSchool)
		if err != nil {
			fmt.Println("Error unmarshling school "+ existingSchool.ID + "\n----->: "+ err.Error())

			if strings.Contains(err.Error(), "unexpected end of JSON input") {
				fmt.Println("No data means existing school found for " + school.ID + ", initializing school")
				schoolBytes, err := json.Marshal(&school)
				if err != nil {
					fmt.Println("Error marshling school: "+school.ID)
					return nil, errors.New("Error marshling School: " + school.ID)
				}
				err = stub.PutState(school.ID, schoolBytes)

				if err == nil {
				 	fmt.Println("created school: " + school.ID)
				 	return nil, nil
				} else	{
				 	fmt.Println("failed to create initialize school "+ school.ID)
				 	return nil, errors.New("failed to initialize school for " + school.ID + " => " + err.Error())
				}
			} else {
				return nil, errors.New("error unmarshling school " + school.ID)
			}

		} else {
			fmt.Println("School already exist for " + school.ID)
			return nil, errors.New("School already exist for " + school.ID)
		}
	} else {
		fmt.Println("No existing school for "+ school.ID + ", initializing school")
		schoolBytes, err := json.Marshal(&school)
		if err != nil {
			fmt.Println("Error marshling school: "+school.ID)
			return nil, errors.New("Error marshling School: " + school.ID)
		}

		fmt.Println("creating school " + school.ID)
		err = stub.PutState(school.ID, schoolBytes)

		if err == nil {
			fmt.Println("created school " + school.ID)
			return nil, nil
		} else {
			fmt.Println("failed to create initialize school for " + school.ID)
			return nil, errors.New("failed to initialize school for " + school.ID)
		}
	}
}


// func (t *SimpleChaincode) registerStudent(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	var err error
// 	var name, id string


// 	//need one arg
// 	if len(args) < 2 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting name and student ID")
// 	}

	// id = args[0]
	// name = args[1]


	// var student = Student{StudentID: id, Name: name, LeftFinger: "", RightFinger: "", FacePhoto: ""}
			
// 	studentBytes, err1 := json.Marshal(&student)
	
// 	if err1 != nil {
// 		fmt.Println("Error marshling the student")
// 		return nil, err1
// 	}

// 	fmt.Println("registering student" + student.StudentID)

// 	err = stub.PutState(id, studentBytes)

// 	if err != nil {
// 		fmt.Println("error registering student" + id)
// 		return nil, errors.New("Error registering student " + id)
// 	}

// 	fmt.Println("Student registered")
// 	return nil, nil

// }

// invoke function to register a student
func (t *SimpleChaincode) registerStudent(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("init reg student")

	//need one arg
	if len(args) < 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting name and student ID")
	}


	var err error

	id := args[0]
	name := args[1]
	school := args[2]

	var student = Student{ID: id, Name: name, School: school}

	// check if registering school exist
	fmt.Println("Get registering school state for ID: " + student.School)
	schoolBytes, err := stub.GetState(student.School)
	if err != nil {
		fmt.Println("Registering school not found: " + student.School)
		return nil, errors.New("Check registering school, dont exist: " + student.School)
	}

	

	fmt.Println("Check if student exist with ID: " + student.ID)
	existingStudentBytes, err := stub.GetState(student.ID)
	if err == nil {
		fmt.Println("Student exist with ID: " + student.ID)

		var regStudent Student
		err = json.Unmarshal(existingStudentBytes, &regStudent)
		if err != nil {
			fmt.Println("Error unmarshling student with ID: "+student.ID + "\n------>: "+err.Error())

			if strings.Contains(err.Error(),"unexpected end of JSON input") {
				fmt.Println("No data means existing student found for "+student.ID + ", initializing student.")
					
				studentBytes, err := json.Marshal(&student)
				if err != nil {
					fmt.Println("Error marshling student")
					return nil, errors.New("Error Invalid student")
				}

				err = stub.PutState(student.ID, studentBytes)
				if err == nil {

					// add student to the school
					// 1. Unmarshal the school
					school := School {}
					err = json.Unmarshal(schoolBytes, &school)
					if err !=nil {
						fmt.Println("Error unmarshling school")
						return nil, errors.New("Error Invalid school")
					}

					school.StudentsIds = append(school.StudentsIds, student.ID)

					// 2.Marshal school after updating
					schoolBytes, err = json.Marshal(&school)
					if err != nil {
						fmt.Println("Error marshling school")
						return nil, errors.New("Error Invalid school")
					}

					// 3. put back the updated school
					err = stub.PutState(school.ID, schoolBytes)
					if err != nil {
						fmt.Println("Error updating school")
						return nil, errors.New("Error updating school")
					}


					fmt.Println("Updated student: "+student.ID)
					return nil, nil
				} else {
					fmt.Println("failed to create initialize entry for student with ID: "+ student.ID)
					return nil,errors.New("failed to initialize student for ID: "+ student.ID)
				}
			} else {
				return nil, errors.New("Error unmarshling existing studet account with ID: "+student.ID)
			}
		} else {
			fmt.Println("Student already exist for ID: "+student.ID)
			return nil, errors.New("Can't reinitialize existing user "+student.ID)
		}
 
	} else {
		fmt.Println("No existing entry for student with ID: "+ student.ID +", initializing student")
		studentBytes, err := json.Marshal(&student)
		if err != nil {
			fmt.Println("Error marshling student")
			return nil, errors.New("Error Invalid student")
		}
		err = stub.PutState(student.ID, studentBytes)

		if err == nil {
			// add student to the school
			// 1. Unmarshal the school
			school := School {}
			err = json.Unmarshal(schoolBytes, &school)
			if err !=nil {
				fmt.Println("Error unmarshling school")
				return nil, errors.New("Error Invalid school")
			}

			school.StudentsIds = append(school.StudentsIds, student.ID)

			// 2.Marshal school after updating
			schoolBytes, err = json.Marshal(&school)
			if err != nil {
				fmt.Println("Error marshling school")
				return nil, errors.New("Error Invalid school")
			}

			// 3. put back the updated school
			err = stub.PutState(school.ID, schoolBytes)
			if err != nil {
				fmt.Println("Error updating school")
				return nil, errors.New("Error updating school")
			}

			fmt.Println("Registered student with ID: "+student.ID + " and enrolled them to to school with ID: " + school.ID)
			return nil, nil
		} else {
			fmt.Println("failed to register initialize student for ID: " + student.ID + " => "+err.Error())
			return nil, errors.New("Failed to initialize a student for ID: " + student.ID + " => " + err.Error())
		}
	}
	

}

// invoke function to transfer a student
func (t *SimpleChaincode) transferStudent(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("Transfering student: " + args[0])


	// need one argument
	if len(args) < 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting a students record")
	}

	studentId := args[0]
	trFrom := args[1]
	trTo := args[2]
	trState := "" //args[3]

	transaction := Transaction{StudentID:studentId, FromSchool:trFrom, ToSchool:trTo, State:trState}

	// fmt.Println("Unmarshal Transaction")
	// err := json.Unmarshal([]byte(args[0]), &transaction)
	// if err == nil {
		
		// get state on student
		fmt.Println("getting state of student: " + transaction.StudentID)
		studentBytes, err := stub.GetState(transaction.StudentID)
		if err == nil {
			
			// check if student is member of school they are transfering fromSchool
			fmt.Println("Student found, check if student is registered in FromSchool")
			fromSchoolBytes, err := stub.GetState(transaction.FromSchool)
			if err == nil {

				var fromSchool School

				fmt.Println("Unmarshalling FromSchool")
				err = json.Unmarshal(fromSchoolBytes, &fromSchool)
				if err == nil {
					
					// check if school to transfer to existing
					fmt.Println("check if transfer to school exist for ID: " + transaction.ToSchool)
					toSchoolBytes, err := stub.GetState(transaction.ToSchool)
					if err == nil {

						var toSchool School

						fmt.Println("Unmarshalling FromSchool")
						err = json.Unmarshal(toSchoolBytes, &toSchool)
						if err == nil {

							toSchool.StudentsIds = append(toSchool.StudentsIds, transaction.StudentID)

							fmt.Println("Marshal toSchool")
							toSchoolBytes, err = json.Marshal(&toSchool)
							if err != nil {
								fmt.Println("Failed to marshal toSchool")
								return nil, errors.New("Failed to marshal toSchool => " + err.Error())
							}

							fmt.Println("Enroll transfered student and Write back toSchool")
							err = stub.PutState(toSchool.ID, toSchoolBytes)
							if err == nil {

								// remove student from transfer from school
								for i := range fromSchool.StudentsIds{
									if fromSchool.StudentsIds[i] == transaction.StudentID {
										fmt.Println("Found transfered student, removing from school")
										fromSchool.StudentsIds = append(fromSchool.StudentsIds[:i], fromSchool.StudentsIds[i+1:]...)
										jsonAsBytes, _ := json.Marshal(&fromSchool)
										err = stub.PutState(fromSchool.ID, jsonAsBytes)
										if err != nil {
											fmt.Println("Failed to write back fromSchool")
											return nil, err
										}
										break

									}
								}

								// fromSchool.StudentsIds = delete(fromSchool.StudentsIds, transaction.StudentID)
								// fmt.Println("Remove student and Write back fromSchool")
								// err = stub.PutState(fromSchool.ID, fromSchoolBytes)

								// change students school 
								// 1. Check if student exist - done above

								// 2. Unmarshal student
								var	student Student
								err = json.Unmarshal(studentBytes, &student)

								if err != nil {
									fmt.Println("Failed to Unmarshal Student")
									return nil, errors.New("Failed to Unmarshal student => " + err.Error())
								}

								// 3. change student school and marshal back
								student.School = toSchool.ID
								studentBytes, err = json.Marshal(&student)
								if err != nil {
									fmt.Println("Failed to Marshal Student")
									return nil, errors.New("Failed to Marshal student => " + err.Error())
								}

								// 4. update and write back
								err = stub.PutState(student.ID, studentBytes)
								if err != nil {
									fmt.Println("Failed to update Student")
									return nil, errors.New("Failed to update student => " + err.Error())
								}


								fmt.Println("Sucessfully transfered Student: " + transaction.StudentID + " From school: " +fromSchool.Name + " To school: " + toSchool.Name)
								return nil, nil
								
									
								// }else {
								// 	fmt.Println("Failed to write back toSchool")
								// 	return nil, errors.New("Failed to write back toSchool " + toSchool.ID + " => " +err.Error())
								// }
							}else {
								fmt.Println("Failed to write back toSchool")
								return nil, errors.New("Failed to write back toSchool " + toSchool.ID + " => " +err.Error())
							}

						} else {
							fmt.Println("Error Unmarshalling ToSchool for ID: " + transaction.ToSchool)
							return nil, errors.New("Error Unmarshalling FromSchool for ID: " + transaction.ToSchool)
						}
						
					} else {
						fmt.Println("ToSchool not found for ID: " + transaction.ToSchool)
						return nil, errors.New("School to transfer to doesnt exist for ID: " + transaction.ToSchool)
					}
				} else {
					fmt.Println("Error Unmarshalling FromSchool for ID: " + transaction.FromSchool)
					return nil, errors.New("Error Unmarshalling FromSchool for ID: " + transaction.FromSchool)
				}

			} else {
				fmt.Println("FromSchool not found for ID: " + transaction.FromSchool)
				return nil, errors.New("School to transfer from doesnt exist for ID: " + transaction.FromSchool)
			} 
		} else {
			fmt.Println("Student not found")
			return nil, errors.New("Student not found for ID: "+ transaction.StudentID)
		}
	// } else {
	// 	fmt.Println("Error Unmarshalling transaction")
	// 	return nil, errors.New("Invalid transaction " + " => " + err.Error())
	// }


}


func getStudent(studentId string, stub *shim.ChaincodeStub) (Student, error) {
	var student Student

	studentBytes, err := stub.GetState(studentId)

	if err != nil {
		fmt.Println("Account not found for : " + studentId)
		return student, errors.New("Error retrieving student: " + studentId)
	}

	err = json.Unmarshal(studentBytes, &student)
	if err != nil {
		fmt.Println("Error unmarshling student: " + studentId)
		return student, errors.New("Error unmarshling student: " + studentId)
	}

	return student, nil
}

func getSchool(schoolId string, stub *shim.ChaincodeStub) (School, error) {
	var school School

	schoolBytes, err := stub.GetState(schoolId)

	if err != nil {
		fmt.Println("School not found for : " + schoolId)
		return school, errors.New("Error retrieving school: " + schoolId)
	}

	err = json.Unmarshal(schoolBytes, &school)
	if err != nil {
		fmt.Println("Error unmarshling student: " + schoolId)
		return school, errors.New("Error unmarshling school: " + schoolId)
	}

	return school, nil
}

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)

	if function == "registerStudent" {
		fmt.Println("Firing registerStudent")
		//Create an asset with some value
		return t.registerStudent(stub, args)
	} else if function == "createSchool" {
		fmt.Println("Firing createSchool")
		//Create an asset with some value
		return t.createSchool(stub, args)
	} else if function == "transferStudent" {
		fmt.Println("Firing transferStudent")
		//Create an asset with some value
		return t.transferStudent(stub, args)
	} else if function == "init" {
		fmt.Println("Firing init")
		//Create an asset with some value
		return t.Init(stub,function, args)
	}



	return nil, errors.New("Received unknown function invocation")
}



// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	//need one arg
	if len(args) < 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ......")
	}

	if function == "getStudentById" {
		fmt.Println("Get a particular student")
		student, err := getStudent(args[0], stub)

		if err != nil {
			fmt.Println("Error getting student")
			return nil, err
		}else {
			studentBytes, err1 := json.Marshal(&student)
			if err1 != nil {
				fmt.Println("Error marshaling the student")
				return nil, err1
			}
			fmt.Println("All success, returning student")
			return studentBytes, nil
		}

	} else if function == "getSchoolById" {
		fmt.Println("Get a particular school")
		school, err := getSchool(args[0], stub)

		if err != nil {
			fmt.Println("Error getting school")
			return nil, err
		}else {
			schoolBytes, err1 := json.Marshal(&school)
			if err1 != nil {
				fmt.Println("Error marshaling the school")
				return nil, err1
			}
			fmt.Println("All success, returning school")
			return schoolBytes, nil
		}
	}

	return nil, errors.New("Unsupported operation")
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error starting Simple chaincode: %s", err)
	}
}


