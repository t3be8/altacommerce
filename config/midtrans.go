package config

import "github.com/midtrans/midtrans-go"

func SetupGlobalMidtransConfigApi() {
	midtrans.ServerKey = "SB-Mid-server-JRM7bi4-TPMjLFDEBmh0w-2h"
	// change value to `midtrans.Production`, if you want change the env to production
	midtrans.Environment = midtrans.Sandbox
}
