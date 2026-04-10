package authentication

import (
	"context"
	"errors"
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/raymondsugiarto/funder-api/config"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
	shared "github.com/raymondsugiarto/funder-api/shared/context"
	"github.com/raymondsugiarto/funder-api/shared/security"
)

const ServiceName = "authenticationService"

type Service interface {
	GetSession(ctx context.Context) (*entity.UserSessionDto, error)
	SignIn(context.Context, *entity.LoginRequestDto) (*entity.LoginDto, error)
}

type service struct {
	userCredentialService usercredential.Service
}

func NewService(
	userCredentialService usercredential.Service,
) Service {
	return &service{
		userCredentialService: userCredentialService,
	}
}

func (s *service) GetSession(ctx context.Context) (*entity.UserSessionDto, error) {
	token := jwtware.FromContext(ctx)
	claims := token.Claims.(jwt.MapClaims)
	userSessionDto := entity.NewUserSessionDtoFromClaims(claims)

	userCredentialDto, err := s.userCredentialService.FindByID(ctx, userSessionDto.ID)
	if err != nil {
		return nil, errors.New("userNotFound")
	}
	userSessionDto.UserCredential = userCredentialDto

	return userSessionDto, nil
}

func (s *service) SignIn(ctx context.Context, request *entity.LoginRequestDto) (*entity.LoginDto, error) {
	log.WithContext(ctx).Infof("sign in started")
	organizationID := fiber.Locals[*entity.OrganizationDto](ctx.(fiber.Ctx), entity.OrganizationKey).ID
	userCredentialModel, err := s.userCredentialService.GetUserCredentialByUsername(ctx, organizationID, request.Username)
	if err != nil {
		return nil, errors.New("userNotFound")
	}

	userCredentialData := shared.UserCredentialData{
		ID:     userCredentialModel.ID,
		UserID: userCredentialModel.UserID,
	}

	pp, _ := security.HashPassword(request.Password)
	log.WithContext(ctx).Infof("password: %s, hash: %s", request.Password, pp)
	if !security.CheckPasswordHash(request.Password, userCredentialModel.Password) {
		return nil, errors.New("invalidPassword")
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userCredentialData.ID
	claims["uid"] = userCredentialData.UserID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	cfg := config.GetConfig()
	t, err := token.SignedString([]byte(cfg.Server.Rest.SecretKey))
	if err != nil {
		return nil, errors.New("errorGeneratetoken")
	}
	return &entity.LoginDto{
		Token:   t,
		Expired: strconv.Itoa(int(claims["exp"].(int64))),
	}, nil
}
