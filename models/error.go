package models

type TutorInfoError struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
}

var NoError = TutorInfoError{Code: 00000, Msg: "no error"}

var NotFoundError = TutorInfoError{Code: 20401, Msg: "Shortening url not found"}

var BadRequestError = TutorInfoError{Code: 40001, Msg: "Bad request"}

var InternalServerError = TutorInfoError{Code: 50001, Msg: "Internal server error"}
