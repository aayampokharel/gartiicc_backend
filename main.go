package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"nhooyr.io/websocket"
)

var (
	mapOfConnection  = make(map[*websocket.Conn]string)
	mapOfConnections = make(map[*websocket.Conn]bool)
	mapForStream     = make(map[*websocket.Conn]bool)
	mutex            = &sync.Mutex{}
	mutex1           = &sync.Mutex{}
	mutex2           = &sync.Mutex{}
	mutex3           = &sync.Mutex{}
	mutex4           = &sync.Mutex{}
	
	fromReturn       = &sync.Mutex{}
	GlobalCurrentName string
	index            int
	ind             int
	countForTimer int
	toogleForTimer bool=true
	listOfAllNames   = []string{}
	nameForMap        =make(chan string)
	afterDeleteFromSlice=make(chan bool)
)

type MessageText struct {
	Name    string `json:"Name"`
	Message string `json:"Message"`
}

func deleteFromSlice(slice []string,name string)[]string{
	
for i:=range slice {
if(slice[i] == name){
	fmt.Print(i,name,"\n\n");
	return append(slice[:i],slice[i+1:]...);
}
}
fmt.Print("\n\nname bhettena ni sahti \n\n")
return slice;
}
func main() {
	//! fn is for sending message with name like aayam:hello and broadcasting 
	//! to all 
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"*"},
		})
		if err != nil {
			
			c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted");
			return;
		}
		 //as the listt will be updated after the mutex2 only.
		

			mapOfConnection[c] = <-nameForMap;
		
		
		for {
			_, channelJson, err := c.Read(r.Context())
			if err != nil {
				fromReturn.Lock();
			mutex2.Lock()
			fmt.Print("❤❤❤")
			listOfAllNames=deleteFromSlice(listOfAllNames,mapOfConnection[c]);
			if(mapOfConnection[c]==GlobalCurrentName){

				toogleForTimer =false;
				fmt.Print("yellow overriidden");
			}
			fmt.Print("❤❤❤")
			mutex2.Unlock();
			delete(mapOfConnection,c);
			
			//mutex5.Unlock();
			//mutex1.Unlock();
			fromReturn.Unlock();
			log.Println(err);
			c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted");
				
			
				fmt.Print("error in channeljson read");
				return;
			}
			
			var msgText MessageText;
			json.Unmarshal(channelJson, &msgText)
			
			
			if msgText.Message != "" {
				mutex1.Lock()
			

				for k := range mapOfConnection {
					k.Write(r.Context(), websocket.MessageText, (channelJson))
					
				}
				mutex1.Unlock()
			} else {
				fmt.Print("error ")
			}
		}
	})

	paint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"*"},
		})
		if err != nil {
		
			c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
			return;
			
			
		}
	
		mapOfConnections[c] = true

		for {
			_, paintJson, err := c.Read(r.Context())
			if err != nil {
				delete(mapOfConnections,c);
			// mutex.Unlock();
			c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
			
				log.Println("error reading message:yaar error aayo ni ")
				return;
			
			}
			mutex.Lock()
			for k := range mapOfConnections {
				k.Write(r.Context(), websocket.MessageText, paintJson)
			}
			mutex.Unlock()
		}
	})
//  	listOfNames := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	 func(){
// 		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
// 			OriginPatterns: []string{"*"},
// 		})
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
		
// 		mapOfConnections[c] = true;//! this is where true is again set . 
		
		
// 		defer c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
		
		
// 		for{
// _,_,err:=	c.Read(r.Context())
// if err!=nil{
// 	fmt.Print("\n\nexit garyo dost le from eeuta for \n\n ");
// 	return;
// }
// 	mutex5.Lock()
// 	listOfNamesJson,_:=json.Marshal(listOfAllNames);
// 	for kk := range mapOfConnections {
		
// 		kk.Write(r.Context(), websocket.MessageText,((listOfNamesJson)));
// 	}
	
// 	mutex5.Unlock();

// }
// //_ ,value,_:=c.Read(r.Context())
// //if string(value)=="true"{

	
// // 	}else {  //! the format is ["name","points"]   and return ["name", "position"];
// // 	jsonHoldList :=make([]string,2);
	
// // 	mutex5.Lock()

// // 	json.Unmarshal(value,&jsonHoldList);
// // 			listOfPoints:=make([]int,len(listOfAllNames));
// // 			for ind,val:=range listOfAllNames{

// // 				if val==jsonHoldList[0]{
// // 					listOfPoints[ind],_=strconv.Atoi(jsonHoldList[1]);
// // 	sort.Sort(sort.Reverse(sort.IntSlice(listOfPoints)));
// // 					for i,point:=range listOfPoints{
// // 						if convertedPoint,_:=strconv.Atoi(jsonHoldList[1]);point==convertedPoint {
// // 							jsonHoldList[1]=strconv.Itoa(i);
// //                             break;
// // 						}<-
// // 					}
// //              break;

// // }

// // 			}


// // 			mutex5.Unlock();
// // 			sendingListJson,_:=json.Marshal(jsonHoldList);

// // 			for kk := range mapOfConnections {
				
// // 				kk.Write(r.Context(), websocket.MessageText,((sendingListJson)));
				
// // 			}

// // }
		


// 	}();
	
// 	}

	check := http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		

			ca, err := websocket.Accept(w, r, &websocket.AcceptOptions{
				OriginPatterns: []string{"*"},
			})
			if err != nil { //
				
				ca.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
				return;
				
			}
			mutex4.Lock()
			mapForStream[ca] = true //!also use lock 
					
			// Send the current player name immediately
			if(len(listOfAllNames)==0){
				return;
			}
			er:=ca.Write(r.Context(), websocket.MessageText, []byte(listOfAllNames[0]))
			mutex4.Unlock()
			if er!=nil{
				fmt.Print("after nil return error check✔✔✔✔✔\n\n\n");
				fromReturn.Lock();
				//mutex3.Unlock();
				//mutex4.Unlock();
				delete(mapForStream,ca);///huncha kki hunna as outsiide loop
				fromReturn.Unlock();
				ca.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
				return;
			}
			
			// Periodically update the player turn
			
			mutex3.Lock(); 
			for {//# esma khasma only broadcast kaam nothing else broadcast of break word for red container and broadcast of currentName on which drawing board is to be displayed.
				print("hello brother  💦💦💦💦");
				
				print("hello brother  💦💦💦💦");
				toogleForTimer=true; 
				for toogleForTimer {

					time.Sleep(time.Second*1)  
					//print(countForTimer,"====----====----====----====----");
				if	countForTimer=countForTimer+1;countForTimer==5{
countForTimer=0;
					break;
				}
				//@ this is the ensure that the timer works nicely. after 5sec completes it automatically exits but if in between the drawer exits then also this will not  makee others wait. toogleForTimer is used as =false in above exit sectiion so this method will exit and again set to true for next iteration.

			}
			ind++;
			for kz := range mapForStream {
				
				kz.Write(r.Context(), websocket.MessageText, []byte("Break"))
				//! at this pause is for the red container at which post is made for getting the 
				//! value for drawing.
			}
		//?problrm here is small one now ::::   jaba break bhanne sab lai pathaisake pachi red container display huncha and if another player gets added then the red cotainer is not displayed as break bhanne keyword pathaudaiina . or eevery second pathauna paryo break keyword . But this is only for that period of time after that everything is NORMAL.
						time.Sleep(time.Second * 5);
					//!pause for red container
				
					if (index + 1) < len(listOfAllNames) {
						index = index + 1
						} else {
							index = 0
						}
						if(len(listOfAllNames)==0){
							return;
						}
					currentName := listOfAllNames[index]
					GlobalCurrentName=currentName;
				
					for kz := range mapForStream {
					
						
						kz.Write(r.Context(), websocket.MessageText, []byte(currentName))
					
					}
					fmt.Print("before nil return error check✔✔✔✔✔\n\n\n");
					if(len(listOfAllNames) == 0) {
						mutex3.Unlock();
						return;
					};
				}
		
	
		
	})

	currentCheck := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //!takes the current player name from psot in the future ko reesponse bhanne bata and then the value is insesrted in the lsit to update the list of playeers . This returnss first player  if yes first player lai dine the yellow wala drawing option . 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")

		var nameOfCurrentPlayer string
		byteName, _ := io.ReadAll(r.Body)
		json.Unmarshal(byteName, &nameOfCurrentPlayer)
		
		mutex2.Lock()
		
		listOfAllNames = append(listOfAllNames, nameOfCurrentPlayer)
		nameForMap<-nameOfCurrentPlayer;
		
		w.Write([]byte(listOfAllNames[0]))
		mutex2.Unlock()
		
	
	})

	listOfNamesInDrawer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //!takes the current player name from psot in the future ko reesponse bhanne bata and then the value is insesrted in the lsit to update the list of playeers . This returnss first player  if yes first player lai dine the yellow wala drawing option . 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		if(len(listOfAllNames)==0){
			return;
		}
		listAsJson,_:=json.Marshal(listOfAllNames);
		
		w.Write([]byte(listAsJson))

	});

	listOfWords:=http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
	
		slices :=[]string{"ball","brother"};
		if ind==len(slices){
         ind=0;//lock system for shared variable.
        }
		sliceJson,_:=json.Marshal(slices[ind]);
		 //returning just a single value from slice

		w.Write((sliceJson)); 
		
	},);
	chirouter := chi.NewRouter()
	chirouter.Get("/", fn)
	chirouter.Get("/listofnames", listOfNamesInDrawer)
	chirouter.Post("/currentcheck", currentCheck)
	chirouter.Get("/paint", paint)
	chirouter.Get("/check", check)
	chirouter.Get("/listofwords", listOfWords)

 http.ListenAndServe("localhost:8080", chirouter)

} 