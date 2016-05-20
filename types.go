package regionid

import (
	"fmt"
)

// Location 三级定位查询结果，无法确定的层级返回nil
type Location interface {
	Country() *Country
	Subdivision() *Subdivision
	City() *City
}

// Iso3166_1 ISO-3166-1国家代码
type Iso3166_1 struct {
	Alpha2  string
	Alpha3  string
	Numeric int
}

// Country 国家
type Country struct {
	ID           int       // 用户系统定义的城市ID，可为0
	Iso          Iso3166_1 // ISO-3166-1定义的国家ID
	Name         string    // 国家名称
	Subdivisions map[string]*Subdivision
}

// Country 返回国家本省
func (country *Country) Country() *Country {
	return country
}

// Subdivision 返回空
func (country *Country) Subdivision() *Subdivision {
	return nil
}

// City 返回空
func (country *Country) City() *City {
	return nil
}

func (country *Country) String() string {
	if country == nil {
		return "未知国家(0/-/-/0)"
	}
	return fmt.Sprintf("%s(%d/%s/%s/%d)", country.Name, country.ID, country.Iso.Alpha2, country.Iso.Alpha3, country.Iso.Numeric)
}

// Subdivision 子区域(省)
type Subdivision struct {
	ID      int    // 用户系统定义
	Iso     string // ISO-3166-2定义的子区域ID
	Name    string // 子区域名称
	country *Country
	Cities  map[string]*City
}

// Country 返回子区域所在国家
func (subdivision *Subdivision) Country() *Country {
	return subdivision.country
}

// Subdivision 返回子区域本身
func (subdivision *Subdivision) Subdivision() *Subdivision {
	return subdivision
}

// City 返回空
func (subdivision *Subdivision) City() *City {
	return nil
}

func (subdivision *Subdivision) String() string {
	if subdivision == nil {
		return "未知子区域(0/-)"
	}
	return fmt.Sprintf("%s(%d/%s)", subdivision.Name, subdivision.ID, subdivision.Iso)
}

// City 城市
type City struct {
	ID          int    // 用户系统定义的城市ID
	Name        string // 城市名称
	subdivision *Subdivision
}

// Country 返回城市所在国家
func (city *City) Country() *Country {
	return city.subdivision.Country()
}

// Subdivision 返回城市所在子区域
func (city *City) Subdivision() *Subdivision {
	return city.subdivision
}

// City 返回城市本身
func (city *City) City() *City {
	return city
}

func (city *City) String() string {
	if city == nil {
		return "未知城市(0)"
	}
	return fmt.Sprintf("%s(%d)", city.Name, city.ID)
}

// ISP 网络服务提供商
type ISP struct {
	ID   int    // ISP ID
	Name string // ISP名称
}

func (isp *ISP) String() string {
	if isp == nil {
		return "未知运营商(0)"
	}
	return fmt.Sprintf("%s(%d)", isp.Name, isp.ID)
}
