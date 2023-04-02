package delivery

import (
	"context"
	"net/http"

	jsonUser "otus/internal/delivery/user"
	"otus/internal/domain"
	"otus/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser
// @Summary Создание пользователя.
// @Description Создание пользователя.
// @Tags user
// @Accept json
// @Produce json
// @Param data body jsonUser.CreateUserRequest true "Body"
// @Success 201			{object} 	jsonUser.UserResponse
// @Failure 400 		{object}    ErrorResponse
// @Failure 500 		{object}    ErrorResponse
// @Router /user/ [put]
func (d *Delivery) CreateUser(c *gin.Context) {

	request := jsonUser.CreateUserRequest{}
	if err := c.ShouldBind(&request); err != nil {
		SetError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := domain.NewUser(request.UserName, request.FirstName, request.LastName, request.Email, request.Phone)
	if err != nil {
		SetError(c, http.StatusBadRequest, err.Error())
	}
	user, err = d.service.CreateUser(context.Background(), user)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, d.toUserResponse(user))
}

// UpdateUser
// @Summary Обновление пользователя.
// @Description Обновление пользователя.
// @Tags user
// @Accept json
// @Produce json
// @Param data body jsonUser.UpdateUserRequest true "Body"
// @Success 200			{object} 	jsonUser.UserResponse
// @Failure 400 		{object}    ErrorResponse
// @Failure 500 		{object}    ErrorResponse
// @Router /user/ [post]
func (d *Delivery) UpdateUser(c *gin.Context) {

	request := jsonUser.UpdateUserRequest{}
	if err := c.ShouldBind(&request); err != nil {
		SetError(c, http.StatusBadRequest, err.Error())
		return
	}

	updateUserParams := &service.UserUpdateParams{
		Id:        request.Id,
		UserName:  request.UserName,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}

	user, err := d.service.UpdateUser(context.Background(), updateUserParams)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, d.toUserResponse(user))
}

// DeleteUser
// @Summary Удаление пользователя.
// @Description Удаление пользователя.
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400 		{object}    ErrorResponse
// @Failure 500 		{object}    ErrorResponse
// @Router /user/{id} [delete]
func (d *Delivery) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		SetError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = d.service.DeleteUser(context.Background(), userId)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// ReadUserById
// @Summary Получаем пользователя по id.
// @Description Получаем пользователя по id.
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200			{object} 	jsonUser.UserResponse
// @Failure 400 		{object}    ErrorResponse
// @Failure 500 		{object}    ErrorResponse
// @Router /user/{id} [get]
func (d *Delivery) ReadUserById(c *gin.Context) {

	id := c.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		SetError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := d.service.ReadUserById(context.Background(), userId)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, d.toUserResponse(user))
}
