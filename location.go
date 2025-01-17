package gpass

import "github.com/Hutchison-Technologies/gpass/walletobjects"

type LatLongPoint struct {
	Latitude  float64
	Longitude float64
}

func (llp *LatLongPoint) toWO() *walletobjects.LatLongPoint {
	return &walletobjects.LatLongPoint{
		Latitude:  llp.Latitude,
		Longitude: llp.Longitude,
	}
}

func locationListToWO(ll []*LatLongPoint) []*walletobjects.LatLongPoint {
	res := make([]*walletobjects.LatLongPoint, len(ll))
	for i, l := range ll {
		res[i] = l.toWO()
	}

	return res
}

func woToLatLongPoint(llp *walletobjects.LatLongPoint) *LatLongPoint {
	if llp == nil {
		return nil
	}

	return &LatLongPoint{
		Latitude:  llp.Latitude,
		Longitude: llp.Longitude,
	}
}
