package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TrigramInfo struct {
	Name        string
	Element     string
	Meaning     string
	Description string
	ImageDesc   string
	Conclusion  string
}

func formatLines(s string) string {
	return strings.ReplaceAll(s, ". ", ".\n")
}

var trigramData = map[string]TrigramInfo{
	"000": {"Réceptivité", "la Terre", "Ancrage", "La qualité d'être ouvert, stable et de soutien.", "le sol fertile qui accueille la semence.", "Dites oui à la vie et laissez-vous porter par le flux."},
	"111": {"Initiation", "le Ciel", "Créativité", "L'étincelle des nouveaux départs et la force pure.", "la voûte céleste, vaste et infinie.", "Osez sortir des sentiers battus et affirmez votre vision."},
	"100": {"Mouvement", "le Tonnerre", "Éveil", "La poussée soudaine d'énergie ou le changement.", "l'ébranlement qui réveille la vie.", "Agissez avec audace mais restez attentif aux signes."},
	"010": {"Profondeur", "l'Eau", "Flux", "La capacité à naviguer à travers les défis.", "le courant qui s'adapte à tous les obstacles.", "Restez persévérant tout en gardant votre calme intérieur."},
	"001": {"Immobilité", "la Montagne", "Stabilité", "Le pouvoir du calme et de la fixation des limites.", "le sommet immobile qui touche les nuages.", "Sachez vous arrêter et cultivez votre jardin intérieur."},
	"110": {"Souplesse", "le Vent", "Influence", "La force subtile qui façonne la réalité.", "le souffle qui pénètre partout sans effort.", "Privilégiez la douceur et l'influence subtile."},
	"101": {"Clarté", "le Feu", "Rayonnement", "La lumière de la compréhension et de la conscience.", "la flamme qui illumine et rend visible.", "Voyez les choses telles qu'elles sont et agissez avec discernement."},
	"011": {"Ouverture", "le Lac", "Interaction", "La joie de la connexion et de l'expression.", "l'étendue d'eau paisible qui reflète le monde.", "Partagez votre enthousiasme et ouvrez-vous aux autres."},
}

func main() {
	basePath := "internal/adapters/file/data/hexagrams"
	trigramPath := "internal/adapters/file/data/trigrams"
	os.MkdirAll(basePath, 0755)
	os.MkdirAll(trigramPath, 0755)

	// Génération des Trigrammes
	for i := 0; i < 8; i++ {
		binaryStr := fmt.Sprintf("%03b", i)
		info := trigramData[binaryStr]

		path := filepath.Join(trigramPath, fmt.Sprintf("trigram_%s.md", binaryStr))
		content := fmt.Sprintf("# %s (%s)\n\n**%s**\n\n%s\nReprésente %s.", info.Name, info.Element, info.Meaning, formatLines(info.Description), info.ImageDesc)

		os.WriteFile(path, []byte(content), 0644)
	}

	// Génération des Hexagrammes
	for i := 0; i < 64; i++ {
		binaryStr := fmt.Sprintf("%06b", i)
		lowerBin := binaryStr[0:3]
		upperBin := binaryStr[3:6]

		lower := trigramData[lowerBin]
		upper := trigramData[upperBin]

		title := fmt.Sprintf("%s sur %s", upper.Element, lower.Element)
		if lowerBin == upperBin {
			title = fmt.Sprintf("%s Pur", lower.Element)
		}

		path := filepath.Join(basePath, fmt.Sprintf("hexagram_%s.md", binaryStr))

		content := fmt.Sprintf("# %s\n\n", title)

		// Section L'Image
		content += "### L'Image\n"
		content += formatLines(fmt.Sprintf("Ici, %s se trouve au-dessus de %s. Cela évoque %s rencontrant %s. Cette rencontre naturelle définit la dynamique visuelle de cet état.\n\n", upper.Element, lower.Element, upper.ImageDesc, lower.ImageDesc))

		// Section Analyse
		content += "### Analyse\n"
		analyse := fmt.Sprintf("Cet état représente l'interaction entre %s (%s) et %s (%s). ", upper.Name, upper.Meaning, lower.Name, lower.Meaning)
		analyse += fmt.Sprintf("La force de %s agit sur la base fournie par %s. ", upper.Meaning, lower.Meaning)
		analyse += "Il s'agit d'une phase de transition où le monde intérieur et les actions extérieures cherchent un point d'équilibre. "
		analyse += "L'influence mutuelle de ces deux forces crée une configuration unique qui demande une attention particulière à la cohérence de vos intentions.\n\n"
		content += formatLines(analyse)

		// Section Guidance
		content += "### Guidance\n"
		guidance := fmt.Sprintf("Dans cette situation, privilégiez l'harmonie entre %s et %s. ", lower.Meaning, upper.Meaning)
		guidance += "Observez comment les conditions extérieures résonnent avec votre structure interne. "
		guidance += "Le succès dépend de votre capacité à rester fidèle à la nature de ces deux éléments sans en favoriser un au détriment de l'autre."
		content += formatLines(guidance)

		// Section Conclusion
		content += "\n\n### Conseil du jour\n"
		// Mélange des deux conseils pour créer une synthèse
		conclusion := fmt.Sprintf("%s %s", lower.Conclusion, upper.Conclusion)
		if lowerBin == upperBin {
			conclusion = lower.Conclusion
		}
		content += conclusion

		os.WriteFile(path, []byte(content), 0644)
	}
	fmt.Println("Génération enrichie avec conseils personnalisés terminée.")
}
