package cdrType

type NgApCause struct { /* Sequence Type */
	Group int64 `ber:"tagNum:0"`
	Value int64 `ber:"tagNum:1"`
}
