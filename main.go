package main

import (
	"encoding/json"
	"fmt"

	"io"

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
	mutex4LockCheck bool=false;//@initially set to false and true when mutex4 acquired./. 
	listOfAllNames   = []string{}
	nameForMap        =make(chan string)
	
)

type MessageText struct {
	Name    string `json:"Name"`
	Message string `json:"Message"`
}

func deleteFromSlice(slice []string,name string)[]string{
	
for i:=range slice {
if(slice[i] == name){

	return append(slice[:i],slice[i+1:]...);
}
}

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
			if c!=nil{
fmt.Print("from error 1");
				c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted");
			}
			return;
			
		}
		 //as the list will be updated after the mutex2 only.
		

			mapOfConnection[c] = <-nameForMap;
		
		
		for {
			_, channelJson, err := c.Read(r.Context())
			if err != nil {
				if(c!=nil){
					fmt.Print("from error 2");
				fromReturn.Lock();
			mutex2.Lock() //? whhy this extra layer of lock? 
			
			listOfAllNames=deleteFromSlice(listOfAllNames,mapOfConnection[c]);
			if(mapOfConnection[c]==GlobalCurrentName){
fmt.Print("✔✔✔match 🔥🔥🔥");
toogleForTimer =false;//@ thisis a good one as green le exit garyo bhane baal chaina  



			}
		
			mutex2.Unlock();
			delete(mapOfConnection,c);
			
		
			fromReturn.Unlock();
		

				c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted");
			}
				
			
				
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
				mutex1.Lock()
			

				for k := range mapOfConnection {
					k.Write(r.Context(), websocket.MessageText, []byte(" "))
					
				}
				mutex1.Unlock()
				
			}
		}
	})

	paint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"*"},
		})
		if err != nil {
		 if c!=nil{
			fmt.Print("from error 3");
			 c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
		 }
			return;
			
			
		}
	
		mapOfConnections[c] = true

		for {
			_, paintJson, err := c.Read(r.Context())
			if err != nil {
				delete(mapOfConnections,c);
			// mutex.Unlock();
			if c!=nil{
				fmt.Print("from error 4");
				c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
			}
			
				
				return;
			
			}
			mutex.Lock()
			for k := range mapOfConnections {
				k.Write(r.Context(), websocket.MessageText, paintJson)
			}
			mutex.Unlock()
		}
	})


	check := http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) { 
		
fmt.Print("from starting of check");
			ca, err := websocket.Accept(w, r, &websocket.AcceptOptions{
				OriginPatterns: []string{"*"},
			})
			if err != nil { //
				if ca!=nil{
					fmt.Print("from error 4");
					if mutex4LockCheck{
	
						defer mutex4.Unlock();
					}
					ca.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
				}
				return;
				
			}
			mutex4.Lock()
			
			fmt.Print("\n\n lock problem ","\n\n");
			mutex4LockCheck=true;
			mapForStream[ca] = true //!also use lock 
					
			// Send the current player name immediately
			if(len(listOfAllNames)==0){
				mutex4.Unlock()
				return;
			}
			er:=ca.Write(r.Context(), websocket.MessageText, []byte(listOfAllNames[0]))
		
			mutex4.Unlock()
			if er!=nil{
				if ca!=nil{
					
					
					fromReturn.Lock();
					
					delete(mapForStream,ca);///huncha kki hunna as outsiide loop
					fromReturn.Unlock();
					ca.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
				}
				return;
			}
			
			// Periodically update the player turn
			
			mutex3.Lock(); 
			defer mutex3.Unlock();
	
			fmt.Print("list empty problem how>? ")
			GlobalCurrentName=listOfAllNames[0];
			for {//# esma khasma only broadcast kaam nothing else broadcast of break word for red container and broadcast of currentName on which drawing 
			
				for toogleForTimer {
					fmt.Print("insiide loop aaecha ta ferii ",countForTimer);
					time.Sleep(time.Second*1)  
					fmt.Print("\n\ntiimer\n\n")
					if	countForTimer=countForTimer+1;countForTimer==7{
						fmt.Print("\n\ncountForTimer set to 0\n\n")
						break;
						} 
					}
					toogleForTimer=true; 
					//@ this is the ensure that the timer works nicely. after 5sec completes it automatically exits but if in between the drawer exits then also this will not  makee others wait. toogleForTimer is used as =false in above exit sectiion so this method will exit and again set to true for next iteration.
					
					countForTimer=0;
					ind++;
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
						
						kz.Write(r.Context(), websocket.MessageText, []byte("Break"))
						//! at this pause is for the red container at which post is made for getting the 
						//! value for drawing.
					}
		//?problrm here is small one now ::::   jaba break bhanne sab lai pathaisake pachi red container display huncha and if another player gets added then the red cotainer is not displayed as break bhanne keyword pathaudaiina . or eevery second pathauna paryo break keyword . But this is only for that period of time after that everything is NORMAL.
		if(len(listOfAllNames) == 0) {
			fmt.Print("length 0 bhayo ni ")
			return;
			};//!yo condition kina mathi lyako hola bhanda to return the lock if ii dont then yo sleep ta 2o sec ho chalcha sure and lock dinu parcha so unlock garna parchha so that aaune le 40 seconds chai atleast chalairakhos . and as long as the timer of below is lesser no problem . baal chaina yes eeuta extrra resource chai consume bhaiirako cha.  
			
			//@ COMPULSORILY yo red container lai we can use channel AS NABHAE SLEEP HUDA ARKO AAUNE LE 10 SEC EXTRA KURNA PARCHA AS LOCK RELEASE BHAKO HUNNA . IF STATEMENT IS EXPENSIVE . 
			time.Sleep(time.Second * 5);
		
			//!pause for red container BREAK CONTAINER
			
			for kz := range mapForStream {
					
						
						kz.Write(r.Context(), websocket.MessageText, []byte(currentName))
					
					}
				fmt.Print("from iinsidie the functioin being ✔✔✔✔✔");
				
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
		if ind>=len(slices){
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

//http.ListenAndServe("0.0.0.0:10000", chirouter)
 http.ListenAndServe(":8080", chirouter)


} 