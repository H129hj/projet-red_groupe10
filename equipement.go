package projetred

func (c *Character) AddEquipement(nom string, valeur int) {
    if c.equipement == nil {
        c.equipement = make(map[string]int)
    }
    c.equipement[nom] = valeur
}

func (c *Character) RemoveEquipement(nom string) {
    delete(c.equipement, nom)
}
