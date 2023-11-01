package http

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandlers struct {
	userUC users.UseCase
}

// NewAuthHandlers Auth handlers constructor
func NewAuthHandlers(userUC users.UseCase) *userHandlers {
	return &userHandlers{userUC: userUC}
	//return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC, logger: log}
}

func (h *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "this is go!!!")
		//	span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "auth.Register")
		//	defer span.Finish()
		//
		//	user := &models.User{}
		//	if err := utils.ReadRequest(c, user); err != nil {
		//		utils.LogResponseError(c, h.logger, err)
		//		return c.JSON(httpErrors.ErrorResponse(err))
		//	}
		//
		//	createdUser, err := h.authUC.Register(ctx, user)
		//	if err != nil {
		//		utils.LogResponseError(c, h.logger, err)
		//		return c.JSON(httpErrors.ErrorResponse(err))
		//	}
		//
		//	sess, err := h.sessUC.CreateSession(ctx, &models.Session{
		//		UserID: createdUser.User.UserID,
		//	}, h.cfg.Session.Expire)
		//	if err != nil {
		//		utils.LogResponseError(c, h.logger, err)
		//		return c.JSON(httpErrors.ErrorResponse(err))
		//	}
		//
		//	c.SetCookie(utils.CreateSessionCookie(h.cfg, sess))
		//
		//	return c.JSON(http.StatusCreated, createdUser)
	}
}
