package usecase

import (
	"github.com/iarsham/teacher-tool-api/configs"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"github.com/iarsham/teacher-tool-api/internal/models"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
)

type templateUsecase struct {
	templateRepository domain.TemplateRepository
	logger             *zap.Logger
	cfg                *configs.Config
}

func NewTemplateUsecase(templateRepository domain.TemplateRepository, logger *zap.Logger, cfg *configs.Config) domain.TemplateUsecase {
	return &templateUsecase{
		templateRepository: templateRepository,
		logger:             logger,
		cfg:                cfg,
	}
}

func (t *templateUsecase) GetObjID(r *http.Request) (uint64, error) {
	id, err := helpers.GetUIDParam(r)
	if err != nil {
		t.logger.Error(err.Error())
		return 0, err
	}
	return id, err
}

func (t *templateUsecase) FindAll() ([]*models.Templates, error) {
	templates, err := t.templateRepository.FindAll()
	if err != nil {
		t.logger.Error(err.Error())
		return nil, err
	}
	return templates, nil
}

func (t *templateUsecase) Create(template *entities.TemplateRequest, userID uint64) (*models.Templates, error) {
	template.UserID = userID
	tmpl, err := t.templateRepository.Create(template)
	if err != nil {
		t.logger.Error(err.Error())
		return nil, err
	}
	return tmpl, nil
}

func (t *templateUsecase) Delete(templateID uint64) error {
	if err := t.templateRepository.Delete(templateID); err != nil {
		t.logger.Error(err.Error())
	}
	return nil
}

func (t *templateUsecase) UploadFile(file multipart.File, folder, fileName string) (string, error) {
	link, err := helpers.UploadAwsS3(t.cfg, file, folder, fileName)
	if err != nil {
		t.logger.Error(err.Error())
		return "", err
	}
	return link, nil
}
