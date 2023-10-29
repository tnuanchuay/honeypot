package honeypot

type Pot struct {
	Path     string `json:"path" form:"path" query:"path"`
	Redirect string `json:"redirect" form:"redirect" query:"redirect"`
	User     string `json:"user" form:"user" query:"user"`
}
