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

var (
	emptyIso3166_2 Iso3166_1
)

// Country represents a country
type Country struct {
	id           int
	iso          Iso3166_1
	name         string
	subdivisions map[string]*Subdivision
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

// ID returns customized ID.
func (country *Country) ID() int {
	if country == nil {
		return 0
	}
	return country.ID()
}

// Name returns country's name.
func (country *Country) Name() string {
	if country == nil {
		return ""
	}
	return country.name
}

// Iso returns ISO-3166-1 ID
func (country *Country) Iso() Iso3166_1 {
	if country == nil {
		return emptyIso3166_2
	}
	return country.iso
}

// String returns a string in format "Name(ID/Iso.Alpha2/Iso.Alpha3/Iso.Numeric)".
func (country *Country) String() string {
	if country == nil {
		return "未知国家(0/-/-/0)"
	}
	return fmt.Sprintf("%s(%d/%s/%s/%d)", country.name, country.id, country.iso.Alpha2, country.iso.Alpha3, country.iso.Numeric)
}

// Subdivision represent a subdivision like province.
type Subdivision struct {
	id      int
	iso     string
	name    string
	cities  map[string]*City
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

// ID returns customized subdivision ID.
func (subdivision *Subdivision) ID() int {
	if subdivision == nil {
		return 0
	}
	return subdivision.id
}

// Iso returns ISO-3166-2 ID.
func (subdivision *Subdivision) Iso() string {
	if subdivision == nil {
		return ""
	}
	return subdivision.iso
}

// Name returns subdivision's name.
func (subdivision *Subdivision) Name() string {
	if subdivision == nil {
		return ""
	}
	return subdivision.name
}

// String return a string in format "Name(ID/Iso)"
func (subdivision *Subdivision) String() string {
	if subdivision == nil {
		return "未知子区域(0/-)"
	}
	return fmt.Sprintf("%s(%d/%s)", subdivision.name, subdivision.id, subdivision.iso)
}

// City represents a city
type City struct {
	id          int
	name        string
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

// ID returns customized city's ID.
func (city *City) ID() int {
	if city == nil {
		return 0
	}
	return city.id
}

// Name returns city's name.
func (city *City) Name() string {
	if city == nil {
		return ""
	}
	return city.name
}

// String returns a string in format "Name(ID)".
func (city *City) String() string {
	if city == nil {
		return "未知城市(0)"
	}
	return fmt.Sprintf("%s(%d)", city.name, city.id)
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
