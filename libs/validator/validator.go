package validator

import (
	"net/http"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	ruTranslations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
	uni        *ut.UniversalTranslator
}

// NewCustomValidator создаёт новый объект кастомного валидатора
func NewCustomValidator() *CustomValidator {
	// Инициализация универсального переводчика с поддержкой английского и русского языков
	ruLocale := ru.New()
	enLocale := en.New()

	uni := ut.New(enLocale, enLocale, ruLocale) // enLocale по умолчанию, ruLocale как дополнительный

	// Получение переводчика для английского языка по умолчанию
	trans, _ := uni.GetTranslator("en")

	cv := &CustomValidator{
		validator:  validator.New(),
		uni:        uni,
		translator: trans,
	}

	// Регистрация переводов для валидатора
	_ = enTranslations.RegisterDefaultTranslations(cv.validator, trans)

	return cv
}

// Validate метод для валидации структуры без контекста (обязателен для Echo)
func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

// ValidateWithContext метод для валидации структуры с учётом локализации из контекста запроса
func (cv *CustomValidator) ValidateWithContext(i any, c echo.Context) error {
	// Определяем локализацию на основе заголовка Accept-Language
	lang := cv.getLocale(c)

	// Настраиваем переводчик на основе выбранного языка
	cv.setTranslator(lang)

	// Выполняем валидацию структуры
	if err := cv.validator.Struct(i); err != nil {
		// Локализуем сообщение об ошибке
		translatedError := cv.translateError(err)

		// Возвращаем локализованное сообщение об ошибке с кодом 400
		return echo.NewHTTPError(http.StatusBadRequest, translatedError)
	}

	return nil
}

// getLocale определяет локализацию на основе заголовка Accept-Language
func (cv *CustomValidator) getLocale(c echo.Context) string {
	tag, _ := language.MatchStrings(language.NewMatcher([]language.Tag{
		language.English,
		language.Russian,
	}), c.Request().Header.Get("Accept-Language"))

	if len(tag.String()) == 0 {
		return language.English.String()
	}

	return tag.String()
}

// setTranslator устанавливает переводчик на основе локализации
func (cv *CustomValidator) setTranslator(lang string) {
	// Получаем переводчик для языка
	trans, found := cv.uni.GetTranslator(lang)
	if found {
		cv.translator = trans
		// Регистрируем переводы для русского языка
		if lang == "ru" {
			_ = ruTranslations.RegisterDefaultTranslations(cv.validator, trans)
		} else {
			_ = enTranslations.RegisterDefaultTranslations(cv.validator, trans)
		}
	}
}

// translateError переводит ошибку на нужный язык
func (cv *CustomValidator) translateError(err error) []string {
	// Преобразуем ошибки в текст с учётом локализации
	errs := err.(validator.ValidationErrors)
	var translatedErrors []string
	for _, e := range errs {
		translatedErrors = append(translatedErrors, e.Translate(cv.translator))
	}
	return translatedErrors
}
