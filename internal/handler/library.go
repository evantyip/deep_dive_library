package handler

import (
  "github.com/evantyip/deep_dive_library/internal/view/library"
  "github.com/labstack/echo/v4"
)

type LibraryHandler struct {}

func (lh *LibraryHandler) HandleShow (c echo.Context) error {
  return render(c, library.Show())
}
