package cdrType

type SessionAMBR struct { /* Sequence Type */
	AmbrUL Bitrate `ber:"tagNum:1"`
	AmbrDL Bitrate `ber:"tagNum:2"`
}
