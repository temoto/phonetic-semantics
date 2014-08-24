package phonosem

import (
	"errors"
)

type FeatureId uint16

type Feature struct {
	Id    FeatureId
	Value int8
}

// FeatureId
// Not using iota here on purpose. The ids must be fixed forever
// for compatibility with external storage and other implementations.
const (
	FeatureIdInvalid         = 0
	FeatureIdGoodBad         = 1
	FeatureIdBigSmall        = 2
	FeatureIdSoftHard        = 3
	FeatureIdFemaleMale      = 4
	FeatureIdLightDark       = 5
	FeatureIdActivePassive   = 6
	FeatureIdSimpleComplex   = 7
	FeatureIdStrongWeak      = 8
	FeatureIdHotCold         = 9
	FeatureIdFastSlow        = 10
	FeatureIdPrettyUgly      = 11
	FeatureIdSmoothRough     = 12
	FeatureIdLightHeavy      = 13
	FeatureIdFunnySad        = 14
	FeatureIdSafeScary       = 15
	FeatureIdMajesticLow     = 16
	FeatureIdBrightDim       = 17
	FeatureIdRoundedAngular  = 18
	FeatureIdJoyfulSorrowful = 19
	FeatureIdLoudQuiet       = 20
	FeatureIdLongShort       = 21
	FeatureIdBraveCowardly   = 22
	FeatureIdKindEvil        = 23
	FeatureIdMightyFeeble    = 24
	FeatureIdMobileSluggish  = 25
)

var featureIds = []FeatureId{
	FeatureIdGoodBad,
	FeatureIdBigSmall,
	FeatureIdSoftHard,
	FeatureIdFemaleMale,
	FeatureIdLightDark,
	FeatureIdActivePassive,
	FeatureIdSimpleComplex,
	FeatureIdStrongWeak,
	FeatureIdHotCold,
	FeatureIdFastSlow,
	FeatureIdPrettyUgly,
	FeatureIdSmoothRough,
	FeatureIdLightHeavy,
	FeatureIdFunnySad,
	FeatureIdSafeScary,
	FeatureIdMajesticLow,
	FeatureIdBrightDim,
	FeatureIdRoundedAngular,
	FeatureIdJoyfulSorrowful,
	FeatureIdLoudQuiet,
	FeatureIdLongShort,
	FeatureIdBraveCowardly,
	FeatureIdKindEvil,
	FeatureIdMightyFeeble,
	FeatureIdMobileSluggish,
}

var featureNameMap = map[string]FeatureId{
	"Хороший-Плохой":           FeatureIdGoodBad,
	"Большой-Маленький":        FeatureIdBigSmall,
	"Нежный-Грубый":            FeatureIdSoftHard,
	"Женственный-Мужественный": FeatureIdFemaleMale,
	"Светлый-Тёмный":           FeatureIdLightDark,
	"Активный-Пассивный":       FeatureIdActivePassive,
	"Простой-Сложный":          FeatureIdSimpleComplex,
	"Сильный-Слабый":           FeatureIdStrongWeak,
	"Горячий-Холодный":         FeatureIdHotCold,
	"Быстрый-Медленный":        FeatureIdFastSlow,
	"Красивый-Отталкивающий":   FeatureIdPrettyUgly,
	"Гладкий-Шероховатый":      FeatureIdSmoothRough,
	"Лёгкий-Тяжёлый":           FeatureIdLightHeavy,
	"Весёлый-Грустный":         FeatureIdFunnySad,
	"Безопасный-Страшный":      FeatureIdSafeScary,
	"Величественный-Низменный": FeatureIdMajesticLow,
	"Яркий-Тусклый":            FeatureIdBrightDim,
	"Округлый-Угловатый":       FeatureIdRoundedAngular,
	"Радостный-Печальный":      FeatureIdJoyfulSorrowful,
	"Громкий-Тихий":            FeatureIdLoudQuiet,
	"Длинный-Короткий":         FeatureIdLongShort,
	"Храбрый-Трусливый":        FeatureIdBraveCowardly,
	"Добрый-Злой":              FeatureIdKindEvil,
	"Могучий-Хилый":            FeatureIdMightyFeeble,
	"Подвижный-Медлительный":   FeatureIdMobileSluggish,
}

func (f Feature) Slug() string {
	return f.Id.Slug()
}

func (fid FeatureId) Slug() string {
	panic("not implemented")
}

func FeatureIdFromName(s string) (FeatureId, error) {
	if id, ok := featureNameMap[s]; ok {
		return id, nil
	}
	return FeatureIdInvalid, errors.New("Unknown feature name: '" + s + "'")
}
