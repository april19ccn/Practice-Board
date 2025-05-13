package converter

// mg <-> g
func MGToG(mg Milligrams) Gram { return Gram(mg / 1000) }

func GToMG(g Gram) Milligrams { return Milligrams(g * 1000) }

// g <-> kg
func GToKG(g Gram) Kilogram { return Kilogram(g / 1000) }

func KGToG(kg Kilogram) Gram { return Gram(kg * 1000) }

// mg <-> kg
func MGToKG(mg Milligrams) Kilogram { return Kilogram(mg / 1000000) }

func KGToMG(kg Kilogram) Milligrams { return Milligrams(kg * 1000000) }
