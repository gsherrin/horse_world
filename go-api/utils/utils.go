package utils

func ClassColorPicker(a string) string {
	classColor := a
	var xColor string
	switch classColor {
	case "WARLOCK":
		warlockColor := "#8788EE"
		c := &xColor
		*c = warlockColor
	case "MAGE":
		mageColor := "#3FC7EB"
		c := &xColor
		*c = mageColor
	case "HUNTER":
		hunterColor := "#AAD372"
		c := &xColor
		*c = hunterColor
	case "PRIEST":
		priestColor := "#FFFFFF"
		c := &xColor
		*c = priestColor
	case "ROGUE":
		rogueColor := "#FFF468"
		c := &xColor
		*c = rogueColor
	case "DRUID":
		druidColor := "#FF7C0A"
		c := &xColor
		*c = druidColor
	case "SHAMAN":
		shamanColor := "#0070DD"
		c := &xColor
		*c = shamanColor
	case "WARRIOR":
		warriorColor := "#C69B6D"
		c := &xColor
		*c = warriorColor
	case "PALADIN":
		paladinColor := "#F48CBA"
		c := &xColor
		*c = paladinColor
	case "DEMONHUNTER":
		dhColor := "#A330C9"
		c := &xColor
		*c = dhColor
	case "NEUTRAL":
		neutralColor := "#7d7d7d"
		c := &xColor
		*c = neutralColor
	}
	return xColor
}

func RarityColorPicker(a string) string {
	rarityColor := a
	var yColor string
	switch rarityColor {
	case "COMMON":
		commonColor := "#9d9d9d"
		c := &yColor
		*c = commonColor
	case "RARE":
		rareColor := "#0070dd"
		c := &yColor
		*c = rareColor
	case "EPIC":
		epicColor := "#a335ee"
		c := &yColor
		*c = epicColor
	case "LEGENDARY":
		legColor := "#ff8000"
		c := &yColor
		*c = legColor
	case "FREE":
		freeColor := "#00ccff"
		c := &yColor
		*c = freeColor
	}
	return yColor
}
