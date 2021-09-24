package types

import (
	"time"

	"github.com/shopspring/decimal"
)

// Article defined 文章
type Article struct {
	// ID defined 主键
	ID int64 `xorm:"bigint64(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// UnansweredCount defined 未答复数目
	UnansweredCount int64 `xorm:"comment('未答复数目') 'unanswered_count'" json:"unanswered_count" form:"unanswered_count" xml:"unanswered_count"`
	// BestAnswerersCount defined 最佳答复数目
	BestAnswerersCount int64 `xorm:"comment('最佳答复数目') 'best_answerers_count'" json:"best_answerers_count" form:"best_answerers_count" xml:"best_answerers_count"`
	// IsSuperTopicVote defined 是否超级话题投票
	IsSuperTopicVote int64 `xorm:"comment('是否超级话题投票') 'is_super_topic_vote'" json:"is_super_topic_vote" form:"is_super_topic_vote" xml:"is_super_topic_vote"`
	// Excerpt defined 摘录
	Excerpt string `xorm:"varchar(512) comment('摘录') 'excerpt'" json:"excerpt" form:"excerpt" xml:"excerpt"`
	// IsVote defined 是否投票
	IsVote int64 `xorm:"comment('是否投票') 'is_vote'" json:"is_vote" form:"is_vote" xml:"is_vote"`
	// IsBlack defined
	IsBlack int64 `xorm:"'is_black'" json:"is_black" form:"is_black" xml:"is_black"`
	// QuestionsCount defined 提问数目
	QuestionsCount int64 `xorm:"comment('提问数目') 'questions_count'" json:"questions_count" form:"questions_count" xml:"questions_count"`
	// Category defined 分类
	Category string `xorm:"varchar(36) comment('分类') 'category'" json:"category" form:"category" xml:"category"`
	// Name defined 标题
	Name string `xorm:"varchar(108) comment('标题') 'name'" json:"name" form:"name" xml:"name"`
	// Introduction defined 简介
	Introduction string `xorm:"varchar(512) comment('简介') 'introduction'" json:"introduction" form:"introduction" xml:"introduction"`
	// URL defined 地址
	URL string `xorm:"varchar(512) comment('地址') 'url'" json:"url" form:"url" xml:"url"`
	// FollowersCount defined 粉丝数
	FollowersCount int64 `xorm:"comment('粉丝数') 'followers_count'" json:"followers_count" form:"followers_count" xml:"followers_count"`
	// Type defined 类别
	Type string `xorm:"varchar(36) comment('类别') 'type'" json:"type" form:"type" xml:"type"`
	// Reward defined 打赏
	Reward decimal.Decimal `xorm:"decimal(6,2) comment('打赏') 'reward'" json:"reward" form:"reward" xml:"reward"`
	// Creater defined 创建人
	Creater int64 `xorm:"bigint64(20) notnull comment('创建人') 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateTime defined 创建时间
	CreateTime time.Time `xorm:"datetime comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// Updater defined 最后更新人
	Updater int64 `xorm:"bigint64(20) notnull comment('最后更新人') 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateTime defined 最后更新时间
	UpdateTime time.Time `xorm:"datetime comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// IsDelete defined 删除标记
	IsDelete int64 `xorm:"notnull comment('删除标记') 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// Remark defined 备注
	Remark string `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// TableName table name of defined Article
func (m *Article) TableName() string {
	return "article"
}
