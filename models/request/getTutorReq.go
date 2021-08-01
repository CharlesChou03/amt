package models

type GetTutorReq struct {
	Tutor string `json:"-" uri:"tutor"`
}

func (r *GetTutorReq) Validate() bool {
	if r.Tutor == "" {
		return false
	}
	return true
}
