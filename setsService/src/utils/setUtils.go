package utils

import (
	"net/url"
	"reflect"
	"strconv"

	"github.com/jedzeins/jlpt_api/setsService/src/models"
	"gopkg.in/mgo.v2/bson"
)

func ParseSetUrlId(reqURL *url.URL) (*models.SetRequestById, *models.SetRequestParamsError) {

	encodedId := reqURL.Query().Get("id")
	decodedId, err := url.QueryUnescape(encodedId)
	if err != nil {
		return nil, &models.SetRequestParamsError{ErrorMessage: "issue decoding the ID"}
	}

	return &models.SetRequestById{Id: decodedId}, nil

}

func ParseSetUrlCheckStrings(reqURL *url.URL) (*models.SetRequestParamsUnParsed, *models.SetRequestParamsError) {

	var params models.SetRequestParamsUnParsed
	var paramsError models.SetRequestParamsError

	encodedId := reqURL.Query().Get("id")
	decodedId, err := url.QueryUnescape(encodedId)
	if err != nil {
		return nil, &models.SetRequestParamsError{ErrorMessage: "issue decoding the ID"}
	}

	params.Id = decodedId

	encodedSetName := reqURL.Query().Get("setName")
	decodedSetName, err := url.QueryUnescape(encodedSetName)
	if err != nil {
		paramsError.ErrorMessage = "issue decoding the setNameParam"
		return nil, &paramsError
	}

	params.SetName = decodedSetName

	isPublicString := reqURL.Query().Get("isPublic")

	if isPublicString != "" {
		_, err := strconv.ParseBool(isPublicString)
		if err != nil {
			paramsError.ErrorMessage = "isPublic must be `true` or `false`"
			return nil, &paramsError
		}
		params.IsPublic = isPublicString
	}

	isMineString := reqURL.Query().Get("isMine")
	if isMineString != "" {
		_, err := strconv.ParseBool(isMineString)
		if err != nil {
			paramsError.ErrorMessage = "isMine must be `true` or `false`"
			return nil, &paramsError
		}
		params.IsMine = isMineString
	}

	limitString := reqURL.Query().Get("limit")
	if limitString != "" {
		_, err := strconv.Atoi(limitString)
		if err != nil {
			paramsError.ErrorMessage = "limit must be an integer"
			return nil, &paramsError
		}
		params.Limit = limitString
	} else {
		params.Limit = "20"
	}

	offsetString := reqURL.Query().Get("offset")
	if offsetString != "" {
		_, err := strconv.Atoi(offsetString)
		if err != nil {
			paramsError.ErrorMessage = "ofsset must be an integer"
			return nil, &paramsError
		}

		params.Offset = offsetString
	} else {
		params.Offset = "20"
	}

	return &params, nil
}

func SetToBson(set *models.Set) (bson.M, error) {
	queryObject := bson.M{}

	v := reflect.ValueOf(*set)
	typeOfURLParams := v.Type()
	for i := 0; i < v.NumField(); i++ {

		name := LowerCaseString(typeOfURLParams.Field(i).Name)
		val := v.Field(i).Interface()

		if name == "iD" {
			continue
		}

		queryObject[name] = val

	}

	return queryObject, nil

}

// func ParseSetUrl(reqURL *url.URL) (*models.SetRequestParams, *models.SetRequestParamsError) {

// 	var params models.SetRequestParams
// 	var paramsError models.SetRequestParamsError

// 	encodedSetName := reqURL.Query().Get("setName")
// 	decodedSetName, err := url.QueryUnescape(encodedSetName)
// 	if err != nil {
// 		paramsError.ErrorMessage = "issue decoding the setNameParam"
// 		return nil, &paramsError
// 	}

// 	params.SetName = decodedSetName

// 	isPublicString := reqURL.Query().Get("isPublic")

// 	if isPublicString != "" {
// 		isPublic, err := strconv.ParseBool(isPublicString)
// 		if err != nil {
// 			paramsError.ErrorMessage = "isPublic must be `true` or `false`"
// 			return nil, &paramsError
// 		}
// 		params.IsPublic = isPublic
// 	}

// 	isMineString := reqURL.Query().Get("isMine")
// 	if isMineString != "" {
// 		isMine, err := strconv.ParseBool(isMineString)
// 		if err != nil {
// 			paramsError.ErrorMessage = "isMine must be `true` or `false`"
// 			return nil, &paramsError
// 		}
// 		params.IsMine = isMine
// 	}

// 	limitString := reqURL.Query().Get("limit")
// 	if limitString != "" {
// 		limit, err := strconv.Atoi(limitString)
// 		if err != nil {
// 			paramsError.ErrorMessage = "limit must be an integer"
// 			return nil, &paramsError
// 		}
// 		params.Limit = limit
// 	} else {
// 		params.Limit = 20
// 	}

// 	offsetString := reqURL.Query().Get("offset")
// 	if offsetString != "" {
// 		offset, err := strconv.Atoi(offsetString)
// 		if err != nil {
// 			paramsError.ErrorMessage = "ofsset must be an integer"
// 			return nil, &paramsError
// 		}

// 		params.Offset = offset
// 	} else {
// 		params.Offset = 20
// 	}

// 	return &params, nil
// }
