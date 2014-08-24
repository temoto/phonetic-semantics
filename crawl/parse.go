package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/temoto/phonetic-semantics/phonosem"
)

var reFeature *regexp.Regexp

func parseValue(s string) (int8, error) {
	s = strings.Replace(s, ",", ".", 1)
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	v := int8(f64*10 - 30)
	return v, nil
}

func ParseFeatures(b []byte) ([]phonosem.Feature, error) {
	s := strings.NewReplacer("\r", "", "\n", "", "\t", "").Replace(string(b))
	// log.Printf("ParseFeatures: input: %s", s)

	features := make([]phonosem.Feature, 0, 25)
	matches := reFeature.FindAllStringSubmatch(s, -1)
	// log.Printf("ParseFeatures: matches: %d", len(matches))
	for _, group := range matches {
		// group[0] - whole match
		// group[1] - capture 1, feature name
		// group[2] - capture 2, value
		featureName := group[1]
		featureValue, err := parseValue(group[2])
		check(err)
		log.Printf("feature %s %d", featureName, featureValue)
		featureId, err := phonosem.FeatureIdFromName(featureName)
		check(err)
		feature := phonosem.Feature{
			Id:    featureId,
			Value: featureValue,
		}
		features = append(features, feature)
	}

	return features, nil
}

func init() {
	reFeature = regexp.MustCompile(`<tr>\s*<td>([^<-]+-[^<-]+)</td>\s*<td>(\d+(?:,\d+)?)</td>`)
}
