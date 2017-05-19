package main
func 

/*
	0 ActorId    string `json:"actorid"`
	1 ActorName  string `json:"actorname"`
	2 Committed  string  `json:"committed"`
	3 Reimbursed string `json:"reimbursed"`
	4 Awarded    string `json:"awarded"`
	5 Spent      string `json:"spent"`
	6 Received   string `json:"received"`
	7 Delegated  string `json:"delegated"`
*/

	// grantor info
	var act1  = make([]string{}, 8,8)
	act1[0] = "ACT-101"                //ActorId
	act1[1] = "PPM Foundation"     //ActorName
	act1[2] = "125000"             //Comiitted
	act1[3] = "19500"              //Reimbursed
	act1[4] = ""                   //Awarded
	act1[5] = ""                   //Spent
    act1[6] = ""                   //Received 
    act1[7] = ""                   //Delegated

	// grantee info
	var act2  = make([]string{}, 8,8)
 	act2[0] = "ACT-102"                     //ActorId
	act2[1] = "Stanford University"      //ActorName
	act2[2] = ""                        //Comiitted
	act2[3] = ""                         //Reimbursed
	act2[4] = "125000"                 //Awarded
	act2[5] = "23000"                  //Spent
    act2[6] = "10000"                 //Received 
    act2[7] = "45000"                 //Delegated

	// sub-grantee info
	var act3  = make([]string{}, 8,8)
 	act3[0] = "ACT-103"                //ActorId
	act3[1] = "John Hopkins University"     //ActorName
	act3[2] = ""               //Comiitted
	act3[3] = ""               //Reimbursed
	act3[4] = "45000"          //Awarded
	act3[5] = "12000"          //Spent
    act3[6] = "9500"          //Received 
    act3[7] = ""               //Delegated

    // Supplier info -- shows spending form all the grantees and sub-grantees
	var act4  = make([]string{}, 8,8)
 	act4[0] = "ACT-104"                //ActorId
	act4[1] = "Dixon consulting"     //ActorName
	act4[2] = ""               //Comiitted
	act4[3] = ""               //Reimbursed
	act4[4] = ""          //Awarded
	act4[5] = ""          //Spent
    act4[6] = "35000"          //Received 
    act4[7] = ""               //Delegated


	t.init_actor(stub, act1)
	t.init_actor(stub, act2)
	t.init_actor(stub, act3)
	t.init_actor(stub, act4)

	//**************************************

	type Award struct {
	AwardId        string `json:"awardid"`
	AwardName      string `json:"awardname"`	
	AwardAmount string `json:"awardamount"`	
	Expenses []Expenditure `json:"expenditures"`
	Reimburses []Reimbursement `json:"reimburses"`
}

   awdId := "AWD-10"
   awardName := "CANCER RESEARCH"
   awardAmount := "125000"
   
   awardStr := `{"awardid": "` + awdId + `", "awardname": "` + awardName + `", "awardamount": "` + awardAmount + 
	       `", {"` +
	       `"expenditures": "` + expStr0 + `", "` +
			                   + expStr1 
			 `"}`



	expStr0 := `{"expenditureid": "` + expId + `", "amount": "` + expAmount + `", "date": "` + expDate + 
	       `", "type": "` + expType + `", "status": "` + expStatus + `", "fromactor": "` + fromActor + 
	       `", "toactor": "` + toActor + `"}`


  	remStr1 := `{"reimbursementid": "` + remId + `", "amount": "` + remAmountStr + `", "fromactor": "` + remFromActor + 
	       `", "toActor": "` + remToActor + `", "date": "` + remDate + `", "expenditureid": "` + remExpId + 
	        `"}` 


//***********************************************
	type Reimbursement struct {
	ReimbursementId string `json:"reimbursementid"`
	Amount          int `json:"amount"`
	FromActor        string `json:"fromactor"`
	ToActor          string `json:"toactor"`
	Date            string `json:"date"`
	ExpenditureId string `json:"expenditureid"`
}

	// Reimbursement
	var rem1  = make([]string, 6,6)
	rem1[0] = "REM-301"                //ReimbursementId
	rem1[1] = "3000"                   //Amount
    rem1[2] = "ACT-101"                //FromActor  
    rem1[3] = "ACT-102"    	           //ToActor
	rem1[4] = "05-02-2017"             //Date	
	rem1[5] = "EXP-201"               //ExpenditureId

 	var rem2  = make([]string, 6,6)
	rem2[0] = "REM-302"                //ReimbursementId
	rem2[1] = "4000"                   //Amount
    rem2[2] = "ACT-101"                //FromActor  
    rem2[3] = "ACT-102"    	           //ToActor
	rem2[4] = "05-04-2017"             //Date	
	rem2[5] = "EXP-203"               //ExpenditureId   

	var rem3  = make([]string, 6,6)
	rem3[0] = "REM-303"                //ReimbursementId
	rem3[1] = "3000"                   //Amount
    rem3[2] = "ACT-101"                //FromActor  
    rem3[3] = "ACT-102"    	           //ToActor
	rem3[4] = "05-09-2017"             //Date	
	rem3[5] = "EXP-204"               //ExpenditureId 

	var rem4  = make([]string, 6,6)
	rem4[0] = "REM-304"                //ReimbursementId
	rem4[1] = "5000"                   //Amount
    rem4[2] = "ACT-101"                //FromActor  
    rem4[3] = "ACT-102"    	           //ToActor
	rem4[4] = "05-11-2017"             //Date	
	rem4[5] = "EXP-205"               //ExpenditureId 

	var rem5  = make([]string, 6,6)
	rem5[0] = "REM-305"                //ReimbursementId
	rem5[1] = "2000"                   //Amount
    rem5[2] = "ACT-102"                //FromActor  
    rem5[3] = "ACT-103"    	           //ToActor
	rem5[4] = "05-12-2017"             //Date	
	rem5[5] = "EXP-206"               //ExpenditureId 


	var rem6  = make([]string, 6,6)
	rem6[0] = "REM-306"                //ReimbursementId
	rem6[1] = "1000"                   //Amount
    rem6[2] = "ACT-102"                //FromActor  
    rem6[3] = "ACT-103"    	           //ToActor
	rem6[4] = "05-16-2017"             //Date	
	rem6[5] = "EXP-208"               //ExpenditureId 

	var rem7  = make([]string, 6,6)
	rem7[0] = "REM-307"                //ReimbursementId
	rem7[1] = "1500"                   //Amount
    rem7[2] = "ACT-102"                //FromActor  
    rem7[3] = "ACT-103"    	           //ToActor
	rem7[4] = "05-19-2017"             //Date	
	rem7[5] = "EXP-209"               //ExpenditureId 


// ============================================================================================================================
// Init account - create a new account, store into chaincode world state, and then append the account index
// ============================================================================================================================
func (t *SimpleChaincode) init_reimbursement(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	fmt.Println("- start init acount")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}



	remId := args[0]

	remAmount, err := strconv.ParseFloat(args[1], 64)   //strings.ToLower(args[1])

	if err != nil {
		return nil, errors.New("2nd argument must be a numeric string")
	}

    remFromActor := arg[2]
    remToActor := arg[3]
	remDate := arg[4]
    remExpId := arg[5]
    

	//check if account already exists
	accountAsBytes, err := stub.GetState(remId)
	if err != nil {
		return nil, errors.New("Failed to get expenditure id")
	}

	rem := Reimbursement{}
	json.Unmarshal(accountAsBytes, &rem)
	if rem.ReimbursementId == remId {
		return nil, errors.New("This reimbursement arleady exists")
	}

	
	remAmountStr: = strconv.FormatFloat(remAmount, 'f', -1, 64)

	//build the expenditure json string 
	str := `{"reimbursementid": "` + remId + `", "amount": "` + remAmountStr + `", "fromactor": "` + remFromActor + 
	       `", "toActor": "` + remToActor + `", "date": "` + remDate + `", "expenditureid": "` + remExpId + 
	        `"}`
	//jsonAsBytesActor, _ := json.Marshal(newActor)
	err = stub.PutState(remId, []byte(str))
	if err != nil {
		return nil, err
	}

	//get the account index
	accountsAsBytes, err := stub.GetState(accountIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get account index")
	}
	var accountIndex []string
	json.Unmarshal(accountsAsBytes, &accountIndex)

	//append the index 
	accountIndex = append(accountIndex, remId)
	jsonAsBytes, _ := json.Marshal(accountIndex)
	err = stub.PutState(accountIndexStr, jsonAsBytes)

	return nil, nil
}

//************************
/*
type Expenditure struct {
	ExpenditureId   string `json:"expenditureid"`
	Amount          float64 `json:"amount"`	
	Date            string `json:"date"`
	Type            string `json:"type"`
	Status          string `json:"status"`	
	FromActor        string `json:"fromactor"`
	ToActor          string `json:"toactor"`
	
}
*/


	// Expense
	var exp1  = make([]string, 7,7)
	exp1[0] = "EXP-201"                //ExpenditureId
	exp1[1] = "3000"     //Amount
	exp1[2] = "05-02-2017"             //Date
	exp1[3] = "Travel"              //Type
	exp1[4] = "Approved"                   //Status
    exp1[5] = "ACT-102"                   //FromActor --Grantee spending 
    exp1[6] = "ACT-104"                   //ToActor   --Supplier receiving the spending


	var exp2  = make([]string, 7,7)
	exp2[0] = "EXP-202"                //ExpenditureId
	exp2[1] = "8000"     //Amount
	exp2[2] = "05-03-2017"             //Date
	exp2[3] = "Equipment"              //Type
	exp2[4] = "Pending"                   //Status
    exp2[5] = "ACT-102"                   //FromActor --Grantee spending 
    exp2[6] = "ACT-104"                   //ToActor   --Supplier receiving the spending

	var exp3 = make([]string, 7,7)
	exp3[0] = "EXP-203"                //ExpenditureId
	exp3[1] = "4000"     //Amount
	exp3[2] = "05-04-2017"             //Date
	exp3[3] = "Training"              //Type
	exp3[4] = "Approved"                   //Status
    exp3[5] = "ACT-102"                   //FromActor --Grantee spending 
    exp3[6] = "ACT-104"  

    var exp4 = make([]string, 7,7)
	exp4[0] = "EXP-204"                //ExpenditureId
	exp4[1] = "3000"     //Amount
	exp4[2] = "05-09-2017"             //Date
	exp4[3] = "Software License"              //Type
	exp4[4] = "Approved"                   //Status
    exp4[5] = "ACT-102"                   //FromActor --Grantee spending 
    exp4[6] = "ACT-104"  

    var exp5 = make([]string, 7,7)
	exp5[0] = "EXP-205"                //ExpenditureId
	exp5[1] = "5000"     //Amount
	exp5[2] = "05-11-2017"             //Date
	exp5[3] = "Specimens"              //Type
	exp5[4] = "Approved"                   //Status
    exp5[5] = "ACT-102"                   //FromActor --Grantee spending 
    exp5[6] = "ACT-104"

    var exp6 = make([]string, 7,7)
	exp6[0] = "EXP-206"                //ExpenditureId
	exp6[1] = "2000"     //Amount
	exp6[2] = "05-12-2017"             //Date
	exp6[3] = "Consultancy"              //Type
	exp6[4] = "Approved"                   //Status
    exp6[5] = "ACT-103"                   //FromActor --Grantee spending 
    exp6[6] = "ACT-104"

    var exp7 = make([]string, 7,7)
	exp7[0] = "EXP-207"                //ExpenditureId
	exp7[1] = "7500"     //Amount
	exp7[2] = "05-15-2017"             //Date
	exp7[3] = "Equipment"              //Type
	exp7[4] = "Pending"                   //Status
    exp7[5] = "ACT-103"                   //FromActor --Grantee spending 
    exp7[6] = "ACT-104"

    var exp8 = make([]string, 7,7)
	exp8[0] = "EXP-208"                //ExpenditureId
	exp8[1] = "1000"     //Amount
	exp8[2] = "05-16-2017"             //Date
	exp8[3] = "Travel"              //Type
	exp8[4] = "Approved"                   //Status
    exp8[5] = "ACT-103"                   //FromActor --Grantee spending 
    exp8[6] = "ACT-104"

    var exp9 = make([]string, 7,7)
	exp9[0] = "EXP-209"                //ExpenditureId
	exp9[1] = "1500"     //Amount
	exp9[2] = "05-19-2017"             //Date
	exp9[3] = "Training"              //Type
	exp9[4] = "Approved"                   //Status
    exp9[5] = "ACT-103"                   //FromActor --Grantee spending 
    exp9[6] = "ACT-104"

	t.init_expenditure(stub, exp1)
	t.init_expenditure(stub, exp2)
	t.init_expenditure(stub, exp3)
	t.init_expenditure(stub, exp4)
	t.init_expenditure(stub, exp5)
	t.init_expenditure(stub, exp6)
	t.init_expenditure(stub, exp7)
	t.init_expenditure(stub, exp8)
	t.init_expenditure(stub, exp9)

****************************

	// ============================================================================================================================
// Init account - create a new account, store into chaincode world state, and then append the account index
// ============================================================================================================================
func (t *SimpleChaincode) init_expenditure(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	//       0        1      2..
	// "accountid", "name",  ...

	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7")
	}

	//input sanitation
	fmt.Println("- start init acount")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}

/*
     exp9[0] = "EXP-209"                //ExpenditureId
	exp9[1] = "1500"     //Amount
	exp9[2] = "05-19-2017"             //Date
	exp9[3] = "Training"              //Type
	exp9[4] = "Approved"                   //Status
    exp9[5] = "ACT-103"                   //FromActor --Grantee spending 
    exp9[6] = "ACT-104"
*/

	expId := args[0]

	expAmount, err := strconv.ParseFloat(args[1], 64)   //strings.ToLower(args[1])

	if err != nil {
		return nil, errors.New("2nd argument must be a numeric string")
	}

    expDate := arg[2]
    expType := arg[3]
    expStatus := arg[4]
    fromActor := arg[5]
    toActor := arg[6]

	
	
	//check if account already exists
	accountAsBytes, err := stub.GetState(expId)
	if err != nil {
		return nil, errors.New("Failed to get expenditure id")
	}

	exp := Expenditure{}
	json.Unmarshal(accountAsBytes, &exp)
	if exp.ExpenditureId == expId {
		return nil, errors.New("This expenditure arleady exists")
	}

	/*
	committedStr := strconv.FormatFloat(committed, 'f', -1, 64)
	reimbursedStr := strconv.FormatFloat(reimbursed, 'f', -1, 64)
	awardedStr := strconv.FormatFloat(awarded, 'f', -1, 64)
	spentStr := strconv.FormatFloat(spent, 'f', -1, 64)
	receivedStr := strconv.FormatFloat(received, 'f', -1, 64)
	delegatedStr := strconv.FormatFloat(delegated, 'f', -1, 64)
  */
	expAmountStr: = strconv.FormatFloat(expAmount, 'f', -1, 64)

	//newActor := Actor{}
	//newActor.ActorId = actorId
	//newActor.ActorName = actorName
	//newActor.Balance = balance

//	   expDate := arg[2]
  ///  expType := arg[3]
   /// expStatus := arg[4]
    //fromActor := arg[5]
    //toActor := arg[6]


	//build the expenditure json string 
	str := `{"expenditureid": "` + expId + `", "amount": "` + expAmount + `", "date": "` + expDate + 
	       `", "type": "` + expType + `", "status": "` + expStatus + `", "fromactor": "` + fromActor + 
	       `", "toactor": "` + toActor + `"}`
	//jsonAsBytesActor, _ := json.Marshal(newActor)
	err = stub.PutState(expId, []byte(str))
	if err != nil {
		return nil, err
	}

	//get the account index
	accountsAsBytes, err := stub.GetState(accountIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get account index")
	}
	var accountIndex []string
	json.Unmarshal(accountsAsBytes, &accountIndex)

	//append the index 
	accountIndex = append(accountIndex, expId)
	jsonAsBytes, _ := json.Marshal(accountIndex)
	err = stub.PutState(accountIndexStr, jsonAsBytes)

	return nil, nil
}

