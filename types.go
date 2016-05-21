package regionid

import (
	"fmt"
)

// Location represents a location object.
// A location includes 3 levels: country, subdivision and city.
type Location interface {
	Country() *Country
	Subdivision() *Subdivision
	City() *City
}

// Iso3166_1 represents ISO-3166-1 country code
type Iso3166_1 struct {
	Alpha2  string
	Alpha3  string
	Numeric int
}

// Country represents a country
type Country struct {
	// Customized country ID
	ID int

	// ISO-3166-1 country ID
	Iso Iso3166_1

	// Name
	Name string

	// All subdivisions with subdivision's name as key and subdivision instance as value
	Subdivisions map[string]*Subdivision
}

// Country returns the country itself.
func (country *Country) Country() *Country {
	return country
}

// Subdivision always returns nil.
func (country *Country) Subdivision() *Subdivision {
	return nil
}

// City always returns nil.
func (country *Country) City() *City {
	return nil
}

// String returns a string in format "Name(ID/Iso.Alpha2/Iso.Alpha3/Iso.Numeric)".
func (country *Country) String() string {
	if country == nil {
		return "未知国家(0/-/-/0)"
	}
	return fmt.Sprintf("%s(%d/%s/%s/%d)", country.Name, country.ID, country.Iso.Alpha2, country.Iso.Alpha3, country.Iso.Numeric)
}

// Subdivision represent a subdivision like province.
type Subdivision struct {
	// Customized subdivision ID
	ID int

	// ISO-3166-2 ID
	Iso string

	// Name
	Name string

	// All cities in the subdivision, with the city's name as key and the city instance as value
	Cities map[string]*City

	country *Country
}

// Country returns the country which the subdivision belongs to.
func (subdivision *Subdivision) Country() *Country {
	return subdivision.country
}

// Subdivision returns the subdivision itself.
func (subdivision *Subdivision) Subdivision() *Subdivision {
	return subdivision
}

// City always returns nil.
func (subdivision *Subdivision) City() *City {
	return nil
}

// String return a string in format "Name(ID/Iso)"
func (subdivision *Subdivision) String() string {
	if subdivision == nil {
		return "未知子区域(0/-)"
	}
	return fmt.Sprintf("%s(%d/%s)", subdivision.Name, subdivision.ID, subdivision.Iso)
}

// City represents a city
type City struct {
	// Customized city ID
	ID int

	// Name
	Name string

	subdivision *Subdivision
}

// Country returns the country of the city.
func (city *City) Country() *Country {
	return city.subdivision.Country()
}

// Subdivision returns the subdivision of the city.
func (city *City) Subdivision() *Subdivision {
	return city.subdivision
}

// City returns the city itself.
func (city *City) City() *City {
	return city
}

// String returns a string in format "Name(ID)".
func (city *City) String() string {
	if city == nil {
		return "未知城市(0)"
	}
	return fmt.Sprintf("%s(%d)", city.Name, city.ID)
}

// ISP represent a Internet service provider
type ISP struct {
	// Customized ID
	ID int

	// Name
	Name string
}

// String returns a string in format "Name(ID)"
func (isp *ISP) String() string {
	if isp == nil {
		return "未知运营商(0)"
	}
	return fmt.Sprintf("%s(%d)", isp.Name, isp.ID)
}
