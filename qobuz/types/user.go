package types

type User struct {
	Id           int         `json:"id"`
	PublicId     string      `json:"publicId"`
	Email        string      `json:"email"`
	Login        string      `json:"login"`
	Firstname    interface{} `json:"firstname"`
	Lastname     interface{} `json:"lastname"`
	DisplayName  string      `json:"display_name"`
	CountryCode  string      `json:"country_code"`
	LanguageCode string      `json:"language_code"`
	Zone         string      `json:"zone"`
	Store        string      `json:"store"`
	Country      string      `json:"country"`
	Avatar       string      `json:"avatar"`
	Genre        string      `json:"genre"`
	Age          int         `json:"age"`
	CreationDate string      `json:"creation_date"`
	Subscription struct {
		Offer            string `json:"offer"`
		Periodicity      string `json:"periodicity"`
		StartDate        string `json:"start_date"`
		EndDate          string `json:"end_date"`
		IsCanceled       bool   `json:"is_canceled"`
		HouseholdSizeMax int    `json:"household_size_max"`
	} `json:"subscription"`
	Credential struct {
		Id          int    `json:"id"`
		Label       string `json:"label"`
		Description string `json:"description"`
		Parameters  struct {
			LossyStreaming          bool  `json:"lossy_streaming"`
			LosslessStreaming       bool  `json:"lossless_streaming"`
			HiresStreaming          bool  `json:"hires_streaming"`
			HiresPurchasesStreaming bool  `json:"hires_purchases_streaming"`
			MobileStreaming         bool  `json:"mobile_streaming"`
			OfflineStreaming        bool  `json:"offline_streaming"`
			HfpPurchase             bool  `json:"hfp_purchase"`
			IncludedFormatGroupIds  []int `json:"included_format_group_ids"`
			ColorScheme             struct {
				Logo string `json:"logo"`
			} `json:"color_scheme"`
			Label      string `json:"label"`
			ShortLabel string `json:"short_label"`
			Source     string `json:"source"`
		} `json:"parameters"`
	} `json:"credential"`
	LastUpdate struct {
		Favorite       int `json:"favorite"`
		FavoriteAlbum  int `json:"favorite_album"`
		FavoriteArtist int `json:"favorite_artist"`
		FavoriteTrack  int `json:"favorite_track"`
		Playlist       int `json:"playlist"`
		Purchase       int `json:"purchase"`
	} `json:"last_update"`
	StoreFeatures struct {
		Download                 bool `json:"download"`
		Streaming                bool `json:"streaming"`
		Editorial                bool `json:"editorial"`
		Club                     bool `json:"club"`
		Wallet                   bool `json:"wallet"`
		Weeklyq                  bool `json:"weeklyq"`
		Autoplay                 bool `json:"autoplay"`
		InappPurchaseSubscripton bool `json:"inapp_purchase_subscripton"`
		OptIn                    bool `json:"opt_in"`
		PreRegisterOptIn         bool `json:"pre_register_opt_in"`
		PreRegisterZipcode       bool `json:"pre_register_zipcode"`
		MusicImport              bool `json:"music_import"`
		Radio                    bool `json:"radio"`
	} `json:"store_features"`
	PlayerSettings struct {
		SonosAudioFormat int `json:"sonos_audio_format"`
	} `json:"player_settings"`
	Externals struct {
	} `json:"externals"`
}
