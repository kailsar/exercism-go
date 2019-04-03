//space is a package that upon being given the number of seconds
//a person has been alive for and a planet, will tell you how many
//years (rotations round the sun) the person has been alive for, for
//that planet.
package space

type Planet string

const earth_years_to_seconds = 31557600

//A map of planets to their year length in Earth years
var year_length = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Earth": 1,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

//Age is a function that, given a number of seconds, calculates the 
//number of years that represents in the year-length of the given planet
func Age(age_in_seconds float64, planet Planet) float64 {
	
	age_in_earth_years := age_in_seconds / earth_years_to_seconds
	age_in_planet_years := age_in_earth_years / year_length[planet]

	return age_in_planet_years
}