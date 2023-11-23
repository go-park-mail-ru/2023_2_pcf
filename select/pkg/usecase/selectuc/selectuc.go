package selectuc

import "AdHub/select/pkg/entities"

type SelectUseCase struct {
}

func New() *SelectUseCase {
	return &SelectUseCase{}
}

func (su SelectUseCase) GetAd(adtarget []*entities.Target, pad *entities.Target) int {
	res := adtarget[0].Id
	for i := 0; i < len(adtarget); i++ {
		ad := adtarget[i]
		if ad.Min_age > pad.Min_age {
			if ad.Max_age < pad.Max_age {
				var count int
				if len(ad.Tags) < len(pad.Tags) {
					for j := 0; j < len(ad.Tags); j++ {
						if ad.Tags[j] == pad.Tags[j] {
							count++
						}

					}
				} else {
					for j := 0; j < len(pad.Tags); j++ {
						if ad.Tags[j] == pad.Tags[j] {
							count++
						}
					}
				}
				if count > 1 {
					res = ad.Id
				}
			}
		}
	}

	return res
}
