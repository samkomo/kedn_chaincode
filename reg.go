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