// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package define

const(
	HttpContentType = "Content-Type"
    HttpDataType = "application/json"
)
type HttpRetCode struct{
	Code             int                   `json:"code"`      // 0,成功，1 失败
	ErrMsg           string                `json:"errMsg"`    // 0 success，1，error msg
}

func NewHttpRetcode(code int, errMsg string) HttpRetCode {
	var retCode HttpRetCode
	retCode.Code = code
	retCode.ErrMsg = errMsg

	return retCode
}