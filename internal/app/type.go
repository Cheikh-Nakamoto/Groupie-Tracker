package app

type Value struct {
	Artists   string `json:"artists"`
	Dates     string `json:"dates"`
	Locations string `json:"locations"`
	Relation  string `json:"relation"`
}

type Global struct {
	Artist       PageData
	DonneRestant LOcDatRel
}
type PageData struct {
	Artist []Artist
}
type LOcDatRel struct {
	Location map[string][]Locations
	Date     map[string][]Dates
	Relation map[string][]Relations
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}
type Relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type LatLng struct { 
	Lat string 
	Lng string
}

type Info struct {
	Artist   Artist
	Date     Dates
	Location Locations
	Coordoones []LatLng
	Relation Relations
}

type ArtistInfo struct {
	ID           int      
	Image        string   
	Name         string  
	Members      []string
	CreationDate int     
	FirstAlbum   string   
	Location []string 
	Dates []string
	Relations map[string][]string
}

type Listeinfo struct{
	Artists   []ArtistInfo
}


type Error struct{
	ErrorMsg string
	Code int
	Image string
}

type Localisation  struct {
	Documentation string `json:"documentation"`
	Licenses      []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"licenses"`
	Rate struct {
		Limit     int `json:"limit"`
		Remaining int `json:"remaining"`
		Reset     int `json:"reset"`
	} `json:"rate"`
	Results []struct {
		Annotations struct {
			Dms struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"DMS"`
			Mgrs       string `json:"MGRS"`
			Maidenhead string `json:"Maidenhead"`
			Mercator   struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"Mercator"`
			Osm struct {
				NoteURL string `json:"note_url"`
				URL     string `json:"url"`
			} `json:"OSM"`
			UnM49 struct {
				Regions struct {
					Africa           string `json:"AFRICA"`
					Sn               string `json:"SN"`
					SUBSAHARANAFRICA string `json:"SUB-SAHARAN_AFRICA"`
					WesternAfrica    string `json:"WESTERN_AFRICA"`
					World            string `json:"WORLD"`
				} `json:"regions"`
				StatisticalGroupings []string `json:"statistical_groupings"`
			} `json:"UN_M49"`
			Callingcode int `json:"callingcode"`
			Currency    struct {
				AlternateSymbols     []string `json:"alternate_symbols"`
				DecimalMark          string   `json:"decimal_mark"`
				DisambiguateSymbol   string   `json:"disambiguate_symbol"`
				Format               string   `json:"format"`
				HTMLEntity           string   `json:"html_entity"`
				IsoCode              string   `json:"iso_code"`
				IsoNumeric           string   `json:"iso_numeric"`
				Name                 string   `json:"name"`
				SmallestDenomination int      `json:"smallest_denomination"`
				Subunit              string   `json:"subunit"`
				SubunitToUnit        int      `json:"subunit_to_unit"`
				Symbol               string   `json:"symbol"`
				SymbolFirst          int      `json:"symbol_first"`
				ThousandsSeparator   string   `json:"thousands_separator"`
			} `json:"currency"`
			Flag     string  `json:"flag"`
			Geohash  string  `json:"geohash"`
			Qibla    float64 `json:"qibla"`
			Roadinfo struct {
				DriveOn string `json:"drive_on"`
				SpeedIn string `json:"speed_in"`
			} `json:"roadinfo"`
			Sun struct {
				Rise struct {
					Apparent     int `json:"apparent"`
					Astronomical int `json:"astronomical"`
					Civil        int `json:"civil"`
					Nautical     int `json:"nautical"`
				} `json:"rise"`
				Set struct {
					Apparent     int `json:"apparent"`
					Astronomical int `json:"astronomical"`
					Civil        int `json:"civil"`
					Nautical     int `json:"nautical"`
				} `json:"set"`
			} `json:"sun"`
			Timezone struct {
				Name         string `json:"name"`
				NowInDst     int    `json:"now_in_dst"`
				OffsetSec    int    `json:"offset_sec"`
				OffsetString string `json:"offset_string"`
				ShortName    string `json:"short_name"`
			} `json:"timezone"`
			What3Words struct {
				Words string `json:"words"`
			} `json:"what3words"`
		} `json:"annotations"`
		Bounds struct {
			Northeast struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"northeast"`
			Southwest struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"bounds"`
		Components struct {
			ISO31661Alpha2 string `json:"ISO_3166-1_alpha-2"`
			ISO31661Alpha3 string `json:"ISO_3166-1_alpha-3"`
			Category       string `json:"_category"`
			Type           string `json:"_type"`
			Continent      string `json:"continent"`
			Country        string `json:"country"`
			CountryCode    string `json:"country_code"`
			State          string `json:"state"`
			StateCode      string `json:"state_code"`
			Town           string `json:"town"`
		} `json:"components"`
		Confidence int    `json:"confidence"`
		Formatted  string `json:"formatted"`
		Geometry   struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	StayInformed struct {
		Blog     string `json:"blog"`
		Mastodon string `json:"mastodon"`
	} `json:"stay_informed"`
	Thanks    string `json:"thanks"`
	Timestamp struct {
		CreatedHTTP string `json:"created_http"`
		CreatedUnix int    `json:"created_unix"`
	} `json:"timestamp"`
	TotalResults int `json:"total_results"`
}