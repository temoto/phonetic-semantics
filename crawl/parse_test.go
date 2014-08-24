package main

import (
	"strings"
	"testing"

	"github.com/temoto/phonetic-semantics/phonosem"
)

func TestParse01(t *testing.T) {
	s := `
		<tr>
		<td>Большой-Маленький</td>
		<td>2,4</td>
		<td><div style='background: blue; width: 21%; height: 20px;'></div></td>
		<td><span id='blue'>Большой</span></td>
	</tr>
		<tr>
		<td>Нежный-Грубый</td>
		<td>4</td>
		<td><div style='background: red; width: 34%; height: 20px;'></div></td>
		<td><span id='red'>Грубый</span></td>
	</tr>
`
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	features, err := ParseFeatures([]byte(s))
	check(err)
	if len(features) != 2 {
		t.Fatalf("Expected len(features) = 2; found %d", len(features))
	}
	if features[0].Id != phonosem.FeatureIdBigSmall {
		t.Fatalf("Expected feature[0] BigSmall (%d); found %d", phonosem.FeatureIdBigSmall, features[0].Id)
	}
	if features[0].Value != -5 {
		t.Fatalf("Expected feature[0] value (%d); found %d", -5, features[0].Value)
	}
	if features[1].Id != phonosem.FeatureIdSoftHard {
		t.Fatalf("Expected feature[1] SoftHard (%d); found %d", phonosem.FeatureIdSoftHard, features[1].Id)
	}
	if features[1].Value != 10 {
		t.Fatalf("Expected feature[1] value (%d); found %d", 10, features[1].Value)
	}
	t.Logf("testParse01 %v", features)
}
