# go-regionid

go-regionid is a library to represent a world in 3-level(country, subdivision and cities) and get a singleton region by name.
It is used by some IP locating library.

This library contains a built-in world for China users. It contains all the country defined by ISO-3166-1 and all the cities (and subdivisions) defined by of China defined by National Bureau of Statistics of China.

##Example

    regionid.LoadBuiltinWorld()
    location := regionid.GetLocation("中国", "上海", "")

The build in world contains:

##WorldData

The world data is represented by multi-line text. Each line represents a region level or a isp.
The country format:

    country,ID,Alpha2,Alpha3,Numeric,Name

The subdivision format:

    subdivision,ID,ISO-3166-2,Name

The city format:

    city,ID,Name

The isp:

    isp,ID,Name

The line must be organized in order like:

    country1
    subdivision1.1
    city1.1.1
    city1.1.2
    ...
    subdivision1.2
    ...
    country2
    ...
    isp1
    isp2

The missing level can be ignored. The builtin.go contains a real world example.
