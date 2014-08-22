package phonosem

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

func (f Feature) Slug() string {
	return f.Id.Slug()
}

func (fid FeatureId) Slug() string {
	panic("not implemented")
}
