package main

import (
	"strconv"
	"time"
)

type ActiveMatchResponse struct {
	Uid         string `json:"uid"`
	Competitors []struct {
		Id        int    `json:"id"`
		LongName  string `json:"longName"`
		ShortName string `json:"shortName"`
		IconUrl   string `json:"iconUrl"`
		Score     int    `json:"score"`
		Color     string `json:"color"`
	} `json:"competitors"`
	Status      string `json:"status"`
	StatusTexts struct {
		LiveStatusText    string `json:"liveStatusText"`
		PendingStatusText string `json:"pendingStatusText"`
	} `json:"statusTexts"`
	TimeToMatch int      `json:"timeToMatch"`
	LinkToMatch string   `json:"linkToMatch"`
	IsEncore    bool     `json:"isEncore"`
	MatchDate   unixTime `json:"matchDate"`
}

type ActiveMatchesResponse struct {
	Data []ActiveMatchResponse `json:"data"`
}

func convertTeams(match Matches) []Team {
	teams := make([]Team, 2)
	for i := 0; i < 2; i++ {
		score := 0
		if len(match.Scores) > i {
			score = match.Scores[i]
		}
		team := Team{
			Name:            match.Competitors[i].Name,
			AbbreviatedName: match.Competitors[i].AbbreviatedName,
			Icon:            match.Competitors[i].Icon,
			Score:           score,
		}
		teams[i] = team
	}
	return teams
}

func convertTeamsColored(match ActiveMatchResponse) []TeamColored {
	teams := make([]TeamColored, 2)
	for i := 0; i < 2; i++ {
		team := TeamColored{
			Team: Team{
				Name:            match.Competitors[i].LongName,
				AbbreviatedName: match.Competitors[i].ShortName,
				Icon:            match.Competitors[i].IconUrl,
				Score:           match.Competitors[i].Score,
			},
			Color: match.Competitors[i].Color,
		}
		teams[i] = team
	}
	return teams
}

func convertMatch(match Matches, teams []Team, event Events) Match {
	return Match{
		ID:     match.ID,
		Status: match.Status,
		Encore: match.IsEncore,
		Teams:  teams,
		Start:  time.Time(match.StartDate),
		End:    time.Time(match.EndDate),
		Event:  event.EventBanner.Title,
	}
}

type unixTime time.Time

func (t unixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *unixTime) UnmarshalJSON(s []byte) (err error) {
	q, err := strconv.ParseInt(string(s), 10, 64)

	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0)
	return
}

type AutoGeneratedResponse struct {
	Data Data `json:"data"`
}
type Competitors struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AbbreviatedName string `json:"abbreviatedName"`
	Icon            string `json:"icon"`
	Logo            string `json:"logo"`
}
type ACL struct {
}
type Link struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}
type PublishDetails struct {
	Environment string    `json:"environment"`
	Locale      string    `json:"locale"`
	Time        time.Time `json:"time"`
	User        string    `json:"user"`
}
type StatusText struct {
	Title          string         `json:"title"`
	Key            string         `json:"key"`
	Value          string         `json:"value"`
	Tags           []any          `json:"tags"`
	Locale         string         `json:"locale"`
	UID            string         `json:"uid"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	ContentTypeUID string         `json:"ContentTypeUid"`
	ACL            ACL            `json:"ACL"`
	Version        int            `json:"Version"`
	InProgress     bool           `json:"InProgress"`
	PublishDetails PublishDetails `json:"publishDetails"`
}
type Tickets struct {
	Version        int            `json:"Version"`
	Locale         string         `json:"locale"`
	UID            string         `json:"uid"`
	ACL            ACL            `json:"ACL"`
	InProgress     bool           `json:"InProgress"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedBy      string         `json:"createdBy"`
	Link           Link           `json:"link"`
	Status         string         `json:"status"`
	StatusText     []StatusText   `json:"statusText"`
	Tags           []any          `json:"tags"`
	Title          string         `json:"title"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	UpdatedBy      string         `json:"updatedBy"`
	ContentTypeUID string         `json:"ContentTypeUid"`
	PublishDetails PublishDetails `json:"publishDetails"`
}
type ColorLogoURL struct {
	Src string `json:"src"`
}
type GrayLogoURL struct {
	Src string `json:"src"`
}
type BroadcastChannels struct {
	Title        string       `json:"title"`
	ColorLogoURL ColorLogoURL `json:"colorLogoUrl"`
	GrayLogoURL  GrayLogoURL  `json:"grayLogoUrl"`
	Linkable     bool         `json:"linkable"`
	Link         Link         `json:"link"`
}
type GeoFence struct {
	Address           string `json:"address"`
	CountryCode       string `json:"countryCode"`
	ExcludeRegionCode bool   `json:"excludeRegionCode"`
	Latitude          int    `json:"latitude"`
	Longitude         int    `json:"longitude"`
	MaxRadius         int    `json:"maxRadius"`
	MinRadius         int    `json:"minRadius"`
	OtherCountryCodes []any  `json:"otherCountryCodes"`
	RegionCode        string `json:"regionCode"`
	VenueName         string `json:"venueName"`
}
type Venue struct {
	Version          int            `json:"Version"`
	Locale           string         `json:"locale"`
	UID              string         `json:"uid"`
	ACL              ACL            `json:"ACL"`
	InProgress       bool           `json:"InProgress"`
	CreatedAt        time.Time      `json:"createdAt"`
	CreatedBy        string         `json:"createdBy"`
	EventPerksEvents []any          `json:"eventPerksEvents"`
	GeoFence         GeoFence       `json:"geoFence"`
	Image            any            `json:"image"`
	Link             Link           `json:"link"`
	Tags             []any          `json:"tags"`
	Title            string         `json:"title"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	UpdatedBy        string         `json:"updatedBy"`
	ContentTypeUID   string         `json:"ContentTypeUid"`
	PublishDetails   PublishDetails `json:"publishDetails"`
	Location         string         `json:"location"`
	Name             string         `json:"name"`
}
type Matches struct {
	ID                int                 `json:"id"`
	Status            string              `json:"status"`
	Link              string              `json:"link"`
	StartDate         unixTime            `json:"startDate"`
	EndDate           unixTime            `json:"endDate"`
	Competitors       []Competitors       `json:"competitors"`
	Tickets           Tickets             `json:"tickets"`
	Scores            []int               `json:"scores"`
	BroadcastChannels []BroadcastChannels `json:"broadcastChannels"`
	IsEncore          bool                `json:"isEncore"`
	Venue             Venue               `json:"venue"`
}
type TeamIconSvg struct {
	Src string `json:"src"`
}
type TeamIconPng struct {
	Src string `json:"src"`
}
type TeamLogo struct {
	TeamIconUsage string      `json:"teamIconUsage"`
	TeamIconSvg   TeamIconSvg `json:"teamIconSvg"`
	TeamIconPng   TeamIconPng `json:"teamIconPng"`
}
type HostingTeam struct {
	TeamID    int        `json:"teamId"`
	LongName  []string   `json:"longName"`
	ShortName string     `json:"shortName"`
	Link      Link       `json:"link"`
	TeamLogo  []TeamLogo `json:"teamLogo"`
}
type Primary struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
type Secondary struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
type Tertiary struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
type Headings struct {
	Primary   Primary   `json:"primary"`
	Secondary Secondary `json:"secondary"`
	Tertiary  Tertiary  `json:"tertiary"`
}
type Ticket struct {
	Title      string     `json:"title"`
	Link       Link       `json:"link"`
	StatusText StatusText `json:"statusText"`
}
type EventBanner struct {
	Title                 string      `json:"title"`
	BackgroundImageURL    string      `json:"backgroundImageUrl"`
	BackgroundVideos      []any       `json:"backgroundVideos"`
	FeaturedImage         string      `json:"featuredImage"`
	HostedBy              string      `json:"hostedBy"`
	HostingTeam           HostingTeam `json:"hostingTeam"`
	Headings              Headings    `json:"headings"`
	Sponsors              []any       `json:"sponsors"`
	Ticket                Ticket      `json:"ticket"`
	Venue                 Venue       `json:"venue"`
	BottomBackgroundColor string      `json:"bottomBackgroundColor"`
}
type Events struct {
	Matches     []Matches   `json:"matches"`
	EventBanner EventBanner `json:"eventBanner"`
}
type Pagination struct {
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
	NextPage     int `json:"nextPage"`
	PreviousPage any `json:"previousPage"`
}
type TableData struct {
	Name        string     `json:"name"`
	Sponsor     []any      `json:"sponsor"`
	PresentedBy string     `json:"presentedBy"`
	StartDate   time.Time  `json:"startDate"`
	Subtitle    string     `json:"subtitle"`
	EndDate     time.Time  `json:"endDate"`
	WeekNumber  int        `json:"weekNumber"`
	Events      []Events   `json:"events"`
	Pagination  Pagination `json:"pagination"`
}
type UndefinedTeamLogo struct {
	Src string `json:"src"`
}
type Assets struct {
	UndefinedTeamLogo UndefinedTeamLogo `json:"undefinedTeamLogo"`
}
type Soldout struct {
	Locale         string         `json:"locale"`
	ACL            ACL            `json:"ACL"`
	ContentTypeUID string         `json:"ContentTypeUid"`
	InProgress     bool           `json:"InProgress"`
	Version        int            `json:"Version"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedBy      string         `json:"createdBy"`
	Tags           []any          `json:"tags"`
	Title          string         `json:"title"`
	UID            string         `json:"uid"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	UpdatedBy      string         `json:"updatedBy"`
	Value          string         `json:"value"`
	PublishDetails PublishDetails `json:"publishDetails"`
}
type Unavailable struct {
	Title          string         `json:"title"`
	Value          string         `json:"value"`
	Tags           []any          `json:"tags"`
	Locale         string         `json:"locale"`
	UID            string         `json:"uid"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	ContentTypeUID string         `json:"ContentTypeUid"`
	ACL            ACL            `json:"ACL"`
	Version        int            `json:"Version"`
	InProgress     bool           `json:"InProgress"`
	Key            string         `json:"key"`
	PublishDetails PublishDetails `json:"publishDetails"`
}
type Normal struct {
	Title          string         `json:"title"`
	Value          string         `json:"value"`
	Tags           []any          `json:"tags"`
	Locale         string         `json:"locale"`
	UID            string         `json:"uid"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	ContentTypeUID string         `json:"ContentTypeUid"`
	ACL            ACL            `json:"ACL"`
	Version        int            `json:"Version"`
	InProgress     bool           `json:"InProgress"`
	Key            string         `json:"key"`
	PublishDetails PublishDetails `json:"publishDetails"`
}
type Postponed struct {
	Locale         string         `json:"locale"`
	ACL            ACL            `json:"ACL"`
	InProgress     bool           `json:"InProgress"`
	Version        int            `json:"Version"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedBy      string         `json:"createdBy"`
	Key            string         `json:"key"`
	Tags           []any          `json:"tags"`
	Title          string         `json:"title"`
	UID            string         `json:"uid"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	UpdatedBy      string         `json:"updatedBy"`
	Value          string         `json:"value"`
	PublishDetails PublishDetails `json:"publishDetails"`
	ContentTypeUID string         `json:"ContentTypeUid"`
}
type Canceled struct {
	Locale         string         `json:"locale"`
	ACL            ACL            `json:"ACL"`
	ContentTypeUID string         `json:"ContentTypeUid"`
	InProgress     bool           `json:"InProgress"`
	Version        int            `json:"Version"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedBy      string         `json:"createdBy"`
	Key            string         `json:"key"`
	Tags           []any          `json:"tags"`
	Title          string         `json:"title"`
	UID            string         `json:"uid"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	UpdatedBy      string         `json:"updatedBy"`
	Value          string         `json:"value"`
	PublishDetails PublishDetails `json:"publishDetails"`
}
type TicketsStatusText struct {
	Soldout     []Soldout     `json:"soldout"`
	Unavailable []Unavailable `json:"unavailable"`
	Normal      []Normal      `json:"normal"`
	Postponed   []Postponed   `json:"postponed"`
	Canceled    []Canceled    `json:"canceled"`
}
type Strings struct {
	UpNext            string            `json:"upNext"`
	UndefinedTeamName string            `json:"undefinedTeamName"`
	Playing           string            `json:"playing"`
	Encore            string            `json:"encore"`
	Live              string            `json:"live"`
	Error             string            `json:"error"`
	LiveNow           string            `json:"liveNow"`
	Header            []string          `json:"header"`
	SpoilerMessage    []string          `json:"spoilerMessage"`
	TicketsStatusText TicketsStatusText `json:"ticketsStatusText"`
	Title2020         string            `json:"title2020"`
	Disclaimer        string            `json:"disclaimer"`
	AllTeams          string            `json:"allTeams"`
	NextWeek          string            `json:"nextWeek"`
	PreviousWeek      string            `json:"previousWeek"`
	Final             string            `json:"final"`
	Now               string            `json:"now"`
	MatchDetails      string            `json:"matchDetails"`
	Watch             string            `json:"watch"`
	Tickets           string            `json:"tickets"`
	SoldOut           string            `json:"soldOut"`
	RegularSeason     string            `json:"regularSeason"`
	AllStartWeekend   string            `json:"allStartWeekend"`
	Playoffs          string            `json:"playoffs"`
	GrandFinals       string            `json:"grandFinals"`
	Weeks             string            `json:"weeks"`
	EncoreText        string            `json:"encoreText"`
	AnalyticsString   string            `json:"analyticsString"`
}

type Dropdowns struct {
	Name string `json:"name"`
}
type Subtabs struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Selected  bool      `json:"selected"`
}
type Tabs struct {
	ID       string    `json:"id,omitempty"`
	Stage    string    `json:"stage,omitempty"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Selected bool      `json:"selected"`
	Subtabs  []Subtabs `json:"subtabs,omitempty"`
}
type Meta struct {
	Assets    Assets      `json:"assets"`
	Strings   Strings     `json:"strings"`
	Dropdowns []Dropdowns `json:"dropdowns"`
	Tabs      []Tabs      `json:"tabs"`
	TeamID    string      `json:"teamId"`
}
type Data struct {
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	TableData TableData `json:"tableData"`
	Meta      Meta      `json:"meta"`
}
