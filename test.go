package main

//		accountSid := "ACaad6ab76876e7822323bbe3f91106810"
//		authToken := "97214b0dd56c113e81d4524cad8545a8"
//		urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
//
//		// Create possible message bodies
//		quotes := "Йоу, вот твой код смс для входа в приложение\n 342121"
//
//		// Set up rand
//		rand.Seed(time.Now().Unix())
//
//		msgData := url.Values{}
//		msgData.Set("To", "+79674781443")
//		msgData.Set("From", "+13607039136")
//		msgData.Set("Body", quotes)
//		msgDataReader := *strings.NewReader(msgData.Encode())
//
//		client := &http.Client{}
//		req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
//		req.SetBasicAuth(accountSid, authToken)
//		req.Header.Add("Accept", "application/json")
//		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//
//		resp, _ := client.Do(req)
//		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
//			var data map[string]interface{}
//			decoder := json.NewDecoder(resp.Body)
//			err := decoder.Decode(&data)
//			if err == nil {
//				fmt.Println(data["sid"])
//			}
//		} else {
//			fmt.Println(resp.Status)
//		}
