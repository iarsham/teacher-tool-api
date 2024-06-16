package usecase

import (
	"fmt"
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
)

type questionsUsecase struct {
	questionsRepository domain.QuestionsRepository
	logger              *zap.Logger
	cfg                 *configs.Config
}

func NewQuestionsUsecase(questionsRepository domain.QuestionsRepository, logger *zap.Logger, cfg *configs.Config) domain.QuestionsUsecase {
	return &questionsUsecase{
		questionsRepository: questionsRepository,
		logger:              logger,
		cfg:                 cfg,
	}
}

func (q *questionsUsecase) FindAll() ([]*models.Questions, error) {
	qs, err := q.questionsRepository.FindAll()
	if err != nil {
		q.logger.Error(err.Error())
		return nil, err
	}
	return qs, nil
}
func (q *questionsUsecase) FindByFile(file *multipart.FileHeader) (*models.Questions, error) {
	path := helpers.DstPath("questions", file.Filename)
	link := helpers.CreateS3Url(q.cfg, path)
	fmt.Println(link)
	qs, err := q.questionsRepository.FindByFile(link)
	if err != nil {
		q.logger.Error(err.Error())
		return nil, err
	}
	return qs, nil
}

func (q *questionsUsecase) FindByID(id uint64) (*models.Questions, error) {
	qs, err := q.questionsRepository.FindByID(id)
	if err != nil {
		q.logger.Error(err.Error())
		return nil, err
	}
	return qs, nil
}

func (q *questionsUsecase) GetObjID(r *http.Request) (uint64, error) {
	id, err := helpers.GetUIDParam(r)
	if err != nil {
		q.logger.Error(err.Error())
		return 0, err
	}
	return id, err
}

func (q *questionsUsecase) GetUserID(r *http.Request) uint64 {
	return helpers.GetUserID(r)
}

func (q *questionsUsecase) Create(question *entities.QuestionRequest, link string, userID uint64) (*models.Questions, error) {
	qs, err := q.questionsRepository.Create(question, link, userID)
	if err != nil {
		q.logger.Error(err.Error())
		return nil, err
	}
	return qs, nil
}

func (q *questionsUsecase) Delete(id uint64) error {
	if err := q.questionsRepository.Delete(id); err != nil {
		q.logger.Error(err.Error())
		return err
	}
	return nil
}

func (q *questionsUsecase) UploadFile(file multipart.File, folder, fileName string) (string, error) {
	link, err := helpers.UploadAwsS3(q.cfg, file, folder, fileName)
	if err != nil {
		q.logger.Error(err.Error())
		return "", err
	}
	return link, nil
}
