package model

type AddUserToGroupReq struct {
	RoleId   string   `json:"role_id"`
	GMIds    []string `json:"gm_ids"`
	GroupIds []string `json:"group_ids"`
}

type UserInfo struct {
	GMId             int64    `json:"gm_id,string"`
	MaxSessionNumber int      `json:"max_session_number"`
	PersonalTagIds   []string `json:"personal_tag_ids"`
}
