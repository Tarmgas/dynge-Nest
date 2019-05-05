// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2018, The The FRED Project
// Copyright (c) 2019, The The DYNGE Project

//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "DYNGE-walletd-session.log"
	logWalletdAllSessionsFilename        = "DYNGE-walletd.log"
	logTurtleCoindCurrentSessionFilename = "DYNGEdaemon-session.log"
	logTurtleCoindAllSessionsFilename    = "DYNGEdaemon.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "DYNGE-walletd"
	turtlecoindCommandName               = "DYNGEdaemon"
)
