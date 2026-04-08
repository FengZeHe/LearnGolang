package fun

import (
	"context"
	"log"
	"time"

	"github.com/basicprojectv2/jobs/events"
	"github.com/basicprojectv2/jobs/repository/dao"
	"github.com/redis/go-redis/v9"
)

func ExecTimeKeeping(taskID uint) (err error) {
	log.Println("正在执行报时任务", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func SayHi(taskID uint) error {
	log.Printf("【任务%v】Hello! 我是 SayHi 任务", taskID)
	return nil
}

func CalcHotList(taskID uint, taskDA0 dao.TaskDAO, rdb redis.Cmdable) (err error) {
	log.Println(taskID, ":正在重新计算热榜")
	pageSize := 1000
	index := 1
	for {
		alist, err := taskDA0.GetArticleIDs(index, pageSize)
		if err != nil {
			log.Println("err", err)
			return err
		}
		if len(alist) == 0 {
			break
		}

		//  每计算完1000篇就写入redis,清空缓存区
		for _, art := range alist {
			//  计算分数
			artScore := events.CalcHotScore(art)

			_, err := rdb.ZAdd(context.Background(), "hotlist/articles/score/", redis.Z{Score: artScore, Member: art.ID}).Result()
			if err != nil {
				log.Println("err", err)
				return err
			}

			_, err = rdb.HSet(context.Background(), "hotlist/articles/"+art.ID, "title", art.Title, "score", 0.1).Result()
			if err != nil {
				log.Println("err", err)
				return err
			}
		}

		index++
	}
	return nil
}
