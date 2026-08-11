package tgn

func TgnToWhosOnFirstPlacetype(tgn_pt string) string {

	var pt string

	switch tgn_pt {
	case "81021/dependent state", "82133/governorate", "81181/territory":
		pt = "dependency"
	case "81161/province", "81145/oblast", "81165/region (administrative division)", "81166/autonomous region", "81126/national division", "81117/department":
		pt = "region"
	case "81260/controlled region":
		pt = "region"
	case "81151/parish (political)", "81225/regional division":
		pt = "county"
	case "81118/local council":
		pt = "localadmin"
	case "81010/nation", "82171/republic":
		pt = "country"
	case "81105/canton":
		pt = "canton"
	case "83044/special municipality":
		pt = "locality"
	case "81170/autonomous community":
		pt = "localadmin"
	case "81157/prefecture":
		pt = "borough"
	case "83002/inhabited place", "83043/municipality":
		pt = "locality"
	case "81100/first level subdivision":
		pt = "region"
	case "83007/locality":
		pt = "locality"
	case "82411/borough":
		pt = "borough"
	case "82402/fourth level subdivision":
		pt = "localadmin"
	case "51212/country house", "51011/building", "51821/church", "51828/temple", "51211/house", "51239/hut", "51511/office building":
		pt = "building"
	case "51556/college", "83371/industrial center", "51554/hospital", "54551/airport", "51557/university", "54253/air base":
		pt = "campus"
	case "29000/continent":
		pt = "continent"
	case "82471/township":
		pt = "localadmin"
	case "81102/country":
		pt = "country"
	case "81114/autonomous city", "83040/city", "83010/hamlet":
		pt = "locality"
	case "81175/state":
		pt = "region"
	case "84251/neighborhood":
		pt = "neighbourhood"
	case "81300/second level subdivision":
		pt = "county"
	case "82401/third level subdivision":
		pt = "macrocounty"
	case "54051/ranch":
		pt = "campus"
	case "81115/county":
		pt = "county"
	default:
		pt = "custom"
	}

	return pt
}
