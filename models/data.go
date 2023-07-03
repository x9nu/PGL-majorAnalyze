package models

type Res struct {
	Code   int
	Msg    string   `json:"message"`
	Result []Result `json:"result"`
}

type Result struct {
	Matchs       []Match     `json:"dtoList"`
	ItemCount    int         `json:"itemCount"`
	CsgoEventDTO interface{} `json:"csgoEventDTO"`
}

type Match struct {
	ID              int64       `json:"id"`                  // 比赛ID
	MatchID         int64       `json:"matchId"`             // 比赛ID
	NamiMatchID     interface{} `json:"namiMatchId"`         // Nami平台比赛ID
	EventID         int64       `json:"eventId"`             // 赛事ID
	StartTime       int64       `json:"startTime"`           // 比赛开始时间（时间戳）
	Team1ID         int64       `json:"team1Id"`             // 队伍1 ID
	Team1DTO        Team        `json:"team1DTO"`            // 队伍1 详细信息
	Team2ID         int64       `json:"team2Id"`             // 队伍2 ID
	Team2DTO        Team        `json:"team2DTO"`            // 队伍2 详细信息
	Score1          int         `json:"score1"`              // 队伍1 得分
	Score2          int         `json:"score2"`              // 队伍2 得分
	BO              string      `json:"bo"`                  // 比赛场次（BO1/BO3/BO5等）
	Star            int         `json:"star"`                // 星级评价（1-5星）
	WinnerTeamID    int64       `json:"winnerTeamId"`        // 获胜队伍 ID
	MatchType       int         `json:"matchType"`           // 比赛类型（1:普通比赛 2:重要比赛 3:大赛决赛等）
	Status          int         `json:"status"`              // 比赛状态（1:未开始 2:进行中 3:已结束 4:已取消）
	StatsDTOList    interface{} `json:"statsDTOList"`        // 战队统计信息列表
	CsgoEventDTO    interface{} `json:"csgoEventDTO"`        // CSGO赛事信息
	MatchDetailDTO  interface{} `json:"matchDetailDTO"`      // 比赛详细信息
	MapBPDTOS       interface{} `json:"mapBPDTOS"`           // 地图选择信息
	SingleMatchData interface{} `json:"singleMatchDataDTOS"` // 单场比赛信息
	SubscribeStatus interface{} `json:"subscribeStatus"`     // 订阅状态
	RoomID          interface{} `json:"roomId"`              // 房间ID
	Platform        interface{} `json:"platform"`            // 平台信息
	StageID         interface{} `json:"stageId"`             // 阶段ID
	NewsID          int64       `json:"newsId"`              // 相关新闻ID
	MatchIDList     interface{} `json:"matchIdList"`         // 相关比赛ID列表
	CompletedStatus interface{} `json:"completedStatus"`     // 完成状态
	HasPredict      bool        `json:"hasPredict"`          // 是否有比赛预测
}

type Team struct {
	ID            int        `json:"id"`        // 球队 ID
	TeamID        int        `json:"teamId"`    // 球队 ID（另一种格式）
	Name          string     `json:"name"`      // 球队名称
	LogoBlack     string     `json:"logoBlack"` // 球队黑色 logo
	LogoWhite     string     `json:"logoWhite"` // 球队白色 logo
	Rank          int        `json:"rank"`      // 球队排名
	Location      string     `json:"location"`  // 球队所在地
	PlayerDTOList []struct { // 球员列表
		ID          int    `json:"id"`          // 球员 ID
		Nickname    string `json:"nickname"`    // 球员昵称
		RealName    string `json:"realName"`    // 球员真实姓名
		Age         int    `json:"age"`         // 球员年龄
		Avatar      string `json:"avatar"`      // 球员头像
		Role        int    `json:"role"`        // 球员角色
		Score       int    `json:"score"`       // 球员评分
		Country     string `json:"country"`     // 球员所属国家
		TeamID      int    `json:"teamId"`      // 所属球队 ID
		TeamName    string `json:"teamName"`    // 所属球队名称
		Birthplace  string `json:"birthplace"`  // 球员出生地
		Description string `json:"description"` // 球员描述
	} `json:"playerDTOList"`
	IsSupportTeam int `json:"isSupportTeam"` // 是否支持球队
}
