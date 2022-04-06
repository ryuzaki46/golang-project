package cart

import (
	"fmt"
	"group-project/limamart/delivery/helper"
	_cartUseCase "group-project/limamart/usecase/cart"
	"net/http"
	"strconv"

	_middlewares "group-project/limamart/delivery/middlewares"
	_entities "group-project/limamart/entities"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase _cartUseCase.CartUseCaseInterface
}

func NewCartHandler(u _cartUseCase.CartUseCaseInterface) CartHandler {
	return CartHandler{
		cartUseCase: u,
	}
}

func (uh *CartHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		
		carts, err := uh.cartUseCase.GetAll(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all carts", carts))
	}
}

func (uh *CartHandler) CreateCartHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		var param _entities.Cart

	err := c.Bind(&param)
	idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		param.UserID = uint(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		carts, err := uh.cartUseCase.CreateCart(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create cart", carts))
	}
}

func (uh *CartHandler) UpdateCartHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		var param _entities.Cart
		id, _ := strconv.Atoi(c.Param("id"))

		getid, err := uh.cartUseCase.GetCartById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}


		fmt.Println("id param user id", getid.UserID)

		if uint(idToken) != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized2"))
		}

		
		err = c.Bind(&param)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		carts, err := uh.cartUseCase.UpdateCart(id, param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update cart", carts))
	}
}

func (uh *CartHandler) DeleteCartHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		
		id, _ := strconv.Atoi(c.Param("id"))

		getid, err := uh.cartUseCase.GetCartById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		fmt.Println("id param user id", getid.UserID)

		if uint(idToken) != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		
		err = uh.cartUseCase.DeleteCart(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success delete cart", err))
	}
}