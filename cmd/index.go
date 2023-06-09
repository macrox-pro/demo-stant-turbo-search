package cmd

//goland:noinspection SpellCheckingInspection
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/blevesearch/bleve/v2/document"

	"github.com/samber/lo"

	cresty "github.com/legion-zver/vss-brain-search/internal/helpers/cached_resty"

	"github.com/legion-zver/vss-brain-search/internal/models/premier"

	"github.com/legion-zver/vss-brain-search/internal/search"

	"github.com/blevesearch/bleve/v2"
	"github.com/corpix/uarand"
	"github.com/spf13/cobra"
)

var (
	cacheDirPath string
)

func init() {
	indexSyncCmd.PersistentFlags().StringVarP(
		&cacheDirPath, "cache:path", "c", "./cache.resty",
		"cache path (default is ./cache.resty)",
	)

	rootCmd.AddCommand(indexInitCmd)
	rootCmd.AddCommand(indexSyncCmd)
	rootCmd.AddCommand(indexCountsCmd)
}

var indexInitCmd = &cobra.Command{
	Use:   "index:init",
	Short: "create search index",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.New(indexPath, search.NewIndexMapping())
		if err != nil {
			log.Fatalln(err)
		}
		defer func(index bleve.Index) {
			_ = index.Close()
		}(index)
		log.Println("Ready!")
	},
}

var indexCountsCmd = &cobra.Command{
	Use:   "index:counts",
	Short: "get counts from search index",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer func(index bleve.Index) {
			_ = index.Close()
		}(index)

		idDict, err := index.FieldDict("_id")
		if err != nil {
			log.Fatalln(err)
		}
		typeCounter := make(map[string]int)
		genresCounter := make(map[string]int)
		personsCounter := make(map[string]int)
		countriesCounter := make(map[string]int)
		for {
			dict, err := idDict.Next()
			if err != nil || dict == nil {
				break
			}
			doc, err := index.Document(dict.Term)
			if err != nil {
				break
			}
			if obj, ok := doc.(*document.Document); ok {
				for _, field := range obj.Fields {
					switch field.Name() {
					case "countries":
						countriesCounter[string(field.Value())]++
					case "persons":
						personsCounter[string(field.Value())]++
					case "genres":
						genresCounter[string(field.Value())]++
					case "type":
						typeCounter[string(field.Value())]++
					}
				}
			}
		}
		_ = idDict.Close()

		fmt.Println("\ntypes:")
		for k, v := range typeCounter {
			if len(k) > 0 {
				fmt.Printf("  - name: %s\n", k)
				fmt.Printf("    count: %d\n", v)
			}
		}
		fmt.Println("\ngenres:")
		for k, v := range genresCounter {
			k = strings.Trim(strings.TrimSpace(k), ".,:'\"")
			if len(k) > 0 {
				fmt.Printf("  - name: %s\n", k)
				fmt.Printf("    count: %d\n", v)
			}
		}
		fmt.Println("\ncountries:")
		for k, v := range countriesCounter {
			k = strings.Trim(strings.TrimSpace(k), ".,:'\"")
			if len(k) > 0 {
				fmt.Printf("  - name: %s\n", k)
				fmt.Printf("    count: %d\n", v)
			}
		}
		fmt.Println("\npersons:")
		for k, v := range personsCounter {
			k = strings.Trim(strings.TrimSpace(k), ".,:'\"")
			if len(k) > 0 {
				if strings.Contains(k, ".") ||
					strings.Contains(k, ",") {
					continue
				}
				if len(strings.Split(k, " ")) > 2 {
					continue
				}
				fmt.Printf("  - name: %s\n", k)
				fmt.Printf("    count: %d\n", v)
			}
		}
	},
}

var indexSyncCmd = &cobra.Command{
	Use:   "index:sync",
	Short: "sync search index",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer func(index bleve.Index) {
			_ = index.Close()
		}(index)

		// Create cache dir
		_ = os.MkdirAll(cacheDirPath, os.ModePerm)

		page, limit, retryCount := 1, 100, 0
		client := cresty.New()
		for {
			resp, err := client.CachedR(cacheDirPath).
				SetDoNotParseResponse(false).
				SetHeader("Accept", "application/json").
				SetHeader("User-Agent", uarand.GetRandom()).
				Get(fmt.Sprintf(
					"https://premier.one/api/metainfo/tv/?page=%d&limit=%d&picture_type=card_group&style=portrait&device=web",
					page, limit,
				))
			if err != nil {
				retryCount++
				if retryCount > 3 {
					log.Fatalln(err)
				}
				time.Sleep(1000)
			}
			var payload premier.Response[premier.TV]
			if err = json.Unmarshal(resp.Body(), &payload); err != nil {
				log.Println(err)
				continue
			}
			if resp.IsCached() {
				fmt.Println("Cached iteration", page, "with", len(payload.Results), "items...")
			} else {
				fmt.Println("Iteration", page, "with", len(payload.Results), "items...")
			}
			batch := index.NewBatch()
			for _, item := range payload.Results {
				if len(item.Externals) < 1 {
					continue
				}
				object := search.IndexObject{
					Service:        "premier.one",
					Slug:           strings.TrimSpace(item.Slug),
					Name:           strings.TrimSpace(item.Name),
					Year:           strings.TrimSpace(item.Year),
					Title:          strings.TrimSpace(item.OriginalTitle),
					Genres:         make([]string, 0),
					Persons:        make([]string, 0),
					Countries:      make([]string, 0),
					YearEnd:        strings.TrimSpace(item.YearEnd),
					Picture:        strings.TrimSpace(item.Picture),
					YearStart:      strings.TrimSpace(item.YearStart),
					Description:    strings.TrimSpace(item.Description),
					AgeRestriction: strings.TrimSpace(item.AgeRestriction),
					IsActive:       item.IsActive,
				}
				if item.Type != nil {
					if len(item.Type.Title) > 0 {
						object.Type = strings.ToLower(item.Type.Title)
					} else {
						object.Type = strings.ToLower(item.Type.Name)
					}
				}
				if item.Provider != nil {
					object.Provider = item.Provider.Name
				}
				if len(item.Genres) > 0 {
					for _, genre := range item.Genres {
						if len(genre.Name) > 0 {
							object.Genres = append(object.Genres, genre.Name)
						}
					}
					if len(object.Genres) > 1 {
						object.Genres = lo.Uniq(object.Genres)
						sort.Strings(object.Genres)
					}
				}
				if len(item.Countries) > 0 {
					for _, country := range item.Countries {
						if len(country.Name) > 0 {
							object.Countries = append(object.Countries, country.Name)
						}
						if len(country.TwoLetter) > 0 {
							object.Countries = append(object.Countries, country.TwoLetter)
						}
					}
					if len(object.Countries) > 1 {
						object.Countries = lo.Uniq(object.Countries)
						sort.Strings(object.Countries)
					}

				}
				// Get all persons
				personPage := 1
				for {
					personResp, err := client.CachedR(cacheDirPath).
						SetDoNotParseResponse(false).
						SetHeader("Accept", "application/json").
						SetHeader("User-Agent", uarand.GetRandom()).
						Get(fmt.Sprintf(
							"https://premier.one/api/metainfo/tv/%d/person?page=%d&limit=%d&picture_type=card_group&style=portrait&device=web",
							item.ID, personPage, limit,
						))
					if err != nil {
						log.Println("\tFail get persons", err)
						break
					}
					var personPayload premier.Response[premier.Person]
					if err = json.Unmarshal(personResp.Body(), &personPayload); err != nil {
						log.Println(err)
						break
					}
					if len(personPayload.Results) < 1 {
						break
					}
					if resp.IsCached() {
						fmt.Println("\tTV", item.ID, "- cached iteration persons", personPage, "with", len(personPayload.Results), "items...")
					} else {
						fmt.Println("\tTV", item.ID, "- iteration persons", personPage, "with", len(personPayload.Results), "items...")
					}
					for _, personItem := range personPayload.Results {
						object.Persons = append(object.Persons, personItem.PersonData.Name)
					}
					if !personPayload.HasNext {
						break
					}
					personPage++
				}
				if len(object.Persons) > 1 {
					object.Persons = lo.Uniq(object.Persons)
					sort.Strings(object.Persons)
				}
				object.HasGenres = len(object.Genres) > 0
				object.HasPersons = len(object.Persons) > 0
				object.HasCountries = len(object.Countries) > 0
				if err := batch.Index(fmt.Sprint(item.ID), object); err != nil {
					log.Println(err)
				}
			}
			if err := index.Batch(batch); err != nil {
				log.Println(err)
			}
			if !payload.HasNext {
				break
			}
			page++
		}
		log.Println("Ready!")
	},
}
