package article

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/basicprojectv2/internal/repository/dao"
	"log"
)

// 定义consumer 接口
type Consumer interface {
	ConsumeReadEvent() error
}

// SaramaConsumer 实现Consumer接口
type SaramaConsumer struct {
	consumer   sarama.Consumer
	articleDAO dao.ArticleDAO
}

// 创建一个Sarama消费者
func NewSaramaConsumer(consumer sarama.Consumer, articleDAO dao.ArticleDAO) Consumer {
	return &SaramaConsumer{
		consumer:   consumer,
		articleDAO: articleDAO,
	}
}

// ConsumerReadEvent消费
func (s *SaramaConsumer) ConsumeReadEvent() error {
	partitions, err := s.consumer.ConsumePartition(TopicReadEvent, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Error consuming partition %d: %v \n", TopicReadEvent, err)
		return err
	}
	defer partitions.Close()

	for msg := range partitions.Messages() {
		var evt ReadEvent
		if err := json.Unmarshal(msg.Value, &evt); err != nil {
			log.Printf("Error unmarshalling message: %v \n", err)
			continue
		}
		log.Println("Consume Read event:", evt, evt.Aid, evt.Uid)
		// 更新 MySQL 的 read 字段加 1
		if err := s.updateArticleReadCount(context.Background(), evt.Aid); err != nil {
			log.Printf("failed to update article read count: %v", err)
		}
	}
	return nil
}

// updateArticleReadCount 更新文章阅读计数
func (s *SaramaConsumer) updateArticleReadCount(ctx context.Context, articleID int64) error {
	idStr := fmt.Sprintf("%d", articleID)
	if err := s.articleDAO.AddArticleCount(ctx, idStr); err != nil {
		return fmt.Errorf("failed to add article count: %w", err)
	}
	return nil
}
