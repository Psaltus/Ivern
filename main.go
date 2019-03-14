package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//types allow for custom variable types, as well as custom structures to the variable
type Output struct {
	SummonerName              string
	ProfileIconID             int
	AccountID                 int
	HighestAchievedSeasonTier string

	Match []struct {
		Name       string
		Lane       string `json:"lane"`
		GameID     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		PlatformID string `json:"platformId"`
		Timestamp  int64  `json:"timestamp"`
		Queue      int    `json:"queue"`
		Role       string `json:"role"`
		Season     int    `json:"season"`
		Image      string
		Stats      struct {
			ParticipantIdentities []struct {
				Player struct {
					CurrentPlatformID string `json:"currentPlatformId"`
					SummonerName      string `json:"summonerName"`
					MatchHistoryURI   string `json:"matchHistoryUri"`
					PlatformID        string `json:"platformId"`
					CurrentAccountID  int    `json:"currentAccountId"`
					ProfileIcon       int    `json:"profileIcon"`
					SummonerID        int    `json:"summonerId"`
					AccountID         int    `json:"accountId"`
				} `json:"player"`
				ParticipantID int `json:"participantId"`
			} `json:"participantIdentities"`
			ParticipantID int
			GameVersion   string `json:"gameVersion"`
			PlatformID    string `json:"platformId"`
			GameMode      string `json:"gameMode"`
			MapID         int    `json:"mapId"`
			GameType      string `json:"gameType"`
			Teams         []struct {
				FirstDragon bool `json:"firstDragon"`
				Bans        []struct {
					PickTurn   int `json:"pickTurn"`
					ChampionID int `json:"championId"`
				} `json:"bans"`
				FirstInhibitor       bool   `json:"firstInhibitor"`
				Win                  string `json:"win"`
				FirstRiftHerald      bool   `json:"firstRiftHerald"`
				FirstBaron           bool   `json:"firstBaron"`
				BaronKills           int    `json:"baronKills"`
				RiftHeraldKills      int    `json:"riftHeraldKills"`
				FirstBlood           bool   `json:"firstBlood"`
				TeamID               int    `json:"teamId"`
				FirstTower           bool   `json:"firstTower"`
				VilemawKills         int    `json:"vilemawKills"`
				InhibitorKills       int    `json:"inhibitorKills"`
				TowerKills           int    `json:"towerKills"`
				DominionVictoryScore int    `json:"dominionVictoryScore"`
				DragonKills          int    `json:"dragonKills"`
			} `json:"teams"`
			Participants []struct {
				Stats struct {
					HighestStreak                   string
					Item1                           int  `json:"item1"`
					TotalPlayerScore                int  `json:"totalPlayerScore"`
					VisionScore                     int  `json:"visionScore"`
					UnrealKills                     int  `json:"unrealKills"`
					Win                             bool `json:"win"`
					ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
					LargestCriticalStrike           int  `json:"largestCriticalStrike"`
					TotalDamageDealt                int  `json:"totalDamageDealt"`
					MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
					LargestMultiKill                int  `json:"largestMultiKill"`
					LargestKillingSpree             int  `json:"largestKillingSpree"`
					QuadraKills                     int  `json:"quadraKills"`
					TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
					MagicalDamageTaken              int  `json:"magicalDamageTaken"`
					LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
					NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
					FirstTowerAssist                bool `json:"firstTowerAssist"`
					NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
					GoldEarned                      int  `json:"goldEarned"`
					Item2                           int  `json:"item2"`
					Item3                           int  `json:"item3"`
					Item0                           int  `json:"item0"`
					Deaths                          int  `json:"deaths"`
					Item6                           int  `json:"item6"`
					WardsPlaced                     int  `json:"wardsPlaced"`
					Item4                           int  `json:"item4"`
					Item5                           int  `json:"item5"`
					TurretKills                     int  `json:"turretKills"`
					TripleKills                     int  `json:"tripleKills"`
					DamageSelfMitigated             int  `json:"damageSelfMitigated"`
					GoldSpent                       int  `json:"goldSpent"`
					MagicDamageDealt                int  `json:"magicDamageDealt"`
					Kills                           int  `json:"kills"`
					DoubleKills                     int  `json:"doubleKills"`
					FirstInhibitorKill              bool `json:"firstInhibitorKill"`
					TrueDamageTaken                 int  `json:"trueDamageTaken"`
					FirstBloodAssist                bool `json:"firstBloodAssist"`
					FirstBloodKill                  bool `json:"firstBloodKill"`
					Assists                         int  `json:"assists"`
					TotalScoreRank                  int  `json:"totalScoreRank"`
					NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
					CombatPlayerScore               int  `json:"combatPlayerScore"`
					VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
					DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
					PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
					PentaKills                      int  `json:"pentaKills"`
					TrueDamageDealt                 int  `json:"trueDamageDealt"`
					TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
					ChampLevel                      int  `json:"champLevel"`
					ParticipantID                   int  `json:"participantId"`
					FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
					WardsKilled                     int  `json:"wardsKilled"`
					FirstTowerKill                  bool `json:"firstTowerKill"`
					TotalHeal                       int  `json:"totalHeal"`
					TotalMinionsKilled              int  `json:"totalMinionsKilled"`
					PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
					DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
					SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
					TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
					TotalUnitsHealed                int  `json:"totalUnitsHealed"`
					InhibitorKills                  int  `json:"inhibitorKills"`
					TotalDamageTaken                int  `json:"totalDamageTaken"`
					KillingSprees                   int  `json:"killingSprees"`
					TimeCCingOthers                 int  `json:"timeCCingOthers"`
					PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
				} `json:"stats"`
				Spell1ID      int `json:"spell1Id"`
				ParticipantID int `json:"participantId"`
				Runes         []struct {
					RuneID int `json:"runeId"`
					Rank   int `json:"rank"`
				} `json:"runes"`
				HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
				Masteries                 []struct {
					MasteryID int `json:"masteryId"`
					Rank      int `json:"rank"`
				} `json:"masteries"`
				Spell2ID   int `json:"spell2Id"`
				TeamID     int `json:"teamId"`
				ChampionID int `json:"championId"`
				Spell1Full string
				Spell2Full string
			} `json:"participants"`

			GameDuration int   `json:"gameDuration"`
			GameCreation int64 `json:"gameCreation"`
		}
	} `json:"matches"`
}

type Profile struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     int    `json:"accountId"`
	ID            int    `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

type playerMatches struct {
	Match []struct {
		Name      string
		Lane      string `json:"lane"`
		Champion  int    `json:"champion"`
		Season    string `json:"season"`
		Region    string `json:"region"`
		MatchID   int64  `json:"matchId"`
		Queue     string `json:"queue"`
		Role      string `json:"role"`
		Timestamp int64  `json:"timestamp"`
	} `json:"matches"`
}

type Champion struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
}

type SpellInfo struct {
	Image struct {
		Full   string `json:"full"`
		Group  string `json:"group"`
		Sprite string `json:"sprite"`
		H      int    `json:"h"`
		W      int    `json:"w"`
	} `json:"image"`
}

//The web template we'll be using for our search webpage.
const internalTempl = `
<title>{{.SummonerName}}'s Stats - Ivern</title>
<style>
body{
	background-color: #393939;
	color: #FFFFFF;
}
h1{
	text-align: center;
}
h1, h2{
    font-family: sans-serif;
}
table{
	font-family: sans-serif;
	padding: 10px;
}
td{
    text-align: center;
    vertical-align: middle;
    padding: 7px;
}
a:link {
    color: LightGreen;
}
a:visited{
	color: LightGreen;
}
</style>
<body>
<form method="POST" action="/search">
     <input type="text" name="Search" placeholder="Summoner Name" autocomplete="off" required/> <input class="submit" type="submit" value="Search Summoner"/>
</form>
<h1>{{ .SummonerName }} <img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/profileicon/{{.ProfileIconID}}.png" height=70px width=70px> </h1>
<h1>Highest Rank this Season: {{ .HighestAchievedSeasonTier }}</h1> <br/><br/>
Here's your match history:<br/>
	
{{range .Match}}
{{$PartID := .Stats.ParticipantID}}
_________________________________________________________________________________________________<br/>
<h2 div="ChampionHeader"><img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/champion/{{.Name}}.png" height="60px" width="60px"></img> <i>{{.Name}}</i></h2> <a href="http://matchhistory.na.leagueoflegends.com/en/#match-details/NA1/{{.GameID}}/{{$.AccountID}}?tab=overview">Link to Official Stats!</a> 
			
			{{range .Stats.Participants}}
				{{if eq $PartID .ParticipantID}}
				<table border=1 {{if .Stats.Win}} bordercolor="GREEN" {{else}} bordercolor="RED" {{end}}>	
					<tr>
						<td rowspan=2><img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/spell/{{ .Spell1Full }}" height="60px" width="60px"></img><br/>
						<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/spell/{{ .Spell2Full }}" height="60px" width="60px"></img></td>
						
						<td> <img src="http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/score.png"></img><br/>
							 <b div="kda"> {{ .Stats.Kills }} / {{ .Stats.Deaths }} / {{ .Stats.Assists }} </b></td>
						
						<td>
						{{if .Stats.Win }}
							<b>Victory!</b> <br/>
							Ranked Draft Mode (5v5)
						{{else}}
							<b>Defeat!</b> <br/>
							Ranked Draft Mode (5v5)
						{{end}}
						</td>
						<td>
							<i><b>{{.Stats.HighestStreak}}</b></i>
						</td>
						<td rowspan=2>
						<img src="http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/items.png"></img> <br/>
							{{if gt .Stats.Item0 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item0 }}.png" width="60px" height="60px"></img>
							{{end}}
							{{if gt .Stats.Item1 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item1 }}.png" width="60px" height="60px"></img>
							{{end}}
							{{if gt .Stats.Item2 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item2 }}.png" width="60px" height="60px"></img>
							{{end}}<br/>
							{{if gt .Stats.Item3 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item3 }}.png" width="60px" height="60px"></img>
							{{end}}
							{{if gt .Stats.Item4 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item4 }}.png" width="60px" height="60px"></img>
							{{end}}
							{{if gt .Stats.Item5 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item5 }}.png" width="60px" height="60px"></img>
							{{end}}
							{{if gt .Stats.Item6 0}}
							<img src="http://ddragon.leagueoflegends.com/cdn/7.16.1/img/item/{{ .Stats.Item6 }}.png" width="60px" height="60px"></img>
							{{end}}
						</td>
					</tr>
					<tr>
						<td>
						<img src="http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/minion.png"></img><br/>
						<b div="kda">{{ .Stats.TotalMinionsKilled }} CS</b>
						</td>
						<td>
						<img src="http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/gold.png"></img><br/>
						{{ .Stats.GoldEarned }}
						</td>
						<td>
						<img src="http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/champion.png"></img><br/>
						Level: {{ .Stats.ChampLevel }}
						</td>
					</tr>
				</table>
				{{end}}
			{{end}}
		
{{ end }}
</body>
`

//API-Key for Riot Developers
const apiKey string

//Our http Client which will be making the API Calls
var client = &http.Client{}

//Main variables
var out Output
var summonerName string

//Temporary variables to filter information for the Main variables
var record Profile
var list playerMatches
var champ Champion
var spell SpellInfo

func main() {

	fmt.Println("Serving Files")
	//HandleFunc in Golang is used to look for the ending of the URL. It's told what the paramaters are, and then executes a function
	//One difference is that if a parameter is sent at the end of the URL, and it's not explicitely listed below, it will execute the function closest to it's call
	// URL.com/s, for instance, will send URL.com/search, unless it's explicitely told not to.
	http.HandleFunc("/", homeFunc)
	http.HandleFunc("/search", searchFunc)

	http.ListenAndServe(":8080", nil)

}

//Most web-based functions will require a ResponseWriter, and a Request argument to be passed through it.
//The request is what is send to the function, while the Writer is what will be used to send information back
//to the webpage.
func homeFunc(response http.ResponseWriter, request *http.Request) {

	//Takes the homepage, index.html, and sends it to the client as a template. This allows for changes in the future if I want
	//to send information to it later.
	t, _ := template.ParseFiles("index.html")

	//Executes the template, but does not send any information for input.
	t.Execute(response, nil)

}

func searchFunc(response http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {

		//Returns the user to the homepage, unless they explicitly search from the form
		http.Redirect(response, request, "/", 301)
		return

	} else {

		//Collects the form information pushed from the homepage, then stores it into out.SummonerName
		request.ParseForm()

		fmt.Println("---")
		fmt.Println("username:", request.FormValue("Search"))
		out.SummonerName = request.FormValue("Search")

		fmt.Println(out.SummonerName)

		//Calls the summonerSearch function, which does not need any input since the valuable information is already stored in out.SummonerName
		summonerSearch()

		if record.ID == 0 {
			//If no user is found, return to the homepage
			fmt.Println("No Summoner Found / Possible Rate Limit hit?")
			http.Redirect(response, request, "/", 301)
			return

		} else {

			//Once every bit of information is extracted and placed into the variable "out," the variable is then executed and input into the template
			t := template.New("main")
			t, _ = t.Parse(internalTempl)
			t.Execute(response, out)

		}

		//reset the ID and SummonerName to nil values
		out.SummonerName = ""
		record.ID = 0

	}

}

func summonerSearch() {

	//This is the URL for the API Call, specifically for getting profile information for the user out.SummonerName
	url := "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/" + out.SummonerName + "?api_key=" + apiKey

	log.Println(url)

	//req will set the paramaters for the API Call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	//resp will enact the API Call, starting a link between the server and the remote server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	//This if statement is just a fancy way to double error check, in case there are no responses from the remote server
	//At the same time, it's breaking down the API call, returned as a JSON object, into the variable "record", which has a Profile type and structure.
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Panic(err)
	}

	//All relevant data will be pulled into the out variable, which is structured above at line 12
	out.ProfileIconID = record.ProfileIconID
	out.SummonerName = record.Name
	out.AccountID = record.AccountID

	if record.ID == 0 {
		return
	}

	//Closes the API call to save resources from piling up
	resp.Body.Close()

	//New API call to the URL below, requesting Match History. As there are API limits in place, and I have only been given access to Ranked matches,
	//I have set the limit to 5 Matches per API Call.  If this is to change in the future, we can update the URL accordingly
	url = "https://na1.api.riotgames.com/lol/match/v3/matchlists/by-account/" + strconv.Itoa(record.AccountID) + "?endIndex=5&beginIndex=0&api_key=" + apiKey

	log.Println(url)

	//Rinse and Repeat above
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	//Directly Decode the important information into out
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		log.Panic(err)
	}

	resp.Body.Close()

	//Since each Match is stored into out.Match[], we will need to pull the information that we want to display.
	//Along with this, the Champion that was played was given to us in an Int. We will need to send that information back
	//in order to get the Champion's name as well. We do this for each match using a for loop for the size (or LENgth) of out.Match[]
	for i := 0; i < len(out.Match); i++ {

		//getChampionName requires the Champion ID, stored in list.Match[i].Champion, and the iteration of the loop, i, to ensure
		//that the information is stored to its coresponding match.
		getChampionName(out.Match[i].Champion, i)

		fmt.Print(out.Match[i].Name + " ")
		fmt.Println(strconv.FormatInt(out.Match[i].GameID, 10))

		//Lastly, we need the Match information for each individual Match.
		//Just like getChampionName, getMatchInfo requires the MatchID as well as the iteration of the loop.
		getMatchInfo(out.Match[i].GameID, i)

	}

}

func getChampionName(x int, i int) {
	//New URL specifically for Champion information.
	url := "https://na1.api.riotgames.com/lol/static-data/v3/champions/" + strconv.Itoa(x) + "?locale=en_US&tags=image&api_key=" + apiKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	//Decode into champ, since we only need the Name of the champion. We don't need to output all of the information to out.
	if err := json.NewDecoder(resp.Body).Decode(&champ); err != nil {
		log.Panic(err)
	}

	//Send the information to the specific Match in out, then close the API Call to save resources
	out.Match[i].Name = champ.Name
	resp.Body.Close()
}

func getMatchInfo(x int64, i int) {

	//Here, we make a call to get previous match stats.  This includes K / D / A, Victory/Defeat, Creep Score, and everything else.
	url := "https://na1.api.riotgames.com/lol/match/v3/matches/" + strconv.FormatInt(x, 10) + "?api_key=" + apiKey

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	//Decode everything into out.Match[i].Stats. While we won't be using all of it, this allows expandability in the future
	//in case I wish to display more stats on the page.
	if err := json.NewDecoder(resp.Body).Decode(&out.Match[i].Stats); err != nil {
		log.Fatal("Error decoding stats")
	}

	var p int

	//Here, we are only looking for the correct iteration that matches the original player we searched for.
	//Since there are 10 players total, we want to search through each of them until we find our specific player.

	//Using a temporary value to hold the correct Participant iteration (p), we subtract 1 because the slice starts at 0, while the ID value starts at 1
	for y := 0; y < len(out.Match[i].Stats.ParticipantIdentities); y++ {
		if out.SummonerName == out.Match[i].Stats.ParticipantIdentities[y].Player.SummonerName {
			out.Match[i].Stats.ParticipantID = out.Match[i].Stats.ParticipantIdentities[y].ParticipantID
			p = out.Match[i].Stats.ParticipantIdentities[y].ParticipantID - 1
			continue
		}
	}

	//Send the HAST to out's basic call to remove any possible errors
	out.HighestAchievedSeasonTier = out.Match[i].Stats.Participants[p].HighestAchievedSeasonTier

	//TotalMinionsKilled only counts basic minions. To display the correct Creep Score,
	//we need to add TotalMinionsKilled as well as NeutralMinionsKilled
	out.Match[i].Stats.Participants[p].Stats.TotalMinionsKilled = out.Match[i].Stats.Participants[p].Stats.TotalMinionsKilled + out.Match[i].Stats.Participants[p].Stats.NeutralMinionsKilled

	//Here, we assign file names to the Summoner Spells so that we may call them later in the webpage via the template.
	//We use p to ensure that we're on the same player as before

	switch out.Match[i].Stats.Participants[p].Spell1ID {
	case 1:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerBoost.png"
	case 3:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerExhaust.png"
	case 4:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerFlash.png"
	case 6:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerHaste.png"
	case 7:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerHeal.png"
	case 11:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerSmite.png"
	case 12:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerTeleport.png"
	case 13:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerMana.png"
	case 14:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerDot.png"
	case 21:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerBarrier.png"
	case 30:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerPoroRecall.png"
	case 31:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerPoroThrow.png"
	case 32:
		out.Match[i].Stats.Participants[p].Spell1Full = "SummonerSnowball.png"
	}

	switch out.Match[i].Stats.Participants[p].Spell2ID {
	case 1:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerBoost.png"
	case 3:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerExhaust.png"
	case 4:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerFlash.png"
	case 6:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerHaste.png"
	case 7:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerHeal.png"
	case 11:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerSmite.png"
	case 12:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerTeleport.png"
	case 13:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerMana.png"
	case 14:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerDot.png"
	case 21:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerBarrier.png"
	case 30:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerPoroRecall.png"
	case 31:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerPoroThrow.png"
	case 32:
		out.Match[i].Stats.Participants[p].Spell2Full = "SummonerSnowball.png"
	}

	//Lastly, we want the player's highest killstreak so that we can display it on the webpage! For bragging rights, of course.
	switch out.Match[i].Stats.Participants[p].Stats.LargestMultiKill {
	default:
		out.Match[i].Stats.Participants[p].Stats.HighestStreak = "No Multikill"
	case 2:
		out.Match[i].Stats.Participants[p].Stats.HighestStreak = "Double Kill"
	case 3:
		out.Match[i].Stats.Participants[p].Stats.HighestStreak = "Triple Kill!"
	case 4:
		out.Match[i].Stats.Participants[p].Stats.HighestStreak = "Quadrakill!"
	case 5:
		out.Match[i].Stats.Participants[p].Stats.HighestStreak = "PENTAKILL!"
	}

	resp.Body.Close()

}
