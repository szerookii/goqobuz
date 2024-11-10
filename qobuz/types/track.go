package types

type Track struct {
	MaximumBitDepth int    `json:"maximum_bit_depth"`
	Copyright       string `json:"copyright"`
	Performers      string `json:"performers"`
	AudioInfo       struct {
		ReplaygainTrackPeak float64 `json:"replaygain_track_peak"`
		ReplaygainTrackGain float64 `json:"replaygain_track_gain"`
	} `json:"audio_info"`
	Performer struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"performer"`
	Album struct {
		Image struct {
			Small     string `json:"small"`
			Thumbnail string `json:"thumbnail"`
			Large     string `json:"large"`
		} `json:"image"`
		MaximumBitDepth int `json:"maximum_bit_depth"`
		MediaCount      int `json:"media_count"`
		Artist          struct {
			Image       interface{} `json:"image"`
			Name        string      `json:"name"`
			Id          int         `json:"id"`
			AlbumsCount int         `json:"albums_count"`
			Slug        string      `json:"slug"`
			Picture     interface{} `json:"picture"`
		} `json:"artist"`
		Upc        string `json:"upc"`
		ReleasedAt int    `json:"released_at"`
		Label      struct {
			Name        string `json:"name"`
			Id          int    `json:"id"`
			AlbumsCount int    `json:"albums_count"`
			SupplierId  int    `json:"supplier_id"`
			Slug        string `json:"slug"`
		} `json:"label"`
		Title           string  `json:"title"`
		QobuzId         int     `json:"qobuz_id"`
		Version         *string `json:"version"`
		Duration        int     `json:"duration"`
		ParentalWarning bool    `json:"parental_warning"`
		TracksCount     int     `json:"tracks_count"`
		Popularity      int     `json:"popularity"`
		Genre           struct {
			Path  []int  `json:"path"`
			Color string `json:"color"`
			Name  string `json:"name"`
			Id    int    `json:"id"`
			Slug  string `json:"slug"`
		} `json:"genre"`
		MaximumChannelCount int         `json:"maximum_channel_count"`
		Id                  string      `json:"id"`
		MaximumSamplingRate float64     `json:"maximum_sampling_rate"`
		Previewable         bool        `json:"previewable"`
		Sampleable          bool        `json:"sampleable"`
		Displayable         bool        `json:"displayable"`
		Streamable          bool        `json:"streamable"`
		StreamableAt        int         `json:"streamable_at"`
		Downloadable        bool        `json:"downloadable"`
		PurchasableAt       interface{} `json:"purchasable_at"`
		Purchasable         bool        `json:"purchasable"`
		ReleaseDateOriginal string      `json:"release_date_original"`
		ReleaseDateDownload string      `json:"release_date_download"`
		ReleaseDateStream   string      `json:"release_date_stream"`
		ReleaseDatePurchase string      `json:"release_date_purchase"`
		Hires               bool        `json:"hires"`
		HiresStreamable     bool        `json:"hires_streamable"`
	} `json:"album"`
	Work     interface{} `json:"work"`
	Composer struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"composer,omitempty"`
	Isrc                string  `json:"isrc"`
	Title               string  `json:"title"`
	Version             *string `json:"version"`
	Duration            int     `json:"duration"`
	ParentalWarning     bool    `json:"parental_warning"`
	TrackNumber         int     `json:"track_number"`
	MaximumChannelCount int     `json:"maximum_channel_count"`
	Id                  int     `json:"id"`
	MediaNumber         int     `json:"media_number"`
	MaximumSamplingRate float64 `json:"maximum_sampling_rate"`
	ReleaseDateOriginal string  `json:"release_date_original"`
	ReleaseDateDownload string  `json:"release_date_download"`
	ReleaseDateStream   string  `json:"release_date_stream"`
	ReleaseDatePurchase string  `json:"release_date_purchase"`
	Purchasable         bool    `json:"purchasable"`
	Streamable          bool    `json:"streamable"`
	Previewable         bool    `json:"previewable"`
	Sampleable          bool    `json:"sampleable"`
	Downloadable        bool    `json:"downloadable"`
	Displayable         bool    `json:"displayable"`
	PurchasableAt       *int    `json:"purchasable_at"`
	StreamableAt        int     `json:"streamable_at"`
	Hires               bool    `json:"hires"`
	HiresStreamable     bool    `json:"hires_streamable"`
}

type Album struct {
	MaximumBitDepth int `json:"maximum_bit_depth"`
	Image           struct {
		Small     string      `json:"small"`
		Thumbnail string      `json:"thumbnail"`
		Large     string      `json:"large"`
		Back      interface{} `json:"back"`
	} `json:"image"`
	MediaCount int `json:"media_count"`
	Artist     struct {
		Image       interface{} `json:"image"`
		Name        string      `json:"name"`
		Id          int         `json:"id"`
		AlbumsCount int         `json:"albums_count"`
		Slug        string      `json:"slug"`
		Picture     interface{} `json:"picture"`
	} `json:"artist"`
	Artists []struct {
		Id    int      `json:"id"`
		Name  string   `json:"name"`
		Roles []string `json:"roles"`
	} `json:"artists"`
	Upc        string `json:"upc"`
	ReleasedAt int    `json:"released_at"`
	Label      struct {
		Name        string `json:"name"`
		Id          int    `json:"id"`
		AlbumsCount int    `json:"albums_count"`
		SupplierId  int    `json:"supplier_id"`
		Slug        string `json:"slug"`
	} `json:"label"`
	Title           string      `json:"title"`
	QobuzId         int         `json:"qobuz_id"`
	Version         interface{} `json:"version"`
	Url             string      `json:"url"`
	Duration        int         `json:"duration"`
	ParentalWarning bool        `json:"parental_warning"`
	Popularity      int         `json:"popularity"`
	TracksCount     int         `json:"tracks_count"`
	Genre           struct {
		Path  []int  `json:"path"`
		Color string `json:"color"`
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Slug  string `json:"slug"`
	} `json:"genre"`
	MaximumChannelCount int           `json:"maximum_channel_count"`
	Id                  string        `json:"id"`
	MaximumSamplingRate float64       `json:"maximum_sampling_rate"`
	Articles            []interface{} `json:"articles"`
	ReleaseDateOriginal string        `json:"release_date_original"`
	ReleaseDateDownload string        `json:"release_date_download"`
	ReleaseDateStream   string        `json:"release_date_stream"`
	Purchasable         bool          `json:"purchasable"`
	Streamable          bool          `json:"streamable"`
	Previewable         bool          `json:"previewable"`
	Sampleable          bool          `json:"sampleable"`
	Downloadable        bool          `json:"downloadable"`
	Displayable         bool          `json:"displayable"`
	PurchasableAt       int           `json:"purchasable_at"`
	StreamableAt        int           `json:"streamable_at"`
	Hires               bool          `json:"hires"`
	HiresStreamable     bool          `json:"hires_streamable"`
}

type FullAlbum struct {
	MaximumBitDepth int `json:"maximum_bit_depth"`
	Image           struct {
		Small     string      `json:"small"`
		Thumbnail string      `json:"thumbnail"`
		Large     string      `json:"large"`
		Back      interface{} `json:"back"`
	} `json:"image"`
	MediaCount int `json:"media_count"`
	Artist     struct {
		Image       interface{} `json:"image"`
		Name        string      `json:"name"`
		Id          int         `json:"id"`
		AlbumsCount int         `json:"albums_count"`
		Slug        string      `json:"slug"`
		Picture     interface{} `json:"picture"`
	} `json:"artist"`
	Artists []struct {
		Id    int      `json:"id"`
		Name  string   `json:"name"`
		Roles []string `json:"roles"`
	} `json:"artists"`
	Upc        string `json:"upc"`
	ReleasedAt int    `json:"released_at"`
	Label      struct {
		Name        string `json:"name"`
		Id          int    `json:"id"`
		AlbumsCount int    `json:"albums_count"`
		SupplierId  int    `json:"supplier_id"`
		Slug        string `json:"slug"`
	} `json:"label"`
	Title           string      `json:"title"`
	QobuzId         int         `json:"qobuz_id"`
	Version         interface{} `json:"version"`
	Url             string      `json:"url"`
	Duration        int         `json:"duration"`
	ParentalWarning bool        `json:"parental_warning"`
	Popularity      int         `json:"popularity"`
	TracksCount     int         `json:"tracks_count"`
	Genre           struct {
		Path  []int  `json:"path"`
		Color string `json:"color"`
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Slug  string `json:"slug"`
	} `json:"genre"`
	MaximumChannelCount int           `json:"maximum_channel_count"`
	Id                  string        `json:"id"`
	MaximumSamplingRate float64       `json:"maximum_sampling_rate"`
	Articles            []interface{} `json:"articles"`
	ReleaseDateOriginal string        `json:"release_date_original"`
	ReleaseDateDownload string        `json:"release_date_download"`
	ReleaseDateStream   string        `json:"release_date_stream"`
	Purchasable         bool          `json:"purchasable"`
	Streamable          bool          `json:"streamable"`
	Previewable         bool          `json:"previewable"`
	Sampleable          bool          `json:"sampleable"`
	Downloadable        bool          `json:"downloadable"`
	Displayable         bool          `json:"displayable"`
	PurchasableAt       int           `json:"purchasable_at"`
	StreamableAt        int           `json:"streamable_at"`
	Hires               bool          `json:"hires"`
	HiresStreamable     bool          `json:"hires_streamable"`
	Awards              []interface{} `json:"awards"`
	Goodies             []interface{} `json:"goodies"`
	Area                interface{}   `json:"area"`
	Catchline           string        `json:"catchline"`
	Composer            struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		Slug        string      `json:"slug"`
		AlbumsCount int         `json:"albums_count"`
		Picture     interface{} `json:"picture"`
		Image       interface{} `json:"image"`
	} `json:"composer"`
	CreatedAt                      int           `json:"created_at"`
	GenresList                     []string      `json:"genres_list"`
	Period                         interface{}   `json:"period"`
	Copyright                      string        `json:"copyright"`
	IsOfficial                     bool          `json:"is_official"`
	MaximumTechnicalSpecifications string        `json:"maximum_technical_specifications"`
	ProductSalesFactorsMonthly     int           `json:"product_sales_factors_monthly"`
	ProductSalesFactorsWeekly      int           `json:"product_sales_factors_weekly"`
	ProductSalesFactorsYearly      int           `json:"product_sales_factors_yearly"`
	ProductType                    string        `json:"product_type"`
	ProductUrl                     string        `json:"product_url"`
	RecordingInformation           string        `json:"recording_information"`
	RelativeUrl                    string        `json:"relative_url"`
	ReleaseTags                    []interface{} `json:"release_tags"`
	ReleaseType                    string        `json:"release_type"`
	Slug                           string        `json:"slug"`
	Subtitle                       string        `json:"subtitle"`
	TrackIds                       []int         `json:"track_ids"`
	Tracks                         struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Total  int `json:"total"`
		Items  []struct {
			MaximumBitDepth int    `json:"maximum_bit_depth"`
			Copyright       string `json:"copyright"`
			Performers      string `json:"performers"`
			AudioInfo       struct {
				ReplaygainTrackPeak float64 `json:"replaygain_track_peak"`
				ReplaygainTrackGain float64 `json:"replaygain_track_gain"`
			} `json:"audio_info"`
			Performer struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"performer"`
			Work     interface{} `json:"work"`
			Composer struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"composer"`
			Isrc                string      `json:"isrc"`
			Title               string      `json:"title"`
			Version             interface{} `json:"version"`
			Duration            int         `json:"duration"`
			ParentalWarning     bool        `json:"parental_warning"`
			TrackNumber         int         `json:"track_number"`
			MaximumChannelCount int         `json:"maximum_channel_count"`
			Id                  int         `json:"id"`
			MediaNumber         int         `json:"media_number"`
			MaximumSamplingRate float64     `json:"maximum_sampling_rate"`
			ReleaseDateOriginal string      `json:"release_date_original"`
			ReleaseDateDownload string      `json:"release_date_download"`
			ReleaseDateStream   string      `json:"release_date_stream"`
			ReleaseDatePurchase string      `json:"release_date_purchase"`
			Purchasable         bool        `json:"purchasable"`
			Streamable          bool        `json:"streamable"`
			Previewable         bool        `json:"previewable"`
			Sampleable          bool        `json:"sampleable"`
			Downloadable        bool        `json:"downloadable"`
			Displayable         bool        `json:"displayable"`
			PurchasableAt       int         `json:"purchasable_at"`
			StreamableAt        int         `json:"streamable_at"`
			Hires               bool        `json:"hires"`
			HiresStreamable     bool        `json:"hires_streamable"`
		} `json:"items"`
	} `json:"tracks"`
	AlbumsSameArtist struct {
		Items []struct {
			MaximumBitDepth int `json:"maximum_bit_depth"`
			Image           struct {
				Small     string      `json:"small"`
				Thumbnail string      `json:"thumbnail"`
				Large     string      `json:"large"`
				Back      interface{} `json:"back"`
			} `json:"image"`
			MediaCount int `json:"media_count"`
			Artist     struct {
				Image       interface{} `json:"image"`
				Name        string      `json:"name"`
				Id          int         `json:"id"`
				AlbumsCount int         `json:"albums_count"`
				Slug        string      `json:"slug"`
				Picture     interface{} `json:"picture"`
			} `json:"artist"`
			Artists []struct {
				Id    int      `json:"id"`
				Name  string   `json:"name"`
				Roles []string `json:"roles"`
			} `json:"artists"`
			Upc        string `json:"upc"`
			ReleasedAt int    `json:"released_at"`
			Label      struct {
				Name        string `json:"name"`
				Id          int    `json:"id"`
				AlbumsCount int    `json:"albums_count"`
				SupplierId  int    `json:"supplier_id"`
				Slug        string `json:"slug"`
			} `json:"label"`
			Title           string      `json:"title"`
			QobuzId         int         `json:"qobuz_id"`
			Version         interface{} `json:"version"`
			Url             string      `json:"url"`
			Duration        int         `json:"duration"`
			ParentalWarning bool        `json:"parental_warning"`
			Popularity      int         `json:"popularity"`
			TracksCount     int         `json:"tracks_count"`
			Genre           struct {
				Path  []int  `json:"path"`
				Color string `json:"color"`
				Name  string `json:"name"`
				Id    int    `json:"id"`
				Slug  string `json:"slug"`
			} `json:"genre"`
			MaximumChannelCount int           `json:"maximum_channel_count"`
			Id                  string        `json:"id"`
			MaximumSamplingRate int           `json:"maximum_sampling_rate"`
			Articles            []interface{} `json:"articles"`
			ReleaseDateOriginal string        `json:"release_date_original"`
			ReleaseDateDownload string        `json:"release_date_download"`
			ReleaseDateStream   string        `json:"release_date_stream"`
			Purchasable         bool          `json:"purchasable"`
			Streamable          bool          `json:"streamable"`
			Previewable         bool          `json:"previewable"`
			Sampleable          bool          `json:"sampleable"`
			Downloadable        bool          `json:"downloadable"`
			Displayable         bool          `json:"displayable"`
			PurchasableAt       int           `json:"purchasable_at"`
			StreamableAt        int           `json:"streamable_at"`
			Hires               bool          `json:"hires"`
			HiresStreamable     bool          `json:"hires_streamable"`
			Description         []interface{} `json:"description"`
		} `json:"items"`
	} `json:"albums_same_artist"`
	Description string `json:"description"`
}
