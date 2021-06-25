package service

import "context"

type AppJSONService struct {
}

func NewAppJSONService() *AppJSONService {
	return &AppJSONService{}
}

func (a *AppJSONService) GetApps(c context.Context) (string, error) {
	return "ok", nil
}

func (a *AppJSONService) GetGMs(c context.Context) (string, error) {
	gms := `{
		"error": {},
		"total": 3,
		"gms": [
			{
				"gm_id": "1234",
				"name": "李丹",
				"belong": "莉莉丝游戏",
				"role": {
						"role_id": "12334",
						"name": "管理员",
						"slug": "admin",
						"description": "拥有项目所有权限"
				},
				"groups": [
					{"name": "小语种客服", "group_id": "18701929304"},
					{"name": "小语种客", "group_id": "18701929305"}
        ],
				"created_at": 150000000
			}
		]
	}`

	return gms, nil
}

func (a *AppJSONService) GetGroups(c context.Context) (string, error) {
	groups := `{
		"error": {},
		"groups": [
			{
				"group_id": "12344",
				"name": "客服小组",
				"app_id": "1233434",
				"tag_ids": [
					"sass"
				],
				"created_at": "162324334"
			},
			{
				"group_id": "12345",
				"name": "客服小",
				"app_id": "1233434",
				"tag_ids": [
					"sass"
				],
				"created_at": "162324335"
			}
		]
	}`

	return groups, nil
}

func (a *AppJSONService) GetUserInfo(c context.Context) (string, error) {
	userInfo := `{
		"error": {},
		"personal_tag_ids": ["1", "2", "3"],
		"app_tags": [
			{
				"id": "1",
				"name": "1语"
			},
			{
				"id": "2",
				"name": "2语"
			},
			{
				"id": "3",
				"name": "3语"
			},
			{
				"id": "4",
				"name": "4语"
			},
			{
				"id": "5",
				"name": "5语"
			}
		],
		"max_session_number": 5
	}`
	return userInfo, nil
}
