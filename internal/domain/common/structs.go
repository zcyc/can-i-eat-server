package common_domain

type BhList []*Bh
type Bh struct {
	Name    string   `json:"name"`
	TagList []string `json:"tagList"`
}
