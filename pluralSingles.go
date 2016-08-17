package nouns

import (
	"errors"
	"regexp"
	"strings"
)

var plural = [][]string{
	[]string{"thou", "you"},
	[]string{"m[ae]n$", "men"},
	[]string{"eaux$", "$0"},
	[]string{"(child)(?:ren)?$", "${1}ren"},
	[]string{"(pe)(?:rson|ople)$", "${1}ople"},
	[]string{"(m|l)(?:ice|ouse)$", "${1}ice"},
	[]string{"(matr|cod|mur|sil|vert|ind|append)(?:ix|ex)$", "${1}ices"},
	[]string{"(x|ch|ss|sh|zz)$", "${1}es"},
	[]string{"([^ch][ieo][ln])ey$", "${1}ies"},
	[]string{"([^aeiouy]|qu)y$", "${1}ies"},
	[]string{"(?:(kni|wi|li)fe|(ar|l|ea|eo|oa|hoo)f)$", "${1}${2}ves"},
	[]string{"sis$", "ses"},
	[]string{"(apheli|hyperbat|periheli|asyndet|noumen|phenomen|criteri|organ|prolegomen|hedr|automat)(?:a|on)$", "${1}a"},
	[]string{"(agend|addend|millenni|dat|extrem|bacteri|desiderat|strat|candelabr|errat|ov|symposi|curricul|automat|quor)(?:a|um)$", "${1}a"},
	[]string{"(her|at|gr)o$", "${1}oes"},
	[]string{"(seraph|cherub)(?:im)?$", "${1}im"},
	[]string{"(alumn|alg|vertebr)(?:a|ae)$", "${1}ae"},
	[]string{"(alumn|syllab|octop|vir|radi|nucle|fung|cact|stimul|termin|bacill|foc|uter|loc|strat)(?:us|i)$", "${1}i"},
	[]string{"([^l]ias|[aeiou]las|[emjzr]as|[iu]am)$", "${1}"},
	[]string{"(e[mn]u)s?$", "${1}s"},
	[]string{"(alias|[^aou]us|tlas|gas|ris)$", "${1}es"},
	[]string{"(ax|test)is$", "${1}es"},
	[]string{"([^aeiou]ese)$", "${1}"},
	[]string{"s?$", "s"},
}

var singular = [][]string{
	[]string{"men$", "man"},
	[]string{"(eau)x?$", "${1}"},
	[]string{"(child)ren$", "${1}"},
	[]string{"(pe)(rson|ople)$", "${1}rson"},
	[]string{"(matr|append)ices$", "${1}ix"},
	[]string{"(cod|mur|sil|vert|ind)ices$", "${1}ex"},
	[]string{"(alumn|alg|vertebr)ae$", "${1}a"},
	[]string{"(apheli|hyperbat|periheli|asyndet|noumen|phenomen|criteri|organ|prolegomen|hedr|automat)a$", "${1}on"},
	[]string{"(agend|addend|millenni|dat|extrem|bacteri|desiderat|strat|candelabr|errat|ov|symposi|curricul|quor)a$", "${1}um"},
	[]string{"(alumn|syllab|octop|vir|radi|nucle|fung|cact|stimul|termin|bacill|foc|uter|loc|strat)(?:us|i)$", "${1}us"},
	[]string{"(cris|test|diagnos)(?:is|es)$", "${1}is"},
	[]string{"(movie|twelve)s$", "${1}"},
	[]string{"(e[mn]u)s?$", "${1}"},
	[]string{"(x|ch|ss|sh|zz|tto|go|cho|alias|[^aou]us|tlas|gas|(?:her|at|gr)o|ris)(?:es)?$", "${1}"},
	[]string{"(seraph|cherub)im$", "${1}"},
	[]string{"(m|l)ice$", "${1}ouse"},
	[]string{"\\b(mon|smil)ies$", "${1}ey"},
	[]string{"\\b([pl]|zomb|(?:neck|cross)?$|coll|faer|food|gen|goon|group|lass|talk|goal|cut)ies$", "${1}ie"},
	[]string{"ies$", "y"},
	[]string{"(ar|(?:wo|[ae])l|[eo][ao])ves$", "${1}f"},
	[]string{"(wi|kni|(?:after|half|high|low|mid|non|night|[^\\w]|^)li)ves$", "${1}fe"},
	[]string{"(^analy)(?:sis|ses)$", "${1}sis"},
	[]string{"((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(?:sis|ses)$", "${1}sis"},
	[]string{"(ss)$", "${1}"},
	[]string{"s$", ""},
}

var irregulars = [][]string{
	[]string{"I", "we"},
	[]string{"me", "us"},
	[]string{"he", "they"},
	[]string{"she", "they"},
	[]string{"them", "them"},
	[]string{"myself", "ourselves"},
	[]string{"yourself", "yourselves"},
	[]string{"itself", "themselves"},
	[]string{"herself", "themselves"},
	[]string{"himself", "themselves"},
	[]string{"themself", "themselves"},
	[]string{"is", "are"},
	[]string{"was", "were"},
	[]string{"has", "have"},
	[]string{"this", "these"},
	[]string{"that", "those"},
	// Words ending in with a consonant and `o`.
	[]string{"echo", "echoes"},
	[]string{"dingo", "dingoes"},
	[]string{"volcano", "volcanoes"},
	[]string{"tornado", "tornadoes"},
	[]string{"torpedo", "torpedoes"},
	[]string{"embargo", "embargoes"},
	[]string{"veto", "vetoes"},
	// Ends with `us`.
	[]string{"genus", "genera"},
	[]string{"viscus", "viscera"},
	// Ends with `ma`.
	[]string{"stigma", "stigmata"},
	[]string{"stoma", "stomata"},
	[]string{"dogma", "dogmata"},
	[]string{"lemma", "lemmata"},
	[]string{"schema", "schemata"},
	[]string{"anathema", "anathemata"},
	// Other irregular rules.
	[]string{"ox", "oxen"},
	[]string{"axe", "axes"},
	[]string{"die", "dice"},
	[]string{"yes", "yeses"},
	[]string{"foot", "feet"},
	[]string{"eave", "eaves"},
	[]string{"goose", "geese"},
	[]string{"tooth", "teeth"},
	[]string{"quiz", "quizzes"},
	[]string{"human", "humans"},
	[]string{"proof", "proofs"},
	[]string{"carve", "carves"},
	[]string{"valve", "valves"},
	[]string{"looey", "looies"},
	[]string{"thief", "thieves"},
	[]string{"groove", "grooves"},
	[]string{"pickaxe", "pickaxes"},
	[]string{"whiskey", "whiskies"},
}

var uncountables = []string{
	"advice",
	"adulthood",
	"agenda",
	"aid",
	"alcohol",
	"ammo",
	"athletics",
	"bison",
	"blood",
	"bream",
	"buffalo",
	"butter",
	"carp",
	"cash",
	"chassis",
	"chess",
	"clothing",
	"commerce",
	"cod",
	"cooperation",
	"corps",
	"digestion",
	"debris",
	"diabetes",
	"energy",
	"equipment",
	"elk",
	"excretion",
	"expertise",
	"flounder",
	"fun",
	"gallows",
	"garbage",
	"graffiti",
	"headquarters",
	"health",
	"herpes",
	"highjinks",
	"homework",
	"housework",
	"information",
	"jeans",
	"justice",
	"kudos",
	"labour",
	"literature",
	"machinery",
	"mackerel",
	"mail",
	"media",
	"mews",
	"moose",
	"music",
	"news",
	"pike",
	"plankton",
	"pliers",
	"pollution",
	"premises",
	"rain",
	"research",
	"rice",
	"salmon",
	"scissors",
	"series",
	"scissors",
	"series",
	"sewage",
	"rice",
	"salmon",
	"scissors",
	"series",
	"sewage",
	"shambles",
	"shrimp",
	"species",
	"staff",
	"swine",
	"trout",
	"traffic",
	"transporation",
	"tuna",
	"wealth",
	"welfare",
	"sewage",
	"rice",
	"salmon",
	"scissors",
	"series",
	"sewage",
	"shambles",
	"shrimp",
	"species",
	"staff",
	"swine",
	"trout",
	"traffic",
	"transporation",
	"tuna",
	"wealth",
	"welfare",
	"shambles",
	"shrimp",
	"species",
	"staff",
	"swine",
	"trout",
	"traffic",
	"transporation",
	"tuna",
	"wealth",
	"welfare",
	"whiting",
	"wildebeest",
	"wildlife",
	"you",
	// Regexes.
	"rice",
	"salmon",
	"scissors",
	"series",
	"sewage",
	"shambles",
	"shrimp",
	"species",
	"staff",
	"swine",
	"trout",
	"traffic",
	"transporation",
	"tuna",
	"wealth",
	"welfare",
	"pox$",
	"ois$",
	"deer$",
	"fish$",
	"sheep$",
	"measles$",
	"[^aeiou]ese$",
}

func transform(old string, rules [][]string, irregularFrom, irregularTo int) (string, error) {

	// Check Iregular cases
	for _, iregular := range irregulars {
		compare := iregular[irregularFrom]

		if compare == strings.ToLower(old) {
			return iregular[irregularTo], nil
		}
	}

	// Check Uncountable cases
	for _, uncountable := range uncountables {
		if uncountable == old {
			return old, nil
		}
		if strings.Contains(uncountable, "$") {
			if match, errs := regexp.MatchString(uncountable, old); match {
				if errs != nil {
					return "", errs
				}
				return old, nil
			}
		}
	}

	// Check the Rule cases
	for _, rule := range rules {
		regex, replace := rule[0], rule[1]

		if match, errs := regexp.MatchString(regex, old); match {
			if errs != nil {
				return "", errs
			}

			regex, errs := regexp.Compile(regex)

			if errs != nil {
				return "", errs
			}
			return regex.ReplaceAllString(old, replace), nil
		}
	}
	return "", errors.New("Could not transform string.")
}

//Pluralize converts a singular word (string) to a plural word
func Pluralize(singular string) (string, error) {
	return transform(singular, plural, 0, 1)
}

//Singularize converts a plural word (string) to a singular word
func Singularize(pural string) (string, error) {
	return transform(pural, singular, 1, 0)
}
