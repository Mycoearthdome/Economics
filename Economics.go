package main

import (
	rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	rand2 "math/rand"
	"time"
)

// Economic-Jihad MacroEconomics.

// Civilians under attacks from their government in full fledge freedoms violations from becomming beligerents
// under a Sovereign's rule to oppress them until they recover their citizenship under a Sovereign's word.
//
// Jihadi's involve themselves in every respect of your laws in having your staff grow their own version of hatrid
// grow the ranks of intelligence and or the army effetively converting them and exposing them to the same challenge
// they carry in their mission.

const WarDuration = 1000 // 1000 year war.

// Jihadi's Ressources
var JiLifetime = 82 // Man life expectancy in Canada
var JiTime = JiLifetime
var JiMoney int = 20000 // 6 months in and out 10k every trip.

var JiRessourcefullness float64 = 0
var JiRessources float64 = 0
var JiDedication float64 = 0
var JiReactions float64 = 0
var JiInterest float64 = 0
var JiActions float64 = 0

// Sovereign's Ressources
var SLifetime = JiLifetime
var STime = SLifetime
var SMoney float64 = 2000000000000        // 2 Trillion a year (CANADA's GDP ... )
var Govt_Expenses float64 = 1200000000000 // 1.2 Trillion a year
var EconomicGrowth float64 = 1.03         // 3% a year
var Population int = 40000000             // ~Canada's Population in 2023
var IntelligenceStaff int = 100000        // an army
var SRessources = IntelligenceStaff

var SPercentGDPDefence float64 = 0.0014 // Nato's requirement. (2%)
var SGovernanceSkills float64 = 0.8
var SInterest float64 = 0.1
var SPolitics float64 = 1
var SDoctrine float64 = 1
var SHealth float64 = 1
var SFaith float64 = 1
var SWord float64 = 1
var SRule float64 = 1

func TerrorScale(Compulsion int) string {
	// Where effort is subtile, seeping and depends on the individual's rigidity (psychorigid/traumatized)
	switch Compulsion {
	case 0:
		return "normal"
	case 1:
		return "deception"
	case 2:
		return "anger"
	case 3:
		return "hate"
	case 4:
		return "rage"
	case 5:
		return "fury"
	case 6:
		return "madness"
	case 7:
		return "fervor"
	case 8:
		return "Transcendence"
	case 9:
		return "Bliss"
	case 10:
		SFaith += 0.01
		SRule += 0.001
		SMoney = SMoney - ReNegate(SMoney*SPercentGDPDefence)
		return "Justice"
	}
	return ""
}

func jiTerrorActivity(Category int) string { // Exploits on an army to engage their actions in hatred against civility.
	switch Category {
	case 0:
		JiRessourcefullness++
		return "reading"
	case 1:
		JiDedication++
		return "planning"
	case 2:
		JiRessources++
		return "acquiring"
	case 3:
		JiDedication++
		return "determined"
	case 4:
		JiReactions++
		return "planting"
	case 5:
		JiInterest++
		return "deceptive"
	case 6:
		JiActions++
		return "arrogant"
	case 7:
		JiActions++
		return "agressive"
	case 8:
		JiRessourcefullness++
		return "in transcendence"
	case 9:
		JiLifetime++
		return "in bliss"
	case 10:
		JiDedication++
		return "in humility"
	}
	return ""
}

func SAttention(Manuevers int) string {
	// Sovereign tries to act in proportionality through it's military manuevers.
	switch Manuevers {
	case 0:
		SGovernanceSkills += 0.1
		return "observe"
	case 1:
		SInterest += 0.1
		return "involve"
	case 2:
		SPolitics += 0.1
		return "act"
	case 3:
		SGovernanceSkills += 0.1
		return "deceive"
	case 4:
		SDoctrine += 0.1
		return "enforce"
	case 5:
		SRule += 0.1
		return "engage"
	case 6:
		SWord += 0.1
		return "is wise"
	case 7:
		SDoctrine += 0.1
		return "determine"
	case 8:
		SFaith += 0.1
		return "is Transcendental"
	case 9:
		SHealth += 0.1
		return "in Bliss"
	case 10:
		SRule += 0.1
		return "is mercyfull"
	}
	return ""
}

func MiRetaliatory(Prerogative int) string {
	// Military's ressourcefullness and creativity in determination over a targeted individual.
	switch Prerogative {
	case 0:
		SGovernanceSkills += 0.01
		return "prepares"
	case 1:
		SDoctrine++
		return "dedicates"
	case 2:
		JiRessources--
		return "undermines"
	case 3:
		JiDedication--
		return "reduces"
	case 4:
		JiInterest--
		return "determines"
	case 5:
		JiActions--
		return "deters"
	case 6:
		JiRessourcefullness--
		return "learns"
	case 7:
		JiActions--
		return "engages"
	case 8:
		SGovernanceSkills += 0.1
		return "reads"
	case 9:
		SFaith++
		return "prays"
	case 10:
		SInterest++
		return "is ready"
	}
	return ""
}

func JiTerrorActivity(Citizen float64) string {
	// Society is composed of members in and under the law that regulate themselves according to the rule of law. When it's better justice to do it yourself
	// a general system failure occurs, there people escape or fight as one.
	// Those are events of terror on the land engaging against security of the persons a year.
	switch {
	case Citizen <= 0.1:
		return "normal" // life carries on people hear about terrorist activity and reflect on their situation.
	case Citizen > 0.1 && Citizen <= 0.2:
		JiRessourcefullness += 0.1
		SPolitics -= 0.1
		return "worried" // talks go about the people.
	case Citizen > 0.2 && Citizen <= 0.3:
		JiRessourcefullness += 0.5
		SPolitics -= 0.2
		return "panicked" // mistakes are made of people becoming agressive against people.
	case Citizen > 0.3 && Citizen <= 0.4:
		JiRessources++
		SFaith -= 0.01
		return "escaping" //flight.
	case Citizen > 0.4 && Citizen <= 0.5:
		SFaith += 0.02
		JiRessourcefullness -= 0.2
		SPolitics += 0.1
		return "returning" //prepare
	case Citizen > 0.5 && Citizen <= 0.6:
		SGovernanceSkills += 0.01
		SInterest += 0.01
		SPolitics -= 0.001
		SDoctrine -= 0.01
		return "preparing"
	case Citizen > 0.6 && Citizen <= 0.7:
		SWord++
		SRule += 0.1
		return "rallying"
	case Citizen > 0.7 && Citizen <= 0.8:
		SHealth -= 0.01
		SPolitics += 0.2
		SDoctrine += 0.01
		SFaith += 0.01
		JiRessourcefullness -= 0.5
		JiRessources -= 0.2
		JiInterest++
		JiActions++
		return "fighting" // warriors are brought of the people.
	case Citizen > 0.8 && Citizen <= 0.9:
		JiRessources -= 0.1
		JiRessourcefullness -= 0.4
		JiDedication += 0.03
		return "friends"
	case Citizen > 0.9 && Citizen <= 1:
		SInterest += 0.01
		SPolitics += 0.1
		SDoctrine += 0.01
		SGovernanceSkills -= 0.01
		JiRessources -= 1
		JiRessourcefullness -= 0.7
		JiDedication += 0.3
		JiReactions++
		JiInterest--
		JiActions += 0.000001
		return "allies"
	case Citizen >= 1:
		JiRessourcefullness++
		JiRessources++
		JiDedication++
		JiReactions++
		JiInterest++
		JiActions++
		JiLifetime++
		SHealth--
		return "losing"
	}
	return ""
}

func SRetaliatory(Resolve float64) string {
	// Sovereign's proportionality in public interests to preserve a sentiment of Peace towards it's people.
	switch {
	case Resolve <= 0.1:
		return "normal" // which is already an act against the context of the law.
	case Resolve > 0.1 && Resolve <= 0.2:
		SInterest += 0.1
		SGovernanceSkills += 0.1
		SPolitics -= 0.01
		return "decisive" // when people start attacking people against Westminster' provision.
	case Resolve > 0.2 && Resolve <= 0.3:
		SRule += 0.1
		SWord += 0.01
		return "involved" // mistakes are made of people becoming agressive against people.
	case Resolve > 0.3 && Resolve <= 0.4:
		SFaith -= 0.001
		SDoctrine -= 0.01
		SInterest += 0.001
		return "accepting" // a general feeling of guilt.
	case Resolve > 0.4 && Resolve <= 0.5:
		SHealth += 0.01
		SPolitics += 0.001
		SFaith += 0.01
		SWord += 0.001
		return "happy"
	case Resolve > 0.5 && Resolve <= 0.6:
		SPolitics += 0.1
		SInterest += 0.01
		SDoctrine += 0.001
		SWord += 0.01
		SRule += 0.1
		SGovernanceSkills += 1
		JiDedication += 0.01
		JiReactions += 0.1
		JiActions += 1
		JiRessources -= 0.5
		JiRessourcefullness += 0.03
		return "mobilizing"
	case Resolve > 0.6 && Resolve <= 0.7:
		SGovernanceSkills += 0.8
		JiReactions -= 0.7
		JiInterest += 0.3
		SRule += 0.1
		SFaith += 0.01
		SWord += 0.1
		SInterest++
		SPolitics -= 0.2
		SHealth += 0.01
		return "instructing"
	case Resolve > 0.7 && Resolve <= 1:
		SFaith += 1
		SHealth += 0.1
		return "well" // a prospect for a Proclamation and/or Peace Treaty be drafted or not.
	case Resolve >= 1:
		JiRessourcefullness++
		JiRessources++
		JiDedication++
		JiReactions--
		JiInterest--
		JiActions++
		SPercentGDPDefence += 0.0001
		SHealth -= 0.5
		SWord += 0.7
		SPolitics++
		SInterest -= 0.4
		SDoctrine++
		SFaith -= 0.3
		SRule += 0.1
		return "worried"
	}
	return ""
}

func cryptoseed() (int64, error) {
	var b [8]byte
	_, err := rand2.Read(b[:])
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

func IntSight(min int, max int) int {
	//min := 0 //start values
	//max := 10 //start values

	rangeBigInt := big.NewInt(int64(max - min + 1))
	randomBigInt, _ := rand.Int(rand.Reader, rangeBigInt)
	randomInt := int(randomBigInt.Int64()) + min

	return randomInt
}

func CitizenResolve(min float64, max float64) float64 {
	//min := 0.0000001 //start values
	//max := 1 //start values

	seed, _ := cryptoseed()
	rand2.Seed(seed)

	return min + rand2.Float64()*(max-min)
}

func ReNegate(capital float64) float64 {
	const week = 168  // hours
	const year = 8760 // hours
	DefenceCapital := capital
	KeptInterestOnTarget := IntSight(2*week, year)                                      // 2weeks to 10 years
	Salary := IntSight(9, 18)                                                           // 75000 Dollars to 150000 Dollars a year
	Involvment := IntSight(3, 100000)                                                   // 3 to 100000
	TimeOnRecord := IntSight(1, year)                                                   // 1 hour to 10 years in hours
	WarEffort := capital - float64(KeptInterestOnTarget*Salary*Involvment*TimeOnRecord) // The cost of that war
	if WarEffort < 0 {
		WarEffort = DefenceCapital
	}
	Attrition := int(WarEffort / float64(Involvment*Salary*TimeOnRecord)) // Attrition strategy considered
	IntelligenceStaff -= Attrition                                        // People we can't pay anymore
	IntelligenceStaff += Attrition                                        // we replace the staff
	Population -= Attrition                                               // The intelligence pool got lower
	if Population < 0 {
		Population = 0
	}
	JiTime += Attrition // To the Glory of Jihad.
	return WarEffort
}

func Strategy(year int, MinInt int, MaxInt int, MinFloat float64, MaxFloat float64) (int, int, float64, float64) {
	HumanInt := IntSight(MinInt, MaxInt)
	MilitaryIntelligence := CitizenResolve(MinFloat, MaxFloat)

	fmt.Sprintf("Year-->%5d, jihadis are %s, the Sovereign %s, the military %s, the people are %s, the Sovereign is %s. Terror levels is %s\n", year, jiTerrorActivity(HumanInt), SAttention(HumanInt), MiRetaliatory(HumanInt), JiTerrorActivity(MilitaryIntelligence), SRetaliatory(MilitaryIntelligence), TerrorScale(HumanInt))

	return HumanInt, MaxInt, MilitaryIntelligence, MaxFloat // temporary
}

func main() {
	var NbBeligerents int
	var KingCosts float64
	var BannerPrinted bool = false
	var MinInt, MaxInt int = 0, 10
	var MinFloat, MaxFloat float64 = 0.0000001, 1
	var CurrentYear = time.Now().Year()
	for year := CurrentYear; year < (CurrentYear + WarDuration); year++ {
		var NewMoney float64 = SMoney
		MinInt, MaxInt, MinFloat, MaxFloat = Strategy(year, MinInt, MaxInt, MinFloat, MaxFloat)
		if MinInt >= MaxInt || MinFloat >= MaxFloat {
			//fmt.Println("-==Peace is restored==-")
			MinInt, MaxInt = 0, 10
			MinFloat, MaxFloat = 0.0000001, 1
			NbBeligerents++ // Caught in the justice system.
		}
		if SMoney > 0 && Population > 0 {
			KingCosts += NewMoney - SMoney
			SMoney += (NewMoney * EconomicGrowth) - Govt_Expenses //carry over from the year before.
			continue
		} else {
			fmt.Printf("------------YEAR-----%d----------RESULTS-----------------------------------------\n", year-CurrentYear) //82 long (41 to center)
			BannerPrinted = true
			break
		}
	}
	if !BannerPrinted {
		fmt.Printf("------WAR--WAGED-FOR---%d-----YEARS--------------------------------RESULTS-\n", WarDuration) //82 long (41 to center)
	}
	fmt.Printf("|%.f%34s | Jihad's Ressourcefullness%16.f|\n", SGovernanceSkills, "Governance", JiRessourcefullness)
	fmt.Printf("|%.f%33s | Jihad's Interest%25.f|\n", SInterest, "Sovereign's Interest", JiInterest)
	fmt.Printf("|%.f%34s | Jihad's Ressources%23.f|\n", SPolitics, "Sovereign's Politics", JiRessources)
	fmt.Printf("|%.f%34s | Jihad's Dedication%23.f|\n", SDoctrine, "Military Doctrine", JiDedication)
	fmt.Printf("|%.f%34s | Jihad's Reactions%24.f|\n", SHealth, "Kingdom's Health", JiReactions)
	fmt.Printf("|%.f%33s | Jihad's Actions%26.f|\n", SFaith, "Kingdom's Faith", JiActions)
	fmt.Printf("|----------------------------------CROWN----------------------------------------|\n")
	if (JiRessourcefullness + JiInterest + JiRessources + JiDedication + JiReactions + JiActions) < (SGovernanceSkills + SInterest + SPolitics + SDoctrine + SHealth + SFaith) {
		fmt.Printf("|%.f%34s | Jihad was DEFEATED!%22.f|\n", SWord, "The King Notoriety", (JiRessourcefullness + JiInterest + JiRessources + JiDedication + JiReactions + JiActions))
	} else {
		fmt.Printf("|%.f%34s | Jihad has WON!%28.f|\n", SWord, "The King Notoriety", (JiRessourcefullness + JiInterest + JiRessources + JiDedication + JiReactions + JiActions))
	}
	if SMoney > 0 && Population > 0 {
		fmt.Printf("|-----------------------------------RULE----------------------------------------|\n")
		fmt.Printf("%19.f%50s            |\n", SRule, "The RULE of his MAJESTY")
		fmt.Printf("%19.f$%48s           |\n", SMoney, "The GDP left for the Colony")
		fmt.Printf("|_______________________________________________________________________________|\n")
	} else {
		fmt.Printf("|----------------------------------COLONY---------------------------------------|\n")
		fmt.Printf("|%18.f$%30s POP(intelligence):%7d|\n", SMoney, "The Colony was DESTROYED", Population)
		fmt.Printf("|-----------------------------ECONOMIC JIHAD WON-----------BELIGERENTS---------%d|\n", NbBeligerents)
		fmt.Printf("|%18.f$%20s against the Beligerent's COST:%8d$|\n", KingCosts, "The KING's Capital", NbBeligerents*JiMoney)
		fmt.Printf("----------------------------------THE-END----------------------------------------\n")
	}

	fmt.Println("  RATIO-->1 Jihad's Martyr  DOLLAR EQUALS", int64(KingCosts/float64(NbBeligerents*JiMoney)), "treasury DOLLARS<--RATIO")
	fmt.Println(" In the end, every Martyr's life was valued to", JiTime, "years by other jihadis\n")
	fmt.Println(" Every Martyr ....lived through abuse of intelligence...an equivalent-->", JiTime/JiLifetime, "lives.\n")

	fmt.Printf("%s\n\n--Magna Carta 1215\n", "'If we, our chief justice, our officials, or any of our servants offend in any \nrespect against any man, or transgress any of the articles of the peace or of\n this security, and the offence is made known to four of the said twenty-five\n barons, they shall come to us – or in our absence from the kingdom to the\n chief justice –  to declare it and claim immediate redress. If we, or in our \nabsence abroad the chief justice, make no redress within forty days, reckoning\n from the day on which the offence was declared to us or to him, the four barons\n shall refer the matter to the rest of the twenty-five barons, who may distrain\n upon and assail us in every way possible, with the support of the whole\n community of the land, by seizing our castles, lands, possessions, or anything\n else saving only our own person and those of the queen and our children, until\n they have secured such redress as they have determined upon. Having secured\n the redress, they may then resume their normal obedience to us.'")
}
