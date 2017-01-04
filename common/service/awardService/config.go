package awardService

const (
	AWARD_ONLINE_READIS_KEY = "award_online_readis_key"
)

func init() {
	//配置抽成
	OnlineConfig.Rate = 30

	//配置最低奖励
	OnlineConfig.BaseAward = 99
	OnlineConfig.LeastLevel = 1
	OnlineConfig.MaxLevel = 15

	//配置时间间隔
	OnlineConfig.TimeConfig = make(map[int32]int64)
	OnlineConfig.TimeConfig[1] = 9
	OnlineConfig.TimeConfig[2] = 60 * 1
	OnlineConfig.TimeConfig[3] = 60 * 2
	OnlineConfig.TimeConfig[4] = 60 * 3
	OnlineConfig.TimeConfig[5] = 60 * 4
	OnlineConfig.TimeConfig[6] = 60 * 6
	OnlineConfig.TimeConfig[7] = 60 * 8
	OnlineConfig.TimeConfig[8] = 60 * 10
	OnlineConfig.TimeConfig[9] = 60 * 15
	OnlineConfig.TimeConfig[10] = 60 * 18
	OnlineConfig.TimeConfig[11] = 60 * 18
	OnlineConfig.TimeConfig[12] = 60 * 20
	OnlineConfig.TimeConfig[13] = 60 * 20
	OnlineConfig.TimeConfig[14] = 60 * 25
	OnlineConfig.TimeConfig[15] = 60 * 25

	//配置奖励
	OnlineConfig.AwardLimieConfig = make(map[int32]int64)
	OnlineConfig.AwardLimieConfig[1] = 1500
	OnlineConfig.AwardLimieConfig[2] = 1500
	OnlineConfig.AwardLimieConfig[3] = 1500
	OnlineConfig.AwardLimieConfig[4] = 1500
	OnlineConfig.AwardLimieConfig[5] = 1500
	OnlineConfig.AwardLimieConfig[6] = 1500
	OnlineConfig.AwardLimieConfig[7] = 1500
	OnlineConfig.AwardLimieConfig[8] = 2000
	OnlineConfig.AwardLimieConfig[9] = 2000
	OnlineConfig.AwardLimieConfig[10] = 2000
	OnlineConfig.AwardLimieConfig[11] = 2000
	OnlineConfig.AwardLimieConfig[12] = 2000
	OnlineConfig.AwardLimieConfig[13] = 2000
	OnlineConfig.AwardLimieConfig[14] = 2000
	OnlineConfig.AwardLimieConfig[15] = 2000

}

var OnlineConfig struct {
	TimeConfig       map[int32]int64 //时间间隔
	AwardLimieConfig map[int32]int64 //奖励的限制
	BaseAward        int64           //最低奖励
	Rate             int64           //抽成
	LeastLevel       int32           //最小的等级
	MaxLevel         int32           //最大的等级
}
