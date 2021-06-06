package constant

type RoleType int32

// role
const (
	Lord RoleType = iota
	Farmer
)

func (t RoleType) ToString() string {
	m := map[RoleType]string{}
	m[Lord] = "lord"
	m[Farmer] = "farmer"
	return m[t]
}

type StateType int32

// user state
const (
	Waiting StateType = iota // 等待
	Already                  // 准备就绪
	Playing                  // 游戏中
)

func (t StateType) ToString() string {
	m := map[StateType]string{}
	m[Waiting] = "waiting"
	m[Already] = "already"
	m[Playing] = "playing"
	return m[t]
}

// hooks
const (
	PLAYER_STATE_CHANGED = "PLAYER_STATE_CHANGED"
)
