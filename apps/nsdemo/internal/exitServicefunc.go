//
// Copyright (c) 2018 Tencent
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package internal

import (
	"fmt"
	"os"
)

func ExitService(err error) {
	Destruct()
	NsplusSdk.LoggingClient.Error(fmt.Sprintf("terminating service with error[%v]", err))
	//close(errChan)
	os.Exit(1)
}
