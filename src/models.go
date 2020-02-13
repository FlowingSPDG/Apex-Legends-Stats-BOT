package src

import (
	"time"
)

type ProfileStatsResponse struct {
	Data struct {
		PlatformInfo struct {
			PlatformSlug           string      `json:"platformSlug"`
			PlatformUserID         string      `json:"platformUserId"`
			PlatformUserHandle     string      `json:"platformUserHandle"`
			PlatformUserIdentifier string      `json:"platformUserIdentifier"`
			AvatarURL              string      `json:"avatarUrl"`
			AdditionalParameters   interface{} `json:"additionalParameters"`
		} `json:"platformInfo"`
		UserInfo struct {
			IsPremium       bool          `json:"isPremium"`
			IsVerified      bool          `json:"isVerified"`
			IsInfluencer    bool          `json:"isInfluencer"`
			CountryCode     interface{}   `json:"countryCode"`
			CustomAvatarURL interface{}   `json:"customAvatarUrl"`
			SocialAccounts  []interface{} `json:"socialAccounts"`
		} `json:"userInfo"`
		Metadata struct {
			CurrentSeason    int    `json:"currentSeason"`
			ActiveLegend     string `json:"activeLegend"`
			ActiveLegendName string `json:"activeLegendName"`
		} `json:"metadata"`
		Segments []struct {
			Type       string `json:"type"`
			Attributes struct {
			} `json:"attributes,omitempty"`
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata,omitempty"`
			ExpiryDate time.Time `json:"expiryDate"`
			Stats      struct {
				Level struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"level"`
				Kills struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"kills"`
				Damage struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"damage"`
				Headshots struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"headshots"`
				SeasonWins struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"seasonWins"`
				SeasonDamage struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"seasonDamage"`
				Season2Wins struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"season2Wins"`
				RankScore struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
						IconURL string `json:"iconUrl"`
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"rankScore"`
			} `json:"stats,omitempty"`
		} `json:"segments"`
		AvailableSegments []struct {
			Type       string `json:"type"`
			Attributes struct {
			} `json:"attributes"`
		} `json:"availableSegments"`
		ExpiryDate time.Time `json:"expiryDate"`
	} `json:"data"`
}
