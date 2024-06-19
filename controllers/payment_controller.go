package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-crud-app/models"
	"net/http"
	"regexp"
)

func ValidateCard(c echo.Context) error {
	card := new(models.Card)
	if err := c.Bind(card); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cardNumberPattern := `^\d{16}$`
	expireDatePattern := `^\d{2}\/\d{2}$`
	securityCodePattern := `^\d{3}$`

	isCardNumberValid, _ := regexp.MatchString(cardNumberPattern, card.Number)
	isExpireDateValid, _ := regexp.MatchString(expireDatePattern, card.ExpireDate)
	isSecurityCodeValid, _ := regexp.MatchString(securityCodePattern, card.SecurityCode)

	if !isCardNumberValid || !isExpireDateValid || !isSecurityCodeValid {
		return c.JSON(http.StatusBadRequest, "Invalid card details")
	}

	return c.JSON(http.StatusOK, "Card details are valid")
}
