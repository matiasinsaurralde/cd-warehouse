package main

type Warehouse struct {
	CDs []*CD
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		CDs: make([]*CD, 0),
	}
}

func (w *Warehouse) AddCD(cd *CD) {
	w.CDs = append(w.CDs, cd)
}
func (w *Warehouse) FindCD(lookupCD *CD) []*CD {
	results := make([]*CD, 0)
	for _, cd := range w.CDs {
		var match bool
		if lookupCD.Artist != "" && lookupCD.Artist == cd.Artist {
			match = true
		}
		if cd.Name != "" && lookupCD.Name == cd.Name {
			match = true
		}
		if !match {
			continue
		}
		results = append(results, cd)
	}
	return results
}
