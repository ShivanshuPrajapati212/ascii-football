package models

type Player struct {
	Name  string `json:"name"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Goals int    `json:"goals"`
}

type Football struct {
	X         int     `json:"x"`
	Y         int     `json:"y"`
	XVelocity float32 `json:"x_velocity"`
	YVelocity float32 `json:"y_velocity"`
}

type Game struct {
	TeamA      []Player `json:"team_a"`
	TeamB      []Player `json:"team_b"`
	Ball       Football `json:"ball"`
	TeamAGoals int      `json:"team_a_goals"`
	TeamBGoals int      `json:"team_b_goals"`
	Target     int      `json:"target"`
}
