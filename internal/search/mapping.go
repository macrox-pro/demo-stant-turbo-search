package search

import (
	"log"

	"github.com/blevesearch/bleve/v2/analysis/lang/en"

	"github.com/blevesearch/bleve/v2/analysis/analyzer/standard"

	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"

	"github.com/blevesearch/bleve/v2/analysis/tokenizer/unicode"

	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/lang/ru"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/token/ngram"
	"github.com/blevesearch/bleve/v2/mapping"
)

func NewIndexMapping() *mapping.IndexMappingImpl {
	m, filters := bleve.NewIndexMapping(), []string{
		lowercase.Name,
		en.StopName,
		ru.StopName,
		ru.SnowballStemmerName,
	}
	m.DefaultAnalyzer = ru.AnalyzerName
	if err := m.AddCustomTokenFilter("ngram_1_2", map[string]interface{}{
		"type": ngram.Name,
		"min":  1,
		"max":  2,
	}); err != nil {
		log.Fatalln(err)
	} else {
		filters = append(filters, "ngram_1_2")
	}
	if err := m.AddCustomAnalyzer("custom_ru", map[string]interface{}{
		"type":          custom.Name,
		"tokenizer":     unicode.Name,
		"token_filters": filters,
	}); err != nil {
		log.Fatalln(err)
	} else {
		ruFieldMapping := bleve.NewTextFieldMapping()
		ruFieldMapping.Analyzer = "custom_ru"

		ignoreFieldMapping := bleve.NewTextFieldMapping()
		ignoreFieldMapping.IncludeTermVectors = false
		ignoreFieldMapping.IncludeInAll = false
		ignoreFieldMapping.Index = false

		standardFieldMapping := bleve.NewTextFieldMapping()
		standardFieldMapping.Analyzer = standard.Name

		booleanFieldMapping := bleve.NewBooleanFieldMapping()
		keywordFieldMapping := bleve.NewKeywordFieldMapping()

		docMapping := bleve.NewDocumentMapping()
		docMapping.DefaultAnalyzer = ru.AnalyzerName

		// RU fields
		docMapping.AddFieldMappingsAt("title", ruFieldMapping)

		// Simple fields
		docMapping.AddFieldMappingsAt("ageRestriction", keywordFieldMapping)
		docMapping.AddFieldMappingsAt("yearStart", keywordFieldMapping)
		docMapping.AddFieldMappingsAt("yearEnd", keywordFieldMapping)
		docMapping.AddFieldMappingsAt("year", keywordFieldMapping)

		// Ignored
		docMapping.AddFieldMappingsAt("picture", ignoreFieldMapping)

		// Boolean fields
		docMapping.AddFieldMappingsAt("isActive", booleanFieldMapping)
		docMapping.AddFieldMappingsAt("hasGenres", booleanFieldMapping)
		docMapping.AddFieldMappingsAt("hasPersons", booleanFieldMapping)
		docMapping.AddFieldMappingsAt("hasCountries", booleanFieldMapping)

		m.DefaultMapping = docMapping
	}
	return m
}
