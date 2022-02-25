package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

//投票功能:
//1.
/*投票的几种情况
direction=1
	1.之前0，现在1
	2.之前-1，现在1
direction=0
	1.之前1，现在0
	2.之前-1，现在0
direction=-1
	1.之前0，现在-1
	2.之前1，现在-1

投票限制:
每个帖子自发表之日一个星期之内允许投票，超过一星期不允许投票
	1. 到期之后将redis中保存的赞成票数及反对票存储在mysql中
	2. 到期之后删除KsyPostVotedZSetPF
*/

//VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug(
		"VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}
