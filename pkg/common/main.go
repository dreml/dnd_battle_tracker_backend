package common

type MonsterInfoResponse struct {
	Count   int           `json:"count"`
	Results []MonsterInfo `json:"results"`
}

type MonsterInfo struct {
	Index string `json:"index"`
	Name  string `json:"name"`
}

type Monster struct {
	Index     string `json:"index"`
	Name      string `json:"name"`
	HitPoints int    `json:"hit_points"`
	Image     string `json:"image"`
}
