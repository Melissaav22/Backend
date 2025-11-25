package validators

import (
	"VetiCare/entities/dto"
	"errors"
)

var (
	ErrInvalidFullNameUser = errors.New("EL nombre completo debe contener solo letras")
	ErrInvalidDUIUser      = errors.New("DUI inválido: formato esperado ########-#")
	ErrInvalidPhoneUser    = errors.New("EL numero de teléfono es inválido, formato esperado ####-####")
	ErrInvalidEmailUser    = errors.New("LA direccion de correo es invalida")
	ErrInvalidPf           = errors.New("El codigo de foto es invalido")
)

func ValidateUserDTO(user dto.UserDTO) error {
	if err := ValidateFullName(user.FullName); err != nil {
		return ErrInvalidFullNameUser
	}
	if err := ValidateDUI(user.DUI); err != nil {
		return ErrInvalidDUIUser
	}
	if err := ValidatePhone(user.Phone); err != nil {
		return ErrInvalidPhoneUser
	}
	if err := ValidateEmail(user.Email); err != nil {
		return ErrInvalidEmailUser
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}
	return nil
}

func ValidateUpdatedUserDTO(user dto.UpdateUserDTO) error {
	if err := ValidateUpdatedName(user.FullName); err != nil {
		return ErrInvalidFullNameUser
	}
	if err := ValidateUpdatedDUI(user.DUI); err != nil {
		return ErrInvalidDUIUser
	}
	if err := ValidateUpdatedPhone(user.Phone); err != nil {
		return ErrInvalidPhoneUser
	}
	if err := ValidateUpdatedEmail(user.Email); err != nil {
		return ErrInvalidEmailUser
	}
	if err := ValidateUpdatedPf(user.Pf); err != nil {
		return ErrInvalidPf
	}
	return nil
}
