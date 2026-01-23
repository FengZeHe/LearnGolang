package events

import (
	"log"
	"math"
	"time"

	"github.com/basicprojectv2/jobs/domain"
)

func ExecTimeKeeping(taskID uint) (err error) {
	log.Println("正在执行报时任务", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

const (
	HotDecayCoefficient = 0.1 // 热衰减系数
	ReadScore           = 1   // 阅读分权重
	LikeScore           = 5   // 点赞分权重
	CollectScore        = 10  //收藏分权重
)

func CalcHotScore(article domain.ArticleWithInteractive) float64 {
	// 基础分数
	baseScore := float64(article.ReadCount)*ReadScore + float64(article.LikeCount)*LikeScore + float64(article.CollectCount)*CollectScore

	layout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, article.CreatedAt, time.Local)
	if err != nil {
		return 0
	}
	pubishedDuration := time.Since(t).Hours()
	decayFactor := math.Exp(-HotDecayCoefficient * pubishedDuration)

	return baseScore * decayFactor
}
