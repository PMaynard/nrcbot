package main

import (
	"context"
	"github.com/carlmjohnson/requests"
	"time"
)

type CrossRefWorksDOI struct {
	Status         string `json:"status"`
	MessageType    string `json:"message-type"`
	MessageVersion string `json:"message-version"`
	Message        struct {
		Indexed struct {
			DateParts [][]int   `json:"date-parts"`
			DateTime  time.Time `json:"date-time"`
			Timestamp int64     `json:"timestamp"`
		} `json:"indexed"`
		UpdateTo []struct {
			Updated struct {
				DateParts [][]int   `json:"date-parts"`
				DateTime  time.Time `json:"date-time"`
				Timestamp int64     `json:"timestamp"`
			} `json:"updated"`
			Doi   string `json:"DOI"`
			Type  string `json:"type"`
			Label string `json:"label"`
		} `json:"update-to"`
		ReferenceCount int    `json:"reference-count"`
		Publisher      string `json:"publisher"`
		Issue          string `json:"issue"`
		License        []struct {
			Start struct {
				DateParts [][]int   `json:"date-parts"`
				DateTime  time.Time `json:"date-time"`
				Timestamp int64     `json:"timestamp"`
			} `json:"start"`
			ContentVersion string `json:"content-version"`
			DelayInDays    int    `json:"delay-in-days"`
			URL            string `json:"URL"`
		} `json:"license"`
		Funder []struct {
			Doi           string   `json:"DOI"`
			Name          string   `json:"name"`
			DoiAssertedBy string   `json:"doi-asserted-by"`
			Award         []string `json:"award"`
		} `json:"funder"`
		ContentDomain struct {
			Domain               []string `json:"domain"`
			CrossmarkRestriction bool     `json:"crossmark-restriction"`
		} `json:"content-domain"`
		Chair []struct {
			Name        string `json:"name"`
			Sequence    string `json:"sequence"`
			Affiliation []any  `json:"affiliation"`
		} `json:"chair"`
		ShortContainerTitle []string `json:"short-container-title"`
		PublishedPrint      struct {
			DateParts [][]int `json:"date-parts"`
		} `json:"published-print"`
		Abstract string `json:"abstract"`
		Doi      string `json:"DOI"`
		Type     string `json:"type"`
		Created  struct {
			DateParts [][]int   `json:"date-parts"`
			DateTime  time.Time `json:"date-time"`
			Timestamp int64     `json:"timestamp"`
		} `json:"created"`
		Page                string   `json:"page"`
		UpdatePolicy        string   `json:"update-policy"`
		Source              string   `json:"source"`
		IsReferencedByCount int      `json:"is-referenced-by-count"`
		Title               []string `json:"title"`
		Prefix              string   `json:"prefix"`
		Volume              string   `json:"volume"`
		ClinicalTrialNumber []struct {
			ClinicalTrialNumber string `json:"clinical-trial-number"`
			Registry            string `json:"registry"`
		} `json:"clinical-trial-number"`
		Author []struct {
			Orcid              string `json:"ORCID"`
			AuthenticatedOrcid bool   `json:"authenticated-orcid"`
			Suffix             string `json:"suffix"`
			Given              string `json:"given"`
			Family             string `json:"family"`
			Sequence           string `json:"sequence"`
			Affiliation        []struct {
				Name string `json:"name"`
			} `json:"affiliation"`
		} `json:"author"`
		Member          string `json:"member"`
		PublishedOnline struct {
			DateParts [][]int `json:"date-parts"`
		} `json:"published-online"`
		Reference []struct {
			Key           string `json:"key"`
			DoiAssertedBy string `json:"doi-asserted-by"`
			Doi           string `json:"DOI"`
		} `json:"reference"`
		ContainerTitle []string `json:"container-title"`
		OriginalTitle  []any    `json:"original-title"`
		Language       string   `json:"language"`
		Link           []struct {
			URL                 string `json:"URL"`
			ContentType         string `json:"content-type"`
			ContentVersion      string `json:"content-version"`
			IntendedApplication string `json:"intended-application"`
		} `json:"link"`
		Deposited struct {
			DateParts [][]int   `json:"date-parts"`
			DateTime  time.Time `json:"date-time"`
			Timestamp int64     `json:"timestamp"`
		} `json:"deposited"`
		Score    int `json:"score"`
		Resource struct {
			Primary struct {
				URL string `json:"URL"`
			} `json:"primary"`
		} `json:"resource"`
		Subtitle   []any `json:"subtitle"`
		ShortTitle []any `json:"short-title"`
		Issued     struct {
			DateParts [][]int `json:"date-parts"`
		} `json:"issued"`
		ReferencesCount int `json:"references-count"`
		JournalIssue    struct {
			Issue           string `json:"issue"`
			PublishedOnline struct {
				DateParts [][]int `json:"date-parts"`
			} `json:"published-online"`
			PublishedPrint struct {
				DateParts [][]int `json:"date-parts"`
			} `json:"published-print"`
		} `json:"journal-issue"`
		URL      string   `json:"URL"`
		Archive  []string `json:"archive"`
		Relation struct {
			IsReplyTo []struct {
				IDType     string `json:"id-type"`
				ID         string `json:"id"`
				AssertedBy string `json:"asserted-by"`
			} `json:"is-reply-to"`
			References []struct {
				IDType     string `json:"id-type"`
				ID         string `json:"id"`
				AssertedBy string `json:"asserted-by"`
			} `json:"references"`
			HasTranslation []struct {
				IDType     string `json:"id-type"`
				ID         string `json:"id"`
				AssertedBy string `json:"asserted-by"`
			} `json:"has-translation"`
		} `json:"relation"`
		Issn     []string `json:"ISSN"`
		IssnType []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"issn-type"`
		Published struct {
			DateParts [][]int `json:"date-parts"`
		} `json:"published"`
		Assertion []struct {
			URL         string `json:"URL,omitempty"`
			Order       int    `json:"order"`
			Name        string `json:"name"`
			Label       string `json:"label"`
			Explanation struct {
				URL string `json:"URL"`
			} `json:"explanation,omitempty"`
			Value string `json:"value,omitempty"`
			Group struct {
				Name  string `json:"name"`
				Label string `json:"label"`
			} `json:"group,omitempty"`
		} `json:"assertion"`
	} `json:"message"`
}

func GetDOI(doi string) (CrossRefWorksDOI, error) {
	var res CrossRefWorksDOI
	err := requests.
		URL("https://api.crossref.org/").
		Pathf("/works/%s", doi).
		ToJSON(&res).
		Fetch(context.Background())
	if err != nil {
		return res, err
	}
	return res, nil
}
