package req

import "github.com/labstack/echo/v4"

func ParseReq(c echo.Context) (echo.Map, error) {
	data := echo.Map{}
	err := c.Bind(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
