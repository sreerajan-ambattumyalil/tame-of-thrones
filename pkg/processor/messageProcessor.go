package processor

import (
	"github.com/sreerajan-ambattumyalil/tame-of-thrones/pkg/repository"
)

type Processor interface {
	Process()
}

type MessageProcessor struct {
	kingdomsRepository repository.KingdomRepository
}

func NewMessageProcessor(kingdomsRepository repository.KingdomRepository) MessageProcessor {
	return MessageProcessor{kingdomsRepository: kingdomsRepository}
}

func (p MessageProcessor) Process(message chan string) {

}
