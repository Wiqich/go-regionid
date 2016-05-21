# go-regionid

go-regionid is a library to represent a world in 3-level(country, subdivision and cities) and get a singleton region by name.
It is used by some IP locating library.

Example:

    regionid.LoadBuiltinWorld()
    location := regionid.GetLocation("中国", "上海", "")
