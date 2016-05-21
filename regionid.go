package regionid

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	worldCountries map[string]*Country
	worldIsps      map[string]*ISP
)

// Initialized check whether the world is initialized.
func Initialized() bool {
	return worldCountries != nil && worldIsps != nil
}

// LoadBuiltinWorld load the build world
// The build-in world only contains all country defined by ISO-3166-1 and all cities (and the subdivisions) of China defined by National Bureau of Statistics of China.
func LoadBuiltinWorld() error {
	return LoadWorld(builtinWorld)
}

// LoadWorld load external world data from a byte slice
func LoadWorld(data []byte) error {
	countries := make(map[string]*Country)
	isps := make(map[string]*ISP)
	var country *Country
	var subdivision *Subdivision
	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if len(text) == 0 {
			continue
		}
		record := strings.Split(text, ",")
		switch record[0] {
		case "country":
			country = parseCountry(record[1:])
			if country == nil {
				return fmt.Errorf("错误的国家ID数据: %v", record)
			}
			countries[country.name] = country
		case "subdivision":
			subdivision = parseSubdivision(record[1:], country)
			if subdivision == nil {
				return fmt.Errorf("错误的子区域ID数据: %v", record)
			}
			country.subdivisions[subdivision.name] = subdivision
		case "city":
			city := parseCity(record[1:], subdivision)
			if city == nil {
				return fmt.Errorf("错误的城市ID数据: %v", record)
			}
			subdivision.cities[city.name] = city
		case "isp":
			isp, aliases := parseISP(record[1:])
			if isp == nil {
				return fmt.Errorf("错误的ISP ID数据: %v", record)
			}
			isps[isp.Name] = isp
			for _, alias := range aliases {
				isps[alias] = isp
			}
		case "":
		default:
			return fmt.Errorf("未知数据行格式: %v", record)
		}
	}
	worldCountries = countries
	worldIsps = isps
	return nil
}

// LoadWorldFromStream load external world data from a stream
func LoadWorldFromStream(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return LoadWorld(data)
}

// LoadWorldFromFile load external world data from a local file
func LoadWorldFromFile(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return LoadWorld(content)
}

// GetLocation get the Location speicified by "country", "subdivision" and "city" from the world.
// If the city is not found, return the subdivision.
// If the subdivision is not found, return the country.
// If the country is not found, return nil.
func GetLocation(country, subdivision, city string) Location {
	if worldCountries == nil {
		return nil
	}
	lv1 := worldCountries[country]
	if lv1 == nil {
		return nil
	}
	lv2 := lv1.subdivisions[subdivision]
	if lv2 == nil {
		return lv1
	}
	lv3 := lv2.cities[city]
	if lv3 == nil {
		return lv2
	}
	return lv3
}

// GetISP get the ISP instance from the world
func GetISP(name string) *ISP {
	return worldIsps[name]
}

func parseCountry(fields []string) *Country {
	if len(fields) != 5 {
		return nil
	}
	id, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return nil
	}
	isoNumeric, err := strconv.ParseInt(fields[3], 10, 32)
	if err != nil {
		return nil
	}
	return &Country{
		id: int(id),
		iso: Iso3166_1{
			Alpha2:  fields[1],
			Alpha3:  fields[2],
			Numeric: int(isoNumeric),
		},
		name:         fields[4],
		subdivisions: make(map[string]*Subdivision),
	}
}

func parseSubdivision(fields []string, country *Country) *Subdivision {
	if len(fields) != 3 || country == nil {
		return nil
	}
	id, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return nil
	}
	return &Subdivision{
		id:      int(id),
		iso:     fields[1],
		name:    fields[2],
		country: country,
		cities:  make(map[string]*City),
	}
}

func parseCity(fields []string, subdivision *Subdivision) *City {
	if len(fields) != 2 {
		return nil
	}
	id, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return nil
	}
	return &City{
		id:          int(id),
		name:        fields[1],
		subdivision: subdivision,
	}
}

func parseISP(fields []string) (*ISP, []string) {
	if len(fields) < 2 {
		return nil, nil
	}
	id, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return nil, nil
	}
	if len(fields) == 2 {
		return &ISP{
			ID:   int(id),
			Name: fields[1],
		}, nil
	}
	return &ISP{
		ID:   int(id),
		Name: fields[1],
	}, fields[2:]
}
