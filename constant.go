package util

import "math"

func GetSPCfactor(samplesize int, SD, RangeSD, RangeCL float64) (A2, D2, D3, D4 float64) {
	switch samplesize {
	case 1:
		A2 = SD * 3 / RangeCL
		D3 = 0.0
		D4 = RangeSD*3/RangeCL + 1
	case 2:
		A2 = 1.880
		D3 = 0.000
		D4 = 3.267
	case 3:
		A2 = 1.023
		D3 = 0.000
		D4 = 2.575
	case 4:
		A2 = 0.729
		D3 = 0.000
		D4 = 2.282
	case 5:
		A2 = 0.577
		D3 = 0.000
		D4 = 2.115
	case 6:
		A2 = 0.483
		D3 = 0.000
		D4 = 2.004
	case 7:
		A2 = 0.419
		D3 = 0.076
		D4 = 1.924
	case 8:
		A2 = 0.373
		D3 = 0.136
		D4 = 1.864
	case 9:
		A2 = 0.337
		D3 = 0.184
		D4 = 1.816
	case 10:
		A2 = 0.308
		D3 = 0.223
		D4 = 1.777
	}

	D2 = Round(3/(A2*math.Sqrt(float64(samplesize))), 3)

	return
}
