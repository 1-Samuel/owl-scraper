package main

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Repository interface {
	Get() ([]Match, error)
	Persist(match Match) error
	PersistActive(activeMatch ActiveMatch) error
	DeleteActiveMatch(uid string) error
	GetActive() (*ActiveMatch, error)
}

const uid = "87034ff4-b307-4c02-82ab-49a97e33b490"

type Scraper struct {
	repo        Repository
	activeMatch *ActiveMatch
}

func New(repo Repository) *Scraper {
	return &Scraper{
		repo: repo,
	}
}

func (s *Scraper) start() {
	log.Printf("Started fetching matches ...")
	matchCount, pageCount := s.fetch(1)
	log.Printf("Fetched %d matches from %d pages", matchCount, pageCount)
}

func (s *Scraper) fetchActiveMatch() {
	url := "https://pk0yccosw3.execute-api.us-east-2.amazonaws.com/production/v2/content-types/match-ticker/?locale=en-us"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("error creating request")
	}

	setHeaders(req)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // TODO log error

	if resp.StatusCode != http.StatusOK {
		s.updateActiveMatch(nil)
		return
	}

	response := new(ActiveMatchesResponse)
	err = json.NewDecoder(resp.Body).Decode(response)

	if err != nil || len(response.Data) == 0 {
		s.updateActiveMatch(nil)
		return
	}

	data := response.Data[0]
	match := &ActiveMatch{
		UID:         uid,
		Teams:       convertTeamsColored(data),
		Status:      data.Status,
		TimeToMatch: 0,
		LinkToMatch: data.LinkToMatch,
		IsEncore:    data.IsEncore,
		MatchDate:   time.Time(data.MatchDate),
	}
	log.Printf("Found active match %s %d:%d %s at %s - %s", match.Teams[0].AbbreviatedName, match.Teams[0].Score, match.Teams[1].Score, match.Teams[1].AbbreviatedName, match.MatchDate, match.Status)
	s.updateActiveMatch(match)

}

func (s *Scraper) updateActiveMatch(match *ActiveMatch) {
	savedMatch, err := s.repo.GetActive()

	if (err != nil || savedMatch != nil) && match == nil {
		log.Printf("No active match found => clearing active match")
		err := s.repo.DeleteActiveMatch(uid)
		if err != nil {
			panic(err)
		}
		s.activeMatch = nil
		go s.start()
		return
	}

	if match != nil && savedMatch != nil && !cmp.Equal(*savedMatch, *match) {
		log.Printf("Active match change detected")
		go s.start()
	}

	err = s.repo.PersistActive(*match)
	if err != nil {
		panic(err)
	}
}

func (s *Scraper) fetch(weekNumber int) (matchCount int, pageCount int) {
	url := s.generateUrl(weekNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("error creating request")
	}

	setHeaders(req)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // TODO log error

	if resp.StatusCode != http.StatusOK {
		panic("could not get matches from api")
	}

	response := new(AutoGeneratedResponse)
	err = json.NewDecoder(resp.Body).Decode(response)

	for _, event := range response.Data.TableData.Events {
		for _, match := range event.Matches {
			teams := convertTeams(match)
			match := convertMatch(match, teams, event)
			err := s.repo.Persist(match)
			if err != nil {
				panic(err)
			} else {
				matchCount++
			}
		}
	}

	pagination := response.Data.TableData.Pagination

	if pagination.NextPage != nil {
		time.Sleep(2 * time.Second)
		fetchedMatches, fetchedPages := s.fetch(*pagination.NextPage)
		return fetchedMatches + matchCount, fetchedPages + 1
	}
	return matchCount, 1
}

func (s *Scraper) generateUrl(weekNumber int) string {
	urlPrefix := "https://pk0yccosw3.execute-api.us-east-2.amazonaws.com/production/v2/content-types/schedule/blt27f16f110b3363f7/week/"
	urlSuffix := "/team/allteams?locale=en-us"
	return urlPrefix + strconv.Itoa(weekNumber) + urlSuffix
}

func (s *Scraper) isMatchActive() bool {
	if s.activeMatch != nil {
		return true
	}
	matches, err := s.repo.Get()
	if err != nil {
		return false
	}
	for _, match := range matches {
		if match.Status != "CONCLUDED" && match.Start.Add(-time.Hour).Before(time.Now()) {
			return true
		}
	}
	return false
}

func setHeaders(req *http.Request) {
	req.Header.Add("Host", "pk0yccosw3.execute-api.us-east-2.amazonaws.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/112.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Referer", "https://overwatchleague.com/")
	req.Header.Add("x-origin", "overwatchleague.com")
	req.Header.Add("Origin", "https://overwatchleague.com")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("TE", "trailers")
}
