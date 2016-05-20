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

func LoadBuildWorld() error {
	return LoadWorld(builtinWorld)
}

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
			countries[country.Name] = country
		case "subdivision":
			subdivision = parseSubdivision(record[1:], country)
			if subdivision == nil {
				return fmt.Errorf("错误的子区域ID数据: %v", record)
			}
			country.Subdivisions[subdivision.Name] = subdivision
		case "city":
			city := parseCity(record[1:], subdivision)
			if city == nil {
				return fmt.Errorf("错误的城市ID数据: %v", record)
			}
			subdivision.Cities[city.Name] = city
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

func LoadWorldFromStream(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return LoadWorld(data)
}

func LoadWorldFromFile(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return LoadWorld(content)
}

func GetLocation(country, subdivision, city string) Location {
	if worldCountries == nil {
		return nil
	}
	if lv1 := worldCountries[country]; lv1 == nil {
		return nil
	} else if lv2 := lv1.Subdivisions[subdivision]; lv2 == nil {
		return lv1
	} else if lv3 := lv2.Cities[city]; lv3 == nil {
		return lv2
	} else {
		return lv3
	}
}

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
		ID: int(id),
		Iso: Iso3166_1{
			Alpha2:  fields[1],
			Alpha3:  fields[2],
			Numeric: int(isoNumeric),
		},
		Name:         fields[4],
		Subdivisions: make(map[string]*Subdivision),
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
		ID:      int(id),
		Iso:     fields[1],
		Name:    fields[2],
		country: country,
		Cities:  make(map[string]*City),
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
		ID:          int(id),
		Name:        fields[1],
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
